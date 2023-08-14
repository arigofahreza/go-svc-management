package models

type LoginModel struct {
	Username string
	Password string
}

type TokenModel struct {
	AccessToken  string
	RefreshToken string
}
