package notionclient

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestChoirClient(t *testing.T) {
	client, err := NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Get all choirs", func(t *testing.T) {

		result, err := client.Choir.GetAll()
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, len(result) > 1)
	})

	t.Run("Find choir by name", func(t *testing.T) {

		name := "Quattro Stagioni"
		result, err := client.Choir.GetByName(name)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, name, result.Name)
	})
}
