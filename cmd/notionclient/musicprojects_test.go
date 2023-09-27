package notionclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestMusicProjectsService(t *testing.T) {
	client, err := NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	t.Run("QUERY all database pages", func(t *testing.T) {
		// empty body
		var body string

		result, err := client.MusicProjects.query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 3)
	})

	t.Run("QUERY database with time filter", func(t *testing.T) {
		var body string
		body = `{ 
				"filter": {
		              "property": "Year",
		              "number": {
		                  "equals": 2020
		              }
				}
			}`
		result, err := client.MusicProjects.query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 3)
		//count := 0
		//for _, musicProject := range result {
		//	t.Log(musicProject)
		//	count++
		//}
		//t.Log(count)
	})

	t.Run("GetById", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		result, err := client.MusicProjects.GetById(projectId)
		assert.Empty(t, err)
		assert.Equal(t, projectId, result.Id)
	})

	t.Run("GetByTitle", func(t *testing.T) {
		title := "Palestrina and Marenzio"
		result, err := client.MusicProjects.GetByTitle(title)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, title, result.Title)
	})

	t.Run("GetByTitle - missing", func(t *testing.T) {
		fakeTitle := "Fake Title"
		_, err := client.MusicProjects.GetByTitle(fakeTitle)
		assert.NotEmpty(t, err)
	})

	t.Run("GetWithStatus on Going", func(t *testing.T) {
		status := MusicProjectsStatusOnGoing
		result, err := client.MusicProjects.GetWithStatus(status)
		assert.Empty(t, err)
		assert.True(t, len(result) > 1)
	})

	t.Run("GetPublished", func(t *testing.T) {
		result, err := client.MusicProjects.GetPublished()
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 2)
	})

	t.Run("Create test project", func(t *testing.T) {
		title := "test"
		choirId := "f7883d7baee1467a88463cc6fdac2ee9"
		status := "On Going"
		year := 1999
		result, err := client.MusicProjects.CreateProject(title, choirId, status, year)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	})
}
