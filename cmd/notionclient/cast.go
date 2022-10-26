package notionclient

import (
	"encoding/json"
	"fmt"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/parsedmodels"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/unparsedmodels"
	"io"
	"net/http"
)

type CastService interface {
	Query(body string) ([]parsedmodels.Contact, error)
	GetAll() ([]parsedmodels.Contact, error)
	GetByProjectId(projectId string) ([]parsedmodels.Contact, error)
}

type CastClient struct {
	apiClient *NotionApiClient
	cfg       config
}

func (s *CastClient) Query(body string) ([]parsedmodels.Contact, error) {
	var castParsed []parsedmodels.Contact
	var err error

	hasMore := true
	count := 0
	nextCursor := ""
	isBodyEmpty := body == ""

	for hasMore {

		var resp *http.Response
		var castUnparsed unparsedmodels.Cast

		if count == 0 {
			resp, err = s.apiClient.request(s.cfg.databases.castID, []byte(body))
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
			resp, err = s.apiClient.request(s.cfg.databases.castID, []byte(newBody))
			if err != nil {
				return nil, err
			}
		}

		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyResp, &castUnparsed)

		for _, contact := range castUnparsed.Results {
			var parsedContact parsedmodels.Contact

			parseContact(&contact, &parsedContact)

			castParsed = append(castParsed, parsedContact)
		}
		nextCursor = castUnparsed.NextCursor
		hasMore = castUnparsed.HasMore
		count++
	}

	return castParsed, nil
}

func (s *CastClient) GetAll() ([]parsedmodels.Contact, error) {
	query, err := s.Query("")
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (s *CastClient) GetByProjectId(projectId string) ([]parsedmodels.Contact, error) {
	query, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	var result []parsedmodels.Contact

	for _, contact := range query {
		for _, projectIdCast := range contact.MusicProject {
			if projectIdCast == projectId {
				result = append(result, contact)
			}
		}
	}

	return result, nil
}

func parseContact(u *unparsedmodels.Contact, p *parsedmodels.Contact) {
	var singer string

	if len(u.Properties.SingerRollup.Rollup.Array) == 0 {
		singer = ""
	} else if len(u.Properties.SingerRollup.Rollup.Array[0].Title) == 0 {
		singer = ""
	} else {
		singer = u.Properties.SingerRollup.Rollup.Array[0].Title[0].PlainText
	}

	var singerId string

	if len(u.Properties.Singer.Relation) == 0 {
		singerId = ""
	} else {
		singerId = u.Properties.Singer.Relation[0].ID
	}

	var note string

	if len(u.Properties.Note.RichText) == 0 {
		note = ""
	} else {
		note = u.Properties.Note.RichText[0].PlainText
	}

	var phone string

	if len(u.Properties.Phone.Rollup.Array) == 0 {
		phone = ""
	} else {
		phone = u.Properties.Phone.Rollup.Array[0].PhoneNumber
	}

	var status string

	status = u.Properties.Status.Select.Name

	var order float64
	order = u.Properties.Order.Number

	var email string

	if len(u.Properties.Email.Rollup.Array) == 0 {
		email = ""
	} else {
		email = u.Properties.Email.Rollup.Array[0].Email
	}

	var musicProject []string

	if len(u.Properties.MusicProject.Relation) == 0 {
		musicProject = []string{}
	} else {
		for _, record := range u.Properties.MusicProject.Relation {
			musicProject = append(musicProject, record.ID)
		}
	}

	var role string

	if len(u.Properties.Role.Title) == 0 {
		role = ""
	} else {
		role = u.Properties.Role.Title[0].PlainText
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Singer = singer
	p.SingerId = singerId
	p.Note = note
	p.Phone = phone
	p.Status = status
	p.Order = order
	p.Email = email
	p.MusicProject = musicProject
	p.Role = role
}
