package unparsedmodels

import "time"

type Date struct {
	Start    string      `json:"start"`
	End      string      `json:"end"`
	TimeZone interface{} `json:"time_zone"`
}

type Select struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Schedule struct {
	Object     string `json:"object"`
	Results    []Task `json:"results"`
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
}

type Task struct {
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
		MusicProject struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Relation []struct {
				ID string `json:"id"`
			} `json:"relation"`
		} `json:"Music Project"`
		Type struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Type"`
		GoalOutcome struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []interface{} `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Goal Outcome"`
		Post struct {
			ID       string        `json:"id"`
			Type     string        `json:"type"`
			Relation []interface{} `json:"relation"`
		} `json:"Post"`
		LocationCityRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type  string `json:"type"`
				Array []struct {
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
				} `json:"array"`
				Function string `json:"function"`
			} `json:"rollup"`
		} `json:"Location City Rollup"`
		Priority struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Priority"`
		PillarsRoutineAndMain struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []interface{} `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Pillars (Routine and Main.)"`
		DoDate struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Date Date   `json:"date"`
		} `json:"Do Date"`
		Duration struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Formula struct {
				Type   string `json:"type"`
				String string `json:"string"`
			} `json:"formula"`
		} `json:"Duration"`
		URL struct {
			ID   string      `json:"id"`
			Type string      `json:"type"`
			URL  interface{} `json:"url"`
		} `json:"url"`
		RoutinesMaintenance struct {
			ID       string        `json:"id"`
			Type     string        `json:"type"`
			Relation []interface{} `json:"relation"`
		} `json:"Routines & Maintenance"`
		Kanban struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Select Select `json:"select"`
		} `json:"Kanban"`
		LocationRollup struct {
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
		} `json:"Location Rollup"`
		DueDate struct {
			ID   string      `json:"id"`
			Type string      `json:"type"`
			Date interface{} `json:"date"`
		} `json:"Due Date"`
		Waiting struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Checkbox bool   `json:"checkbox"`
		} `json:"Waiting"`
		Ante struct {
			ID       string        `json:"id"`
			Type     string        `json:"type"`
			Relation []interface{} `json:"relation"`
		} `json:"Ante"`
		Done struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Checkbox bool   `json:"checkbox"`
		} `json:"Done"`
		Location struct {
			ID       string `json:"id"`
			Type     string `json:"type"`
			Relation []struct {
				ID string `json:"id"`
			} `json:"relation"`
		} `json:"Location"`
		PillarsGroupingRoutMain struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []interface{} `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Pillars Grouping (Rout. & Main.)"`
		ChoirRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type  string `json:"type"`
				Array []struct {
					Type     string        `json:"type"`
					Relation []interface{} `json:"relation"`
				} `json:"array"`
				Function string `json:"function"`
			} `json:"rollup"`
		} `json:"Choir Rollup"`
		Projects struct {
			ID       string        `json:"id"`
			Type     string        `json:"type"`
			Relation []interface{} `json:"relation"`
		} `json:"Projects"`
		Task struct {
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
		} `json:"Task"`
		Notes struct {
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
		} `json:"Notes"`
	} `json:"properties"`
	URL string `json:"url"`
}

type AutoGenerated struct {
}
