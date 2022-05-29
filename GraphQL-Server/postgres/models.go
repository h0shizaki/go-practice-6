package postgres

import (
	"log"
)

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func (d *Db) GetUserByName(name string) []User {
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")

	if err != nil {
		log.Println("GetUserByName Preparation Err: ", err)
	}

	rows, err := stmt.Query(name)

	if err != nil {
		log.Println("GetUserByName Query Err: ", err)
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

	return users
}
