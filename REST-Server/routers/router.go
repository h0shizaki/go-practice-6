package routers

import (
	"net/http"
	"rest-server/data/status"
	"rest-server/utils/cor"

	"github.com/julienschmidt/httprouter"
)

func CreateRoutes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/status", status.StatusHandler)

	return cor.EnableCORS(router)
}
