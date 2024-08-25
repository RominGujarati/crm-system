package models

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Company  string `json:"company" bson:"company"`
	Status   string `json:"status" bson:"status"`
	Notes    string `json:"notes" bson:"notes"`
}
