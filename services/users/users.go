package users

import "sinappsebackend/app"

type User struct {
	Id uint32 `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email string `db:"email" json:"email"`
}

func GetByEmail(email string) (*User, error) {
	var u User
	err := app.DB.Get(&u, "SELECT * FROM users WHERE email=$1 LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetById(id uint32) (*User, error) {
	var u User
	err := app.DB.Get(&u, "SELECT * FROM users WHERE id=$1 LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(u User) (uint32, error) {
	rows, err := app.DB.NamedQuery("INSERT INTO users (username, email) VALUES (:username, :email) RETURNING id", &u)
	if err != nil {
		return 0, err
	}
	rows.Next()
	err = rows.StructScan(&u)
	if err != nil {
		return 0, err
	}
	return u.Id, nil
}