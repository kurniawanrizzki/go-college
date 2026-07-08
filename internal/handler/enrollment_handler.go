package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"

	"github.com/rs/zerolog"
)

// CreateEnrollment godoc
//
//	@Summary		Create a new enrollment
//	@Description	Enroll a college (NIM) into a course for a semester
//	@Tags			enrollments
//	@Accept			json
//	@Produce		json
//	@Param			enrollment	body		dto.CreateEnrollmentRequest	true	"Enrollment data"
//	@Success		201			{object}	dto.HttpSuccessResp{data=entity.Enrollment}
//	@Failure		400			{object}	dto.HTTPErrorResp
//	@Failure		500			{object}	dto.HTTPErrorResp
//	@Router			/enrollment/create [post]
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

// UpdateEnrollment godoc
//
//	@Summary		Update an enrollment
//	@Description	Update semester and grade for the enrollment identified by NIM and course code
//	@Tags			enrollments
//	@Accept			json
//	@Produce		json
//	@Param			nim			path		string						true	"College NIM"
//	@Param			course		path		string						true	"Course code"
//	@Param			enrollment	body		dto.UpdateEnrollmentRequest	true	"Enrollment data"
//	@Success		200			{object}	dto.HttpSuccessResp
//	@Failure		400			{object}	dto.HTTPErrorResp
//	@Failure		404			{object}	dto.HTTPErrorResp
//	@Router			/enrollment/{nim}/{course} [put]
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

// DeleteEnrollment godoc
//
//	@Summary		Delete an enrollment
//	@Description	Delete an enrollment identified by its id
//	@Tags			enrollments
//	@Produce		json
//	@Param			id	path		int	true	"Enrollment id"
//	@Success		200	{object}	dto.HttpSuccessResp
//	@Failure		400	{object}	dto.HTTPErrorResp
//	@Failure		404	{object}	dto.HTTPErrorResp
//	@Router			/enrollment/{id} [delete]
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

// GetEnrollmentsByNim godoc
//
//	@Summary		List enrollment details by NIM
//	@Description	Retrieve all enrollments (with course detail) for a given college NIM
//	@Tags			enrollments
//	@Produce		json
//	@Param			nim	path		string	true	"College NIM"
//	@Success		200	{object}	dto.HttpSuccessResp{data=[]entity.EnrollmentDetail}
//	@Failure		400	{object}	dto.HTTPErrorResp
//	@Router			/enrollment/nim/{nim} [get]
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
