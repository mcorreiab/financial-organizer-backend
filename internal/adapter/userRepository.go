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
	_, err := ur.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func (ur UserRepository) FindUserByUsername(username string) (*entities.User, error) {
	var u entities.User
	err := ur.db.QueryRow("SELECT * from users where username = $1", username).Scan(&u.Username, &u.Password)

	if err == nil {
		return &u, nil
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return nil, err
}
