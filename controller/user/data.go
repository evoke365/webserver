package user

// LoginResponse defines the response data structure of /login
type LoginResponse struct {
	Token string `json:"token"`
}

// VerifyUserResponse defines the response data structure of /verify
type VerifyUserResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
