package content

const (
	CONTENT_INPUT = `
{{define "ContentInput"}}
mutation CreateDocuments {
	handles: createDocuments(
		userId: "{{ .UserId }}",
		token: "{{ .Token }}",
		documents: [
			{{range .Content.Documents}}
			{
      			name: "{{IdentifyDocument .Metadata.Identification.Identifier .Title }}",
				userId: "{{ .Metadata.Identification.UserId }}",
				documentId: "{{IdentifyDocument .Metadata.Identification.Identifier .Title }}",
      			data: [
        			["title", "{{ .Title }}"],
        			["slug", "{{Slugify .Title}}"],
        			["publish", "{{ .Publish }}"],
        			["body", {{ .Body | Stringigy }}],
        			["langue", "{{ .Langue }}"],
        			["niveau", "{{ .Niveau }}"],
        			["filtre_visuel", "{{ .FiltreVisuel }}"],
        			["identifier", "{{ .Metadata.Identification.Identifier }}"],
					{{range .Metadata.Identification.Tags}}
					["{{.}}", "{{.}}"],
					{{end}}
      			]
			},
			{{end}}
		]
	)
}
{{end}}
`
	CONTENT_UPDATE_INPUT = `
{{define "ContentInput"}}
mutation UpdateDocuments {
	handles: updateDocuments(
		userId: "{{ .UserId }}",
		token: "{{ .Token }}",
		documents: [
			{{range .Content.Documents}}
			{
      			name: "{{IdentifyDocument .Metadata.Identification.Identifier .Title }}",
				userId: "{{ .Metadata.Identification.UserId }}",
				documentId: "{{IdentifyDocument .Metadata.Identification.Identifier .Title }}",
      			data: [
        			["title", "{{ .Title }}"],
        			["slug", "{{Slugify .Title}}"],
        			["publish", "{{ .Publish }}"],
        			["body", "{{ .Body }}"],
        			["langue", "{{ .Langue }}"],
        			["niveau", "{{ .Niveau }}"],
        			["filtre_visuel", "{{ .FiltreVisuel }}"],
        			["identifier", "{{ .Metadata.Identification.Identifier }}"],
					{{range .Metadata.Identification.Tags}}
					["{{.}}", "{{.}}"],
					{{end}}
      			]
			},
			{{end}}
		]
	)
}
{{end}}
`
	CONTENT_OUTPUT = `
{{define "ContentOutput"}}
query RetrieveDocuments {
	documents: getLesChoses(
		userId: "{{ .UserId }}",
		token: "{{ .Token }}",
		names: [
			{{range .ContentHandles.ItemIds}}
			"{{ .Identifier }}",
			{{end}}
		]
	){
    thing {
      name
      userId
      thingId
      version
      score
      createdAt
      updatedAt
    }
    data {
      dataId
      thingId
      key
      value
    }
  }
}
{{end}}
`
	CONTENT_DELETE_INPUT = `
{{define "ContentDeleteInput"}}
mutation DeleteDocuments {
	handles: deleteDocuments(
		userId: "{{ .UserId }}",
		token: "{{ .Token }}",
		documentIds: [
			{{range .ContentHandles.ItemIds}}
			"{{ .Identifier }}",
			{{end}}
		]
	)
}
{{end}}
`
)
