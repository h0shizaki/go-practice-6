package users

import (
	"fmt"
	"net/http"
	"server/log"
	"server/utils/write"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Log.LogRequest(r)

	users, err := AllUsers()

	if err != nil {
		log.Log.LogDanger("Error getting users info ")
		return
	}

	err = write.WriteJSON(w, 200, users)

	if err != nil {
		log.Log.LogDanger(fmt.Sprintf("%v", err))
		return
	}

}
