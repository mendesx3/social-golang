package response

//user social
type UserResponse struct {
	ID   uint64 `json:id,omitempty`
	Name string `json:name,omitempty`
}
