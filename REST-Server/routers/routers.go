package routers

import (
	"net/http"
	"server/models"
	"server/utils/cor"

	"github.com/julienschmidt/httprouter"
)

func CreateRoute(modelHander models.Models) http.Handler {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/users", modelHander.DB.GetAllUserHandler)

	return cor.EnableCORS(router)

}
