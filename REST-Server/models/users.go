package models

import (
	"context"
	"log"
	"time"
)

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func (m *DBModel) DeleteUser(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := m.DB.Prepare("DELETE FROM users where id = $1")

	if err != nil {
		log.Println("Delete user Preparation Err:", err)
		return err
	}

	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		log.Println("Delete user execute Err:", err)
		return err
	}

	return nil
}

func (m *DBModel) InsertUser(payload User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := m.DB.Prepare("INSERT INTO users(name,age,profession,friendly) values($1,$2,$3,$4) ;")

	if err != nil {
		log.Println("Insert user Preparation Err:", err)
		return err
	}

	_, err = stmt.ExecContext(ctx, payload.Name, payload.Age, payload.Profession, payload.Friendly)

	if err != nil {
		log.Println("Insert user execute Err:", err)
		return err
	}

	return nil
}

func (m *DBModel) GetUser(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := m.DB.Prepare("SELECT * FROM users where id = $1")

	if err != nil {
		log.Println("GetOne Preparation Err:", err)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, id)

	var user User

	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.Profession,
		&user.Friendly,
	)

	if err != nil {
		log.Println("GetOne Err:", err)
		return nil, err
	}

	return &user, nil

}

func (m *DBModel) AllUser() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt, err := m.DB.Prepare("SELECT * FROM users ")

	if err != nil {
		log.Println("GetAll Preparation Err:", err)
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx)

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
			log.Println("Error scanning rows Err:", err)
		}
		users = append(users, r)
	}

	return users, nil
}
