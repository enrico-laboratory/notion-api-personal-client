package unparsedmodels

import "time"

// MusicProjects Commented out the properties presently not needed
type MusicProjects struct {
	Object     string         `json:"object"`
	Results    []MusicProject `json:"results"`
	NextCursor string         `json:"next_cursor"`
	HasMore    bool           `json:"has_more"`
}

type MusicProjectCreateResponse struct {
	Object          string       `json:"object"`
	ID              string       `json:"id"`
	CreatedTime     time.Time    `json:"created_time"`
	LastEditedTime  time.Time    `json:"last_edited_time"`
	CreatedBy       ObjectAndId  `json:"created_by"`
	LastEditedBy    ObjectAndId  `json:"last_edited_by"`
	Cover           interface{}  `json:"cover"`
	Icon            interface{}  `json:"icon"`
	Parent          Parent       `json:"parent"`
	Properties      MusicProject `json:"properties"`
	Archived        bool         `json:"archived"`
	URL             string       `json:"url"`
	PublicURL       any          `json:"public_url"`
	DeveloperSurvey string       `json:"developer_survey"`
}

type MusicProject struct {
	Object         string      `json:"object"`
	ID             string      `json:"id"`
	CreatedTime    time.Time   `json:"created_time"`
	LastEditedTime time.Time   `json:"last_edited_time"`
	CreatedBy      ObjectAndId `json:"created_by"`
	LastEditedBy   ObjectAndId `json:"last_edited_by"`
	Cover          interface{} `json:"cover"`
	Icon           interface{} `json:"icon"`
	Parent         Parent      `json:"parent"`
	Archived       bool        `json:"archived"`
	Properties     struct {
		CreatedTime CreatedTime `json:"Created Time"`
		StartDate   Date        `json:"Start Date"`
		Year        NumberInt   `json:"Year"`
		Status      Select      `json:"Status"`
		Repertoire  Relation    `json:"Repertoire"`
		Cancelled   Checkbox    `json:"Cancelled"`
		Task        Relation    `json:"Task"`
		Description RichText    `json:"Description"`
		Completed   Checkbox    `json:"Completed"`
		Notes       Relation    `json:"Notes"`
		Cast        Relation    `json:"Cast"`
		Choir       Relation    `json:"Choir"`
		MediaVault  Relation    `json:"Media Vault"`
		TaskRollup  struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string  `json:"type"`
				Array    []Title `json:"array"`
				Function string  `json:"function"`
			} `json:"rollup"`
		} `json:"Task Rollup"`
		ChoirRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string  `json:"type"`
				Array    []Title `json:"array"`
				Function string  `json:"function"`
			} `json:"rollup"`
		} `json:"Choir Rollup"`
		EndDate   Date     `json:"End Date"`
		Published Checkbox `json:"Published"`
		Projects  Relation `json:"Projects"`
		Title     Title    `json:"Title"`
		Poster    Url      `json:"Poster"`
		Excerpt   RichText `json:"Excerpt"`
	} `json:"properties"`
	URL string `json:"url"`
}
