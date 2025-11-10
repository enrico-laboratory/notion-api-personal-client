package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/enrico-laboratory/notion-api-personal-client/client/models/parsedmodels"
	"github.com/enrico-laboratory/notion-api-personal-client/client/models/unparsedmodels"
	"io"
	"net/http"
)

type ChoirService interface {
	query(body string) ([]parsedmodels.Choir, error)
	GetByName(name string) (*parsedmodels.Choir, error)
	GetAll() ([]parsedmodels.Choir, error)
}

type ChoirClient struct {
	apiClient *NotionApiClient
	cfg       config
}

func (s *ChoirClient) GetAll() ([]parsedmodels.Choir, error) {
	query, err := s.query("")
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (s *ChoirClient) GetByName(name string) (*parsedmodels.Choir, error) {
	body := fmt.Sprintf(`{ 
				"filter": {
		              "property": "Choir",
		              "title": {
		                  "equals": "%v"
		              }
				}
			}`, name)
	query, err := s.query(body)
	if err != nil {
		return nil, err
	}
	if len(query) == 0 {
		return nil, errors.New(fmt.Sprintf("choir with name %v does not exist", name))
	}
	if len(query) > 1 {
		return nil, errors.New(fmt.Sprintf("found multiple choirs with name %v", name))
	}

	return &query[0], nil
}

func (s *ChoirClient) query(body string) ([]parsedmodels.Choir, error) {
	var choirParsed []parsedmodels.Choir
	var err error

	hasMore := true
	count := 0
	nextCursor := ""
	isBodyEmpty := body == ""

	for hasMore {

		var resp *http.Response
		var ChoirUnparsed unparsedmodels.Choirs

		if count == 0 {
			resp, err = s.apiClient.databaseQuery(s.cfg.databases.ChoirID, []byte(body))
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
			resp, err = s.apiClient.databaseQuery(s.cfg.databases.ChoirID, []byte(newBody))
			if err != nil {
				return nil, err
			}
		}

		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyResp, &ChoirUnparsed)

		for _, c := range ChoirUnparsed.Results {
			var parsedChoir parsedmodels.Choir

			parseChoir(&c, &parsedChoir)

			choirParsed = append(choirParsed, parsedChoir)
		}
		nextCursor = ChoirUnparsed.NextCursor
		hasMore = ChoirUnparsed.HasMore
		count++
	}

	return choirParsed, nil
}

func parseChoir(u *unparsedmodels.Choir, p *parsedmodels.Choir) {
	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Name = u.Properties.Choir.Title[0].PlainText
}
