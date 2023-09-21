package notionclient

import (
	"encoding/json"
	"fmt"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/parsedmodels"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/unparsedmodels"
	"io"
	"net/http"
)

type LocationsService interface {
	Query(body string) ([]parsedmodels.Location, error)
}

type LocationsClient struct {
	apiClient *NotionApiClient
	cfg       config
}

func (s *LocationsClient) Query(body string) ([]parsedmodels.Location, error) {
	var locationsParsed []parsedmodels.Location
	var err error

	hasMore := true
	count := 0
	nextCursor := ""
	isBodyEmpty := body == ""

	for hasMore {

		var resp *http.Response
		var locationsUnparsed unparsedmodels.Locations

		if count == 0 {
			resp, err = s.apiClient.databaseQuery(s.cfg.databases.locationID, []byte(body))
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
			resp, err = s.apiClient.databaseQuery(s.cfg.databases.locationID, []byte(newBody))
			if err != nil {
				return nil, err
			}
		}

		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyResp, &locationsUnparsed)

		for _, location := range locationsUnparsed.Results {
			var parsedLocation parsedmodels.Location

			parseLocation(&location, &parsedLocation)

			locationsParsed = append(locationsParsed, parsedLocation)
		}
		nextCursor = locationsUnparsed.NextCursor
		hasMore = locationsUnparsed.HasMore
		count++
	}

	return locationsParsed, nil
}

func parseLocation(u *unparsedmodels.Location, p *parsedmodels.Location) {
	var contact string

	if len(u.Properties.Contact.Relation) == 0 {
		contact = ""
	} else {
		contact = u.Properties.Contact.Relation[0].ID
	}

	var phone string

	if len(u.Properties.Phone.Rollup.Array) == 0 {
		phone = ""
	} else {
		phone = u.Properties.Phone.Rollup.Array[0].PhoneNumber
	}

	var city string

	if len(u.Properties.City.RichText) == 0 {
		city = ""
	} else {
		city = u.Properties.City.RichText[0].PlainText
	}

	var email string

	if len(u.Properties.Email.Rollup.Array) == 0 {
		email = ""
	} else {
		email = u.Properties.Email.Rollup.Array[0].Email
	}

	var tasks []string

	if len(u.Properties.Tasks.Relation) == 0 {
		tasks = []string{}
	} else {
		for _, record := range u.Properties.Tasks.Relation {
			tasks = append(tasks, record.ID)
		}
	}

	var purpose []string

	if len(u.Properties.Purpose.MultiSelect) == 0 {
		purpose = []string{}
	} else {
		for _, record := range u.Properties.Purpose.MultiSelect {
			purpose = append(tasks, record.Name)
		}
	}

	var address string

	if len(u.Properties.Address.RichText) == 0 {
		address = ""
	} else {
		address = u.Properties.Address.RichText[0].PlainText
	}

	var location string

	if len(u.Properties.Location.Title) == 0 {
		location = ""
	} else {
		location = u.Properties.Location.Title[0].PlainText
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Contact = contact
	p.Phone = phone
	p.City = city
	p.Email = email
	p.Task = tasks
	p.Purpose = purpose
	p.Address = address
	p.Location = location

}
