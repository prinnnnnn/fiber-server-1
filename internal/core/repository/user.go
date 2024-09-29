package repository

/*

UserRepository implements port.UserRepository interface (port's spec)
and provide an access to the postgres database

*/

import (
	"context"
	"errors"
	"fiber-server-1/internal/core/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	result := ur.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil

}

func (ur *UserRepository) GetUserById(ctx context.Context, id uint) (*models.User, error) {

	var user models.User
	result := ur.db.WithContext(ctx).First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func (ur *UserRepository) GetUserFriends(ctx context.Context, id uint) ([]models.User, error) {
	var friendships []models.Friendship

	result := ur.db.Preload("User1").Preload("User2").
		Where("user_id1 = ? OR user_id2 = ?", id, id).
		Find(&friendships)

	if result.Error != nil {
		return nil, result.Error
	}

	userFriends := make([]models.User, 0)

	for _, fsh := range friendships {
		freind := fsh.User1

		if freind.ID == id {
			freind = fsh.User2
		}

		userFriends = append(userFriends, freind)
	}

	return userFriends, nil
}

func (ur *UserRepository) ToggleFriendStatus(ctx context.Context, id, friendId uint) ([]models.User, error) {

	friendship, err := ur.getFriendship(id, friendId)

	if err != nil {
		return nil, err
	}

	if friendship != nil {
		// They're friends => remove
		err := ur.deleteFriendship(id, friendId)

		if err != nil {
			return nil, err
		}

	} else {
		// They are not friends => add
		err := ur.createFriendship(id, friendId)

		if err != nil {
			return nil, err
		}

	}

	return ur.GetUserFriends(ctx, id)

}

// GetFriendship checks if a friendship exists
func (ur *UserRepository) getFriendship(userID1, userID2 uint) (*models.Friendship, error) {
	var friendship models.Friendship
	result := ur.db.Where("(user_id1 = ? AND user_id2 = ?) OR (user_id1 = ? AND user_id2 = ?)",
		userID1, userID2, userID2, userID1).First(&friendship)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // Friendship doesn't exist
		}
		return nil, result.Error
	}
	return &friendship, nil
}

func (ur *UserRepository) createFriendship(userID1, userID2 uint) error {
	friendship := models.Friendship{
		UserID1: userID1,
		UserID2: userID2,
	}
	return ur.db.Create(&friendship).Error
}

func (ur *UserRepository) deleteFriendship(userID1, userID2 uint) error {

	result := ur.db.Where("(user_id1 = ? AND user_id2 = ?) OR (user_id1 = ? AND user_id2 = ?)",
		userID1, userID2, userID2, userID1).Delete(&models.Friendship{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("friendship not found")
	}
	return nil
}
