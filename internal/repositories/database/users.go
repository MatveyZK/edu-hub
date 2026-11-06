package database

import (
	"database/sql"
	"fmt"

	"github.com/MatveyZK/edu-hub/internal/entities"
	"github.com/jmoiron/sqlx"
)

type UserStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) GetAll() ([]entities.User, error) {
	var users []entities.User

	query := `
	SELECT id, first_name, last_name, email, created_at, password_hash, role_id
	FROM users;`

	err := s.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserStore) GetByRoleID(role_id int) ([]entities.User, error) {
	var students []entities.User

	query := `
	SELECT id, first_name, last_name, email, created_at, password_hash, role_id
	FROM users
	WHERE role_id = $1;`

	err := s.db.Select(&students, query, role_id)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf(`users with role_id %d not found`, role_id)
	}

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *UserStore) GetStudents() ([]entities.User, error) {
	var students []entities.User

	query := `
	SELECT id, first_name, last_name, email, created_at, password_hash, role_id
	FROM users
	WHERE role_id = $1;`

	err := s.db.Select(&students, query, 1)

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *UserStore) GetByID(id int) (*entities.User, error) {
	var user entities.User

	query := `
	SELECT id, first_name, last_name, email, created_at, password_hash, role_id
	FROM users
	WHERE role_id = $1;`

	err := s.db.Get(&user, query, id)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf(`user with id %d not found`, id)
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) Create(input entities.CreateUserInput) (*entities.User, error) {
	var user entities.User

	query := `
	INSERT INTO users 
	(first_name, last_name, email, password_hash, role_id)
	VALUES ($1, $2, $3, $4, $5)`

	err := s.db.QueryRowx(query, input.FirstName, input.LastName, input.Email, input.PasswordHash, input.RoleID).StructScan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserStore) Update(id int, input entities.UpdateUserInput) (*entities.User, error) {
	user, err := s.GetByID(id)

	if err != nil {
		return nil, err
	}

	query := `
	UPDATE users 
	SET first_name=$1, last_name=$2, email=$3, password_hash=$4
	WHERE id=$5;`

	_, err = s.db.Exec(query, user.FirstName, user.LastName, user.Email, user.PasswordHash, id)

	if err != nil {
		return nil, err
	}

	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		user.LastName = *input.LastName
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.PasswordHash != nil {
		user.PasswordHash = *input.PasswordHash
	}

	return user, nil
}

func (s *UserStore) Delete(id int) error {
	query := `DELETE FROM users WHERE id=$1;`

	result, err := s.db.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf(`user with id %d noy found`, id)
	}

	return nil
}
