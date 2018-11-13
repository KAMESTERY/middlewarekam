package utils

import (
	"encoding/json"
	"fmt"
	"github.com/Machiel/slugify"
)

func Slugify(title string) (slug string) {
	slug = slugify.Slugify(title)
	return
}

func Truncate(s string) string {
	var numRunes = 0
	for index, _ := range s {
		numRunes++
		if numRunes > 240 {
			return s[:index] + "[...]"
		}
	}
	return s
}

func IdentifyDocument(identifier, title string) string {
	documentId := fmt.Sprintf("%s.%s.%s", Namespace, identifier, Slugify(title))
	return documentId
}

func Stringigy(s string) (clean_string string) {
	b, _ := json.Marshal(s)
	clean_string = string(b)
	return
}
