package parsedmodels

import "time"

type MusicProject struct {
	Id             string    `bson:"_id" json:"id"`
	CreatedTime    time.Time `bson:"created_time" json:"created_time"`
	LastEditedTime time.Time `bson:"last_edited_time" json:"last_edited_time"`
	Title          string    `bson:"title" json:"title"`
	Year           int       `bson:"year" json:"year"`
	Status         string    `bson:"status" json:"status"`
	Description    string    `bson:"description" json:"description"`
	ChoirRollup    string    `bson:"choir_rollup" json:"choir_rollup"`
}
