package rest

import (
	"encoding/json"
	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"
	"net/http"
	"strconv"

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
//	@Summary		List all colleges
//	@Description	Retrieve all college records
//	@Tags			colleges
//	@Produce		json
//	@Success		200	{object}	dto.HttpSuccessResp{data=[]entity.College}
//	@Failure		500	{object}	dto.HTTPErrorResp
//	@Router			/college/all [get]
func (e *rest) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	colleges, err := e.svc.College.FindAll(ctx)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, colleges)
}

// FindByNim godoc
//
//	@Summary		Get a college by NIM
//	@Description	Retrieve a single college by its NIM
//	@Tags			colleges
//	@Produce		json
//	@Param			nim	path		string	true	"College NIM"
//	@Success		200	{object}	dto.HttpSuccessResp{data=entity.College}
//	@Failure		400	{object}	dto.HTTPErrorResp
//	@Failure		404	{object}	dto.HTTPErrorResp
//	@Router			/college/{nim} [get]
func (e *rest) FindByNim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	nim := r.PathValue("nim")
	if nim == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_college_id")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_id"))
		return
	}

	college, err := e.svc.College.FindByNim(ctx, nim)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, college)
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

// FindByName godoc
//
//	@Summary		Find colleges by name
//	@Description	Retrieve colleges matching the given name
//	@Tags			colleges
//	@Produce		json
//	@Param			name	path		string	true	"College name"
//	@Success		200		{object}	dto.HttpSuccessResp{data=[]entity.College}
//	@Failure		400		{object}	dto.HTTPErrorResp
//	@Router			/college/name/{name} [get]
func (e *rest) FindByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.PathValue("name")
	if name == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_college_name")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_name"))
		return
	}

	colleges, err := e.svc.College.FindByName(ctx, name)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, colleges)
}

// FindBySemester godoc
//
//	@Summary		Find colleges by semester
//	@Description	Retrieve colleges in the given semester
//	@Tags			colleges
//	@Produce		json
//	@Param			semester	path		int	true	"Semester"
//	@Success		200			{object}	dto.HttpSuccessResp{data=[]entity.College}
//	@Failure		400			{object}	dto.HTTPErrorResp
//	@Router			/college/semester/{semester} [get]
func (e *rest) FindBySemester(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	semester := r.PathValue("semester")
	if semester == "" {
		zerolog.Ctx(ctx).Error().Msg("invalid_college_semester")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_semester"))
		return
	}

	parsed, err := strconv.Atoi(semester)

	if err != nil {
		zerolog.Ctx(ctx).Error().Msg("invalid_college_semester")
		e.httpRespError(w, r, appErr.WrapWithCode(nil, appErr.CodeHTTPBadRequest, "invalid_college_semester"))
		return
	}

	colleges, err := e.svc.College.FindBySemester(ctx, parsed)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, colleges)
}
