package parsedmodels

import "time"

type Location struct {
	Id             string    `bson:"_id" json:"id"`
	CreatedTime    time.Time `bson:"created_time" json:"created_time"`
	LastEditedTime time.Time `bson:"last_edited_time" json:"last_edited_time"`
	Contact        string    `json:"contact"`
	Phone          string    `json:"phone"`
	City           string    `json:"city"`
	Email          string    `json:"email"`
	Task           []string  `json:"task"`
	Purpose        []string  `json:"purpose"`
	Address        string    `json:"address"`
	Location       string    `json:"location"`
}
