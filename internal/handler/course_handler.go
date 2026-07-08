package rest

import (
	"encoding/json"
	"net/http"

	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"

	"github.com/rs/zerolog"
)

// CreateCourse godoc
//
//	@Summary		Create a new course
//	@Description	Create a new course record
//	@Tags			courses
//	@Accept			json
//	@Produce		json
//	@Param			course	body		dto.CreateCourseRequest	true	"Course data"
//	@Success		201		{object}	dto.HttpSuccessResp{data=entity.Course}
//	@Failure		400		{object}	dto.HTTPErrorResp
//	@Failure		500		{object}	dto.HTTPErrorResp
//	@Router			/course/create [post]
func (e *rest) CreateCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CreateCourseRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("invalid_request_body")
		e.httpRespError(w, r, appErr.WrapWithCode(err, appErr.CodeHTTPUnmarshal, "invalid_request_body"))
		return
	}

	course, err := e.svc.Course.Create(ctx, req)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusCreated, course)
}

// UpdateCourse godoc
//
//	@Summary		Update a course
//	@Description	Update an existing course identified by code
//	@Tags			courses
//	@Accept			json
//	@Produce		json
//	@Param			code	path		string					true	"Course code"
//	@Param			course	body		dto.UpdateCourseRequest	true	"Course data"
//	@Success		200		{object}	dto.HttpSuccessResp{data=entity.Course}
//	@Failure		400		{object}	dto.HTTPErrorResp
//	@Failure		404		{object}	dto.HTTPErrorResp
//	@Router			/course/{code} [put]
func (e *rest) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCourseRequest
	ctx := r.Context()

	code := r.PathValue("code")
	if code == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_course_code")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_id"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("invalid_request_body")
		e.httpRespError(w, r, appErr.WrapWithCode(err, appErr.CodeHTTPUnmarshal, "invalid_request_body"))
		return
	}

	course, err := e.svc.Course.Update(ctx, code, &req)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, course)
}

// DeleteCourse godoc
//
//	@Summary		Delete a course
//	@Description	Delete a course identified by code
//	@Tags			courses
//	@Produce		json
//	@Param			code	path		string	true	"Course code"
//	@Success		200		{object}	dto.HttpSuccessResp
//	@Failure		400		{object}	dto.HTTPErrorResp
//	@Failure		404		{object}	dto.HTTPErrorResp
//	@Router			/course/{code} [delete]
func (e *rest) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := r.PathValue("code")
	if code == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_course_code")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_nim"))
		return
	}

	err := e.svc.Course.Delete(ctx, code)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, nil)
}

// GetCourseByCode godoc
//
//	@Summary		Get a course by code
//	@Description	Retrieve a single course by its code
//	@Tags			courses
//	@Produce		json
//	@Param			code	path		string	true	"Course code"
//	@Success		200		{object}	dto.HttpSuccessResp{data=entity.Course}
//	@Failure		400		{object}	dto.HTTPErrorResp
//	@Failure		404		{object}	dto.HTTPErrorResp
//	@Router			/course/{code} [get]
func (e *rest) GetCourseByCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := r.PathValue("code")
	if code == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_course_code")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_id"))
		return
	}

	course, err := e.svc.Course.FindByCode(ctx, code)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, course)
}

// GetAllCourses godoc
//
//	@Summary		List all courses
//	@Description	Retrieve all course records
//	@Tags			courses
//	@Produce		json
//	@Success		200	{object}	dto.HttpSuccessResp{data=[]entity.Course}
//	@Failure		500	{object}	dto.HTTPErrorResp
//	@Router			/course/all [get]
func (e *rest) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	courses, err := e.svc.Course.FindAll(ctx)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, courses)
}
