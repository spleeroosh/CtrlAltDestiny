package entity

import (
	"time"
)

//go:generate go run github.com/mailru/easyjson/easyjson

//easyjson:json
type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Age       uint64    `json:"age" db:"age"`
	Social    string    `json:"social" db:"social"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
