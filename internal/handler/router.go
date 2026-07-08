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
	e.mux.Handle("GET "+preference.RouteGetColleges, handler(http.HandlerFunc(e.FindAll)))
	e.mux.Handle("GET "+preference.RouteCollegeByNIM, handler(http.HandlerFunc(e.FindByNim)))
	e.mux.Handle("GET "+preference.RouteCollegeByName, handler(http.HandlerFunc(e.FindByName)))
	e.mux.Handle("GET "+preference.RouteCollegeBySemester, handler(http.HandlerFunc(e.FindBySemester)))
	e.mux.Handle("PUT "+preference.RouteCollegeByNIM, handler(http.HandlerFunc(e.UpdateCollege)))
	e.mux.Handle("DELETE "+preference.RouteCollegeByNIM, handler(http.HandlerFunc(e.DeleteCollege)))
	e.mux.Handle("POST "+preference.RouteCreateCourse, handler(http.HandlerFunc(e.CreateCourse)))
	e.mux.Handle("GET "+preference.RouteGetCourses, handler(http.HandlerFunc(e.GetAllCourses)))
	e.mux.Handle("GET "+preference.RouteCourseByCode, handler(http.HandlerFunc(e.GetCourseByCode)))
	e.mux.Handle("PUT "+preference.RouteCourseByCode, handler(http.HandlerFunc(e.UpdateCourse)))
	e.mux.Handle("DELETE "+preference.RouteCourseByCode, handler(http.HandlerFunc(e.DeleteCourse)))
	e.mux.Handle("POST "+preference.RouteCreateEnrollment, handler(http.HandlerFunc(e.CreateEnrollment)))
	e.mux.Handle("GET "+preference.RouteEnrollmentByNIM, handler(http.HandlerFunc(e.GetEnrollmentsByNim)))
	e.mux.Handle("PUT "+preference.RouteEnrollmentByNimCode, handler(http.HandlerFunc(e.UpdateEnrollment)))
	e.mux.Handle("DELETE "+preference.RouteEnrollmentByID, handler(http.HandlerFunc(e.DeleteEnrollment)))
}
