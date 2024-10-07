package models

import "time"

type Post struct {
	GormModel
	UserID uint `gorm:"not null" json:"userId"`
	// User        User   `gorm:"foreignKey:UserID"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PicturePath string `json:"picturePath"`
	Description string `json:"description"`
}

type Like struct {
	LikedAt time.Time `json:"likedAt"`
	UserID  uint      `gorm:"primaryKey"`
	PostID  uint      `gorm:"primaryKey"`
}
