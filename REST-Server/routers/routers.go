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
	router.HandlerFunc(http.MethodGet, "/api/user/:id", modelHander.DB.GetOneUserHandler)
	router.HandlerFunc(http.MethodPost, "/api/user", modelHander.DB.InsertUserHandler)
	router.HandlerFunc(http.MethodDelete, "/api/user/:id", modelHander.DB.DeleteUserHandler)

	return cor.EnableCORS(router)

}
