package notionclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestRepertoireService(t *testing.T) {
	client, err := NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	t.Run("QUERY all database pages", func(t *testing.T) {
		// empty body
		var body string

		result, err := client.Repertoire.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
		//count := 0
		//log.Println(result)
		//for _, task := range result {
		//	t.Log(task)
		//	count++
		//}
		//t.Log(count)
	})

	t.Run("QUERY database with time filter", func(t *testing.T) {
		var body string
		body = `{ 
				"filter": {
		              "property": "Order",
		              "title": {
		                  "contains": "01"
		              }
				}
			}`
		result, err := client.Repertoire.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 1)
		//count := 0
		//log.Println(result)
		//for _, task := range result {
		//	t.Log(task)
		//	count++
		//}
		//t.Log(count)
	})

	t.Run("GetByProjectId", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		result, err := client.Repertoire.GetByProjectId(projectId)
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
	})

}
