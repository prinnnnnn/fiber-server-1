package repository

/*

UserRepository implements port.UserRepository interface (port's spec)
and provide an access to the postgres database

*/

import (
	"context"
	"errors"
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

	userFriends := make([]models.UserResponse, 0, len(friendships))

	var wg sync.WaitGroup
	var mu sync.Mutex
	var friendId, friendId2 uint

	for rows.Next() {

		if err := rows.Scan(&friendId, &friendId2); err != nil {
			return nil, models.ErrInternalServer
		}

		if friendId == id {
			friendId = friendId2
		}

		wg.Add(1)
		go func(userId uint) {
			defer wg.Done()
			var user models.User
			query := `SELECT * FROM users WHERE id = $1`

			err := ur.db.QueryRow(ctx, query, userId).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.FirstName,
				&user.LastName, &user.Email, &user.Password, &user.Location,
				&user.Occupation, &user.ViewedProfile, &user.Impressions)

			if err != nil {
				return
			}

			mu.Lock()
			userFriends = append(userFriends, *models.MapToResponse(&user))
			mu.Unlock()

		}(friendId)

	}

	wg.Wait()

	return userFriends, nil
}

func (ur *UserRepository) ToggleFriendStatus(ctx context.Context, id, friendId uint) ([]models.UserResponse, error) {

	friendship, err := ur.getFriendship(ctx, id, friendId)

	if errors.Is(err, models.ErrInternalServer) {
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
