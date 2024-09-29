package service

import (
	"context"
	"fiber-server-1/internal/core/models"
	"fiber-server-1/internal/core/port"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserService struct {
	repo port.UserRepository
}

/* Create UserService Instance */
func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Register(ctx context.Context, user *models.User) (*models.User, error) {

	saltRound := 6
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), saltRound)

	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	user.ViewedProfile = rand.Intn(10000)
	user.Impressions = rand.Intn(10000)

	newUser, err := us.repo.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	newUser.Password = ""

	return newUser, nil

}

func (us *UserService) GetUserInfo(ctx context.Context, id uint) (*models.User, error) {

	// user, err :=

	// if err != nil {
	// 	return nil, err
	// }

	return us.repo.GetUserById(ctx, id)
}

func (us *UserService) GetUserFriends(ctx context.Context, id uint) ([]models.User, error) {

	friendsList, err := us.repo.GetUserFriends(ctx, id)

	if err != nil {
		return nil, err
	}

	return friendsList, nil

}

func (us *UserService) AddRemoveFriend(ctx context.Context, id uint, freindId uint) ([]models.User, error) {
	return us.repo.ToggleFriendStatus(ctx, id, freindId)
}
