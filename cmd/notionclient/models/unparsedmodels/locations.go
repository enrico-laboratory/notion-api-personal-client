package unparsedmodels

import "time"

type Locations struct {
	Object     string     `json:"object"`
	Results    []Location `json:"results"`
	NextCursor string     `json:"next_cursor"`
	HasMore    bool       `json:"has_more"`
}

type Location struct {
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
		Contact Relation `json:"Contact"`
		Phone   struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []PhoneNumber `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Phone"`
		City  RichText `json:"City"`
		Email struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string  `json:"type"`
				Array    []Email `json:"array"`
				Function string  `json:"function"`
			} `json:"rollup"`
		} `json:"Email"`
		Tasks    Relation    `json:"Tasks"`
		Purpose  MultiSelect `json:"Purpose"`
		Address  RichText    `json:"Address"`
		Location Title       `json:"Location"`
	} `json:"properties"`
	URL string `json:"url"`
}
