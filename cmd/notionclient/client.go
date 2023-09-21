package notionclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type config struct {
	databases struct {
		musicProjectsID string
		repertoireID    string
		scheduleID      string
		castID          string
		locationID      string
		ChoirID         string
	}
	apiVersion    string
	notionVersion string
	token         string
}

type NotionApiClient struct {
	config        config
	Schedule      ScheduleService
	MusicProjects MusicProjectsService
	Locations     LocationsService
	Cast          CastService
	Repertoire    RepertoireService
	Choir         ChoirService
}

func NewClient(token string) (*NotionApiClient, error) {

	if os.Getenv("NOTION_TOKEN") == "" {
		return nil, errors.New("NOTION_TOKEN not found in the env variables")
	}

	var cfg config
	cfg.databases.musicProjectsID = musicProjectDatabaseId
	cfg.databases.repertoireID = repertoireDatabaseId
	cfg.databases.scheduleID = taskDatabaseId
	cfg.databases.castID = castDatabaseId
	cfg.databases.locationID = locationDatabaseId
	cfg.databases.ChoirID = choirDabaseId
	cfg.apiVersion = apiVersion
	cfg.notionVersion = notionVersion
	cfg.token = token

	client := &NotionApiClient{
		config: cfg,
	}

	client.Schedule = &ScheduleClient{apiClient: client, cfg: cfg}
	client.MusicProjects = &MusicProjectsClient{apiClient: client, cfg: cfg}
	client.Locations = &LocationsClient{apiClient: client, cfg: cfg}
	client.Cast = &CastClient{apiClient: client, cfg: cfg}
	client.Repertoire = &RepertoireClient{apiClient: client, cfg: cfg}
	client.Choir = &ChoirClient{apiClient: client, cfg: cfg}

	return client, nil
}

func (c *NotionApiClient) databaseQuery(databaseID string, body []byte) (*http.Response, error) {
	u := fmt.Sprintf("databases/%v/query", databaseID)
	return request(u, c.config.token, c.config.notionVersion, c.config.apiVersion, body)
}

func (c *NotionApiClient) pages(body []byte) (*http.Response, error) {
	u := "pages"
	return request(u, c.config.token, c.config.notionVersion, c.config.apiVersion, body)
}

func request(path, bearer, notionVersion, apiVersion string, body []byte) (*http.Response, error) {
	var u url.URL
	u.Scheme = "https"
	u.Host = "api.notion.com"
	u.Path = fmt.Sprintf("%v/%v", apiVersion, path)
	r, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer))
	r.Header.Add("Notion-Version", notionVersion)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var apiErr Error
		err = json.NewDecoder(res.Body).Decode(&apiErr)
		if err != nil {
			return nil, err
		}
		return nil, &apiErr
	}

	return res, nil
}
