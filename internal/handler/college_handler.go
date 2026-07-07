package rest

import (
	"encoding/json"
	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"
	"net/http"
	"strconv"

	"github.com/rs/zerolog"
)

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

func (e *rest) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	colleges, err := e.svc.College.FindAll(ctx)

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, colleges)
}

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
