package models

type AccountModel struct {
	Id          string `json:"_id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	Address     string `json:"address" bson:"address"`
	Email       string `json:"email" bson:"email"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
	Gender      string `json:"gender" bson:"gender"`
}

type AccountView struct {
	Id          string `json:"_id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	Address     string `json:"address" bson:"address"`
	Email       string `json:"email" bson:"email"`
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
	Gender      string `json:"gender" bson:"gender"`
}
