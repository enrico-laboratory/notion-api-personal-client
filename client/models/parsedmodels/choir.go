package parsedmodels

import "time"

type Choir struct {
	Id             string    `bson:"_id" json:"id"`
	CreatedTime    time.Time `bson:"created_time" json:"created_time"`
	LastEditedTime time.Time `bson:"last_edited_time" json:"last_edited_time"`
	Name           string    `json:"name"`
}
