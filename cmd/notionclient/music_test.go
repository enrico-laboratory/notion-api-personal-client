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

	var musicId string

	t.Run("Create", func(t *testing.T) {
		properties := &CreateMusicRequestProperties{
			Voices: "SATB",
			Score:  "www.example.com",
			//Media:       "",
			//Recording:   "sdf",
			Composer:    "Test Composer",
			Length:      2.5,
			Instruments: []string{"vln", "clav"},
			Solo:        "SA",
			Title:       "Test Delete music",
		}

		result, err := client.Music.CreateMusic(properties)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, result)

		musicId = result
	})

	t.Run("Delete", func(t *testing.T) {
		err := client.Music.DeleteMusicById(musicId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete not existent", func(t *testing.T) {
		fakeMusicId := "6dc5174a123a4decb8a8a1ea01221fak"
		err := client.Music.DeleteMusicById(fakeMusicId)
		assert.NotEmpty(t, err)
	})

}
