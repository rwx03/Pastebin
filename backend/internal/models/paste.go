package models

import "time"

type Paste struct {
	ID        int       `json:"id"`
	PasteID   string    `json:"paste_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatorID int       `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
	Views     int       `json:"views"`
}
