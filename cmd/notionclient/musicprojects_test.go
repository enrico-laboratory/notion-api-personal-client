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

		result, err := client.MusicProjects.Query(body)
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
		result, err := client.MusicProjects.Query(body)
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

	t.Run("GetWithStatus on Going", func(t *testing.T) {
		status := MusicProjectsStatusOnGoing
		result, err := client.MusicProjects.GetWithStatus(status)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) == 3)
	})

}
