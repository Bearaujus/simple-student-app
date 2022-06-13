package init

import (
	"net/http"

	handler "github.com/Bearaujus/simple-student-app/internal/handler"
	studentHandler "github.com/Bearaujus/simple-student-app/internal/handler/student"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type route struct {
	pattern string
	methods []string
	handler http.Handler
}

func InitRouter(sh studentHandler.HandlerItf) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.MethodNotAllowed(http.NotFound)

	var routes []route = []route{
		{
			pattern: "/student/{sid}",
			methods: []string{http.MethodGet, http.MethodOptions},
			handler: handler.Handler(sh.HandleGetStudentByID),
		},
		{
			pattern: "/students",
			methods: []string{http.MethodGet, http.MethodOptions},
			handler: handler.Handler(sh.HandleGetStudents),
		},
		{
			pattern: "/student",
			methods: []string{http.MethodPost, http.MethodOptions},
			handler: handler.Handler(sh.HandleCreateStudent),
		},
		{
			pattern: "/student/{sid}",
			methods: []string{http.MethodPut, http.MethodPatch, http.MethodOptions},
			handler: handler.Handler(sh.HandleUpdateStudent),
		},
		{
			pattern: "/student/{sid}",
			methods: []string{http.MethodDelete, http.MethodOptions},
			handler: handler.Handler(sh.HandleDeleteStudent),
		},
	}

	for _, route := range routes {
		createRoutes(router, route.pattern, route.methods, route.handler)
	}

	return router
}

func createRoutes(router *chi.Mux, pattern string, methods []string, handler http.Handler) {
	for _, method := range methods {
		router.Method(method, pattern, handler)
	}
}
