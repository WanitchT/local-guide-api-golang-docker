package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type Comments struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Guideid       string             `json:"guideid,omitempty" bson:"guideid,omitempty"`
	UseridComment string             `json:"useridcomment" bson:"useridcomment,omitempty"`
	Useremail     string             `json:"useremail" bson:"useremail,omitempty"`
	Avatar        string             `json:"avatar" bson:"avatar,omitempty"`
	Author        string             `json:"author" bson:"author,omitempty"`
	DateComment   string             `json:"datecomment" bson:"datecomment,omitempty"`
	Message       string             `json:"message" bson:"message,omitempty"`
}
