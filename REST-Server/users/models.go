package users

import (
	"server/config"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Profession string `json:"profession"`
	Friendly   bool   `json:"friendly"`
}

func AllUsers() ([]*User, error) {

	rows, err := config.DB.Query("SELECT * FROM users ;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		u := User{}
		err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Age,
			&u.Profession,
			&u.Friendly,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
