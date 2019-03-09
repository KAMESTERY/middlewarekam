package utils

import (
	"bytes"
	"text/template"
)

var utilsrendering_logger = NewLogger("utilsrendering")

var (
	tpl = template.New("").Funcs(template.FuncMap{
		"Slugify":             Slugify,
		"Truncate":            Truncate,
		"IdentifyDocument":    IdentifyDocument,
		"ToQualifiedCategory": ToQualifiedCategory,
		"EncodeBase64":        EncodeBase64,
		"ToNamespace":         ToNamespace,
		"CategorizeDocument":  CategorizeDocument,
		"Stringify":           Stringify,
		"ToJsonString":        ToJsonString,
	})
)

func RenderString(tmpl, name string, data interface{}) string {
	tpl.Parse(tmpl)
	buf := new(bytes.Buffer)
	tpl.ExecuteTemplate(buf, name, data)
	writtenString := buf.String()
	utilsrendering_logger.Debugf(
		"Template with name %s rendered to:::: \n%s",
		name,
		writtenString,
	)
	return writtenString
}
