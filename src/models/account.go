package models

type AccountModel struct {
	Id       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
	Address  string `json:"address,omitempty" bson:"addres"`
	Email    string `json:"email,omitempty" bson:"email"`
}

type AccountView struct {
	Id       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Name     string `json:"name" bson:"name"`
	Address  string `json:"address,omitempty" bson:"addres"`
	Email    string `json:"email,omitempty" bson:"email"`
}
