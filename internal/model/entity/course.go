package entity

import "time"

type Course struct {
	Code      string
	Name      string
	SKS       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
