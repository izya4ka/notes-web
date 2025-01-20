package models

type LogPassReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenReq struct {
	Token string `json:"token"`
}
