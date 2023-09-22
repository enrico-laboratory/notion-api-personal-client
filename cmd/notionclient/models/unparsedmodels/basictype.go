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
	ID       string `json:"id"`
	Type     string `json:"type"`
	Relation []struct {
		ID string `json:"id"`
	} `json:"relation"`
}

type RichText struct {
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
}

type Select struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Select struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"select"`
}

type NumberFloat struct {
	ID     string  `json:"id"`
	Type   string  `json:"type"`
	Number float64 `json:"number"`
}

type Title struct {
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
	ID          string `json:"id"`
	Type        string `json:"type"`
	MultiSelect []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	} `json:"multi_select"`
}
type Url struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	URL  string `json:"url"`
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
