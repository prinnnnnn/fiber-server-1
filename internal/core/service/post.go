package service

import (
	"context"
	"fiber-server-1/internal/core/models"
	"fiber-server-1/internal/core/port"
)

type PostService struct {
	repo port.PostRepository
}

/* Create post service */
func NewPostService(postRepo port.PostRepository) *PostService {
	return &PostService{
		repo: postRepo,
	}
}

/* Create a new Post */
func (ps *PostService) CreatePost(ctx context.Context, Post *models.Post) (*models.Post, error) {
	return nil, nil
}

/* Get a post's info by its id */
func (ps *PostService) GetPostInfo(ctx context.Context, id uint) (*models.Post, error) {
	return ps.repo.GetPostById(ctx, id)
}

/* Get all posts of a user */
func (ps *PostService) GetPostsByUser(ctx context.Context, userId uint) ([]models.Post, error) {
	return ps.repo.GetPostsByUserId(ctx, userId)
}
