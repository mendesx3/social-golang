package request

import "time"

// UserRequest struct UserRequest
type UserRequest struct {
	ID       int       `json:"id,omitempty,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}
