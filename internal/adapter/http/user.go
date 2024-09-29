package http

import (
	"context"
	"fiber-server-1/internal/core/port"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Ctx struct {
	fiber.Ctx
	ctx context.Context
}

// Done implements context.Context.
func (c *Ctx) Done() <-chan struct{} {
	if c.ctx != nil {
		return c.ctx.Done()
	}
	return nil
}

// Err implements context.Context.
func (c *Ctx) Err() error {
	if c.ctx != nil {
		return c.ctx.Err()
	}
	return nil
}

// Value implements context.Context.
func (c *Ctx) Value(key any) any {
	if c.ctx != nil {
		return c.ctx.Value(key)
	}
	return c.Locals(key.(string))
}

// Deadline implements context.Context.
func (c *Ctx) Deadline() (deadline time.Time, ok bool) {
	if c.ctx != nil {
		return c.ctx.Deadline()
	}
	return time.Time{}, false
}

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

// type Context =

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
		return nil
	}

	customCtx := Ctx{
		Ctx: *ctx,
		ctx: context.Background(),
	}
	user, err := uh.svc.GetUserInfo(&customCtx, uint(userId))

	if err != nil {
		ctx.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
		return err
	}

	ctx.Status(fiber.StatusOK).JSON(user)
	return nil
}

func (uh *UserHandler) GetUserFriends(ctx *fiber.Ctx) error {

	// parsedId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// id := uint(parsedId)

	// friends, err := uh.svc.GetUserFriends(ctx, id)

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, friends)
	return nil

}

func (uh *UserHandler) AddRemoveFriend(ctx *fiber.Ctx) error {

	// parsedId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// id := uint(parsedId)

	// parsedfriendId, err := strconv.ParseUint(ctx.Param("friendId"), 10, 64)

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// friendId := uint(parsedfriendId)

	// friends, err := uh.svc.AddRemoveFriend(ctx, id, friendId)

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, friends)
	return nil

}
