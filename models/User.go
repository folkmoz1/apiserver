package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	Content string
	TextFrom string
	Image string `default:"nil"`
	Timestamp time.Time
}

type User struct {
	ID	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Messages []Message `json:"messages,omitempty" bson:"messages,omitempty"`
}

