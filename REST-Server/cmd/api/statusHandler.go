package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/log"
	"server/utils/env"
)

type Status struct {
	Status      string
	Environment string
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	currentStatus := GetStatus()

	js, err := json.MarshalIndent(currentStatus, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	log.Log.LogRequest(r)
	w.Write(js)
}

func GetStatus() *Status {
	env.CheckENV()

	return &Status{
		Status:      "Available",
		Environment: env.MustGet("ENV"),
	}

}
