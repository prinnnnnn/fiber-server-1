package port

import (
	"context"
	"fiber-server-1/internal/core/models"
)

type UserRepository interface {

	/* Create a new user */
	CreateUser(ctx context.Context, user *models.User) (*models.UserResponse, error)
	/* Get a user's info*/
	GetUserById(ctx context.Context, id uint) (*models.UserResponse, error)
	/* Get a user's friends */
	GetUserFriends(ctx context.Context, id uint) ([]models.UserResponse, error)
	/* Update friendship status between users */
	ToggleFriendStatus(ctx context.Context, id, freindId uint) ([]models.UserResponse, error)
}

type UserService interface {

	/* Register a new user */
	Register(ctx context.Context, user *models.User) (*models.UserResponse, error)
	/* Get user info with the specified ID */
	GetUserInfo(ctx context.Context, id uint) (*models.UserResponse, error)
	/* Get all friends of a particular user */
	GetUserFriends(ctx context.Context, id uint) ([]models.UserResponse, error)
	/* Add/Remove friend */
	AddRemoveFriend(ctx context.Context, id, freindId uint) ([]models.UserResponse, error)
}
