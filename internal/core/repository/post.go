package repository

import (
	"context"
	"errors"
	"fiber-server-1/internal/core/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{
		db,
	}
}

/* Create a new post */
func (pr *PostRepository) CreatePost(ctx context.Context, Post *models.Post) (*models.Post, error) {
	return nil, nil
}

/* Get a post's info by its id */
func (pr *PostRepository) GetPostById(ctx context.Context, id uint) (*models.Post, error) {

	var post models.Post
	query := `SELECT * FROM posts WHERE id = $1`

	err := pr.db.QueryRow(ctx, query, id).Scan(
		&post.ID, &post.CreatedAt,
		&post.CreatedAt, &post.UserID,
		&post.FirstName, &post.LastName,
		&post.PicturePath, &post.Description,
	)

	if err != nil {

		if err == pgx.ErrNoRows {
			return nil, models.ErrDataNotFound
		}

		return nil, models.ErrInternalServer
	}

	return &post, nil

}

/* Get all posts of a user */
func (pr *PostRepository) GetPostsByUserId(ctx context.Context, userId uint) ([]models.Post, error) {

	query := `
		SELECT * FROM posts WHERE user_id = $1
	`

	rows, err := pr.db.Query(ctx, query, userId)
	if err != nil {
		return nil, models.ErrInternalServer
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		if err := rows.Scan(
			&post.ID, &post.CreatedAt,
			&post.CreatedAt, &post.UserID,
			&post.FirstName, &post.LastName,
			&post.PicturePath, &post.Description,
		); err != nil {
			return nil, models.ErrInternalServer
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, models.ErrInternalServer
	}

	return posts, nil

}

func (pr *PostRepository) LikePost(ctx context.Context, userId, postId uint) (*models.Post, error) {

	err := pr.getLikeRecord(ctx, userId, postId)

	if errors.Is(err, models.ErrInternalServer) {
		return nil, err
	}

	var post *models.Post

	if errors.Is(err, models.ErrDataNotFound) {
		// user did not like post
		post, err = pr.likePost(ctx, userId, postId)
		if err != nil {
			return nil, err
		}

	} else {
		// user liked post
		post, err = pr.dislikePost(ctx, userId, postId)

		if err != nil {
			return nil, err
		}

	}

	return post, nil

}

// GetFriendship checks if a friendship exists
func (pr *PostRepository) getLikeRecord(ctx context.Context, userId, postId uint) error {
	var like models.Like

	query := `
		SELECT user_id, post_id 
        FROM likes 
        WHERE (user_id = $1 AND post_id = $2);
	`

	row := pr.db.QueryRow(ctx, query, userId, postId)
	err := row.Scan(&like.UserID, &like.PostID)

	if err != nil {
		// Handle the "not found" case
		if err == pgx.ErrNoRows {
			return models.ErrDataNotFound
		}
		return models.ErrInternalServer
	}

	return nil
}

func (pr *PostRepository) likePost(ctx context.Context, userId, postId uint) (*models.Post, error) {

	query := `
		INSERT INTO likes (user_id, post_id)
		VALUES ($1, $2)
	`

	_, err := pr.db.Exec(ctx, query, userId, postId)
	if err != nil {
		return nil, models.ErrInternalServer
	}

	return pr.GetPostById(ctx, postId)
}

func (pr *PostRepository) dislikePost(ctx context.Context, userId, postId uint) (*models.Post, error) {

	query := `
		DELETE FROM likes
		WHERE (user_id = $1 AND post_id = $2)
	`

	result, err := pr.db.Exec(ctx, query, userId, postId)
	if errors.Is(err, models.ErrInternalServer) {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, models.ErrDataNotFound
	}

	return pr.GetPostById(ctx, postId)
}
