package handlers

import "time"

type ErrorResponse struct {
	Error string `json:"error" example:"usuario nao encontrado"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type MessageResponse struct {
	Message string `json:"message" example:"post archived"`
}

type UserSummary struct {
	ID        uint      `json:"id" example:"1"`
	Username  string    `json:"username" example:"joaogs"`
	Name      string    `json:"name" example:"Joao Santiago"`
	Email     string    `json:"email" example:"joao@joao.com"`
	Biography string    `json:"biography" example:"Desenvolvedor Go e APIs REST"`
	CreatedAt time.Time `json:"created_at" swaggertype:"string" example:"2026-04-04T21:10:41-03:00"`
	UpdatedAt time.Time `json:"updated_at" swaggertype:"string" example:"2026-04-04T21:10:41-03:00"`
}

type PostSummary struct {
	ID        uint      `json:"id" example:"10"`
	UserID    uint      `json:"user_id" example:"1"`
	Text      string    `json:"text" example:"Meu primeiro post no Starti!"`
	Archived  bool      `json:"archived" example:"false"`
	CreatedAt time.Time `json:"created_at" swaggertype:"string" example:"2026-04-04T21:10:41-03:00"`
	UpdatedAt time.Time `json:"updated_at" swaggertype:"string" example:"2026-04-04T21:10:41-03:00"`
}

type CommentSummary struct {
	ID        uint      `json:"id" example:"3"`
	UserID    uint      `json:"user_id" example:"1"`
	PostID    uint      `json:"post_id" example:"10"`
	Message   string    `json:"message" example:"Otimo post, curti bastante!"`
	CreatedAt time.Time `json:"created_at" swaggertype:"string" example:"2026-04-04T21:22:16-03:00"`
	UpdatedAt time.Time `json:"updated_at" swaggertype:"string" example:"2026-04-04T21:22:16-03:00"`
}

type PostResponse struct {
	ID        uint        `json:"id" example:"10"`
	UserID    uint        `json:"user_id" example:"1"`
	Text      string      `json:"text" example:"Meu primeiro post no Starti!"`
	Archived  bool        `json:"archived" example:"false"`
	CreatedAt time.Time   `json:"created_at" swaggertype:"string" example:"2026-04-04T21:10:41-03:00"`
	UpdatedAt time.Time   `json:"updated_at" swaggertype:"string" example:"2026-04-04T21:10:41-03:00"`
	User      UserSummary `json:"user"`
}

type CommentResponse struct {
	ID        uint        `json:"id" example:"3"`
	UserID    uint        `json:"user_id" example:"1"`
	PostID    uint        `json:"post_id" example:"10"`
	Message   string      `json:"message" example:"Otimo post, curti bastante!"`
	CreatedAt time.Time   `json:"created_at" swaggertype:"string" example:"2026-04-04T21:22:16-03:00"`
	UpdatedAt time.Time   `json:"updated_at" swaggertype:"string" example:"2026-04-04T21:22:16-03:00"`
	User      UserSummary `json:"user"`
	Post      PostSummary `json:"post"`
}
