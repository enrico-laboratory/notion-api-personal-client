package unparsedmodels

import "time"

type Choirs struct {
	Object          string   `json:"object"`
	Results         []Choir  `json:"results"`
	NextCursor      string   `json:"next_cursor"`
	HasMore         bool     `json:"has_more"`
	Type            string   `json:"type"`
	PageOrDatabase  struct{} `json:"page_or_database"`
	DeveloperSurvey string   `json:"developer_survey"`
}

type Choir struct {
	Object         string      `json:"object"`
	ID             string      `json:"id"`
	CreatedTime    time.Time   `json:"created_time"`
	LastEditedTime time.Time   `json:"last_edited_time"`
	CreatedBy      ObjectAndId `json:"created_by"`
	LastEditedBy   ObjectAndId `json:"last_edited_by"`
	Cover          any         `json:"cover"`
	Icon           any         `json:"icon"`
	Parent         Parent      `json:"parent"`
	Archived       bool        `json:"archived"`
	Properties     struct {
		Hired    Checkbox `json:"Hired"`
		Projects Relation `json:"Projects"`
		Choir    Title    `json:"Choir"`
	} `json:"properties"`
	URL       string `json:"url"`
	PublicURL any    `json:"public_url"`
}
