package main

import (
	"fmt"
	"net/http"
	"server/log"
	"server/utils/env"
	"server/utils/write"
)

type Status struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

func CheckStatus(w http.ResponseWriter, r *http.Request) {
	log.Log.LogRequest(r)
	currentStatus := GetStatus()

	err := write.WriteJSON(w, 200, currentStatus)

	if err != nil {
		log.Log.LogDanger(fmt.Sprintf("%v", err))
	}

}

func GetStatus() *Status {
	env.CheckENV()

	return &Status{
		Status:      "Available",
		Environment: env.MustGet("ENV"),
	}

}
