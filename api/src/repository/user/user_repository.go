package user

import (
	"api/src/config"
	model "api/src/model/request"
	"api/src/model/response"
	"database/sql"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

func Create(request model.UserRequest) int64 {
	db, _ := config.InitMySql()
	repo := create(db)
	u, err := repo.Create(request)
	if err != nil {
		return 0
	}
	fmt.Println("usuario criado com sucesso", u)
	return u
}

func create(db *sql.DB) *userRepository {
	return &userRepository{db}
}
func (u userRepository) Create(usuario model.UserRequest) (int64, error) {
	attmt, err := u.db.Prepare(
		"insert into user (name, password) values (?, ?)",
	)
	if err != nil {
		fmt.Println(err)
	}
	defer attmt.Close()

	result, err := attmt.Exec(usuario.Name, usuario.Password)
	if err != nil {
		fmt.Println(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	return int64(id), nil
}

func GetAll() []response.UserResponse {
	db, _ := config.InitMySql()
	repo := create(db)
	u, err := repo.GetAll()
	if err != nil {
		return nil
	}
	fmt.Println("usuario criado com sucesso", u)
	return u
}

func (u userRepository) GetAll() ([]response.UserResponse, []response.UserResponse) {
	rows, err := u.db.Query("SELECT * FROM dual")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var userResponse []response.UserResponse

	for rows.Next() {
		var user response.UserResponse

		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, nil
		}

		userResponse = append(userResponse, user)

		if err = rows.Err(); err != nil {
			return nil, nil
		}
	}
	return userResponse, nil
}
