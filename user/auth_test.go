package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthTestSuite struct {
	suite.Suite
}

func (ats *AuthTestSuite) SetupSuite() {
	// Do nothing
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

	email := "hhhh@hhhh.hhh"
	password := "hhhhhhhh"

	claims := Authenticate(context.Background(), email, password)

	assert.True(claims.Email == email)
	assert.True(claims.Role < 0)
	assert.True(len(claims.UserId) > 0)
	assert.True(len(claims.Token) > 0)
}

func (ats *AuthTestSuite) TestEnroll() {

	t := ats.T()

	assert := assert.New(t)

	username := "tttt"
	email := "tttt@tttt.ttt"
	password := "tttttttt"

	ok := Enroll(context.Background(), username, email, password)

	assert.True(!ok)
}
