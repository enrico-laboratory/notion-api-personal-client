package parsedmodels

import (
	"time"
)

type Piece struct {
	Id             string    `bson:"_id" json:"id"`
	CreatedTime    time.Time `bson:"created_time" json:"created_time"`
	LastEditedTime time.Time `bson:"last_edited_time" json:"last_edited_time"`
	Order          string    `json:"order"`
	Num1           string    `json:"1"`
	Num2           string    `json:"2"`
	Num3           string    `json:"3"`
	Num4           string    `json:"4"`
	Num5           string    `json:"5"`
	Num6           string    `json:"6"`
	Num7           string    `json:"7"`
	Num8           string    `json:"8"`
	Num9           string    `json:"9"`
	Num10          string    `json:"10"`
	Num11          string    `json:"11"`
	Num12          string    `json:"12"`
	Note           string    `json:"note"`
	Voicing        string    `json:"voicing"`
	Solo           string    `json:"solo"`
	Music          string    `json:"music"`
	Selected       bool      `json:"selected"`
	MusicProject   []string  `json:"music_project"`
	Media          string    `json:"media"`
	Score          string    `json:"score"`
	Recording      string    `json:"recording"`
	Instrument     []string  `json:"instrument"`
	Length         float64   `json:"length"`
	Composer       string    `json:"composer"`
}
