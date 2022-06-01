package main

import (
	"net/http"
	"server/users"
	"server/utils/cor"

	"github.com/julienschmidt/httprouter"
)

func CreateRoutes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/status", CheckStatus)

	router.HandlerFunc(http.MethodGet, "/api/users", users.GetAllUsersHandler)
	return cor.EnableCORS(router)
}
