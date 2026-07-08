package rest

import (
	"encoding/json"
	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"
	"net/http"
	"strconv"

	"github.com/rs/zerolog"
)

func (e *rest) CreateEnrollment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CreateEnrollmentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("invalid_request_body")
		e.httpRespError(w, r, appErr.WrapWithCode(err, appErr.CodeHTTPUnmarshal, "invalid_request_body"))
		return
	}

	enrollment, err := e.svc.Enrollment.Create(ctx, req)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusCreated, enrollment)
}

func (e *rest) UpdateEnrollment(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateEnrollmentRequest
	ctx := r.Context()

	nim := r.PathValue("nim")
	course := r.PathValue("course")
	if nim == "" || course == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_enrollment_key")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_enrollment_key"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("invalid_request_body")
		e.httpRespError(w, r, appErr.WrapWithCode(err, appErr.CodeHTTPUnmarshal, "invalid_request_body"))
		return
	}

	if err := e.svc.Enrollment.Update(ctx, nim, course, &req); err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, nil)
}

func (e *rest) DeleteEnrollment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("invalid_enrollment_id")
		e.httpRespError(w, r, appErr.WrapWithCode(err, appErr.CodeHTTPBadRequest, "invalid_enrollment_id"))
		return
	}

	if err := e.svc.Enrollment.Delete(ctx, id); err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, nil)
}

func (e *rest) GetEnrollmentsByNim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	nim := r.PathValue("nim")
	if nim == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_enrollment_nim")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_enrollment_nim"))
		return
	}

	details, err := e.svc.Enrollment.FindDetailByNim(ctx, nim)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, details)
}
