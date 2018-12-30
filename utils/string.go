package utils

import (
	"encoding/base64"
	"encoding/json"
	"github.com/Machiel/slugify"
	"strings"
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

func ToNamespace(chunks ...string) (ns string) {
	ns = strings.Join(chunks, NS_SEP)
	return
}

func CategorizeDocument(category string) string {
	categoryName := ToNamespace(Namespace, Slugify(category))
	return categoryName
}

func IdentifyDocument(identifier, title string) string {
	documentId := ToNamespace(Namespace, Slugify(identifier), Slugify(title))
	return documentId
}

func ToQualifiedCategory(document_handle string) (qualified_category string) {
	ns_identifier_chunks := strings.Split(document_handle, NS_SEP)
	qualified_category = ToNamespace(ns_identifier_chunks[:len(ns_identifier_chunks) -1]...)
	return
}

func EncodeBase64(text string) string {
	text64 := base64.URLEncoding.EncodeToString([]byte(text))
	return text64
}

func DecodeBase64(text64 string) (string, error) {
	textBytes, err := base64.URLEncoding.DecodeString(text64)
	if err != nil {
		return "", err
	}
	text := string(textBytes)
	return text, nil
}

func Stringify(s string) (clean_string string) {
	b, _ := json.Marshal(s)
	clean_string = string(b)
	return
}
