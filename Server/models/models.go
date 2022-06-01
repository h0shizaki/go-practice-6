package models

import (
	"context"
	"server/config"
	"time"
)

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func GetAllUsers() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query, err := config.DB.Prepare(`SELECT * FROM users ;`)

	if err != nil {
		return nil, err
	}

	rows, err := query.QueryContext(ctx)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []*User

	for rows.Next() {
		var u User
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

	return users, nil
}
