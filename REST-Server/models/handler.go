package models

import (
	"encoding/json"
	"log"
	"net/http"
	writer "server/utils/write"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (m *DBModel) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	param := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(param.ByName("id"))

	if err != nil {
		log.Println("Invalid id parameter")
		writer.ErrorJson(w, err)
		return
	}

	err = m.DeleteUser(id)

	if err != nil {
		writer.ErrorJson(w, err)
		return
	}

	err = writer.WriteJSON(w, http.StatusOK, response{
		Status:  "OK",
		Message: "User deleted successfully",
	})
	if err != nil {
		writer.ErrorJson(w, err)
		return
	}

}

func (m *DBModel) InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload User

	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		log.Println("Error prasing JSON", err)
		writer.ErrorJson(w, err)
		return
	}
	err = m.InsertUser(payload)

	if err != nil {
		writer.ErrorJson(w, err)
		return
	}

	err = writer.WriteJSON(w, 200, response{
		Status:  "OK",
		Message: "Data added successfully",
	})

	if err != nil {
		writer.ErrorJson(w, err)
		return
	}

}

func (m *DBModel) GetOneUserHandler(w http.ResponseWriter, r *http.Request) {
	param := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(param.ByName("id"))

	if err != nil {
		log.Println("Invalid id parameter")
		writer.ErrorJson(w, err)
		return
	}

	user, err := m.GetUser(id)

	if err != nil {
		writer.ErrorJson(w, err)
		return
	}

	err = writer.WriteJSON(w, http.StatusOK, user)
	if err != nil {
		writer.ErrorJson(w, err)
		return
	}

}

func (m *DBModel) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := m.AllUser()

	if err != nil {
		writer.ErrorJson(w, err)
		return
	}

	writer.WriteJSON(w, 200, users)
}
