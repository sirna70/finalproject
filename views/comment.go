package views

import "time"

type CommentCreate struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	PhotoID   uint       `json:"photo_id"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type CommentUpdate struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type GetComments struct {
	ID        uint         `json:"id"`
	Message   string       `json:"message"`
	PhotoID   uint         `json:"photo_id"`
	UserID    uint         `json:"user_id"`
	UpdatedAt *time.Time   `json:"updated_at"`
	CreatedAt *time.Time   `json:"created_at"`
	User      UserComment  `json:"user"`
	Photo     PhotoComment `json:"photo"`
}

type UserComment struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoComment struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
}
