package http

import (
	"fiber-server-1/internal/core/port"
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

// type Context =

func (uh *UserHandler) Register() {

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

}

func (uh *UserHandler) GetUserInfo() {

	// parsedId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	// id := uint(parsedId)

	// errChan := make(chan error, 1)
	// userChan := make(chan *models.User, 1)

	// var user *models.User
	// go uh.svc.GetUserInfo(ctx, id)

	// select {
	// case err := <-errChan:
	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// case user = <-userChan:
	// 	ctx.JSON(http.StatusOK, user)
	// }

}

func (uh *UserHandler) GetUserFriends() {

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

}

func (uh *UserHandler) AddRemoveFriend() {

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

}
