
package content

import (
	"context"
	"github.com/KAMESTERY/middlewarekam/auth"
	"github.com/KAMESTERY/middlewarekam/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ContentTestSuite struct {
	suite.Suite
	ContentKamClient
}

func (ats *ContentTestSuite) SetupSuite() {
	ats.ContentKamClient = NewContentKamClient()
}

func (ats *ContentTestSuite) TearDownSuite() {
	// Do nothing
}

func TestRunContentTestSuite(t *testing.T) {
	ats := new(ContentTestSuite)
	ats.SetT(t)
	suite.Run(t, ats)
}

func (ats *ContentTestSuite) TestCreate() {

	t := ats.T()

	assert := assert.New(t)

	//////// Data Setup

	authKamClient := auth.NewAuthKamClient()

	user_creds_req := auth.UserCredentialsReq{
		"hhhh@hhhh.hhh",
		"hhhhhhhh",
	}

	auth_claims_resp, err := authKamClient.Authenticate(context.Background(), &user_creds_req)

	token := auth_claims_resp.Token
	category := BOGUS_TOPIC

	a_tags := []string{"some", "cool", "things", "two words"}
	b_tags := []string{"some", "ok", "chose", "wise words"}

	content := GenerateContent(auth_claims_resp.Email, 2)
	content.Documents[0].Metadata.Identification.Tags = append(
		content.Documents[0].Metadata.Identification.Tags,
		a_tags...,
	)
	content.Documents[1].Metadata.Identification.Tags = append(
		content.Documents[1].Metadata.Identification.Tags,
		b_tags...,
	)

	//////// Verify Content
	assert.True(content.Ok())

	//////// API Calls and Assertions

	content_handles, err := ats.ContentKamClient.Create(context.Background(), auth_claims_resp.Email, token, content)

	assert.True(err == nil)
	assert.True(content_handles.Message == "")
	assert.True(len(content_handles.GetDocs().ItemIds) == 2)
	assert.True(len(content_handles.GetMedia().ItemIds) == 4)

	defer func() {
		// Delete
		deleted_content_handles, err := ats.Delete(context.Background(), auth_claims_resp.Email, token, content_handles)

		assert.True(err == nil)
		assert.True(deleted_content_handles.Message == "")
	}()

	// Retrieve Documents
	retrieved_content, err := ats.Get(context.Background(), auth_claims_resp.Email, token, content_handles.GetDocs())

	assert.True(err == nil)
	assert.True(len(retrieved_content.Documents) > 0)
	for _, doc := range retrieved_content.Documents {
		assert.True(doc.Ok())
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	// Retrieve Media
	retrieved_media, err := ats.Get(context.Background(), auth_claims_resp.Email, token, content_handles.GetMedia())

	assert.True(err == nil)
	assert.True(len(retrieved_media.MediaItems) > 0)
	for _, media := range retrieved_media.MediaItems {
		assert.True(media.Ok())
		assert.True(len(media.Metadata.Identification.Tags) > 0)
	}

	// Query
	qry1 := NewQuery(utils.CategorizeDocument(category))
	queried_content1, err := ats.Query(context.Background(), auth_claims_resp.Email, token, qry1)

	assert.True(err == nil)
	assert.True(len(queried_content1.Documents) > 0)
	for _, doc := range queried_content1.Documents {
		assert.True(doc.Ok())
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	// Query with Tag(s)
	limit := int32(1)
	qry2 := NewQuery(utils.CategorizeDocument(category)).
		WithTags("cool", "chose").
		WithLimit(limit)
	queried_content2, err := ats.Query(context.Background(), auth_claims_resp.Email, token, qry2)

	assert.True(err == nil)
	assert.True(len(queried_content2.Documents) > 0)
	assert.True(len(queried_content2.Documents) == int(limit))
	for _, doc := range queried_content2.Documents {
		assert.True(doc.Ok())
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	// Update
	content_to_update := queried_content2
	for _, doc := range content_to_update.Documents {
		doc.Niveau = Document_HIGH
	}

	updated_content_handles, err := ats.Update(context.Background(), auth_claims_resp.Email, token, content_to_update)

	assert.True(err == nil)
	assert.True(updated_content_handles.Message == "")

	retrieved_updated_content, err := ats.Get(context.Background(), auth_claims_resp.Email, token, updated_content_handles.GetDocs())

	assert.True(err == nil)
	assert.True(len(retrieved_updated_content.Documents) > 0)
	for _, doc := range retrieved_updated_content.Documents {
		assert.True(doc.Ok())
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
		assert.True(doc.Niveau == Document_HIGH)
	}

	// Query Latest
	queried_content3, err := ats.Latest(context.Background(), category, 0)

	assert.True(err == nil)
	assert.True(len(queried_content3.Documents) == 2)
	for _, doc := range queried_content3.Documents {
		assert.True(doc.Ok())
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	// Query by Tag(s)
	queried_content4, err := ats.ByTag(context.Background(), category, 0, "cool", "chose")

	assert.True(err == nil)
	assert.True(len(queried_content4.Documents) == 2)
	for _, doc := range queried_content4.Documents {
		assert.True(doc.Ok())
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	// Query One
	queried_content_one_document, err := ats.One(context.Background(), content_handles.GetDocs().ItemIds[1].Identifier)

	assert.True(err == nil)
	assert.True(len(queried_content_one_document.Documents) == 1)
	for _, doc := range queried_content_one_document.Documents {
		assert.True(doc.Ok())
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}
}
