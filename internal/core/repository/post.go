package repository

import (
	"context"
	"fiber-server-1/internal/core/models"
	"sync"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
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
	result := pr.db.WithContext(ctx).First(&post, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil

}

/* Get all posts of a user */
func (pr *PostRepository) GetPostsByUserId(ctx context.Context, userId uint) ([]models.Post, error) {

	var userIds []struct {
		postId uint
		userId uint
	}

	result := pr.db.Model(&models.Post{}).Select("id, user_id").Find(&userIds)

	if result.Error != nil {
		return nil, result.Error
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	posts := make([]models.Post, 0)

	for _, row := range userIds {
		if row.userId == userId {
			wg.Add(1)
			go func(id uint) {
				defer wg.Done()

				var post models.Post
				if err := pr.db.First(&post, id).Error; err != nil {
					return
				}

				mu.Lock()
				posts = append(posts, post)
				mu.Unlock()

			}(row.postId)
		}
	}

	wg.Wait()

	return posts, nil

}
