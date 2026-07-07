package entity

import "time"

type College struct {
	NIM       string    `json:"nim"`
	Name      string    `json:"name"`
	Semester  int       `json:"semester"`
	SKS       int       `json:"sks"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CollegeEnrollments struct {
	College
	Enrollments []EnrollmentDetail `json:"enrollments"`
}
