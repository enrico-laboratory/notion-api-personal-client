package unparsedmodels

import "time"

type Music struct {
	NextCursor     string        `json:"next_cursor"`
	Object         string        `json:"object"`
	Results        []SingleMusic `json:"results"`
	HasMore        bool          `json:"has_more"`
	Type           string        `json:"type"`
	PageOrDatabase struct {
	} `json:"page_or_database"`
	DeveloperSurvey string `json:"developer_survey"`
}

type SingleMusic struct {
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
		Voices              Select   `json:"Voices"`
		Score               Url      `json:"Score"`
		Media               Url      `json:"Media"`
		Recording           Url      `json:"Recording"`
		RepertoireOrder     Relation `json:"Repertoire Order"`
		Composer            RichText `json:"Composer"`
		MusicProjectsRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string     `json:"type"`
				Array    []Relation `json:"array"`
				Function string     `json:"function"`
			} `json:"rollup"`
		} `json:"Music Projects Rollup"`
		Length      NumberFloat `json:"Length"`
		Instruments MultiSelect `json:"Instruments"`
		Solo        Select      `json:"Solo"`
		Music       Title       `json:"Music"`
		URL         string      `json:"url"`
		PublicURL   any         `json:"public_url"`
	} `json:"properties"`
}
