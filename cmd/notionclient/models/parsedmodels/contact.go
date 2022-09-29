package parsedmodels

import "time"

type Contact struct {
	Id             string    `json:"id"`
	CreatedTime    time.Time `bson:"created_time" json:"created_time"`
	LastEditedTime time.Time `bson:"last_edited_time" json:"last_edited_time"`
	Singer         string    `json:"singer"`
	SingerId       string    `json:"singerId"`
	Note           string    `json:"note"`
	Phone          string    `json:"phone"`
	Status         string    `json:"status,omitempty"`
	Order          float64   `json:"order"`
	Email          string    `json:"email,omitempty"`
	MusicProject   []string  `json:"music_project"`
	Role           string    `json:"role"`
}
