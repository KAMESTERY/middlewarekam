package user

import (
	"context"
	"fmt"
	"github.com/KAMESTERY/middlewarekam/utils"
	"net/http"
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
	req, err := http.NewRequest("POST", utils.BackendGQL, utils.NewQuery(qryData))
	if err != nil {
		user_logger.Fatalf("AUTH_ERROR:::: %+v", err)
		return
	}
	req.Header.Set(utils.CONTEN_TYPE, utils.APPLICATION_JSON)

	resp, err := utils.ProcessRequest(ctx, req)
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
	req, err := http.NewRequest("POST", utils.BackendGQL, utils.NewQuery(qryData))
	if err != nil {
		user_logger.Fatalf("AUTH_ERROR:::: %+v", err)
		ok = false
		return
	}
	req.Header.Set(utils.CONTEN_TYPE, utils.APPLICATION_JSON)

	resp, err := utils.ProcessRequest(ctx, req)
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
