package rest

import (
	"encoding/json"
	"fmt"
	"go-college/internal/model/dto"
	appErr "go-college/internal/model/errors"
	"go-college/internal/preference"
	"net/http"
	"time"
)

func (e *rest) httpRespSuccess(w http.ResponseWriter, r *http.Request, statusCode int, resp any) {
	meta := dto.Meta{
		Path:       r.URL.Path,
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.RequestURI, statusCode, http.StatusText(statusCode)),
		Error:      nil,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	httpResp := &dto.HttpSuccessResp{
		Meta:       meta,
		Data:       any(resp),
	}

	writeJSON(w, statusCode, httpResp)
}

func (e *rest) httpRespError(w http.ResponseWriter, r *http.Request, err error) {
	lang := preference.LANG_ID

	appLangHeader := http.CanonicalHeaderKey(preference.APP_LANG)
	if vals := r.Header[appLangHeader]; len(vals) > 0 && vals[0] == preference.LANG_EN {
		lang = preference.LANG_EN
	}

	statusCode, displayError := appErr.Compile(appErr.COMMON, err, lang, true)
	statusStr := http.StatusText(statusCode)

	jsonErrResp := &dto.HTTPErrorResp{
		Meta: dto.Meta{
			Path:       r.URL.Path,
			StatusCode: statusCode,
			Status:     statusStr,
			Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.RequestURI, statusCode, http.StatusText(statusCode)),
			Error:      &displayError,
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}

	writeJSON(w, statusCode, jsonErrResp)
}

func writeJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(v)
}
