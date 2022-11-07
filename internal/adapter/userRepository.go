package adapter

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

func (ur UserRepository) SaveUser(user entities.User) (string, error) {
	return query{
		Table:      "users (username, password)",
		Parameters: []any{user.Username, user.Password},
		db:         ur.db,
	}.Insert()
}

func (ur UserRepository) FindUserByUsername(username string) (*entities.User, error) {
	var u entities.User
	err := ur.db.QueryRow("SELECT * from users where username = $1", username).
		Scan(&u.Id, &u.Username, &u.Password)

	if err == nil {
		return &u, nil
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return nil, err
}

func (ur UserRepository) FindById(id string) (*entities.User, error) {
	var u entities.User
	err := ur.db.QueryRow("SELECT * from users where id = $1", id).
		Scan(&u.Id, &u.Username, &u.Password)

	if err == nil {
		return &u, nil
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return nil, err
}
