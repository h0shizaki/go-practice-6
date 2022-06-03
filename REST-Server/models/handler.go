package models

import (
	"fmt"
	"net/http"
	writer "server/utils/write"
)

func (m *DBModel) GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := m.AllUser()

	if err != nil {
		fmt.Println(err)
	}

	writer.WriteJSON(w, 200, users)

}
