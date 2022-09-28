package parsedmodels

import "time"

type Task struct {
	Id               string    `bson:"_id" json:"id"`
	CreatedTime      time.Time `bson:"created_time" json:"created_time"`
	LastEditedTime   time.Time `bson:"last_edited_time" json:"last_edited_time"`
	MusicProject     []string  `json:"music_project"`
	Type             string    `json:"type"`
	Priority         string    `json:"priority"`
	StartDateAndTime string    `json:"start_date_and_time"`
	EndDateAndTime   string    `json:"end_date_and_time"`
	Duration         string    `json:"duration"`
	Kanban           string    `json:"kanban"`
	Location         string    `json:"location"`
	City             string    `json:"city"`
	IsDone           bool      `json:"isDone"`
	LocationId       []string  `json:"locationId"`
	Choir            string    `json:"choir"`
	Title            string    `json:"title"`
}
