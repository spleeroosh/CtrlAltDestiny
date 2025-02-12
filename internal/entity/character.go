package entity

import (
	"time"
)

//go:generate go run github.com/mailru/easyjson/easyjson

//easyjson:json
type Character struct {
	ID              int       `json:"id" db:"id"`
	UserID          int       `json:"user_id" db:"user_id"`
	Name            string    `json:"name" db:"name"`
	Age             uint64    `json:"age" db:"age"`
	Profession      string    `json:"profession" db:"profession"`
	BurnoutLevel    int       `json:"burnoutLevel" db:"burnout_level"`
	MotivationLevel int       `json:"motivationLevel" db:"motivation_level"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}
