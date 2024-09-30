package port

import (
	"context"
	"fiber-server-1/internal/core/models"
)

type PostRepository interface {

	/* Create a new Post */
	CreatePost(ctx context.Context, Post *models.Post) (*models.Post, error)
	/* Get a Post's info*/
	GetPostById(ctx context.Context, id uint) (*models.Post, error)
	/* Get Posts by userId */
	GetPostsByUserId(ctx context.Context, userId uint) ([]models.Post, error)
	/* Toggle like status of post by userId */
	// ToggleFriendStatus(ctx context.Context, userId, postId uint) ([]models.Post, error)
}

type PostService interface {

	/* Create a new Post */
	CreatePost(ctx context.Context, Post *models.Post) (*models.Post, error)
	/* Get Post info with the specified ID */
	GetPostInfo(ctx context.Context, id uint) (*models.Post, error)
	/* Get all Posts of a user (by userId) */
	GetPostsByUser(ctx context.Context, userId uint) ([]models.Post, error)
	/* Add/Remove friend */
	// AddRemoveFriend(ctx context.Context, id, freindId uint) ([]models.Post, error)
}
