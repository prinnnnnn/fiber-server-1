package port

import (
	"context"
	"fiber-server-1/internal/core/models"
)

type UserRepository interface {

	/* Create a new user */
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	/* Get a user's info*/
	GetUserById(ctx context.Context, id uint) (*models.User, error)
	/* Get a user's friends */
	GetUserFriends(ctx context.Context, id uint) ([]models.User, error)
	/* Update friendship status between users */
	ToggleFriendStatus(ctx context.Context, id, freindId uint) ([]models.User, error)
}

type UserService interface {

	/* Register a new user */
	Register(ctx context.Context, user *models.User) (*models.User, error)
	/* Get user info with the specified ID */
	GetUserInfo(ctx context.Context, id uint) (*models.User, error)
	/* Get all friends of a particular user */
	GetUserFriends(ctx context.Context, id uint) ([]models.User, error)
	/* Add/Remove friend */
	AddRemoveFriend(ctx context.Context, id, freindId uint) ([]models.User, error)
}
