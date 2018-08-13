package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"github.com/s4kibs4mi/zeroserver/log"
)

var router = chi.NewRouter()

// Router returns the api router
func Router() http.Handler {
	router.Use(middleware.DefaultLogger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			code: http.StatusOK,
			Data: "App Running",
		}
		resp.ServeJSON(w)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			code: http.StatusNotFound,
			Data: "Not found",
		}
		resp.ServeJSON(w)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			code: http.StatusMethodNotAllowed,
			Data: "Method not allowed",
		}
		resp.ServeJSON(w)
	})

	router.Post("/auth", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		log.Logger().Infoln(string(b))
		log.Logger().Infoln(r.Header)

		resp := response{
			code: http.StatusOK,
			Data: "Ok",
		}
		resp.ServeJSON(w)
	})

	router.Post("/acl", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		log.Logger().Infoln(string(b))
		log.Logger().Infoln(r.Header)

		resp := response{
			code: http.StatusOK,
			Data: "Unauthorized",
		}
		resp.ServeJSON(w)
	})

	router.Get("/acl", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		log.Logger().Infoln(string(b))

		resp := response{
			code: http.StatusOK,
			Data: "Unauthorized",
		}
		resp.ServeJSON(w)
	})

	router.Get("/hook", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		log.Logger().Infoln(string(b))

		resp := response{
			code: http.StatusOK,
			Data: "Unauthorized",
		}
		resp.ServeJSON(w)
	})
	router.Post("/hook", func(w http.ResponseWriter, r *http.Request) {
		b, _ := ioutil.ReadAll(r.Body)
		log.Logger().Infoln(string(b))

		resp := response{
			code: http.StatusOK,
			Data: "Unauthorized",
		}
		resp.ServeJSON(w)
	})
	return router
}
