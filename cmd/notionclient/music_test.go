package notionclient

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMusicService(t *testing.T) {
	client, err := NewClient(os.Getenv("NOTION_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("get by title", func(t *testing.T) {
		title := "A quiet place"
		result, err := client.Music.GetByTile(title)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, title, result.Title)
	})
}
