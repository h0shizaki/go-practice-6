package models

import "log"

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func (m *DBModel) AllUser() ([]User, error) {
	stmt, err := m.DB.Prepare("SELECT * FROM users ")

	if err != nil {
		log.Println("GetAll Preparation Err: ", err)
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		log.Println("GetAll Err: ", err)
		return nil, err
	}

	var r User

	users := []User{}

	for rows.Next() {
		if err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		); err != nil {
			log.Println("Error scanning rows Err: ", err)
		}
		users = append(users, r)
	}

	return users, nil
}
