package notionclient

import (
	"encoding/json"
	"fmt"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/parsedmodels"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/unparsedmodels"
	"io"
	"net/http"
)

type ScheduleService interface {
	Query(body string) ([]parsedmodels.Task, error)
	GetAll() ([]parsedmodels.Task, error)
	GetByProjectId(projectId string) ([]parsedmodels.Task, error)
	GetByProjectIdAndType(projectId string, t ...string) ([]parsedmodels.Task, error)
}

type ScheduleClient struct {
	apiClient *NotionApiClient
	cfg       config
}

func (s *ScheduleClient) Query(body string) ([]parsedmodels.Task, error) {
	var scheduleParsed []parsedmodels.Task
	var err error

	hasMore := true
	count := 0
	nextCursor := ""
	isBodyEmpty := body == ""

	for hasMore {

		var resp *http.Response
		var scheduleUnparsed unparsedmodels.Schedule

		if count == 0 {
			resp, err = s.apiClient.request(s.cfg.databases.scheduleID, []byte(body))
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
			resp, err = s.apiClient.request(s.cfg.databases.scheduleID, []byte(newBody))
			if err != nil {
				return nil, err
			}
		}

		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyResp, &scheduleUnparsed)

		for _, task := range scheduleUnparsed.Results {
			var parsedTask parsedmodels.Task

			parseTask(&task, &parsedTask)

			scheduleParsed = append(scheduleParsed, parsedTask)
		}
		nextCursor = scheduleUnparsed.NextCursor
		hasMore = scheduleUnparsed.HasMore
		count++
	}

	return scheduleParsed, nil
}

func (s *ScheduleClient) GetAll() ([]parsedmodels.Task, error) {
	query, err := s.Query("")
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (s *ScheduleClient) GetByProjectId(projectId string) ([]parsedmodels.Task, error) {
	query, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	var result []parsedmodels.Task

	for _, task := range query {
		for _, projectIdTask := range task.MusicProject {
			if projectIdTask == projectId {
				result = append(result, task)
			}
		}
	}

	return result, nil
}

func (s *ScheduleClient) GetByProjectIdAndType(projectId string, t ...string) ([]parsedmodels.Task, error) {

	if len(t) == 0 {
		query, err := s.GetAll()
		if err != nil {
			return nil, err
		}

		var result []parsedmodels.Task

		for _, task := range query {
			for _, projectIdTask := range task.MusicProject {
				if projectIdTask == projectId {
					result = append(result, task)
				}
			}
		}

		return result, nil
	}

	var or []Or
	for _, value := range t {
		o := Or{
			Property: "Type",
			Select: struct {
				Equals string `json:"equals"`
			}{Equals: value},
		}
		or = append(or, o)
	}
	var filterOrTypeObject FilterOrSelect

	filterOrTypeObject.Filter.Or = or

	by, err := json.Marshal(&filterOrTypeObject)
	if err != nil {
		return nil, err
	}

	query, err := s.Query(string(by))
	if err != nil {
		return nil, err
	}

	var result []parsedmodels.Task

	for _, task := range query {
		for _, projectIdTask := range task.MusicProject {
			if projectIdTask == projectId {
				result = append(result, task)
			}
		}
	}

	return result, nil
}

func parseTask(u *unparsedmodels.Task, p *parsedmodels.Task) {
	var musicProject []string

	if len(u.Properties.MusicProject.Relation) == 0 {
		musicProject = []string{}
	} else {
		for _, record := range u.Properties.MusicProject.Relation {
			musicProject = append(musicProject, record.ID)
		}
	}

	var _type string

	_type = u.Properties.Type.Select.Name

	var locationRollup string

	if len(u.Properties.LocationRollup.Rollup.Array) == 0 {
		locationRollup = ""
	} else if len(u.Properties.LocationRollup.Rollup.Array[0].Title) == 0 {
		locationRollup = ""
	} else {
		locationRollup = u.Properties.LocationRollup.Rollup.Array[0].Title[0].PlainText
	}

	var cityRollup string

	if len(u.Properties.LocationCityRollup.Rollup.Array) == 0 {
		cityRollup = ""
	} else if len(u.Properties.LocationCityRollup.Rollup.Array[0].RichText) == 0 {
		cityRollup = ""
	} else {
		cityRollup = u.Properties.LocationCityRollup.Rollup.Array[0].RichText[0].PlainText
	}

	var priority string

	priority = u.Properties.Priority.Select.Name

	var startDateAndTime string
	var endDateAndTime string
	emptyTaskDate := unparsedmodels.Date{}
	emptyDate := ""

	if u.Properties.DoDate.Date == emptyTaskDate {
		startDateAndTime = emptyDate
		endDateAndTime = emptyDate
	} else if u.Properties.DoDate.Date.Start == emptyDate &&
		u.Properties.DoDate.Date.End != emptyDate {
		startDateAndTime = emptyDate
		endDateAndTime = u.Properties.DoDate.Date.Start
	} else if u.Properties.DoDate.Date.Start != emptyDate &&
		u.Properties.DoDate.Date.End == emptyDate {
		startDateAndTime = u.Properties.DoDate.Date.Start
		endDateAndTime = emptyDate
	} else {
		startDateAndTime = u.Properties.DoDate.Date.Start
		endDateAndTime = u.Properties.DoDate.Date.End
	}

	var duration string

	duration = u.Properties.Duration.Formula.String

	var kanban string
	emptySelect := unparsedmodels.Select{}

	if u.Properties.Kanban.Select == emptySelect {
		kanban = ""
	} else {
		kanban = u.Properties.Kanban.Select.Name
	}

	var isDone bool

	isDone = u.Properties.Done.Checkbox

	var locationId []string

	if len(u.Properties.Location.Relation) == 0 {
		locationId = []string{}
	} else {
		for _, record := range u.Properties.Location.Relation {
			locationId = append(locationId, record.ID)
		}
	}

	var taskTitle string

	if len(u.Properties.Task.Title) == 0 {
		taskTitle = ""
	} else {
		taskTitle = u.Properties.Task.Title[0].PlainText
	}

	var notes string
	if len(u.Properties.Notes.RichText) == 0 {
		notes = ""
	} else {
		notes = u.Properties.Notes.RichText[0].PlainText
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.MusicProject = musicProject
	p.Type = _type
	p.Location = locationRollup
	p.City = cityRollup
	p.Priority = priority
	p.StartDateAndTime = startDateAndTime
	p.EndDateAndTime = endDateAndTime
	p.Duration = duration
	p.Kanban = kanban
	p.IsDone = isDone
	p.LocationId = locationId
	p.Title = taskTitle
	p.Notes = notes
}
