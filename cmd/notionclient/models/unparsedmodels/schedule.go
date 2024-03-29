package unparsedmodels

import "time"

type Schedule struct {
	Object     string `json:"object"`
	Results    []Task `json:"results"`
	NextCursor string `json:"next_cursor"`
	HasMore    bool   `json:"has_more"`
}

type Task struct {
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
		MusicProject Relation `json:"Music Project"`
		Type         Select   `json:"Type"`
		GoalOutcome  struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []interface{} `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Goal Outcome"`
		Post               Relation `json:"Post"`
		LocationCityRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string     `json:"type"`
				Array    []RichText `json:"array"`
				Function string     `json:"function"`
			} `json:"rollup"`
		} `json:"Location City Rollup"`
		Priority              Select `json:"Priority"`
		PillarsRoutineAndMain struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []interface{} `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Pillars (Routine and Main.)"`
		DoDate   Date `json:"Do Date"`
		Duration struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Formula struct {
				Type   string `json:"type"`
				String string `json:"string"`
			} `json:"formula"`
		} `json:"Duration"`
		URL                 Url      `json:"url"`
		RoutinesMaintenance Relation `json:"Routines & Maintenance"`
		Kanban              Select   `json:"Kanban"`
		LocationRollup      struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string  `json:"type"`
				Array    []Title `json:"array"`
				Function string  `json:"function"`
			} `json:"rollup"`
		} `json:"Location Rollup"`
		DueDate                 Date     `json:"Due Date"`
		Waiting                 Checkbox `json:"Waiting"`
		Ante                    Relation `json:"Ante"`
		Done                    Checkbox `json:"Done"`
		Location                Relation `json:"Location"`
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
				Type     string     `json:"type"`
				Array    []Relation `json:"array"`
				Function string     `json:"function"`
			} `json:"rollup"`
		} `json:"Choir Rollup"`
		Projects Relation `json:"Projects"`
		Task     Title    `json:"Task"`
		Notes    RichText `json:"Notes"`
	} `json:"properties"`
	URL string `json:"url"`
}

type AutoGenerated struct {
}
