package notionclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/parsedmodels"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/unparsedmodels"
	"io"
	"net/http"
)

type MusicService interface {
	query(body string) ([]parsedmodels.Music, error)
	GetByTile(title string) (*parsedmodels.Music, error)
	GetByTileAndComposer(title, composer string) (*parsedmodels.Music, error)
	CreateMusic(properties *CreateMusicRequestProperties) (string, error)
	DeleteMusicById(musicId string) error
}

type MusicClient struct {
	apiClient *NotionApiClient
	cfg       config
}

func (m *MusicClient) query(body string) ([]parsedmodels.Music, error) {
	var musicParsed []parsedmodels.Music
	var err error

	hasMore := true
	count := 0
	nextCursor := ""
	isBodyEmpty := body == ""

	for hasMore {

		var resp *http.Response
		var musicUnparsed unparsedmodels.Music

		if count == 0 {
			resp, err = m.apiClient.databaseQuery(m.cfg.databases.MusicID, []byte(body))
			if err != nil {
				return nil, err
			}
		} else {
			startCursor := fmt.Sprintf(`"start_cursor": "%v"`, nextCursor)
			var newBody string
			if isBodyEmpty {
				newBody = fmt.Sprintf("{%v}", startCursor)
			} else {
				newBody = fmt.Sprintf("%v%v,%v", body[:1], startCursor, body[1:])
			}
			resp, err = m.apiClient.databaseQuery(m.cfg.databases.repertoireID, []byte(newBody))
			if err != nil {
				return nil, err
			}
		}

		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyResp, &musicUnparsed)

		for _, piece := range musicUnparsed.Results {
			var parsedMusic parsedmodels.Music

			parseMusic(&piece, &parsedMusic)

			musicParsed = append(musicParsed, parsedMusic)
		}
		nextCursor = musicUnparsed.NextCursor
		hasMore = musicUnparsed.HasMore
		count++
	}

	return musicParsed, nil
}

func (m *MusicClient) GetByTile(title string) (*parsedmodels.Music, error) {
	body := fmt.Sprintf(`{
    "filter": {
        "property": "Music",
        "title": {
            "equals": "%v"
        	}
    	}
	}`, title)
	query, err := m.query(body)
	if err != nil {
		return nil, err
	}
	if len(query) == 0 {
		return nil, errors.New(fmt.Sprintf("music with name %v does not exist", title))
	}
	if len(query) > 1 {
		return nil, errors.New(fmt.Sprintf("found multiple music with name %v", title))
	}

	return &query[0], nil
}

func (m *MusicClient) GetByTileAndComposer(title, composer string) (*parsedmodels.Music, error) {
	if title == "" || composer == "" {
		return nil, errors.New("tile and composer must be specified")
	}
	body := fmt.Sprintf(`{
"filter": {
  "and": [
    {
      "property": "Music",
      "title": {
        "equals": "%v"
      }
    },
    {
      "property": "Composer",
      "title": {
        "equals": "%v"
      }
    }
  ]
}
}`, title, composer)
	query, err := m.query(body)
	if err != nil {
		return nil, err
	}
	if len(query) == 0 {
		return nil, errors.New(fmt.Sprintf("music with name '%v' and composer '%v' does not exist", title, composer))
	}
	if len(query) > 1 {
		return nil, errors.New(fmt.Sprintf("found multiple music with name '%v' and composer '%v'", title, composer))
	}

	return &query[0], nil
}

func (m *MusicClient) DeleteMusicById(musicId string) error {
	_, err := m.apiClient.pagesDelete(musicId)
	if err != nil {
		return err
	}
	return nil
}

type CreateMusicRequestProperties struct {
	Title       string
	Voices      string
	Score       string
	Media       string
	Recording   string
	Composer    string
	Length      float64
	Instruments []string
	Solo        string
}

func (m *MusicClient) CreateMusic(properties *CreateMusicRequestProperties) (string, error) {
	if properties.Title == "" || properties.Composer == "" {
		return "", errors.New("one of Tile or Composer is missing")
	}
	type createMusicRequest struct {
		Parent struct {
			DatabaseId string `json:"database_id"`
		} `json:"parent"`
		//Properties map[string]any `json:"properties"`
		Properties struct {
			Voices      map[string]unparsedmodels.SelectProperty        `json:"Voices,omitempty"`
			Score       map[string]string                               `json:"Score,omitempty"`
			Media       map[string]string                               `json:"Media,omitempty"`
			Recording   map[string]string                               `json:"Recording,omitempty"`
			Composer    unparsedmodels.RichText                         `json:"Composer"`
			Length      map[string]float64                              `json:"Length,omitempty"`
			Instruments map[string][]unparsedmodels.MultiSelectProperty `json:"Instruments,omitempty"`
			Solo        map[string]unparsedmodels.SelectProperty        `json:"Solo,omitempty"`
			Title       unparsedmodels.Title                            `json:"Music"`
		} `json:"properties"`
	}

	req := &createMusicRequest{}
	req.Parent.DatabaseId = musicDatabaseId

	var titleProperty unparsedmodels.TitleProperty
	titleProperty.Text.Content = properties.Title

	req.Properties.Title.Title = []unparsedmodels.TitleProperty{titleProperty}

	if properties.Voices != "" {
		voices := make(map[string]unparsedmodels.SelectProperty)
		voicesProperties := unparsedmodels.SelectProperty{}
		voicesProperties.Name = properties.Voices
		voices["select"] = voicesProperties
		req.Properties.Voices = voices
	}

	if properties.Score != "" {
		score := make(map[string]string)
		score["url"] = properties.Score
		req.Properties.Score = score
	}
	if properties.Media != "" {
		media := make(map[string]string)
		media["url"] = properties.Media
		req.Properties.Media = media
	}
	if properties.Recording != "" {
		recording := make(map[string]string)
		recording["url"] = properties.Recording
		req.Properties.Recording = recording
	}
	if properties.Solo != "" {
		solo := make(map[string]unparsedmodels.SelectProperty)
		soloProperties := unparsedmodels.SelectProperty{}
		soloProperties.Name = properties.Solo
		solo["select"] = soloProperties
		req.Properties.Solo = solo
	}

	richTextProperty := unparsedmodels.RichTextProperty{}
	richTextProperty.Text.Content = properties.Composer
	req.Properties.Composer.RichText = []unparsedmodels.RichTextProperty{richTextProperty}

	if properties.Length != 0 {
		length := make(map[string]float64)
		length["number"] = properties.Length
		req.Properties.Length = length
	}

	if len(properties.Instruments) > 0 {
		var multiSelectPropertyList []unparsedmodels.MultiSelectProperty
		for _, i := range properties.Instruments {
			multiSelectProperty := unparsedmodels.MultiSelectProperty{Name: i}
			multiSelectPropertyList = append(multiSelectPropertyList, multiSelectProperty)
		}
		instruments := make(map[string][]unparsedmodels.MultiSelectProperty)
		instruments["multi_select"] = multiSelectPropertyList
		req.Properties.Instruments = instruments
	}

	body, err := json.Marshal(&req)
	if err != nil {
		return "", err
	}

	resp, err := m.apiClient.pages(body)
	if err != nil {
		return "", err
	}

	type ResponseID struct {
		ID string `json:"id"`
	}

	var id ResponseID
	err = json.NewDecoder(resp.Body).Decode(&id)
	if err != nil {
		return "", err
	}
	musicId := id.ID
	return musicId, nil
}

func parseMusic(u *unparsedmodels.SingleMusic, p *parsedmodels.Music) {

	var composer string

	if len(u.Properties.Composer.RichText) == 0 {
		composer = ""
	} else {
		composer = u.Properties.Composer.RichText[0].PlainText
	}

	var instruments []string

	if len(u.Properties.Instruments.MultiSelect) == 0 {
		instruments = append(instruments, "")
	} else {
		for _, instrument := range u.Properties.Instruments.MultiSelect {
			instruments = append(instruments, instrument.Name)
		}
	}

	var title string

	if len(u.Properties.Music.Title) == 0 {
		title = ""
	} else {
		title = u.Properties.Music.Title[0].PlainText
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Voices = u.Properties.Voices.Select.Name
	p.Score = u.Properties.Score.URL
	p.Media = u.Properties.Media.URL
	p.Recording = u.Properties.Recording.URL
	p.Composer = composer
	p.Length = u.Properties.Length.Number
	p.Instruments = instruments
	p.Solo = u.Properties.Solo.Select.Name
	p.Title = title
}
