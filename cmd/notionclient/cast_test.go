package notionclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestCastService(t *testing.T) {
	client, err := NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	t.Run("QUERY all database pages", func(t *testing.T) {
		// empty body
		var body string

		result, err := client.Cast.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 3)
	})

	t.Run("QUERY database with time filter", func(t *testing.T) {
		var body string
		body = `{ 
				"filter": {
		              "property": "Role",
		              "title": {
		                  "equals": "A"
		              }
				}
			}`
		result, err := client.Cast.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
		//count := 0
		//for _, musicProject := range result {
		//	t.Log(musicProject)
		//	count++
		//}
		//t.Log(count)
	})

	t.Run("GetByProjectId", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		result, err := client.Cast.GetByProjectId(projectId)
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
	})

	t.Run("GetByProjectIdAndStatus - Confirmed", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		status := CastStatusConfirmed
		result, err := client.Cast.GetByProjectIdAndStatus(projectId, status)
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
	})

	t.Run("GetByProjectIdAndStatus - Confirmed and Waiting", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		status := []string{CastStatusConfirmed, CastStatusWaiting}
		result, err := client.Cast.GetByProjectIdAndStatus(projectId, status...)
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
	})

}
