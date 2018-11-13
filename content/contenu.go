package content

import (
	"context"
	"github.com/KAMESTERY/middlewarekam/utils"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/timestamp"
	"net/http"
	"strconv"
)

var contentcontenu_logger = utils.NewLogger("contentcontenu")

type Document_Langue int32

const (
	Document_ENGLISH    Document_Langue = 0
	Document_FRENCH     Document_Langue = 1
	Document_PORTUGUESE Document_Langue = 2
	Document_SPANISH    Document_Langue = 3
	Document_SWAHILI    Document_Langue = 4
)

var Document_Langue_name = map[Document_Langue]string{
	0: "ENGLISH",
	1: "FRENCH",
	2: "PORTUGUESE",
	3: "SPANISH",
	4: "SWAHILI",
}

var Document_Langue_value = map[string]int32{
	"ENGLISH":    0,
	"FRENCH":     1,
	"PORTUGUESE": 2,
	"SPANISH":    3,
	"SWAHILI":    4,
}

type Document_Niveau int32

const (
	Document_LOW      Document_Niveau = 0
	Document_MEDIUM   Document_Niveau = 1
	Document_HIGH     Document_Niveau = 2
	Document_CRITICAL Document_Niveau = 3
)

var Document_Niveau_name = map[int32]string{
	0: "LOW",
	1: "MEDIUM",
	2: "HIGH",
	3: "CRITICAL",
}

var Document_Niveau_value = map[string]int32{
	"LOW":      0,
	"MEDIUM":   1,
	"HIGH":     2,
	"CRITICAL": 3,
}

type Document_FiltreVisuel int32

const (
	Document_IG_WILLOW    Document_FiltreVisuel = 0
	Document_IG_WALDEN    Document_FiltreVisuel = 1
	Document_IG_VALENCIA  Document_FiltreVisuel = 2
	Document_IG_TOASTER   Document_FiltreVisuel = 3
	Document_IG_SUTRO     Document_FiltreVisuel = 4
	Document_IG_SIERRA    Document_FiltreVisuel = 5
	Document_IG_RISE      Document_FiltreVisuel = 6
	Document_IG_NASHVILLE Document_FiltreVisuel = 7
	Document_IG_MAYFAIR   Document_FiltreVisuel = 8
	Document_IG_LOFI      Document_FiltreVisuel = 9
	Document_IG_KELVIN    Document_FiltreVisuel = 10
	Document_IG_INKWELL   Document_FiltreVisuel = 11
	Document_IG_HUDSON    Document_FiltreVisuel = 12
	Document_IG_HEFE      Document_FiltreVisuel = 13
	Document_IG_EARLYBIRD Document_FiltreVisuel = 14
	Document_IG_BRANNAN   Document_FiltreVisuel = 15
	Document_IG_AMARO     Document_FiltreVisuel = 16
	Document_IG_1977      Document_FiltreVisuel = 17
)

var Document_FiltreVisuel_name = map[int32]string{
	0:  "IG_WILLOW",
	1:  "IG_WALDEN",
	2:  "IG_VALENCIA",
	3:  "IG_TOASTER",
	4:  "IG_SUTRO",
	5:  "IG_SIERRA",
	6:  "IG_RISE",
	7:  "IG_NASHVILLE",
	8:  "IG_MAYFAIR",
	9:  "IG_LOFI",
	10: "IG_KELVIN",
	11: "IG_INKWELL",
	12: "IG_HUDSON",
	13: "IG_HEFE",
	14: "IG_EARLYBIRD",
	15: "IG_BRANNAN",
	16: "IG_AMARO",
	17: "IG_1977",
}

var Document_FiltreVisuel_value = map[string]int32{
	"IG_WILLOW":    0,
	"IG_WALDEN":    1,
	"IG_VALENCIA":  2,
	"IG_TOASTER":   3,
	"IG_SUTRO":     4,
	"IG_SIERRA":    5,
	"IG_RISE":      6,
	"IG_NASHVILLE": 7,
	"IG_MAYFAIR":   8,
	"IG_LOFI":      9,
	"IG_KELVIN":    10,
	"IG_INKWELL":   11,
	"IG_HUDSON":    12,
	"IG_HEFE":      13,
	"IG_EARLYBIRD": 14,
	"IG_BRANNAN":   15,
	"IG_AMARO":     16,
	"IG_1977":      17,
}

type Media_Categorie int32

const (
	Media_AUDIO    Media_Categorie = 0
	Media_DOCUMENT Media_Categorie = 1
	Media_IMAGE    Media_Categorie = 2
	Media_VIDEO    Media_Categorie = 3
)

var Media_Categorie_name = map[int32]string{
	0: "AUDIO",
	1: "DOCUMENT",
	2: "IMAGE",
	3: "VIDEO",
}

var Media_Categorie_value = map[string]int32{
	"AUDIO":    0,
	"DOCUMENT": 1,
	"IMAGE":    2,
	"VIDEO":    3,
}

type Content struct {
	Documents  []*Document `json:"documents,omitempty"`
	MediaItems []*Media    `json:"media_items,omitempty"`
}

func (m *Content) Reset() { *m = Content{} }

func (m *Content) GetDocuments() []*Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

func (m *Content) GetMediaItems() []*Media {
	if m != nil {
		return m.MediaItems
	}
	return nil
}

type ContentHandles struct {
	ItemIds []*Identification `json:"item_ids,omitempty"`
	Message string            `json:"message,omitempty"`
}

func (m *ContentHandles) Reset() { *m = ContentHandles{} }

func (m *ContentHandles) GetItemIds() []*Identification {
	if m != nil {
		return m.ItemIds
	}
	return nil
}

func (m *ContentHandles) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Document struct {
	Title        string                `json:"title,omitempty"`
	Slug         string                `json:"slug,omitempty"`
	Publish      bool                  `json:"publish,omitempty"`
	Body         string                `json:"body,omitempty"`
	Langue       Document_Langue       `json:"langue,omitempty"`
	Niveau       Document_Niveau       `json:"niveau,omitempty"`
	FiltreVisuel Document_FiltreVisuel `json:"filtre_visuel,omitempty"`
	Metadata     *MetaData             `json:"metadata,omitempty"`
}

func (m *Document) Reset() { *m = Document{} }

func (m *Document) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Document) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Document) GetPublish() bool {
	if m != nil {
		return m.Publish
	}
	return false
}

func (m *Document) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Document) GetLangue() Document_Langue {
	if m != nil {
		return m.Langue
	}
	return Document_ENGLISH
}

func (m *Document) GetNiveau() Document_Niveau {
	if m != nil {
		return m.Niveau
	}
	return Document_LOW
}

func (m *Document) GetFiltreVisuel() Document_FiltreVisuel {
	if m != nil {
		return m.FiltreVisuel
	}
	return Document_IG_WILLOW
}

func (m *Document) GetMetadata() *MetaData {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Media struct {
	Name      string          `json:"name,omitempty"`
	Categorie Media_Categorie `json:"categorie,omitempty"`
	FileUrl   string          `json:"file_url,omitempty"`
	Metadata  *MetaData       `json:"metadata,omitempty"`
}

func (m *Media) Reset() { *m = Media{} }

func (m *Media) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Media) GetCategorie() Media_Categorie {
	if m != nil {
		return m.Categorie
	}
	return Media_AUDIO
}

func (m *Media) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *Media) GetMetadata() *MetaData {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type MetaData struct {
	Identification *Identification     `json:"identification,omitempty"`
	Timestamps     *TimeStamps         `json:"timestamps,omitempty"`
	M              map[string]*any.Any `json:"m,omitempty"`
}

func (m *MetaData) Reset() { *m = MetaData{} }

func (m *MetaData) GetIdentification() *Identification {
	if m != nil {
		return m.Identification
	}
	return nil
}

func (m *MetaData) GetTimestamps() *TimeStamps {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

func (m *MetaData) GetM() map[string]*any.Any {
	if m != nil {
		return m.M
	}
	return nil
}

type Identification struct {
	UserId     string   `json:"user_id,omitempty"`
	Identifier string   `json:"identifier,omitempty"`
	Tags       []string `json:"tags,omitempty"`
}

func (m *Identification) Reset() { *m = Identification{} }

func (m *Identification) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Identification) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *Identification) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TimeStamps struct {
	Created *timestamp.Timestamp `json:"created,omitempty"`
	Updated *timestamp.Timestamp `json:"updated,omitempty"`
}

func (m *TimeStamps) Reset() { *m = TimeStamps{} }

func (m *TimeStamps) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *TimeStamps) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

// ContentKamClient is the client API for ContentKam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentKamClient interface {
	// Create
	Create(ctx context.Context, userId, token string, in *Content) (*ContentHandles, error)
	// Update
	Update(ctx context.Context, userId, token string, in *Content) (*ContentHandles, error)
	// Get
	Get(ctx context.Context, userId, token string, in *ContentHandles) (*Content, error)
	// Delete
	Delete(ctx context.Context, userId, token string, in *ContentHandles) (*ContentHandles, error)
}

type contentKamClient struct{}

func NewContentKamClient() ContentKamClient {
	return &contentKamClient{}
}

func (c *contentKamClient) Create(ctx context.Context, userId, token string, in *Content) (*ContentHandles, error) {
	payload := struct {
		UserId  string
		Token   string
		Content Content
	}{
		userId,
		token,
		*in,
	}
	return c.processPayload(ctx, CONTENT_INPUT, "ContentInput", userId, token, payload)
}

func (c *contentKamClient) Update(ctx context.Context, userId, token string, in *Content) (*ContentHandles, error) {
	payload := struct {
		UserId  string
		Token   string
		Content Content
	}{
		userId,
		token,
		*in,
	}
	return c.processPayload(ctx, CONTENT_UPDATE_INPUT, "ContentUpdateInput", userId, token, payload)
}

func (c *contentKamClient) Get(ctx context.Context, userId, token string, in *ContentHandles) (*Content, error) {
	out := new(Content)
	qryData := utils.RenderString(
		CONTENT_OUTPUT,
		"ContentOutput",
		struct {
			UserId         string
			Token          string
			ContentHandles ContentHandles
		}{
			userId,
			token,
			*in,
		},
	)
	req, err := http.NewRequest("POST", utils.BackendGQL, utils.NewQuery(qryData))
	if err != nil {
		contentcontenu_logger.Fatalf("CONTENT_GET_ERROR:::: %+v", err)
		return nil, err
	}
	req.Header.Set(utils.CONTEN_TYPE, utils.APPLICATION_JSON)

	resp, err := utils.ProcessRequest(ctx, req)
	if err == nil {
		resp_struct := &struct {
			Data struct {
				Items []thingOutput `json:"documents"`
			} `json:"data"`
		}{}
		utils.DecodeJson(resp.Body, resp_struct)
		for _, item := range resp_struct.Data.Items {
			doc := &Document{
				Slug: item.Thing.Name,
				Metadata: &MetaData{
					Identification: &Identification{
						Identifier: item.Thing.ThingId,
						UserId:     item.Thing.UserId,
					},
					Timestamps: &TimeStamps{},
				},
			}
			for _, data := range item.Data {
				if data.Key == "body" {
					doc.Body = data.Value
				} else if data.Key == "filtre_visuel" {
					i, _ := strconv.ParseInt(data.Value, 10, 32) // TODO: Revisit this!!!!
					doc.FiltreVisuel = Document_FiltreVisuel(i)
				} else if data.Key == "langue" {
					i, _ := strconv.ParseInt(data.Value, 10, 32) // TODO: Revisit this!!!!
					doc.Langue = Document_Langue(i)
				} else if data.Key == "niveau" {
					i, _ := strconv.ParseInt(data.Value, 10, 32) // TODO: Revisit this!!!!
					doc.Niveau = Document_Niveau(i)
				} else if data.Key == "publish" {
					p, _ := strconv.ParseBool(data.Value)
					doc.Publish = p
				} else if data.Key == "slug" {
					doc.Slug = data.Value
				} else if data.Key == "title" {
					doc.Title = data.Value
				} else if data.Key == "identifier" {
					doc.Metadata.Identification.Identifier = data.Value
				} else {
					doc.Metadata.Identification.Tags = append(doc.Metadata.Identification.Tags, data.Value)
				}
			}
			out.Documents = append(out.Documents, doc)
		}
	}

	return out, nil
}

func (c *contentKamClient) Delete(ctx context.Context, userId, token string, in *ContentHandles) (*ContentHandles, error) {
	payload := struct {
		UserId  string
		Token   string
		ContentHandles ContentHandles
	}{
		userId,
		token,
		*in,
	}
	return c.processPayload(ctx, CONTENT_DELETE_INPUT, "ContentDeleteInput", userId, token, payload)
}

// Incoming

type thing struct {
	Name      string   `json:"name,omitempty"`
	UserId    string   `json:"userId,omitempty"`
	ThingId   string   `json:"thingId,omitempty"`
	Version   int32    `json:"version,omitempty"`
	Score     int32    `json:"score,omitempty"`
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
	DataIds   []string `json:"dataIds,omitempty"`
}

type data struct {
	DataId  string `json:"dataId,omitempty"`
	ThingId string `json:"thingId,omitempty"`
	Key     string `json:"key,omitempty"`
	Value   string `json:"value,omitempty"`
}

type thingOutput struct {
	Thing thing  `json:"thing,omitempty"`
	Data  []data `json:"data,omitempty"`
}
