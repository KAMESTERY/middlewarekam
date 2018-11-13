package utils

import (
	"bytes"
	"text/template"
)

var utilsrendering_logger = NewLogger("utilsrendering")

var (
	tpl = template.New("").Funcs(template.FuncMap{
		"Slugify":          Slugify,
		"Truncate":         Truncate,
		"IdentifyDocument": IdentifyDocument,
		"Stringigy":        Stringigy,
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
