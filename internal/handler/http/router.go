package rest

import (
	"net/http"
	"sync"

	"go-college/internal/middleware"
	"go-college/internal/preference"
	"go-college/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type rest struct {
	mux  *http.ServeMux
	mw   middleware.Middleware
	svc  *service.Service
	sql0 *pgxpool.Pool
}

var onceRestHandler = &sync.Once{}

func InitHttpHandler(mux *http.ServeMux, mw middleware.Middleware, svc *service.Service, sql0 *pgxpool.Pool) {
	var e *rest

	onceRestHandler.Do(func() {
		e = &rest{
			mux:  mux,
			mw:   mw,
			svc:  svc,
			sql0: sql0,
		}

		e.Serve()
	})
}

func (e *rest) Serve() {
	handler := e.mw.Handler()

	e.mux.Handle("POST "+preference.RouteCreateCollege, handler(http.HandlerFunc(e.CreateCollege)))
}
