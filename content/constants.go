package content

const (
	UPPER_LIMIT int32 = 100
)

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
      			name: "{{CategorizeDocument .Metadata.Identification.Identifier }}",
				userId: "{{ .Metadata.Identification.UserId }}",
				documentId: "{{IdentifyDocument .Metadata.Identification.Identifier .Title }}",
      			data: [
        			["title", "{{ .Title }}"],
        			["slug", "{{Slugify .Title}}"],
        			["publish", "{{ .Publish }}"],
        			["body", "{{ .Body | EncodeBase64 }}"],
        			["langue", "{{ .Langue }}"],
        			["niveau", "{{ .Niveau }}"],
        			["filtre_visuel", "{{ .FiltreVisuel }}"],
        			["identifier", "{{ .Metadata.Identification.Identifier }}"]
      			],
                tags: [
                    {{range .Metadata.Identification.Tags}}
					"{{ Slugify . }}",
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
      			name: "{{CategorizeDocument .Metadata.Identification.Identifier }}",
				userId: "{{ .Metadata.Identification.UserId }}",
				documentId: "{{IdentifyDocument .Metadata.Identification.Identifier .Title }}",
      			data: [
        			["title", "{{ .Title }}"],
        			["slug", "{{Slugify .Title}}"],
        			["publish", "{{ .Publish }}"],
        			["body", "{{ .Body | EncodeBase64 }}"],
        			["langue", "{{ .Langue }}"],
        			["niveau", "{{ .Niveau }}"],
        			["filtre_visuel", "{{ .FiltreVisuel }}"],
        			["identifier", "{{ .Metadata.Identification.Identifier }}"]
      			],
                tags: [
                    {{range .Metadata.Identification.Tags}}
					"{{ Slugify . }}",
					{{end}}
                ]
			},
			{{end}}
		]
	)
}
{{end}}
`
	QUERY_OUTPUT = `
{{define "QueryOutput"}}
query QueryDocuments {
	documents: queryLesChoses(
		userId: "{{ .UserId }}",
		token: "{{ .Token }}",
        {{ $limit := .Query.Limit }}{{ if gt $limit 0 }}
        limit: {{ $limit }},
        {{end}}
        attrNames: [
          ["#n" "Name"]
          {{ $length := len .Query.Tags }} {{ if gt $length 0 }}
            ["#tags" "Tags"]
          {{end}}
        ], 
        {{ $length := len .Query.Tags }} {{ if gt $length 0 }}
          filterExpr: "{{range $index, $tag := .Query.Tags}} {{if $index}}OR{{end}} contains(#tags, :Tag_{{$tag}}) {{end}}"
        {{end}}
        keyConditionExpr: "#n = :cat",
        data: [
			[":cat" "{{ .Query.Partition }}"]
            {{range .Query.Tags}}
            [":Tag_{{.}}" "{{.}}"]
            {{end}}
		]
	){
    thing {
      name
      userId
      thingId
      version
      score
      tags
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
	QUERY_PUBLIC_OUTPUT = `
{{define "QueryPublicOutput"}}
query QueryPublicDocuments {
	documents: queryPublicChoses(
        {{ $limit := .Query.Limit }}{{ if gt $limit 0 }}
        limit: {{ $limit }},
        {{end}}
        attrNames: [
          ["#n" "Name"]
          {{ $length := len .Query.Tags }} {{ if gt $length 0 }}
            ["#tags" "Tags"]
          {{end}}
        ], 
        {{ $length := len .Query.Tags }} {{ if gt $length 0 }}
          filterExpr: "{{range $index, $tag := .Query.Tags}} {{if $index}}OR{{end}} contains(#tags, :Tag_{{$tag}}) {{end}}"
        {{end}}
        keyConditionExpr: "#n = :cat",
        data: [
			[":cat" "{{ .Query.Partition }}"]
            {{range .Query.Tags}}
            [":Tag_{{.}}" "{{.}}"]
            {{end}}
		]
	){
    thing {
      name
      userId
      thingId
      version
      score
      tags
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
	CONTENT_OUTPUT = `
{{define "ContentOutput"}}
query RetrieveDocuments {
	documents: getLesChoses(
		userId: "{{ .UserId }}",
		token: "{{ .Token }}",
        data: [
			{{range .ContentHandles.ItemIds}}
			["{{ToQualifiedCategory .Identifier }}", "{{ .Identifier }}"],
			{{end}}
		]
	){
    thing {
      name
      userId
      thingId
      version
      score
      tags
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
	CONTENT_PUBLIC_OUTPUT = `
{{define "ContentPublicOutput"}}
query RetrievePublicDocuments {
	documents: getPublicChoses(
        data: [
			{{range .ContentHandles.ItemIds}}
			["{{ToQualifiedCategory .Identifier }}", "{{ .Identifier }}"],
			{{end}}
		]
	){
    thing {
      name
      userId
      thingId
      version
      score
      tags
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
		data: [
			{{range .ContentHandles.ItemIds}}
			["{{ToQualifiedCategory .Identifier }}", "{{ .Identifier }}"],
			{{end}}
		]
	)
}
{{end}}
`
)

const (
	DEVELOPMENT_TOPIC = "development"
	HISTORY_TOPIC     = "history"
	AFRICA_TOPIC      = "africa"
	LANGUAGE_TOPIC    = "language"
	EDUCATION_TOPIC   = "education"
)

var TOPICS = [...]string{
	DEVELOPMENT_TOPIC,
	HISTORY_TOPIC, 
	AFRICA_TOPIC,
	LANGUAGE_TOPIC,
	EDUCATION_TOPIC,
}

