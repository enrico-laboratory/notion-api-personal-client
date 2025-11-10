package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestScheduleService(t *testing.T) {
	client, err := NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}

	t.Run("QUERY all database pages", func(t *testing.T) {
		// empty body
		var body string

		result, err := client.Schedule.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 400)
	})

	t.Run("QUERY database with time filter", func(t *testing.T) {
		var body string
		body = `{ 
				"filter": {
		              "property": "Do Date",
		              "date": {
		                  "before": "2022-01-01"
		              }
				}
			}`
		result, err := client.Schedule.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 100)

	})

	t.Run("QUERY database with date filter and type", func(t *testing.T) {
		body := fmt.Sprintf(`{ 
				"filter": {
		              "property": "Do Date",
		              "date": {
		                  "on_or_after": "%v"
		              }
				}
			}`, time.Now().Format(time.RFC3339))
		result, err := client.Schedule.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 3)

	})

	t.Run("GetByProjectId", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		result, err := client.Schedule.GetByProjectId(projectId)
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
	})

	t.Run("GetByProjectIdAndType - no Type ", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		result, err := client.Schedule.GetByProjectIdAndType(projectId)
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
	})

	t.Run("GetByProjectIdAndType - Concert and Rehearsal ", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		result, err := client.Schedule.GetByProjectIdAndType(projectId, ScheduleTypeConcert, ScheduleTypeRehearsal)
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
	})

}
