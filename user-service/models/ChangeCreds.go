package models

type UserChangeCreds struct {
	Username string `json:"username"`
	NewUsername string `json:"new_username"`
	NewPassword string `json:"new_password"`
}