package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCastService(t *testing.T) {
	client, err := NewClient()
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

}
