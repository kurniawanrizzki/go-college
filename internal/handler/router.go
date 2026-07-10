package rest

import (
	"net/http"
	"sync"

	"go-college/internal/middleware"
	"go-college/internal/preference"
	"go-college/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

// HTTP method prefixes used in ServeMux route patterns ("METHOD /path").
const (
	methodPost   = http.MethodPost + " "
	methodGet    = http.MethodGet + " "
	methodPut    = http.MethodPut + " "
	methodDelete = http.MethodDelete + " "
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

	e.mux.Handle(methodPost+preference.RouteCreateCollege, handler(http.HandlerFunc(e.CreateCollege)))
	e.mux.Handle(methodGet+preference.RouteGetColleges, handler(http.HandlerFunc(e.FindAll)))
	e.mux.Handle(methodGet+preference.RouteCollegeByNIM, handler(http.HandlerFunc(e.FindByNim)))
	e.mux.Handle(methodGet+preference.RouteCollegeByName, handler(http.HandlerFunc(e.FindByName)))
	e.mux.Handle(methodGet+preference.RouteCollegeBySemester, handler(http.HandlerFunc(e.FindBySemester)))
	e.mux.Handle(methodPut+preference.RouteCollegeByNIM, handler(http.HandlerFunc(e.UpdateCollege)))
	e.mux.Handle(methodDelete+preference.RouteCollegeByNIM, handler(http.HandlerFunc(e.DeleteCollege)))
	e.mux.Handle(methodPost+preference.RouteCreateCourse, handler(http.HandlerFunc(e.CreateCourse)))
	e.mux.Handle(methodGet+preference.RouteGetCourses, handler(http.HandlerFunc(e.GetAllCourses)))
	e.mux.Handle(methodGet+preference.RouteCourseByCode, handler(http.HandlerFunc(e.GetCourseByCode)))
	e.mux.Handle(methodPut+preference.RouteCourseByCode, handler(http.HandlerFunc(e.UpdateCourse)))
	e.mux.Handle(methodDelete+preference.RouteCourseByCode, handler(http.HandlerFunc(e.DeleteCourse)))
	e.mux.Handle(methodPost+preference.RouteCreateEnrollment, handler(http.HandlerFunc(e.CreateEnrollment)))
	e.mux.Handle(methodGet+preference.RouteEnrollmentByNIM, handler(http.HandlerFunc(e.GetEnrollmentsByNim)))
	e.mux.Handle(methodPut+preference.RouteEnrollmentByNimCode, handler(http.HandlerFunc(e.UpdateEnrollment)))
	e.mux.Handle(methodDelete+preference.RouteEnrollmentByID, handler(http.HandlerFunc(e.DeleteEnrollment)))
}
