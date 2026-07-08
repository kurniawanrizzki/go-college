package entity

import "time"

type Course struct {
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	SKS       int       `json:"sks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
