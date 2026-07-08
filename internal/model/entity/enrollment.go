package entity

import "time"

type Enrollment struct {
	ID        int       `json:"id"`
	NIM       string    `json:"nim"`
	Course    string    `json:"course_code"`
	Semester  int       `json:"semester"`
	Grade     string    `json:"grade"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EnrollmentDetail struct {
	Course
	Semester  int       `json:"semester"`
	Grade     string    `json:"grade"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
