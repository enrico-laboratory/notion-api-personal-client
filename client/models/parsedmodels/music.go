package parsedmodels

import "time"

type Music struct {
	Id             string    `bson:"_id" json:"id"`
	CreatedTime    time.Time `bson:"created_time" json:"created_time"`
	LastEditedTime time.Time `bson:"last_edited_time" json:"last_edited_time"`
	Voices         string    `json:"voices"`
	Score          string    `json:"score"`
	Media          string    `json:"media"`
	Recording      string    `json:"recording"`
	Composer       string    `json:"composer"`
	Length         float64   `json:"length"`
	Instruments    []string  `json:"instruments"`
	Solo           string    `json:"solo"`
	Title          string    `json:"title"`
}
