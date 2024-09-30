package http

import (
	"fiber-server-1/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	ps port.PostService
}

func NewPostHandler(service port.PostService) *PostHandler {
	return &PostHandler{
		ps: service,
	}
}

func (ph *PostHandler) CreateNewPost(ctx *fiber.Ctx) error {
	return nil
}

func (ph *PostHandler) GetPostInfo(ctx *fiber.Ctx) error {
	return nil
}

func (ph *PostHandler) GetUsersPosts(ctx *fiber.Ctx) error {
	return nil
}
