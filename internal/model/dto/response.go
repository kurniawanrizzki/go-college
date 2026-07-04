package dto

import "go-college/internal/model/errors"

type Meta struct {
	Error      *errors.AppError `json:"error,omitempty" swaggertype:"primitive,object" extensions:"x-order=4"`
	Path       string           `json:"path" extensions:"x-order=0"`
	Status     string           `json:"status" extensions:"x-order=2"`
	Message    string           `json:"message" extensions:"x-order=3"`
	Timestamp  string           `json:"timestamp" extensions:"x-order=5"`
	StatusCode int              `json:"status_code" extensions:"x-order=1"`
}

type HttpSuccessResp struct {
	Meta Meta `json:"metadata" extensions:"x-order=0"`
	Data any  `json:"data,omitempty" extensions:"x-order=1"`
}

type HTTPErrorResp struct {
	Meta Meta `json:"metadata"`
}
