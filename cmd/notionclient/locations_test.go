package notionclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestLocationsService(t *testing.T) {
	client, err := NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	t.Run("QUERY all database pages", func(t *testing.T) {
		// empty body
		var body string

		result, err := client.Locations.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) > 3)
	})

	t.Run("QUERY database with time filter", func(t *testing.T) {
		var body string
		body = `{ 
				"filter": {
		              "property": "Location",
		              "title": {
		                  "contains": "Noorderkerk"
		              }
				}
			}`
		result, err := client.Locations.Query(body)
		t.Log(fmt.Sprintf("Count results: %v", len(result)))
		assert.Empty(t, err)
		assert.True(t, len(result) == 1)
		//count := 0
		//for _, musicProject := range result {
		//	t.Log(musicProject)
		//	count++
		//}
		//t.Log(count)
	})

}
