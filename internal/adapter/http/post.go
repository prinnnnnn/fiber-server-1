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

	// ph

	return nil
}

func (ph *PostHandler) GetPostInfo(ctx *fiber.Ctx) error {

	postId, err := ctx.ParamsInt("id")

	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code).SendString("Parameter is not provided")
		return err
	}

	post, err := ph.ps.GetPostInfo(NewContext(ctx), uint(postId))

	if err != nil {
		ctx.Status(errorStatusMap[err]).SendString(err.Error())
		return nil
	}

	ctx.Status(fiber.StatusOK).JSON(post)

	return nil
}

func (ph *PostHandler) GetUsersPosts(ctx *fiber.Ctx) error {
	userId, err := ctx.ParamsInt("userId")

	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code).SendString("Parameter is not provided")
		return err
	}

	posts, err := ph.ps.GetPostsByUser(NewContext(ctx), uint(userId))

	if err != nil {
		ctx.Status(errorStatusMap[err]).SendString(err.Error())
		return nil
	}

	ctx.Status(fiber.StatusOK).JSON(posts)

	return nil
}

func (ph *PostHandler) AddRemoveFriend(ctx *fiber.Ctx) error {

	userId, err := ctx.ParamsInt("userId")
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code).SendString("Parameter named 'userId' is not provided")
		return err
	}

	postId, err := ctx.ParamsInt("postId")
	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code).SendString("Parameter named 'postId' is not provided")
		return err
	}

	post, err := ph.ps.AddRemoveFriend(NewContext(ctx), uint(userId), uint(postId))

	if err != nil {
		ctx.Status(errorStatusMap[err]).SendString(err.Error())
		return nil
	}

	ctx.Status(fiber.StatusOK).JSON(post)

	return nil
}
