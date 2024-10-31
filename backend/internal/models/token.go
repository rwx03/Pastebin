package models

type Token struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}
