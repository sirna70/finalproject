package views

import "time"

type UserUpdate struct {
	ID        uint       `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	Age       int        `json:"age"`
	UpdatedAt *time.Time `json:"updated_at"`
}
