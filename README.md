
[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/wesovilabs/templatizer)](https://github.com/wesovilabs/templatizer/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/wesovilabs/templatizer.svg)](https://pkg.go.dev/github.com/wesovilabs/templatizer)
[![go.mod](https://img.shields.io/github/go-mod/go-version/wesovilabs/templatizer)](go.mod)
[![LICENSE](https://img.shields.io/github/license/wesovilabs/templatizer)](LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/wesovilabs/templatizer/build)](https://github.com/wesovilabs/templatizer/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/wesovilabs/templatizer)](https://goreportcard.com/report/github.com/wesovilabs/templatizer)
[![Codecov](https://codecov.io/gh/wesovilabs/templatizer/branch/main/graph/badge.svg)](https://codecov.io/gh/wesovilabs/templatizer)
[![CodeQL](https://github.com/wesovilabs/templatizer/actions/workflows/codeql-analysis.yml/badge.svg?branch=main)](https://github.com/wesovilabs/templatizer/actions/workflows/codeql-analysis.yml)
---
# {{.Templatizer}}

The intention of **Templatizer** is to provide a handy and powerful mechanism to create custom projects from templates.

Gir repositories engines such as Github or Gitlab claim that they support repositories as templates. Actually,  they only permit us to tag repositories as templates; but we will need to replace the "dynamic values" after copying/clonnig the templates.  Sincerely,  this is far to be a teamplate mechanism from my point of view.

## Getting started

Templatizer takes advantage of existing template engines. So far, Templatizer supports Go Template but It's on the roadmap to provide other flavours such as Jinja.

Templatizer is meant to be executed as an executable file from your local machine. Thus,  the communication with the repositories will be established on your own machine and the credentials won't be sent over the Internet.

### Template

A template can contain varibales in the content of the files but also in the name of folders and files.

To define the variables in the templates we will use the specified format by Go Template. Variables are defined as `{{.variable}}`. See the following example taken from a Go file.

```go
package main

import (
	"embed"
	"io/fs"
	"net/http"
	log "{{.logger}}"
)

//go:embed {{.sitePath}}
var content embed.FS

func clientHandler() http.Handler {
	fsys := fs.FS(content)
	contentStatic, _ := fs.Sub(fsys, "{{.sitePath}}")
	return http.FileServer(http.FS(contentStatic))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", clientHandler())
	if err := http.ListenAndServe(":{{.serverPort}}", mux); err != nil {
		log.Fatal(err)
	}
}
```

The below screenshot shows an example of how we can use variabls in the name of the folders and the files.

![Templatizer](docs/screenshot.png)

To make a template available for other people that want to use it, we just need to define a  file that contains the definition of the variables in the template. We just need to enum the variables as shown in the below example::

```yml
version: v1
mode: goTemplate
variables:
  - name: logger
    default: github.com/sirupsen/logrus
  - name: serverPort
    default: 3000
  - name: sitePath
    description: Path to the static embedded foler
  - name: organization
    description: Name of the GH organization
```



The attributes `version` and `mode` are ignored in this version of Templatizer, so we could remove them from the configuration.

Regarding the variables, only the attribute `name`. Anyway, the usage of the attributes `description` and `default` will help us to create handier and more useful templates.

By convection this configuration file will be in the root of your repository and It will be named `templatizer.yml`.

The templates can be hosted  on any web repositories. You can find some examples of Templates on the following repositories:

- [Go template layout]()
- [React template layout]()
- [Terraform template]()
## Installation

**Download the artifact**

**Brew tap**

**From source code**
## Create a repository from a template
### Generate params file
To create a repository from a template we need to pass an input file with the
required variables. Instead of creating this file by hand we can take advantage
of the following command

```bash
templatizer setup --from github.com/ivancorrales/seed
```
Additionally we can pass the following flags:

```yaml
--username: Github|gitlab user handle
--password: The password for the username
```
### Create files

```bash
templatizer create --from github.com/ivancorrales/seed
```

By default, a named file `templatize-params.yml` will be used. If this file
doesn't exist we will be required to pass the input file with flag `--params`

```bash
templatizer create --from github.com/ivancorrales/seed --param customize-template.yml
```

## Provide a template

By default the templatizer tool will inspect all the files in the repository template and It will create a named file `templatizer-params.yml` with the variables to be populated. Additionally we could

## Contributing

Contributions are welcome, but before doing it I hardly recommend you to have a look at the [CONTRIBUTING.md](CONTRIBUTING.md) and the [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

### Pre-requisites

if you work with MacOS you can take advantage of [homebrew](https://brew.sh/index_es) o setup your environment. You just need to run `make setup`.

In other case you would need to setup manually the following tools:

- **Hadolint**: Visit the [official repository](https://github.com/hadolint/hadolint)
- **pre-commit**: Visit the [official site])(https://pre-commit.com/).
