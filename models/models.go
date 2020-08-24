package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type News struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	guideid       string             `json:"guideid,omitempty" bson:"guideid,omitempty"`
	useridcomment string             `json:"useridcomment" bson:"useridcomment,omitempty"`
	useremail     string             `json:"useremail" bson:"useremail,omitempty"`
	avatar        string             `json:"avatar" bson:"avatar,omitempty"`
	author        string             `json:"author" bson:"author,omitempty"`
	datecomment   string             `json:"datecomment" bson:"datecomment,omitempty"`
	message       string             `json:"message" bson:"message,omitempty"`
}
