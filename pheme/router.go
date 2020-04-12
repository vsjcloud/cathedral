package pheme

import (
	"net/http"

	"github.com/go-chi/chi"

	"cathedral/pheme/middleware"
)

func (c *Cathedral) buildRouter() chi.Router {
	router := chi.NewRouter()
	if c.Config.Mode != "production" {
		router.Use(middleware.Logger(c.Logger))
	}
	router.Use(middleware.Recovery(c.Logger))
	router.Route(c.Config.HTTP.BasePath, func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Hello, World!"))
		})
	})
	return router
}
