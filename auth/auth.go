package auth

import (
	"fmt"
	"github.com/KAMESTERY/middlewarekam/utils"
	"context"
	"net/http"
)

var authauth_logger = utils.NewLogger("authauth")

// The authentication request message containing the user's credentials
type UserCredentialsReq struct {
	Email                string   `json:"email,omitempty"`
	Password             string   `json:"password,omitempty"`
}

func (m *UserCredentialsReq) Reset()         { *m = UserCredentialsReq{} }

func (m *UserCredentialsReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserCredentialsReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The response message containing the user's authentication claims
type AuthClaimsResp struct {
	Token                string   `json:"token,omitempty"`
	UserId               string   `json:"userId,omitempty"`
	Email                string   `json:"email,omitempty"`
	Role                 int32    `json:"role,omitempty"`
}

func (m *AuthClaimsResp) Reset()         { *m = AuthClaimsResp{} }

func (m *AuthClaimsResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthClaimsResp) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthClaimsResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthClaimsResp) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

// The enrollment request message containing the user's details
type UserEnrollReq struct {
	Username             string   `json:"username,omitempty"`
	Email                string   `json:"email,omitempty"`
	Password             string   `json:"password,omitempty"`
}

func (m *UserEnrollReq) Reset()         { *m = UserEnrollReq{} }

func (m *UserEnrollReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserEnrollReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserEnrollReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// The response containing the user's enrollment status
type EnrollStatusResp struct {
	Success              bool     `json:"success,omitempty"`
	Message              string   `json:"message,omitempty"`
}

func (m *EnrollStatusResp) Reset()         { *m = EnrollStatusResp{} }

func (m *EnrollStatusResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *EnrollStatusResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// AuthKamClient is the client API for AuthKam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthKamClient interface {
	// Authenticate
	Authenticate(ctx context.Context, in *UserCredentialsReq) (*AuthClaimsResp, error)
	// Enroll
	Enroll(ctx context.Context, in *UserEnrollReq) (*EnrollStatusResp, error)
}

type authKamClient struct {}

func NewAuthKamClient() AuthKamClient {
	return &authKamClient{}
}

func (c *authKamClient) Authenticate(ctx context.Context, in *UserCredentialsReq) (*AuthClaimsResp, error) {
	out := new(AuthClaimsResp)
	qryData := fmt.Sprintf(
		authenticateQuery,
		in.Email,
		in.Email,
		in.Password,
	)
	req, err := http.NewRequest("POST", utils.BackendGQL, utils.NewQuery(qryData))
	if err != nil {
		authauth_logger.Fatalf("AUTH_ERROR:::: %+v", err)
		return nil, err
	}
	req.Header.Set(utils.CONTEN_TYPE, utils.APPLICATION_JSON)

	resp, err := utils.ProcessRequest(ctx, req)
	if err == nil {
		resp_struct := &struct {
			Data struct {
				AuthClaimsResp AuthClaimsResp `json:"authenticate"`
			} `json:"data"`
		}{}
		utils.DecodeJson(resp.Body, resp_struct)
		out = &resp_struct.Data.AuthClaimsResp
	}
	return out, nil
}

func (c *authKamClient) Enroll(ctx context.Context, in *UserEnrollReq) (*EnrollStatusResp, error) {
	out := new(EnrollStatusResp)
	qryData := fmt.Sprintf(
		enrollQuery,
		in.Email,
		in.Email,
		in.Username,
		in.Password,
	)
	req, err := http.NewRequest("POST", utils.BackendGQL, utils.NewQuery(qryData))
	if err != nil {
		authauth_logger.Fatalf("AUTH_ERROR:::: %+v", err)
		return nil, err
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

		if status := resp_struct.Data.Success; status != "" {
			out.Message = resp_struct.Data.Success
			out.Success = true
		}
	}
	return out, nil
}
