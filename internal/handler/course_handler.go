package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// GetAllCourses godoc
//
//	@Summary		List courses
//	@Description	Retrieve course records with optional filtering, sorting and pagination
//	@Tags			courses
//	@Produce		json
//	@Param			code		query		string	false	"Filter by code (partial match)"
//	@Param			name		query		string	false	"Filter by name (partial match)"
//	@Param			sks			query		int		false	"Filter by SKS"
//	@Param			sort_by		query		string	false	"Sort field (code, name, sks)"
//	@Param			sort_dir	query		string	false	"Sort direction (asc, desc)"
//	@Param			page		query		int		false	"Page number"
//	@Param			per_page	query		int		false	"Items per page"
//	@Success		200	{object}	dto.HttpSuccessResp{data=dto.PaginatedResp}
//	@Failure		500	{object}	dto.HTTPErrorResp
//	@Router			/course/all [get]
func (e *rest) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := parseCourseFilter(r)

	courses, pagination, err := e.svc.Course.FindAll(ctx, filter)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, dto.PaginatedResp{
		Items:      courses,
		Page:       pagination.Page,
		PerPage:    pagination.PerPage,
		PageCount:  pagination.PageCount,
		TotalCount: pagination.TotalCount,
	})
}

func parseCourseFilter(r *http.Request) *dto.CourseFilter {
	q := r.URL.Query()

	filter := &dto.CourseFilter{
		Code:    q.Get("code"),
		Name:    q.Get("name"),
		SortBy:  q.Get("sort_by"),
		SortDir: q.Get("sort_dir"),
	}

	if sks, err := strconv.ParseInt(q.Get("sks"), 10, 64); err == nil {
		filter.SKS = sks
	}
	if page, err := strconv.ParseInt(q.Get("page"), 10, 64); err == nil {
		filter.Page = page
	}
	if perPage, err := strconv.ParseInt(q.Get("per_page"), 10, 64); err == nil {
		filter.PerPage = perPage
	}

	return filter
}
