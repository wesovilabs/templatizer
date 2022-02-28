module github.com/{{.repository.organization}}/{{.repository.name}}/build

go {{.goVersion}}

require (
	github.com/golangci/golangci-lint v1.44.2
	github.com/goreleaser/goreleaser v1.4.1
)
