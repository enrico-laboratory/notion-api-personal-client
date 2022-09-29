package unparsedmodels

import "time"

type Cast struct {
	Object     string    `json:"object"`
	Results    []Contact `json:"results"`
	NextCursor string    `json:"next_cursor"`
	HasMore    bool      `json:"has_more"`
}

type Contact struct {
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
	Cover  interface{} `json:"cover"`
	Icon   interface{} `json:"icon"`
	Parent struct {
		Type       string `json:"type"`
		DatabaseID string `json:"database_id"`
	} `json:"parent"`
	Archived   bool `json:"archived"`
	Properties struct {
		SingerRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type  string `json:"type"`
				Array []struct {
					Type  string `json:"type"`
					Title []struct {
						Type string `json:"type"`
						Text struct {
							Content string      `json:"content"`
							Link    interface{} `json:"link"`
						} `json:"text"`
						Annotations struct {
							Bold          bool   `json:"bold"`
							Italic        bool   `json:"italic"`
							Strikethrough bool   `json:"strikethrough"`
							Underline     bool   `json:"underline"`
							Code          bool   `json:"code"`
							Color         string `json:"color"`
						} `json:"annotations"`
						PlainText string      `json:"plain_text"`
						Href      interface{} `json:"href"`
					} `json:"title"`
				} `json:"array"`
				Function string `json:"function"`
			} `json:"rollup"`
		} `json:"Singer Rollup"`
		Singer struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Relation []struct {
				ID string `json:"id"`
			} `json:"relation"`
		} `json:"Singer"`
		Note struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			RichText []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"rich_text"`
		} `json:"Note"`
		Phone struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type  string `json:"type"`
				Array []struct {
					Type        string `json:"type"`
					PhoneNumber string `json:"phone_number"`
				} `json:"array"`
				Function string `json:"function"`
			} `json:"rollup"`
		} `json:"Phone"`
		Status struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Status"`
		Order struct {
			ID     string  `json:"id"`
			Type   string  `json:"type"`
			Number float64 `json:"number"`
		} `json:"Order"`
		Email struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type  string `json:"type"`
				Array []struct {
					Type  string `json:"type"`
					Email string `json:"email"`
				} `json:"array"`
				Function string `json:"function"`
			} `json:"rollup"`
		} `json:"Email"`
		MusicProject struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Relation []struct {
				ID string `json:"id"`
			} `json:"relation"`
		} `json:"Music Project"`
		Role struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			Title []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"title"`
		} `json:"Role"`
	} `json:"properties"`
	URL string `json:"url"`
}
