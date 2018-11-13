package content

import (
	"context"
	"github.com/KAMESTERY/middlewarekam/utils"
	"net/http"
)

func (c *contentKamClient) processPayload(ctx context.Context, tmpl, name, userId, token string, payload interface{}) (*ContentHandles, error) {
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
	req.Header.Set(utils.CONTEN_TYPE, utils.APPLICATION_JSON)

	resp, err := utils.ProcessRequest(ctx, req)
	if err == nil {
		resp_struct := &struct {
			Data struct {
				Handles []string `json:"handles"`
			} `json:"data"`
		}{}
		utils.DecodeJson(resp.Body, resp_struct)

		handles := resp_struct.Data.Handles
		for _, handle := range handles {
			out.ItemIds = append(out.ItemIds, &Identification{
				UserId:     userId,
				Identifier: handle,
			})
		}
	}
	return out, nil
}
