package repository

import (
	"database/sql"

	"github.com/LeoCunha98/urubu-do-pix/internal/domain/entity"
)

type UserRepositoryMySql struct {
	DB *sql.DB
}

func NewUserRepositoryMySql(db *sql.DB) *UserRepositoryMySql {
	return &UserRepositoryMySql{DB: db}
}

func (r *UserRepositoryMySql) Create(user *entity.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, name) VALUES (?, ?)", user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryMySql) FindAll() ([]*entity.User, error) {
	rows, err := r.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
