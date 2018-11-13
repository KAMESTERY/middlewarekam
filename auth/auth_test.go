package auth

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthTestSuite struct {
	suite.Suite
	AuthKamClient
}

func (ats *AuthTestSuite) SetupSuite() {
	ats.AuthKamClient = NewAuthKamClient()
}

func (ats *AuthTestSuite) TearDownSuite() {
	// Do nothing
}

func TestRunAuthTestSuite(t *testing.T) {
	ats := new(AuthTestSuite)
	ats.SetT(t)
	suite.Run(t, ats)
}

func (ats *AuthTestSuite) TestLogin() {

	t := ats.T()

	assert := assert.New(t)

	user_creds_req := UserCredentialsReq{
		"hhhh@hhhh.hhh",
		"hhhhhhhh",
	}

	auth_claims_resp, err := ats.AuthKamClient.Authenticate(context.Background(), &user_creds_req)

	assert.True(err == nil)
	assert.True(auth_claims_resp.Email == user_creds_req.Email)
	assert.True(auth_claims_resp.Role < 0)
	assert.True(len(auth_claims_resp.UserId) > 0)
	assert.True(len(auth_claims_resp.Token) > 0)
}

func (ats *AuthTestSuite) TestEnroll() {

	t := ats.T()

	assert := assert.New(t)

	user_enroll_req := UserEnrollReq{
		"tttt",
		"tttt@tttt.ttt",
		"tttttttt",
	}

	enroll_status_resp, err := ats.AuthKamClient.Enroll(context.Background(), &user_enroll_req)

	assert.True(err == nil)
	assert.True(!enroll_status_resp.Success)
	assert.True(enroll_status_resp.Message == "")
}
