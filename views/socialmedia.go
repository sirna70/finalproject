package views

import "time"

type SocialMediaCreate struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         uint       `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type SocialMedias struct {
	ID             uint            `json:"id"`
	Name           string          `json:"name"`
	SocialMediaUrl string          `json:"social_media_url"`
	UserID         uint            `json:"user_id"`
	CreatedAt      *time.Time      `json:"created_at"`
	UpdatedAt      *time.Time      `json:"updated_at"`
	User           SocialMediaUser `json:"user"`
}

type SocialMediaUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type SocialMediaUpdate struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserID         uint       `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
