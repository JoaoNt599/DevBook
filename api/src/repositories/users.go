package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	results, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lstId, erro := results.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lstId), nil
}
