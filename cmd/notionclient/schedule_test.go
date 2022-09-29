package notionclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestScheduleService(t *testing.T) {
	client, err := NewClient()
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
		//count := 0
		//log.Println(result)
		//for _, task := range result {
		//	t.Log(task)
		//	count++
		//}
		//t.Log(count)
	})

}
