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
