package main

import (
	"fmt"
	"net/http"
	"server/log"
	"server/utils/env"
)

func main() {
	log.Log.LogInfo("Starting Server")

	env.CheckENV()
	port := env.MustGet("PORT")

	if port == "" {
		port = "4040"
	}

	log.Log.LogInfo(fmt.Sprintf("Running on port %v", port))

	router := CreateRoutes()
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)

}
