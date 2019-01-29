package content

import (
	"github.com/KAMESTERY/middlewarekam/utils"
	"github.com/bxcodec/faker"
)

var contentContentGenLogger = utils.NewLogger("content_contentgen")

// GenerateContent Generate Content using Fake Data
func GenerateContent(email string, docCount int) (cnt *Content) {
	cnt = &Content{}
	tags := []string{"generated", "fake", "bogus", "fakealicious"}

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
			doc.Metadata.Identification.Tags = tags
			cnt.Documents = append(cnt.Documents, doc)
			contentContentGenLogger.Debugf("Generated Document_%+v with Title: %+v", i, doc.Title)
		}
		for j := 1; j <= docCount; j++ {
			img := &Media{}
			err := faker.FakeData(img)
			if err != nil {
				contentContentGenLogger.Errorf("CONTENT_GENERATION_ERROR:::: %+v", err)
			} else {
				img.Categorie = Media_IMAGE
				img.Metadata.Identification.UserId = email
				img.Metadata.Identification.Identifier = doc.Title
				img.Metadata.Identification.Tags = tags
				cnt.MediaItems = append(cnt.MediaItems, img)
				contentContentGenLogger.Debugf("Generated Media_%+v%+v with Name: %+v", i, j, img.Name)
			}
		}
	}

	return
}
