package entity

import "time"

type Enrollment struct {
	ID        int
	NIM       string
	Course    string
	Semester  int
	Grade     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type EnrollmentDetail struct {
	Course
	Semester  int
	Grade     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
