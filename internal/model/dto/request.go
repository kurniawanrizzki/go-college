package dto

type CreateCollegeRequest struct {
	NIM      string
	Name     string
	Semester int
	SKS      int
	Active   bool
}

type UpdateCollegeRequest struct {
	Name     string
	Semester int
	SKS      int
	Active   *bool
}

type CreateCourseRequest struct {
	Code string
	Name string
	SKS int
}

type UpdateCourseRequest struct {
	Name string
	SKS int
}
