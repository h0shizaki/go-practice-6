package status

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-server/utils/env"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := GetStatus()

	js, err := json.MarshalIndent(currentStatus, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func GetStatus() *Status {
	env.CheckENV()

	return &Status{
		Status:      "Available",
		Environment: env.MustGet("ENV"),
	}

}
