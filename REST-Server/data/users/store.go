package users

import (
	"context"
	"database/sql"
	"time"
)

//CRUD

type Store struct {
	db *sql.DB
}

func (s *Store) GetAllUser() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query, err := s.db.Prepare(`SELECT * FROM users ;`)

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
