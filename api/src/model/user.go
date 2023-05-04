package model

import (
	"api/src/model/request"
	"api/src/model/response"
	"time"
)

// User struct User
type User struct {
	ID       int       `json:"id,omitempty,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

func UserRequestToUser(userRequest request.UserRequest) *User {
	return &User{
		ID:       userRequest.ID,
		Name:     userRequest.Name,
		Nick:     userRequest.Nick,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Created:  userRequest.Created,
	}
}

func UserToUserRequest(user User) *response.UserResponse {
	return &response.UserResponse{
		Name:     user.Name,
		Nick:     user.Nick,
		Email:    user.Email,
		Password: user.Password,
	}
}
