package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Text      string    `json:"text" gorm:"type:text;not null"`
	Archived  bool      `json:"archived" gorm:"default:false;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}
