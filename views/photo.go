package views

import "time"

type PhotoView struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      UserView
}

type UserView struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoCreateView struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type PhotoUpdateView struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    uint       `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}
