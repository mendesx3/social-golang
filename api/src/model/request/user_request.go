package request

import "time"

// user social
type UserRequest struct {
	ID       uint64    `json:id,omitempty`
	Name     string    `json:name,omitempty`
	Password string    `json:password,omitempty`
	Active   *bool     `json:active,omitempty,default:true`
	Created  time.Time `json:created,omitempty,`
}

/*
{
	"id":1,
	"name":"andre",
	"password":"andre123"
}
*/
