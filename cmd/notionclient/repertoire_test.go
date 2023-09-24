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

	})

	t.Run("GetByProjectId", func(t *testing.T) {
		projectId := "a890db2a-12a5-4606-886d-fb35283250c6"
		result, err := client.Repertoire.GetByProjectId(projectId)
		assert.Empty(t, err)
		assert.True(t, len(result) > 5)
	})

	var pieceId string
	t.Run("Create", func(t *testing.T) {
		newPiece := &CreatePieceRequestProperties{
			Order:     "100",
			MusicID:   "6885b92c3331472da846fa14d7543fa5",
			ProjectID: "830e429de257411b9bd87bf3635459a2",
		}

		result, err := client.Repertoire.Create(newPiece)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, result)
		pieceId = result
	})

	t.Run("Delete", func(t *testing.T) {
		err := client.Repertoire.DeleteById(pieceId)
		assert.Empty(t, err)
	})

}
