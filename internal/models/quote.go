package models

import "time"

type Quote struct {
	ID         int       `json:"id"`
	Author_ID  int       `json:"author_id"`
	Text       string    `json:"quote"`
	Created_at time.Time `json:"created_at"`
}
