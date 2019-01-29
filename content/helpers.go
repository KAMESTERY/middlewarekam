package content

import (
	"context"
	"github.com/KAMESTERY/middlewarekam/utils"
	"net/http"
	"strconv"
)

func (c *contentKamClient) processPayload(ctx context.Context, tmpl, name, userId string, payload interface{}) (*ContentHandles, error) {
	out := new(ContentHandles)
	qryData := utils.RenderString(
		tmpl,
		name,
		payload,
	)
	req, err := http.NewRequest("POST", utils.BackendGQL, utils.NewQuery(qryData))
	if err != nil {
		contentcontenu_logger.Fatalf("PAYLOAD_PROCESSING_ERROR:::: %+v", err)
		return nil, err
	}
	req.Header.Set(utils.CONTENT_TYPE, utils.APPLICATION_JSON)

	resp, err := utils.ProcessRequest(ctx, req)
	if err == nil {
		resp_struct := &struct {
			Data struct {
				Handles []string `json:"handles"`
				MediaHandles []string `json:"media_handles"`
			} `json:"data"`
		}{}
		utils.DecodeJson(resp.Body, resp_struct)

		handles := resp_struct.Data.Handles
		for _, handle := range handles {
			out.ItemIds = append(out.ItemIds, &Identification{
				UserId:     userId,
				Identifier: handle,
				Type: Document_HANDLE,
			})
		}
		media_handles := resp_struct.Data.MediaHandles
		for _, mhandle := range media_handles {
			out.ItemIds = append(out.ItemIds, &Identification{
				UserId:     userId,
				Identifier: mhandle,
				Type: Media_HANDLE,
			})
		}
	}
	return out, nil
}

func (c *contentKamClient) processQuery(ctx context.Context, tmpl, name string, payload interface{}) (*Content, error) {
	qryData := utils.RenderString(
		tmpl,
		name,
		payload,
	)
	req, err := http.NewRequest("POST", utils.BackendGQL, utils.NewQuery(qryData))
	if err != nil {
		contentcontenu_logger.Fatalf("CONTENT_GET_ERROR:::: %+v", err)
		return nil, err
	}
	req.Header.Set(utils.CONTENT_TYPE, utils.APPLICATION_JSON)

	resp, err := utils.ProcessRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	out := new(Content)
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
					Tags: item.Thing.Tags,
				},
				Timestamps: &TimeStamps{},
			},
		}
		media := &Media{
			Name: item.Thing.Name,
			Metadata: &MetaData{
				Identification: &Identification{
					Identifier: item.Thing.ThingId,
					UserId:     item.Thing.UserId,
					Tags: item.Thing.Tags,
				},
				Timestamps: &TimeStamps{},
			},
		}
		for _, data := range item.Data {
			if data.Key == "body" {
				//doc.Body = data.Value
				doc.Body, _ = utils.DecodeBase64(data.Value)
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
			} else if data.Key == "categorie" {
				i, _ := strconv.ParseInt(data.Value, 10, 32) // TODO: Revisit this!!!!
				media.Categorie = Media_Categorie(i)
			} else if data.Key == "fileurl" {
				media.FileUrl = data.Value
			} else {
				doc.Metadata.Identification.Tags = append(doc.Metadata.Identification.Tags, data.Value)
				media.Metadata.Identification.Tags = append(media.Metadata.Identification.Tags, data.Value)
			}
		}
		if doc.Ok() {
			out.Documents = append(out.Documents, doc)
		}
		if media.Ok() {
			out.MediaItems = append(out.MediaItems, media)
		}
	}

	return out, nil
}
