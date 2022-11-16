package model

import "time"

//user social
type User struct {
	ID       uint64    `json:id,omitempty`
	Name     string    `json:name,omitempty`
	Password string    `json:password,omitempty`
	Active   *bool     `json:active,omitempty,default:true`
	created  time.Time `json:created,omitempty`
}
