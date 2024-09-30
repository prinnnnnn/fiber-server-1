package http

import (
	"fiber-server-1/internal/core/port"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

type RegisterUserRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Location   string `json:"location"`
	Occupation string `json:"occupation"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
}

func (uh *UserHandler) Register(ctx *fiber.Ctx) error {

	// var body *RegisterUserRequest

	// if err := ctx.ShouldBindJSON(&body); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// user := &models.User{
	// 	Email:      body.Email,
	// 	Password:   body.Password,
	// 	FirstName:  body.FirstName,
	// 	LastName:   body.LastName,
	// 	Occupation: body.Occupation,
	// 	Location:   body.Location,
	// }

	// user, err := uh.svc.Register(ctx, user)

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, user)
	return nil

}

func (uh *UserHandler) GetUserInfo(ctx *fiber.Ctx) error {

	userId, err := ctx.ParamsInt("id")

	if err != nil {
		ctx.Status(http.StatusBadRequest).SendString("A parameter named 'id' is not provided")
		return err
	}

	user, err := uh.svc.GetUserInfo(NewContext(ctx), uint(userId))

	if err != nil {
		ctx.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(user)
	return nil
}

func (uh *UserHandler) GetUserFriends(ctx *fiber.Ctx) error {

	userId, err := ctx.ParamsInt("id")

	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
		return err
	}

	friends, err := uh.svc.GetUserFriends(NewContext(ctx), uint(userId))

	if err != nil {
		ctx.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(friends)

	return nil

}

func (uh *UserHandler) AddRemoveFriend(ctx *fiber.Ctx) error {

	userId, err := ctx.ParamsInt("id")

	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code).SendString("userId is not provided")
		return err
	}

	friendId, err := ctx.ParamsInt("friendId")

	if err != nil {
		ctx.Status(fiber.ErrBadRequest.Code).SendString("friendId is not provided")
	}

	friends, err := uh.svc.AddRemoveFriend(NewContext(ctx), uint(userId), uint(friendId))

	if err != nil {
		ctx.Status(fiber.ErrInternalServerError.Code).SendString("Fail to fetch data")
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(friends)
	return nil

}
