package models

import "time"

type Author struct {
	ID         int       `json:"id"`
	Name       string    `json:"author"`
	Created_at time.Time `json:"created_at"`
}
