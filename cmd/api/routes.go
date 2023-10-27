package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/partner/v2/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/partner/v2/origins/:origin_id", app.getOriginsById)

	return router
}
