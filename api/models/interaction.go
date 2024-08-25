package models

type Interaction struct {
	ID           string `json:"id" bson:"_id"`
	CustomerID   string `json:"customer_id" bson:"customer_id"`
	InteractionType string `json:"interaction_type" bson:"interaction_type"`
	Status       string `json:"status" bson:"status"`
	Details      string `json:"details" bson:"details"`
}
