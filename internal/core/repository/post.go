package repository

import (
	"context"
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
	// var wg sync.WaitGroup
	// var mu sync.Mutex

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
