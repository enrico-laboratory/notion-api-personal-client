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

type MusicProjectsService interface {
	query(body string) ([]parsedmodels.MusicProject, error)
	GetAll() ([]parsedmodels.MusicProject, error)
	GetById(projectId string) (parsedmodels.MusicProject, error)
	GetWithStatus(status string) ([]parsedmodels.MusicProject, error)
	GetWithStatusNot(status string) ([]parsedmodels.MusicProject, error)
	GetPublished() ([]parsedmodels.MusicProject, error)
	CreateProject(title, choirId, status string, year int) (int, error)
	GetByTitle(title string) (*parsedmodels.MusicProject, error)
}

type MusicProjectsClient struct {
	apiClient *NotionApiClient
	cfg       config
}

func (s *MusicProjectsClient) query(body string) ([]parsedmodels.MusicProject, error) {
	var scheduleParsed []parsedmodels.MusicProject
	var err error

	hasMore := true
	count := 0
	nextCursor := ""
	isBodyEmpty := body == ""

	for hasMore {

		var resp *http.Response
		var musicProjectsUnparsed unparsedmodels.MusicProjects

		if count == 0 {
			resp, err = s.apiClient.databaseQuery(s.cfg.databases.musicProjectsID, []byte(body))
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
			resp, err = s.apiClient.databaseQuery(s.cfg.databases.musicProjectsID, []byte(newBody))
			if err != nil {
				return nil, err
			}
		}

		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyResp, &musicProjectsUnparsed)

		for _, musicProject := range musicProjectsUnparsed.Results {
			var parsedMusicProject parsedmodels.MusicProject

			parseMusicProject(&musicProject, &parsedMusicProject)

			scheduleParsed = append(scheduleParsed, parsedMusicProject)
		}
		nextCursor = musicProjectsUnparsed.NextCursor
		hasMore = musicProjectsUnparsed.HasMore
		count++
	}

	return scheduleParsed, nil
}

func (s *MusicProjectsClient) GetAll() ([]parsedmodels.MusicProject, error) {
	query, err := s.query("")
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (s *MusicProjectsClient) GetById(projectId string) (parsedmodels.MusicProject, error) {
	query, err := s.query("")
	if err != nil {
		return parsedmodels.MusicProject{}, err
	}

	for _, project := range query {
		if project.Id == projectId {
			return project, nil
		}
	}

	return parsedmodels.MusicProject{}, errors.New(fmt.Sprintf("the project with id %v does not exist", projectId))
}

func (s *MusicProjectsClient) GetWithStatus(status string) ([]parsedmodels.MusicProject, error) {
	body := fmt.Sprintf(`{ 
				"filter": {
		              "property": "Status",
		              "select": {
		                  "equals": "%v"
		              }
				}
			}`, status)
	query, err := s.query(body)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (s *MusicProjectsClient) GetWithStatusNot(status string) ([]parsedmodels.MusicProject, error) {
	body := fmt.Sprintf(`{ 
				"filter": {
		              "property": "Status",
		              "select": {
		                  "does_not_equal": "%v"
		              }
				}
			}`, status)
	query, err := s.query(body)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (s *MusicProjectsClient) GetPublished() ([]parsedmodels.MusicProject, error) {
	body := fmt.Sprintf(`{ 
				"filter": {
		              "property": "Published",
		              "checkbox": {
		                  "equals": true
		              }
				}
			}`)
	query, err := s.query(body)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (s *MusicProjectsClient) GetByTitle(title string) (*parsedmodels.MusicProject, error) {
	body := fmt.Sprintf(`{ 
				"filter": {
		              "property": "Title",
		              "title": {
		                  "equals": "%v"
		              }
				}
			}`, title)
	query, err := s.query(body)
	if err != nil {
		return nil, err
	}
	if len(query) == 0 {
		return nil, errors.New(fmt.Sprintf("choir with name %v does not exist", title))
	}
	if len(query) > 1 {
		return nil, errors.New(fmt.Sprintf("found multiple choirs with name %v", title))
	}

	return &query[0], nil
}

func parseMusicProject(u *unparsedmodels.MusicProject, p *parsedmodels.MusicProject) {

	var description string

	if len(u.Properties.Description.RichText) == 0 {
		description = ""
	} else {
		description = u.Properties.Description.RichText[0].PlainText
	}

	var excerpt string
	if len(u.Properties.Excerpt.RichText) == 0 {
		excerpt = ""
	} else {
		excerpt = u.Properties.Excerpt.RichText[0].PlainText
	}
	var choirRollup string

	if len(u.Properties.ChoirRollup.Rollup.Array) == 0 {
		choirRollup = ""
	} else {
		choirRollup = u.Properties.ChoirRollup.Rollup.Array[0].Title[0].PlainText
	}
	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Title = u.Properties.Title.Title[0].PlainText
	p.Year = u.Properties.Year.Number
	p.Status = u.Properties.Status.Select.Name
	p.Description = description
	p.ChoirRollup = choirRollup
	p.Excerpt = excerpt
	p.Poster = u.Properties.Poster.URL
}

type createProjectRequest struct {
	Parent struct {
		DatabaseId string `json:"database_id"`
	} `json:"parent"`
	*CreateProjectRequestProperties `json:"properties"`
}

type CreateProjectRequestProperties struct {
	Title struct {
		Title []*Title `json:"title"`
	} `json:"title"`
	Year struct {
		Number int `json:"number"`
	} `json:"Year"`
	Choir struct {
		Relation []*Relation `json:"relation"`
	} `json:"Choir"`
	Status struct {
		Select struct {
			Name string `json:"name"`
		} `json:"select"`
	} `json:"Status"`
}

type Title struct {
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

type Relation struct {
	ID string `json:"id"`
}

func (s *MusicProjectsClient) CreateProject(title, choirId, status string, year int) (int, error) {
	req := &createProjectRequest{}
	req.Parent.DatabaseId = musicProjectDatabaseId

	properties := &CreateProjectRequestProperties{}
	properties.Year.Number = year

	titleObj := &Title{}
	titleObj.Text.Content = title
	var titleList []*Title
	titleList = append(titleList, titleObj)
	properties.Title.Title = titleList

	choirRelationObj := &Relation{}
	choirRelationObj.ID = choirId
	var choirRelationList []*Relation
	choirRelationList = append(choirRelationList, choirRelationObj)
	properties.Choir.Relation = choirRelationList

	properties.Status.Select.Name = status

	req.CreateProjectRequestProperties = properties

	body, err := json.Marshal(&req)
	if err != nil {
		return 0, err
	}

	resp, err := s.apiClient.pages([]byte(body))
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
