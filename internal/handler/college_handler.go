package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"

	"github.com/rs/zerolog"
)

// CreateCollege godoc
//
//	@Summary		Create a new college
//	@Description	Create a new college record
//	@Tags			colleges
//	@Accept			json
//	@Produce		json
//	@Param			college	body		dto.CreateCollegeRequest	true	"College data"
//	@Success		201		{object}	dto.HttpSuccessResp{data=entity.College}
//	@Failure		400		{object}	dto.HTTPErrorResp
//	@Failure		500		{object}	dto.HTTPErrorResp
//	@Router			/college/create [post]
func (e *rest) CreateCollege(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.CreateCollegeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("invalid_request_body")
		e.httpRespError(w, r, appErr.WrapWithCode(err, appErr.CodeHTTPUnmarshal, "invalid_request_body"))
		return
	}

	college, err := e.svc.College.Create(ctx, req)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusCreated, college)
}

// FindAll godoc
//
//	@Summary		List colleges
//	@Description	Retrieve college records with optional filtering, sorting and pagination
//	@Tags			colleges
//	@Produce		json
//	@Param			name		query		string	false	"Filter by name (partial match)"
//	@Param			semester	query		int		false	"Filter by semester"
//	@Param			sort_by		query		string	false	"Sort field (nim, name, semester, sks)"
//	@Param			sort_dir	query		string	false	"Sort direction (asc, desc)"
//	@Param			page		query		int		false	"Page number"
//	@Param			per_page	query		int		false	"Items per page"
//	@Success		200	{object}	dto.HttpSuccessResp{data=dto.PaginatedResp}
//	@Failure		500	{object}	dto.HTTPErrorResp
//	@Router			/college/all [get]
func (e *rest) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := parseCollegeFilter(r)

	colleges, pagination, err := e.svc.College.FindAll(ctx, filter)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, dto.PaginatedResp{
		Items:      colleges,
		Page:       pagination.Page,
		PerPage:    pagination.PerPage,
		PageCount:  pagination.PageCount,
		TotalCount: pagination.TotalCount,
	})
}

func parseCollegeFilter(r *http.Request) *dto.CollegeFilter {
	q := r.URL.Query()

	filter := &dto.CollegeFilter{
		NIM:     q.Get("nim"),
		Name:    q.Get("name"),
		SortBy:  q.Get("sort_by"),
		SortDir: q.Get("sort_dir"),
	}

	if semester, err := strconv.ParseInt(q.Get("semester"), 10, 64); err == nil {
		filter.Semester = semester
	}
	if page, err := strconv.ParseInt(q.Get("page"), 10, 64); err == nil {
		filter.Page = page
	}
	if perPage, err := strconv.ParseInt(q.Get("per_page"), 10, 64); err == nil {
		filter.PerPage = perPage
	}

	return filter
}

// UpdateCollege godoc
//
//	@Summary		Update a college
//	@Description	Update an existing college identified by NIM
//	@Tags			colleges
//	@Accept			json
//	@Produce		json
//	@Param			nim		path		string						true	"College NIM"
//	@Param			college	body		dto.UpdateCollegeRequest	true	"College data"
//	@Success		200		{object}	dto.HttpSuccessResp{data=entity.College}
//	@Failure		400		{object}	dto.HTTPErrorResp
//	@Failure		404		{object}	dto.HTTPErrorResp
//	@Router			/college/{nim} [put]
func (e *rest) UpdateCollege(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCollegeRequest
	ctx := r.Context()

	nim := r.PathValue("nim")
	if nim == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_college_id")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_id"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("invalid_request_body")
		e.httpRespError(w, r, appErr.WrapWithCode(err, appErr.CodeHTTPUnmarshal, "invalid_request_body"))
		return
	}

	college, err := e.svc.College.Update(ctx, nim, &req)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, college)
}

// DeleteCollege godoc
//
//	@Summary		Delete a college
//	@Description	Delete a college identified by NIM
//	@Tags			colleges
//	@Produce		json
//	@Param			nim	path		string	true	"College NIM"
//	@Success		200	{object}	dto.HttpSuccessResp
//	@Failure		400	{object}	dto.HTTPErrorResp
//	@Failure		404	{object}	dto.HTTPErrorResp
//	@Router			/college/{nim} [delete]
func (e *rest) DeleteCollege(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	nim := r.PathValue("nim")
	if nim == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_college_nim")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_nim"))
		return
	}

	err := e.svc.College.Delete(ctx, nim)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, nil)
}
