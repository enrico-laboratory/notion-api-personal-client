package unparsedmodels

import "time"

type ObjectAndId struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type Parent struct {
	Type       string `json:"type"`
	DatabaseID string `json:"database_id"`
}

type Relation struct {
	ID       string             `json:"id,omitempty"`
	Type     string             `json:"type,omitempty"`
	Relation []RelationProperty `json:"relation,omitempty"`
}

type RelationProperty struct {
	ID string `json:"id,omitempty"`
}

type RichText struct {
	ID       string             `json:"id,omitempty"`
	Type     string             `json:"type,omitempty"`
	RichText []RichTextProperty `json:"rich_text,omitempty"`
}

type RichTextProperty struct {
	Type string `json:"type,omitempty"`
	Text struct {
		Content string      `json:"content,omitempty"`
		Link    interface{} `json:"link,omitempty"`
	} `json:"text,omitempty"`
	Annotations struct {
		Bold          bool   `json:"bold,omitempty"`
		Italic        bool   `json:"italic,omitempty"`
		Strikethrough bool   `json:"strikethrough,omitempty"`
		Underline     bool   `json:"underline,omitempty"`
		Code          bool   `json:"code,omitempty"`
		Color         string `json:"color,omitempty"`
	} `json:"annotations,omitempty"`
	PlainText string      `json:"plain_text,omitempty"`
	Href      interface{} `json:"href,omitempty"`
}

type Select struct {
	ID     string `json:"id,omitempty"`
	Type   string `json:"type,omitempty"`
	Select struct {
		ID    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Color string `json:"color,omitempty"`
	} `json:"select,omitempty"`
}

type NumberFloat struct {
	ID     string  `json:"id,omitempty"`
	Type   string  `json:"type,omitempty"`
	Number float64 `json:"number,omitempty"`
}

type Title struct {
	ID    string          `json:"id,omitempty"`
	Type  string          `json:"type,omitempty"`
	Title []TitleProperty `json:"title,omitempty"`
}

type TitleProperty struct {
	Type string `json:"type,omitempty"`
	Text struct {
		Content string      `json:"content,omitempty"`
		Link    interface{} `json:"link,omitempty"`
	} `json:"text,omitempty"`
	Annotations struct {
		Bold          bool   `json:"bold,omitempty"`
		Italic        bool   `json:"italic,omitempty"`
		Strikethrough bool   `json:"strikethrough,omitempty"`
		Underline     bool   `json:"underline,omitempty"`
		Code          bool   `json:"code,omitempty"`
		Color         string `json:"color,omitempty"`
	} `json:"annotations,omitempty"`
	PlainText string      `json:"plain_text,omitempty"`
	Href      interface{} `json:"href,omitempty"`
}

type Email struct {
	Type  string `json:"type"`
	Email string `json:"email"`
}

type PhoneNumber struct {
	Type        string `json:"type"`
	PhoneNumber string `json:"phone_number"`
}

type Checkbox struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Checkbox bool   `json:"checkbox"`
}

type MultiSelect struct {
	ID          string                `json:"id,omitempty"`
	Type        string                `json:"type,omitempty"`
	MultiSelect []MultiSelectProperty `json:"multi_select,omitempty"`
}

type MultiSelectProperty struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}
type Url struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

type CreatedTime struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	CreatedTime time.Time `json:"created_time"`
}

type Date struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Date struct {
		Start    string `json:"start"`
		End      string `json:"end"`
		TimeZone string `json:"time_zone"`
	} `json:"date"`
}

type NumberInt struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Number int    `json:"number"`
}
