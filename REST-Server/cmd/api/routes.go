package main

import (
	"net/http"
	"server/utils/cor"

	"github.com/julienschmidt/httprouter"
)

func CreateRoutes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/status", StatusHandler)

	return cor.EnableCORS(router)
}
