package models

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionModel struct {
	SessionId string `json:"session_id"`
}

type TokenModel struct {
	ActiveToken  string `json:"active_token"`
	RefreshToken string `json:"refresh_token"`
}

