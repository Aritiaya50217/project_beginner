package utils

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`	
	Name  string `json:"name"`
}
