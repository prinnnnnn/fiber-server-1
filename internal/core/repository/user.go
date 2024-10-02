package repository

/*

UserRepository implements port.UserRepository interface (port's spec)
and provide an access to the postgres database

*/

import (
	"context"
	"errors"
	"fiber-server-1/internal/core/models"
	"fmt"

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

	// Execute the query with named arguments to fetch the user details from the database.
	err := ur.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.FirstName,
		&user.LastName, &user.Email, &user.Password, &user.Location,
		&user.Occupation, &user.ViewedProfile, &user.Impressions)

	if err != nil {

		// handle not found
		if err.Error() == "no rows in result set" {
			return nil, errors.New("user not found")
		}

		// Other errors
		return nil, fmt.Errorf("could not fetch user: %w", err)
	}

	return models.MapToResponse(&user), nil

}

func (ur *UserRepository) GetUserFriends(ctx context.Context, id uint) ([]models.UserResponse, error) {
	var friendships []models.Friendship

	// SQL Query: Select friendships where the user is either in user_id1 or user_id2
	query := `
        SELECT user_id1, user_id2 
        FROM friendships 
        WHERE user_id1 = $1 OR user_id2 = $1;
    `

	rows, err := ur.db.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("could not fetch friendships: %w", err)
	}
	defer rows.Close()

	// Iterate over rows and build the friendships slice
	for rows.Next() {
		var friendship models.Friendship
		if err := rows.Scan(&friendship.UserID1, &friendship.UserID2); err != nil {
			return nil, fmt.Errorf("could not scan friendship: %w", err)
		}
		friendships = append(friendships, friendship)
	}

	// Check for row iteration errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading friendship rows: %w", err)
	}

	userFriends := make([]models.UserResponse, 0, len(friendships))

	// Retrieve friend user details based on friendships
	var friendId uint
	for _, fsh := range friendships {

		if fsh.UserID1 == id {
			friendId = fsh.UserID2
		} else {
			friendId = fsh.UserID1
		}

		user, err := ur.GetUserById(ctx, friendId)

		if err != nil {
			return nil, err
		}

		userFriends = append(userFriends, *user)

	}

	return userFriends, nil
}

func (ur *UserRepository) ToggleFriendStatus(ctx context.Context, id, friendId uint) ([]models.UserResponse, error) {

	friendship, err := ur.getFriendship(ctx, id, friendId)

	if err != nil {
		return nil, err
	}

	if friendship != nil {
		// They're friends => remove
		// fmt.Println("Deleting friendship...")
		err := ur.deleteFriendship(ctx, id, friendId)

		if err != nil {
			return nil, err
		}

	} else {
		// They are not friends => add
		// fmt.Println("Creating friendship...")
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
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, fmt.Errorf("could not fetch friendships: %w", err)
	}

	return &friendship, nil
}

func (ur *UserRepository) createFriendship(ctx context.Context, userID1, userID2 uint) error {
	// Define the SQL insert query
	query := `
		INSERT INTO friendships (user_id1, user_id2)
		VALUES ($1, $2)
	`

	// Execute the insert query
	_, err := ur.db.Exec(ctx, query, userID1, userID2)
	if err != nil {
		return fmt.Errorf("could not create friendship: %w", err)
	}

	return nil
}

func (ur *UserRepository) deleteFriendship(ctx context.Context, userID1, userID2 uint) error {

	query := `
		DELETE FROM friendships
		WHERE (user_id1 = $1 AND user_id2 = $2) OR (user_id1 = $2 AND user_id2 = $1)
	`
	// Execute the delete query
	result, err := ur.db.Exec(ctx, query, userID1, userID2)
	if err != nil {
		return fmt.Errorf("could not delete friendship: %w", err)
	}

	// Check if any row was affected (deleted)
	if result.RowsAffected() == 0 {
		return errors.New("friendship not found")
	}

	return nil
}
