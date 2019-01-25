package content

import (
	"github.com/KAMESTERY/middlewarekam/utils"
	"github.com/bxcodec/faker"
)

var contentContentGenLogger = utils.NewLogger("content_contentgen")

// GenerateContent Generate Content using Fake Data
func GenerateContent(email string, docCount int) (cnt *Content) {
	cnt = &Content{}

	for i := 1; i <= docCount; i++ {
		doc := &Document{}
		err := faker.FakeData(doc)
		if err != nil {
			contentContentGenLogger.Errorf("CONTENT_GENERATION_ERROR:::: %+v", err)
		} else {
			doc.FiltreVisuel = Document_IG_HEFE
			doc.Langue = Document_SWAHILI
			doc.Niveau = Document_HIGH
			doc.Publish = true
			doc.Metadata.Identification.UserId = email
			doc.Metadata.Identification.Identifier = BOGUS_TOPIC
			doc.Metadata.Identification.Tags = []string{"generated", "fake", "bogus", "fakealicious"}
			cnt.Documents = append(cnt.Documents, doc)
		}
	}

	return
}
