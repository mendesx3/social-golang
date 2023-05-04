package user

import (
	"api/src/config"
	"api/src/model"
	"database/sql"
	"fmt"
	"log"
)

type RepositoryUser interface {
	Create(request model.User) (int64, error)
	Update(request model.User) (*model.User, error)
	Delete(id int) error
	GetUserById(id int) (*model.User, error)
	NewRepositoryUser(db *sql.DB) (*Repository, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepositoryUser() (*Repository, error) {
	db, err := config.InitMySql()
	if err != nil {
		log.Println("Error InitMySql: ", err.Error())
		return nil, err
	}
	return &Repository{db}, nil
}

func (u Repository) Create(request model.User) (int64, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	stmt, err := tx.Prepare("INSERT into user (name, nick, email, password) values(?,?,?,?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(request.Name, request.Nick, request.Email, request.Password)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %v", err)
	}

	log.Printf("User created successfully with ID %d", id)

	return id, nil
}

func (u Repository) Update(request model.User) (*model.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	stmt, err := tx.Prepare("UPDATE user SET name = ?, nick = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(request.Name, request.Nick, request.Email, request.Password, request.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get rows affected %v", err)
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	log.Printf("User updated successfully with ID %d", request.ID)

	return &request, nil
}

func (u Repository) GetUserById(id int) (*model.User, error) {
	stmt, err := u.db.Prepare("SELECT * FROM user WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var user model.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Password, &user.Created)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, sql.ErrNoRows
	}

	return &user, nil
}

func (u Repository) Delete(id int) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	stmt, err := tx.Prepare("delete from user where id = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("failed to execute statement %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected %v", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	log.Printf("User deleted successfully with ID %d", id)

	return nil
}
