package unparsedmodels

import "time"

type Repertoire struct {
	Object     string  `json:"object"`
	Results    []Piece `json:"results"`
	NextCursor string  `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
}

type Piece struct {
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
		Num2         RichText `json:"2"`
		Num3         RichText `json:"3"`
		Num4         RichText `json:"4"`
		Num5         RichText `json:"5"`
		Num6         RichText `json:"6"`
		Num7         RichText `json:"7"`
		Num8         RichText `json:"8"`
		Num9         RichText `json:"9"`
		Num10        RichText `json:"10"`
		Num11        RichText `json:"11"`
		Num12        RichText `json:"12"`
		Music        Relation `json:"Music"`
		Note         RichText `json:"Note"`
		VoicesRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string   `json:"type"`
				Array    []Select `json:"array"`
				Function string   `json:"function"`
			} `json:"rollup"`
		} `json:"Voices Rollup"`
		OneTopVoice RichText `json:"1 - Top Voice"`
		MusicRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string  `json:"type"`
				Array    []Title `json:"array"`
				Function string  `json:"function"`
			} `json:"rollup"`
		} `json:"Music Rollup"`
		Selected     Checkbox `json:"Selected"`
		MusicProject Relation `json:"Music Project"`
		MediaRollup  struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string `json:"type"`
				Array    []Url  `json:"array"`
				Function string `json:"function"`
			} `json:"rollup"`
		} `json:"Media Rollup"`
		ScoreRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string `json:"type"`
				Array    []Url  `json:"array"`
				Function string `json:"function"`
			} `json:"rollup"`
		} `json:"Score Rollup"`
		InstrumentsRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []MultiSelect `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Instruments Rollup"`
		LengthRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string        `json:"type"`
				Array    []NumberFloat `json:"array"`
				Function string        `json:"function"`
			} `json:"rollup"`
		} `json:"Length Rollup"`
		ComposerRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string     `json:"type"`
				Array    []RichText `json:"array"`
				Function string     `json:"function"`
			} `json:"rollup"`
		} `json:"Composer Rollup"`
		Length struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Formula struct {
				Type   string  `json:"type"`
				Number float64 `json:"number"`
			} `json:"formula"`
		} `json:"Length"`
		RecordingRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string `json:"type"`
				Array    []Url  `json:"array"`
				Function string `json:"function"`
			} `json:"rollup"`
		} `json:"Recording Rollup"`
		SoloRollup struct {
			ID     string `json:"id"`
			Type   string `json:"type"`
			Rollup struct {
				Type     string   `json:"type"`
				Array    []Select `json:"array"`
				Function string   `json:"function"`
			} `json:"rollup"`
		} `json:"Solo Rollup"`
		Order           Title    `json:"Order"`
		NotesDivisi     RichText `json:"Notes Divisi"`
		NotesRepertoire RichText `json:"Notes Repertoire"`
	} `json:"properties"`
	URL string `json:"url"`
}
