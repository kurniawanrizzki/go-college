package rest

import (
	"encoding/json"
	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"
	"net/http"

	"github.com/rs/zerolog"
)

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

func (e *rest) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	courses, err := e.svc.Course.FindAll(ctx)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, courses)
}
