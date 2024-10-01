package models

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

/* Full User atrributes */
type User struct {
	GormModel
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email" gorm:"unique"`
	Password      string `json:"password"`
	PicturePath   string `json:"picturePath"`
	Location      string `json:"location"`
	Occupation    string `json:"occupation"`
	ViewedProfile int    `json:"viewedProfile"`
	Impressions   int    `json:"impressions"`
}

/* User for Response (exclude password */
type UserResponse struct {
	GormModel
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email" gorm:"unique"`
	PicturePath   string `json:"picturePath"`
	Location      string `json:"location"`
	Occupation    string `json:"occupation"`
	ViewedProfile int    `json:"viewedProfile"`
	Impressions   int    `json:"impressions"`
}

type Friendship struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID1   uint      `gorm:"primaryKey"`
	UserID2   uint      `gorm:"primaryKey"`
	User1     User      `gorm:"foreignKey:UserID1"`
	User2     User      `gorm:"foreignKey:UserID2"`
}

func (Friendship) TableName() string {
	return "friendships"
}

func (f *Friendship) BeforeCreate(tx *gorm.DB) error {
	if f.UserID1 > f.UserID2 {
		f.UserID1, f.UserID2 = f.UserID2, f.UserID1
	}
	return nil
}

func MapToResponse(user *User) *UserResponse {
	return &UserResponse{
		GormModel:     user.GormModel,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		PicturePath:   user.PicturePath,
		Location:      user.Location,
		Occupation:    user.Occupation,
		ViewedProfile: user.ViewedProfile,
		Impressions:   user.Impressions,
	}
}
