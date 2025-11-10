package unparsedmodels

import "time"

type Cast struct {
	Object     string    `json:"object"`
	Results    []Contact `json:"results"`
	NextCursor string    `json:"next_cursor"`
	HasMore    bool      `json:"has_more"`
}

type Contact struct {
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
		SingerRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string  `json:"type"`
				Array    []Title `json:"array"`
				Function string  `json:"function"`
			} `json:"rollup"`
		} `json:"Singer Rollup"`
		Singer Relation `json:"Singer"`
		Note   RichText `json:"Note"`
		Phone  struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []PhoneNumber `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Phone"`
		Status Select      `json:"Status"`
		Order  NumberFloat `json:"Order"`
		Email  struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string  `json:"type"`
				Array    []Email `json:"array"`
				Function string  `json:"function"`
			} `json:"rollup"`
		} `json:"Email"`
		MusicProject Relation `json:"Music Project"`
		Role         Title    `json:"Role"`
	} `json:"properties"`
	URL string `json:"url"`
}
