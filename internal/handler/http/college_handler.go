package rest

import (
	"encoding/json"
	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"
	"net/http"

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
