package utils

import (
	"bytes"
	"net/http"
	"strings"
	"context"
)

var utilshttp_logger = NewLogger("utilshttp")

func ProcessRequest(ctx context.Context, r *http.Request) (resp *http.Response, err error) {
	c := &http.Client{
		Timeout: HTTP_CLIENT_TIMEOUT,
	}
	req := r.WithContext(ctx)
	resp, err = c.Do(req)
	return
}

func NewQuery(queryString string) *strings.Reader {
	buffer := new(bytes.Buffer)
	query := struct {
		Query string `json:"query"`
	}{
		Query: queryString,
	}
	EncodeJson(buffer, query)
	json := buffer.String()
	utilshttp_logger.Debugf("GQL_QUERY:::: %s", json)
	return strings.NewReader(json)
}
