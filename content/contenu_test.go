
package content

import (
	"context"
	"github.com/KAMESTERY/middlewarekam/auth"
	"github.com/KAMESTERY/middlewarekam/utils"
	"github.com/golang/protobuf/ptypes/any"
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

	authKamClient := auth.NewAuthKamClient()

	user_creds_req := auth.UserCredentialsReq{
		"hhhh@hhhh.hhh",
		"hhhhhhhh",
	}

	auth_claims_resp, err := authKamClient.Authenticate(context.Background(), &user_creds_req)

	token := auth_claims_resp.Token
	category := "a cool category"

	a_tags := []string{"some", "cool", "things"}
	b_tags := []string{"some", "ok", "chose"}

	content := Content{
		[]*Document{
			&Document{
				Slug: "AAAAAAA",
				Body: `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent pulvinar aliquam purus non volutpat. Sed varius ornare risus pretium ultrices. Donec placerat at libero id tincidunt. Quisque elementum ipsum non arcu convallis vehicula. Ut mattis ullamcorper dui luctus venenatis. Vivamus blandit mauris felis, sed aliquam urna consectetur eu. Phasellus vitae bibendum risus. Suspendisse dapibus efficitur porta. Proin imperdiet rhoncus condimentum. Aliquam erat volutpat. Pellentesque vitae tellus tortor.

Suspendisse fermentum nunc ut arcu ultricies commodo in nec nulla. Mauris luctus bibendum magna, ut hendrerit lorem ornare id. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras eu lacinia felis, non blandit arcu. Cras pharetra molestie auctor. Aliquam dictum condimentum erat non facilisis. Mauris congue, quam bibendum accumsan varius, elit mi iaculis tortor, vel iaculis ex leo sit amet felis.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris vitae nisi mollis, viverra eros eu, accumsan diam. Donec sed augue in nisi sagittis eleifend. Quisque turpis mi, sodales a rutrum ut, sodales eu orci. In iaculis faucibus ligula eget sodales. Mauris at lacus quis neque aliquam laoreet et eget turpis. Nulla pellentesque ullamcorper tortor a luctus. Mauris ultrices vulputate consectetur. Phasellus rhoncus dui eget enim gravida feugiat. Etiam vehicula, ipsum a imperdiet tincidunt, augue libero tristique quam, vitae feugiat sem arcu eu velit. Curabitur dictum urna a leo scelerisque, eget fringilla nulla porta.

Aenean porttitor elementum enim ut fermentum. Vivamus pretium, urna id posuere mattis, nunc leo feugiat dui, nec commodo odio diam sagittis lacus. Aenean auctor nisi nibh, ac sagittis enim porta id. Fusce pellentesque feugiat urna in facilisis. Donec condimentum sapien eget eros luctus blandit. Nulla euismod sapien a euismod tristique. Donec a nunc eu nisi molestie congue at id quam.

Ut tincidunt vestibulum lectus vitae semper. Vestibulum non quam diam. Maecenas placerat pharetra neque. In tristique, ligula eu tempus tincidunt, metus dui elementum est, id dapibus nulla turpis aliquam lacus. Duis at odio elementum, porta purus ac, fermentum arcu. Donec sollicitudin porta purus nec luctus. Sed scelerisque vitae velit et fermentum. Suspendisse malesuada ullamcorper arcu, at dapibus nisl rutrum ut. Nullam venenatis elit quam, eget consequat dolor sollicitudin ut. Donec molestie urna vel ultrices imperdiet. Praesent elit augue, ultrices malesuada nulla nec, tincidunt feugiat purus. Quisque mattis ac enim non fringilla.
`,
				FiltreVisuel: Document_IG_HEFE,
				Langue: Document_SWAHILI,
				Niveau: Document_HIGH,
				Publish: false,
				Title: "One New Piece of Content",
				Metadata: &MetaData{
					&Identification{
						auth_claims_resp.Email,
						category,
						a_tags,
					},
					&TimeStamps{},
					make(map[string]*any.Any),
				},
			},
			&Document{
				Slug: "BBBBBBB",
				Body: `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent pulvinar aliquam purus non volutpat. Sed varius ornare risus pretium ultrices. Donec placerat at libero id tincidunt. Quisque elementum ipsum non arcu convallis vehicula. Ut mattis ullamcorper dui luctus venenatis. Vivamus blandit mauris felis, sed aliquam urna consectetur eu. Phasellus vitae bibendum risus. Suspendisse dapibus efficitur porta. Proin imperdiet rhoncus condimentum. Aliquam erat volutpat. Pellentesque vitae tellus tortor.

Suspendisse fermentum nunc ut arcu ultricies commodo in nec nulla. Mauris luctus bibendum magna, ut hendrerit lorem ornare id. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras eu lacinia felis, non blandit arcu. Cras pharetra molestie auctor. Aliquam dictum condimentum erat non facilisis. Mauris congue, quam bibendum accumsan varius, elit mi iaculis tortor, vel iaculis ex leo sit amet felis.

Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris vitae nisi mollis, viverra eros eu, accumsan diam. Donec sed augue in nisi sagittis eleifend. Quisque turpis mi, sodales a rutrum ut, sodales eu orci. In iaculis faucibus ligula eget sodales. Mauris at lacus quis neque aliquam laoreet et eget turpis. Nulla pellentesque ullamcorper tortor a luctus. Mauris ultrices vulputate consectetur. Phasellus rhoncus dui eget enim gravida feugiat. Etiam vehicula, ipsum a imperdiet tincidunt, augue libero tristique quam, vitae feugiat sem arcu eu velit. Curabitur dictum urna a leo scelerisque, eget fringilla nulla porta.

Aenean porttitor elementum enim ut fermentum. Vivamus pretium, urna id posuere mattis, nunc leo feugiat dui, nec commodo odio diam sagittis lacus. Aenean auctor nisi nibh, ac sagittis enim porta id. Fusce pellentesque feugiat urna in facilisis. Donec condimentum sapien eget eros luctus blandit. Nulla euismod sapien a euismod tristique. Donec a nunc eu nisi molestie congue at id quam.

Ut tincidunt vestibulum lectus vitae semper. Vestibulum non quam diam. Maecenas placerat pharetra neque. In tristique, ligula eu tempus tincidunt, metus dui elementum est, id dapibus nulla turpis aliquam lacus. Duis at odio elementum, porta purus ac, fermentum arcu. Donec sollicitudin porta purus nec luctus. Sed scelerisque vitae velit et fermentum. Suspendisse malesuada ullamcorper arcu, at dapibus nisl rutrum ut. Nullam venenatis elit quam, eget consequat dolor sollicitudin ut. Donec molestie urna vel ultrices imperdiet. Praesent elit augue, ultrices malesuada nulla nec, tincidunt feugiat purus. Quisque mattis ac enim non fringilla.
`,
				FiltreVisuel: Document_IG_1977,
				Langue: Document_FRENCH,
				Niveau: Document_LOW,
				Publish: true,
				Title: "Another Piece of Content",
				Metadata: &MetaData{
					&Identification{
						auth_claims_resp.Email,
						category,
						b_tags,
					},
					&TimeStamps{},
					make(map[string]*any.Any),
				},
			},
		},
		[]*Media{
			&Media{},
		},
	}

	content_handles, err := ats.ContentKamClient.Create(context.Background(), auth_claims_resp.Email, token, &content)

	assert.True(err == nil)
	assert.True(content_handles.Message == "")

	retrieved_content, err := ats.Get(context.Background(), auth_claims_resp.Email, token, content_handles)

	assert.True(err == nil)
	assert.True(len(retrieved_content.Documents) > 0)
	for _, doc := range retrieved_content.Documents {
		assert.True(len(doc.Body) > 0)
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	qry1 := NewQuery(utils.CategorizeDocument(category))
	queried_content1, err := ats.Query(context.Background(), auth_claims_resp.Email, token, qry1)

	assert.True(err == nil)
	assert.True(len(queried_content1.Documents) > 0)
	for _, doc := range queried_content1.Documents {
		assert.True(len(doc.Body) > 0)
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	limit := int32(1)
	qry2 := NewQuery(utils.CategorizeDocument(category)).
		WithTags("cool", "chose").
		WithLimit(limit)
	queried_content2, err := ats.Query(context.Background(), auth_claims_resp.Email, token, qry2)

	assert.True(err == nil)
	assert.True(len(queried_content2.Documents) > 0)
	assert.True(len(queried_content2.Documents) == int(limit))
	for _, doc := range queried_content2.Documents {
		assert.True(len(doc.Body) > 0)
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	queried_content3, err := ats.Latest(context.Background(), category, 0)

	assert.True(err == nil)
	assert.True(len(queried_content3.Documents) == 2)
	for _, doc := range queried_content3.Documents {
		assert.True(len(doc.Body) > 0)
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	queried_content4, err := ats.ByTag(context.Background(), category, 0, "cool", "chose")

	assert.True(err == nil)
	assert.True(len(queried_content4.Documents) == 2)
	for _, doc := range queried_content4.Documents {
		assert.True(len(doc.Body) > 0)
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	queried_content_one_document, err := ats.One(context.Background(), content_handles.ItemIds[1].Identifier)

	assert.True(err == nil)
	assert.True(len(queried_content_one_document.Documents) == 1)
	for _, doc := range queried_content_one_document.Documents {
		assert.True(len(doc.Body) > 0)
		assert.True(len(doc.Metadata.Identification.Tags) > 0)
	}

	deleted_content_handles, err := ats.Delete(context.Background(), auth_claims_resp.Email, token, content_handles)

	assert.True(err == nil)
	assert.True(deleted_content_handles.Message == "")
}
