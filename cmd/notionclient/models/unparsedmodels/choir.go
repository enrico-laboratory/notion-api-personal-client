package unparsedmodels

import "time"

type Choirs struct {
	Object         string  `json:"object"`
	Results        []Choir `json:"results"`
	NextCursor     string  `json:"next_cursor"`
	HasMore        bool    `json:"has_more"`
	Type           string  `json:"type"`
	PageOrDatabase struct {
	} `json:"page_or_database"`
	DeveloperSurvey string `json:"developer_survey"`
}

type Choir struct {
	Object         string    `json:"object"`
	ID             string    `json:"id"`
	CreatedTime    time.Time `json:"created_time"`
	LastEditedTime time.Time `json:"last_edited_time"`
	CreatedBy      struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		ID     string `json:"id"`
	} `json:"last_edited_by"`
	Cover  any `json:"cover"`
	Icon   any `json:"icon"`
	Parent struct {
		Type       string `json:"type"`
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Archived   bool `json:"archived"`
	Properties struct {
		Hired struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Checkbox bool   `json:"checkbox"`
		} `json:"Hired"`
		Projects struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Relation []struct {
				ID string `json:"id"`
			} `json:"relation"`
			HasMore bool `json:"has_more"`
		} `json:"Projects"`
		Choir struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			Title []struct {
				Type string `json:"type"`
				Text struct {
					Content string `json:"content"`
					Link    any    `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string `json:"plain_text"`
				Href      any    `json:"href"`
			} `json:"title"`
		} `json:"Choir"`
	} `json:"properties"`
	URL       string `json:"url"`
	PublicURL any    `json:"public_url"`
}
