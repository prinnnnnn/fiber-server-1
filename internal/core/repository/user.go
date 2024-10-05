package repository

/*

UserRepository implements port.UserRepository interface (port's spec)
and provide an access to the postgres database

*/

import (
	"context"
	"fiber-server-1/internal/core/models"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.UserResponse, error) {

	return models.MapToResponse(user), nil

}

func (ur *UserRepository) GetUserById(ctx context.Context, id uint) (*models.UserResponse, error) {

	var user models.User
	query := `SELECT * FROM users WHERE id = $1`

	err := ur.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.FirstName,
		&user.LastName, &user.Email, &user.Password, &user.Location,
		&user.Occupation, &user.ViewedProfile, &user.Impressions)

	if err != nil {

		// handle not found
		if err == pgx.ErrNoRows {
			return nil, models.ErrDataNotFound
		}

		// Other errors
		return nil, models.ErrInternalServer
	}

	return models.MapToResponse(&user), nil

}

func (ur *UserRepository) GetUserFriends(ctx context.Context, id uint) ([]models.UserResponse, error) {
	var friendships []models.Friendship

	// Select friendships where the user is either in user_id1 or user_id2
	query := `
        SELECT user_id1, user_id2 
        FROM friendships 
        WHERE user_id1 = $1 OR user_id2 = $1;
    `

	rows, err := ur.db.Query(ctx, query, id)
	if err != nil {
		return nil, models.ErrInternalServer
	}
	defer rows.Close()

	// Iterate rows for each friendships
	for rows.Next() {

		var friendship models.Friendship
		if err := rows.Scan(&friendship.UserID1, &friendship.UserID2); err != nil {
			return nil, models.ErrInternalServer
		}
		friendships = append(friendships, friendship)

	}

	if err := rows.Err(); err != nil {
		return nil, models.ErrInternalServer
	}

	userFriends := make([]models.UserResponse, 0, len(friendships))

	var wg sync.WaitGroup
	var mu sync.Mutex

	var friendId uint
	for _, fsh := range friendships {

		fsh_ := fsh

		wg.Add(1)
		go func() {
			defer wg.Done()

			if fsh_.UserID1 == id {
				friendId = fsh_.UserID2
			} else {
				friendId = fsh_.UserID1
			}

			user, err := ur.GetUserById(ctx, friendId)

			if err != nil {
				return
			}

			mu.Lock()
			userFriends = append(userFriends, *user)
			mu.Unlock()

		}()

		// if fsh.UserID1 == id {
		// 	friendId = fsh.UserID2
		// } else {
		// 	friendId = fsh.UserID1
		// }

		// user, err := ur.GetUserById(ctx, friendId)

		// if err != nil {
		// 	return nil, err
		// }

		// userFriends = append(userFriends, *user)

	}

	wg.Wait()

	return userFriends, nil
}

func (ur *UserRepository) ToggleFriendStatus(ctx context.Context, id, friendId uint) ([]models.UserResponse, error) {

	friendship, err := ur.getFriendship(ctx, id, friendId)

	if err != nil {
		return nil, err
	}

	if friendship != nil {
		// They're friends => remove
		err := ur.deleteFriendship(ctx, id, friendId)

		if err != nil {
			return nil, err
		}

	} else {
		// They are not friends => add
		err := ur.createFriendship(ctx, id, friendId)

		if err != nil {
			return nil, err
		}

	}

	return ur.GetUserFriends(ctx, id)

}

// GetFriendship checks if a friendship exists
func (ur *UserRepository) getFriendship(ctx context.Context, userID1, userID2 uint) (*models.Friendship, error) {
	var friendship models.Friendship

	query := `
		SELECT user_id1, user_id2 
        FROM friendships 
        WHERE (user_id1 = $1 AND user_id2 = $2) OR (user_id2 = $1 AND user_id1 = $2);
	`

	row := ur.db.QueryRow(ctx, query, userID1, userID2)
	err := row.Scan(&friendship.UserID1, &friendship.UserID2)

	if err != nil {
		// Handle the "not found" case
		if err == pgx.ErrNoRows {
			return nil, models.ErrDataNotFound
		}
		return nil, models.ErrInternalServer
	}

	return &friendship, nil
}

func (ur *UserRepository) createFriendship(ctx context.Context, userID1, userID2 uint) error {

	query := `
		INSERT INTO friendships (user_id1, user_id2)
		VALUES ($1, $2)
	`

	_, err := ur.db.Exec(ctx, query, userID1, userID2)
	if err != nil {
		return models.ErrInternalServer
	}

	return nil
}

func (ur *UserRepository) deleteFriendship(ctx context.Context, userID1, userID2 uint) error {

	query := `
		DELETE FROM friendships
		WHERE (user_id1 = $1 AND user_id2 = $2) OR (user_id1 = $2 AND user_id2 = $1)
	`

	result, err := ur.db.Exec(ctx, query, userID1, userID2)
	if err != nil {
		return models.ErrInternalServer
	}

	if result.RowsAffected() == 0 {
		return models.ErrDataNotFound
	}

	return nil
}
