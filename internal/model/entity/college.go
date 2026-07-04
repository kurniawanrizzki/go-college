package entity

import "time"

type College struct {
	NIM       string
	Name      string
	Semester  int
	SKS       int
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CollegeEnrollments struct {
	College
	Enrollments []EnrollmentDetail
}
