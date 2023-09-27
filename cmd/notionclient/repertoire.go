package notionclient

import (
	"encoding/json"
	"fmt"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/parsedmodels"
	"github.com/enrico-laboratory/notion-api-personal-client/cmd/notionclient/models/unparsedmodels"
	"io"
	"log"
	"net/http"
	"strings"
)

type RepertoireService interface {
	Query(body string) ([]parsedmodels.Piece, error)
	GetAll() ([]parsedmodels.Piece, error)
	GetByProjectId(projectId string) ([]parsedmodels.Piece, error)
	GetByProjectIdAndSelected(projectId string) ([]parsedmodels.Piece, error)
	Create(properties *CreatePieceRequestProperties) (string, error)
	DeleteById(pieceId string) error
}

type RepertoireClient struct {
	apiClient *NotionApiClient
	cfg       config
}

func (s *RepertoireClient) Query(body string) ([]parsedmodels.Piece, error) {
	var repertoireParsed []parsedmodels.Piece
	var err error

	hasMore := true
	count := 0
	nextCursor := ""
	isBodyEmpty := body == ""

	for hasMore {

		var resp *http.Response
		var repertoireUnparsed unparsedmodels.Repertoire

		if count == 0 {
			resp, err = s.apiClient.databaseQuery(s.cfg.databases.repertoireID, []byte(body))
			if err != nil {
				return nil, err
			}
		} else {
			startCursor := fmt.Sprintf(`"start_cursor": "%v"`, nextCursor)
			var newBody string
			if isBodyEmpty {
				newBody = fmt.Sprintf("{%v}", startCursor)
			} else {
				newBody = fmt.Sprintf("%v%v,%v", body[:1], startCursor, body[1:])
			}
			resp, err = s.apiClient.databaseQuery(s.cfg.databases.repertoireID, []byte(newBody))
			if err != nil {
				return nil, err
			}
		}

		bodyResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyResp, &repertoireUnparsed)

		for _, piece := range repertoireUnparsed.Results {
			var parsedPiece parsedmodels.Piece

			parsePiece(&piece, &parsedPiece)

			repertoireParsed = append(repertoireParsed, parsedPiece)
		}
		nextCursor = repertoireUnparsed.NextCursor
		hasMore = repertoireUnparsed.HasMore
		count++
	}

	return repertoireParsed, nil
}

func (s *RepertoireClient) GetAll() ([]parsedmodels.Piece, error) {

	query, err := s.Query("")
	if err != nil {
		return nil, err
	}
	return query, nil
}

func (s *RepertoireClient) GetByProjectId(projectId string) ([]parsedmodels.Piece, error) {
	query, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	var result []parsedmodels.Piece

	for _, piece := range query {
		for _, projectIdRepertoire := range piece.MusicProject {
			if projectIdRepertoire == projectId {
				result = append(result, piece)
			}
		}
	}

	return result, nil
}

func (s *RepertoireClient) GetByProjectIdAndSelected(projectId string) ([]parsedmodels.Piece, error) {
	body := `{
    "filter": {
        "property": "Selected",
        "checkbox": {
            "equals": true
        	}
    	}
	}`
	query, err := s.Query(body)
	if err != nil {
		return nil, err
	}

	var result []parsedmodels.Piece

	for _, piece := range query {
		for _, projectIdRepertoire := range piece.MusicProject {
			if projectIdRepertoire == projectId {
				result = append(result, piece)
			}
		}
	}

	return result, nil
}

type CreatePieceRequestProperties struct {
	Order     string
	Selected  bool
	MusicID   string
	ProjectID string
}

func (s *RepertoireClient) Create(properties *CreatePieceRequestProperties) (string, error) {
	type createPieceRequest struct {
		Parent struct {
			DatabaseId string `json:"database_id"`
		} `json:"parent"`
		Properties struct {
			Order        unparsedmodels.Title    `json:"Order"`
			Music        unparsedmodels.Relation `json:"Music"`
			MusicProject unparsedmodels.Relation `json:"Music Project"`
			Selected     unparsedmodels.Checkbox `json:"Selected"`
		} `json:"properties"`
	}

	req := &createPieceRequest{}
	req.Parent.DatabaseId = repertoireDatabaseId

	var orderProperty unparsedmodels.TitleProperty
	orderProperty.Text.Content = properties.Order
	req.Properties.Order.Title = []unparsedmodels.TitleProperty{orderProperty}

	var musicProperty unparsedmodels.RelationProperty
	musicProperty.ID = properties.MusicID
	req.Properties.Music.Relation = []unparsedmodels.RelationProperty{musicProperty}

	var musicProjectProperty unparsedmodels.RelationProperty
	musicProjectProperty.ID = properties.ProjectID
	req.Properties.MusicProject.Relation = []unparsedmodels.RelationProperty{musicProjectProperty}

	var selected unparsedmodels.Checkbox
	selected.Checkbox = properties.Selected
	req.Properties.Selected = selected

	body, err := json.Marshal(&req)
	if err != nil {
		return "", err
	}
	log.Println(string(body))

	resp, err := s.apiClient.pages(body)
	if err != nil {
		return "", err
	}

	type ResponseID struct {
		ID string `json:"id"`
	}

	var id ResponseID
	err = json.NewDecoder(resp.Body).Decode(&id)
	if err != nil {
		return "", err
	}
	repId := id.ID
	return repId, nil
}

func (s *RepertoireClient) DeleteById(pieceId string) error {
	_, err := s.apiClient.pagesDelete(pieceId)
	if err != nil {
		return err
	}
	return nil
}

func parsePiece(u *unparsedmodels.Piece, p *parsedmodels.Piece) {

	var order string

	if len(u.Properties.Order.Title) == 0 {
		order = ""
	} else {
		order = u.Properties.Order.Title[0].PlainText
	}

	var num1 string

	if len(u.Properties.OneTopVoice.RichText) == 0 {
		num1 = ""
	} else {
		num1 = u.Properties.OneTopVoice.RichText[0].PlainText
	}

	var num2 string

	if len(u.Properties.Num2.RichText) == 0 {
		num2 = ""
	} else {
		num2 = u.Properties.Num2.RichText[0].PlainText
	}

	var num3 string

	if len(u.Properties.Num3.RichText) == 0 {
		num3 = ""
	} else {
		num3 = u.Properties.Num3.RichText[0].PlainText
	}

	var num4 string

	if len(u.Properties.Num4.RichText) == 0 {
		num4 = ""
	} else {
		num4 = u.Properties.Num4.RichText[0].PlainText
	}

	var num5 string

	if len(u.Properties.Num5.RichText) == 0 {
		num5 = ""
	} else {
		num5 = u.Properties.Num5.RichText[0].PlainText
	}
	var num6 string

	if len(u.Properties.Num6.RichText) == 0 {
		num6 = ""
	} else {
		num6 = u.Properties.Num6.RichText[0].PlainText
	}

	var num7 string

	if len(u.Properties.Num7.RichText) == 0 {
		num7 = ""
	} else {
		num7 = u.Properties.Num7.RichText[0].PlainText
	}

	var num8 string

	if len(u.Properties.Num8.RichText) == 0 {
		num8 = ""
	} else {
		num8 = u.Properties.Num8.RichText[0].PlainText
	}

	var num9 string

	if len(u.Properties.Num9.RichText) == 0 {
		num9 = ""
	} else {
		num9 = u.Properties.Num9.RichText[0].PlainText
	}

	var num10 string

	if len(u.Properties.Num10.RichText) == 0 {
		num10 = ""
	} else {
		num10 = u.Properties.Num10.RichText[0].PlainText
	}

	var num11 string

	if len(u.Properties.Num11.RichText) == 0 {
		num11 = ""
	} else {
		num11 = u.Properties.Num11.RichText[0].PlainText
	}

	var num12 string

	if len(u.Properties.Num12.RichText) == 0 {
		num12 = ""
	} else {
		num12 = u.Properties.Num12.RichText[0].PlainText
	}

	var note string

	if len(u.Properties.Note.RichText) == 0 {
		note = ""
	} else {
		note = u.Properties.Note.RichText[0].PlainText
	}

	var voicing string

	if len(u.Properties.VoicesRollup.Rollup.Array) == 0 {
		voicing = ""
	} else {
		voicing = u.Properties.VoicesRollup.Rollup.Array[0].Select.Name
	}

	var solo string

	if len(u.Properties.SoloRollup.Rollup.Array) == 0 {
		solo = ""
	} else {
		solo = u.Properties.SoloRollup.Rollup.Array[0].Select.Name
	}

	var music string

	if len(u.Properties.MusicRollup.Rollup.Array) == 0 {
		music = ""
	} else {
		music = u.Properties.MusicRollup.Rollup.Array[0].Title[0].PlainText
	}

	var musicProjectsIds []string

	if len(u.Properties.MusicProject.Relation) == 0 {
		musicProjectsIds = append(musicProjectsIds, "")
	} else {
		for _, musicProjectsId := range u.Properties.MusicProject.Relation {
			musicProjectsIds = append(musicProjectsIds, musicProjectsId.ID)
		}
	}

	var media string

	if len(u.Properties.MediaRollup.Rollup.Array) == 0 {
		media = ""
	} else {
		media = u.Properties.MediaRollup.Rollup.Array[0].URL
	}

	var score string

	if len(u.Properties.ScoreRollup.Rollup.Array) == 0 {
		score = ""
	} else {
		score = u.Properties.ScoreRollup.Rollup.Array[0].URL
	}

	var recording string

	if len(u.Properties.RecordingRollup.Rollup.Array) == 0 {
		recording = ""
	} else {
		recording = u.Properties.RecordingRollup.Rollup.Array[0].URL
	}

	var instruments []string

	if len(u.Properties.InstrumentsRollup.Rollup.Array) == 0 {
		instruments = append(instruments, "")
	} else {
		for _, instrument := range u.Properties.InstrumentsRollup.Rollup.Array[0].MultiSelect {
			instruments = append(instruments, instrument.Name)
		}
	}

	var length float64

	if len(u.Properties.LengthRollup.Rollup.Array) == 0 {
		length = 0
	} else {
		length = u.Properties.LengthRollup.Rollup.Array[0].Number
	}

	var composer string

	if len(u.Properties.ComposerRollup.Rollup.Array) == 0 {
		composer = ""
	} else if len(u.Properties.ComposerRollup.Rollup.Array[0].RichText) == 0 {
		composer = ""
	} else {
		composer = strings.TrimSpace(u.Properties.ComposerRollup.Rollup.Array[0].RichText[0].PlainText)
	}

	var notesDivisi string
	if len(u.Properties.NotesDivisi.RichText) == 0 {
		notesDivisi = ""
	} else {
		notesDivisi = u.Properties.NotesDivisi.RichText[0].PlainText
	}
	var notesRepertoire string
	if len(u.Properties.NotesRepertoire.RichText) == 0 {
		notesRepertoire = ""
	} else {
		notesRepertoire = u.Properties.NotesRepertoire.RichText[0].PlainText
	}

	p.Id = u.ID
	p.CreatedTime = u.CreatedTime
	p.LastEditedTime = u.LastEditedTime
	p.Order = order
	p.Num1 = num1
	p.Num2 = num2
	p.Num3 = num3
	p.Num4 = num4
	p.Num5 = num5
	p.Num6 = num6
	p.Num7 = num7
	p.Num8 = num8
	p.Num9 = num9
	p.Num10 = num10
	p.Num11 = num11
	p.Num12 = num12
	p.Note = note
	p.Voicing = voicing
	p.Solo = solo
	p.Music = music
	p.Selected = u.Properties.Selected.Checkbox
	p.MusicProject = musicProjectsIds
	p.Media = media
	p.Score = score
	p.Recording = recording
	p.Instrument = instruments
	p.Length = length
	p.Composer = composer
	p.NotesDivisi = notesDivisi
	p.NotesRepertoire = notesRepertoire
}
