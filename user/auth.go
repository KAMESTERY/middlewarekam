package user

import (
	"bytes"
	"context"
	"fmt"
	"github.com/KAMESTERY/middlewarekam/utils"
	"net/http"
	"strings"
)

const (
	enrollQuery = `
mutation EnrollUser {
  enroll(
    userId: "%s",
    email: "%s",
    username: "%s",
    password: "%s"
  )
}
`
	authenticateQuery = `
query authenticateUser{
  authenticate(
    userId: "%s",
    email: "%s",
    password: "%s"
  ) {
    token,
    userId,
    email,
    role
  }
}
`
)

var user_logger = utils.NewLogger("userauth")

type AuthClaims struct {
	Token string `json:"token"`
	UserId string `json:"userId"`
	Email  string `json:"email"`
	Role   int32 `json:"role"`
}

func Authenticate(ctx context.Context, email, password string) (authClaims AuthClaims) {
	qryData := fmt.Sprintf(
		authenticateQuery,
		email,
		email,
		password,
	)
	req, err := http.NewRequest("POST", utils.BackendGQL, newQuery(qryData))
	if err != nil {
		user_logger.Fatalf("AUTH_ERROR:::: %+v", err)
		return
	}
	req.Header.Set(CONTEN_TYPE, APPLICATION_JSON)

	resp, err := processRequest(ctx, req)
	if err == nil {
		resp_struct := &struct {
			Data struct {
				AuthClaims AuthClaims `json:"authenticate"`
			} `json:"data"`
		}{}
		utils.DecodeJson(resp.Body, resp_struct)
		authClaims = resp_struct.Data.AuthClaims
	}
	return
}

func Enroll(ctx context.Context, username, email, password string) (ok bool) {
	ok = true
	qryData := fmt.Sprintf(
		enrollQuery,
		email,
		email,
		username,
		password,
	)
	req, err := http.NewRequest("POST", utils.BackendGQL, newQuery(qryData))
	if err != nil {
		user_logger.Fatalf("AUTH_ERROR:::: %+v", err)
		ok = false
		return
	}
	req.Header.Set(CONTEN_TYPE, APPLICATION_JSON)

	resp, err := processRequest(ctx, req)
	if err == nil {
		resp_struct := &struct {
			Data struct {
				Success string `json:"enroll"`
			} `json:"data"`
		}{}
		utils.DecodeJson(resp.Body, resp_struct)
		if status := resp_struct.Data.Success; status == "" {
			ok = false
		}
	}
	return
}

func processRequest(ctx context.Context, r *http.Request) (resp *http.Response, err error) {
	c := &http.Client{
		Timeout: HTTP_CLIENT_TIMEOUT,
	}
	req := r.WithContext(ctx)
	resp, err = c.Do(req)
	return
}

func newQuery(queryString string) *strings.Reader {
	buffer := new(bytes.Buffer)
	query := struct {
		Query string `json:"query"`
	}{
		Query: queryString,
	}
	utils.EncodeJson(buffer, query)
	json := buffer.String()
	user_logger.Debugf("GQL_QUERY:::: %s", json)
	return strings.NewReader(json)
}
