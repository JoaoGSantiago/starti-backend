package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Name      string    `json:"name" gorm:"not null;size:100"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null;size:150"`
	Password  string    `json:"-" gorm:"not null"` // json:"-" impede que o hash da senha apareça em qualquer resposta
	Biography string    `json:"biography" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// OnDelete:CASCADE apaga automaticamente os posts e comentários ao deletar o usuário
	Posts    []Post    `json:"posts,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
