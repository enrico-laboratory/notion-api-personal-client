package notionclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type config struct {
	databases struct {
		musicProjectsID string
		repertoireID    string
		scheduleID      string
		castID          string
		locationID      string
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
}

func NewClient() (*NotionApiClient, error) {

	if os.Getenv("NOTION_TOKEN") == "" {
		return nil, errors.New("NOTION_TOKEN not found in the env variables")
	}

	var cfg config
	cfg.databases.musicProjectsID = os.Getenv("MUSIC_PROJECT_DATABASE_ID")
	cfg.databases.repertoireID = os.Getenv("REPERTOIRE_DATABASE_ID")
	cfg.databases.scheduleID = os.Getenv("TASK_DATABASE_ID")
	cfg.databases.castID = os.Getenv("CAST_DATABASE_ID")
	cfg.databases.locationID = os.Getenv("LOCATION_DATABASE_ID")
	cfg.apiVersion = os.Getenv("API_VERSION")
	cfg.notionVersion = os.Getenv("NOTION_VERSION")
	cfg.token = os.Getenv("NOTION_TOKEN")

	client := &NotionApiClient{
		config: cfg,
	}

	client.Schedule = &ScheduleClient{apiClient: client, cfg: cfg}
	client.MusicProjects = &MusicProjectsClient{apiClient: client, cfg: cfg}
	client.Locations = &LocationsClient{apiClient: client, cfg: cfg}
	client.Cast = &CastClient{apiClient: client, cfg: cfg}
	client.Repertoire = &RepertoireClient{apiClient: client, cfg: cfg}

	return client, nil
}

func (c *NotionApiClient) request(databaseID string, body []byte) (*http.Response, error) {

	url := fmt.Sprintf("https://api.notion.com/%v/databases/%v/query", c.config.apiVersion, databaseID)

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.token))
	r.Header.Add("Notion-Version", c.config.notionVersion)
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
