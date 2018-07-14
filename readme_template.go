package main

const ReadmeTemplate = `
### My Github Stars

### List

{{- range $i, $repo := .}}
* [{{$repo.GetFullName}}]({{$repo.HTMLURL}}) - {{$repo.Description}}
{{- end}}

### Table

| Repo | Description | Homepage | HTMLURL | Language | ForksCount | StargazersCount | OpenIssuesCount |
| ------ | ------ | ------ | ------ | ------ | ------ | ------ | ------ |
{{- range $i, $repo := .}}
| {{$repo.GetFullName}} | {{$repo.Description}} | {{$repo.Homepage}} | {{$repo.HTMLURL}} | {{$repo.GetLanguage}} | {{$repo.ForksCount}} | {{$repo.StargazersCount}} | {{$repo.OpenIssuesCount}} |
{{- end}}
`