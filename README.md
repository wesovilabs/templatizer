
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

The intention of **Templatizer** is to provide a handy and powerful mechanism to create a custom project boilerplate from existing template repositories.

Git repositories engines such as Github or Gitlab claim that they support template repositories. However, they only provide us a way to tag repositories as templates; but we will need to replace values" after copying/cloning the templates.  Sincerely,  this is far to be a template mechanism from my point of view.

## Getting started

Templatizer takes advantage of existing template engines. So far, Templatizer supports Go Template but It's on the roadmap to provide other flavors such as Jinja.

Templatizer is meant to be executed as an executable file from your local machine. Thus,  the communication with the repositories will be established on your own machine and the credentials won't be sent over the Internet.

Templatizer is composed of two parts: a server and a web application but distributed as a single application. The server listens to the requests from the web application on port 16917 (This should be parameterized in upcoming releases). The port used by the web application can be any available port on your local machine. Templatizer will find an available port and open the web app on your browser. The port used by the web application can't be specified so far but It will be soon ([issue](https://github.com/wesovilabs/templatizer/issues/5)).

## Installation

As It was mentioned in the above section, Templatizer is distributed as a single application. Let's see the different available options to install Templatizer on your local machine.
### Homebrew

```bash
brew tap wesovilabs/tools
brew install templatizer
```
### Snap

```bash
snap install templatizer
```
### Download executable files

Visit the [releases](https://github.com/wesovilabs/templatizer/releases) to find the compilation that works best  for you.

### Build Templatizer from the code

The executables of Templates can be found in the folder `dist`, after running the following commands.

```bash
git clone git@github.com:wesovilabs/templatizer.git
cd templatizer
make buildFrontend build
```
### Run from the code

Templatizer can launched from the code as following:

```bash
git clone git@github.com:wesovilabs/templatizer.git
cd templatizer
make buildFrontend run
```

## Templatizer in action

1. The web browser will be automatically opened when we execute `templatizer`. As was mentioned earlier, the port,  in which we can access the web application, can be whichever available port on the machine.

2. We must enter the values of the template repository. We can configure private repositories.

![Templatizer](docs/templatizer-step1.png)

* SSH connection will be supported in [the next release](https://github.com/wesovilabs/templatizer/issues/3).

1. Click on the button `Next` and fill the defined variables for this template.

![Templatizer](docs/templatizer-step3.png)

4. Once the variables are filled click on Process template and a tarball named `templatizer.tar.gz` will be downloaded.





## Define your own templates

The template is the main piece used by Templatizer. A template is a Git repository  hosted  on any web repositories. The templates  will contain values tp be dynamically replaced (variables). The varibales can be used in the content of the files but also in the name of folders and files.

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

Templatizer requires that the Tempalte reqpository host a file with the specification of the variables. This is not rocket science, we just need to enum the variables as shown in the below example::

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
  - name: DBPassword
	description: DB Password for integration tests
	secret: true
  - name: organization
    description: Name of the GH organization
```

The attributes `version` and `mode` could be omited since they are ignored in this version of Templatizer. Regarding the variables, only the attribute `name`. Anyway, the usage of the attributes `description` and `default` will help us to create handier and more useful templates. Apart from that we can also use the attribute `secret` to configure variables as passwords.

By convection the name of this file is `.templatizer.yml` and It's in the root of your repository.


- [Go template layout]()
- [React template layout]()
- [Terraform template]()

## Contributing

Contributions are welcome, but before doing it I hardly recommend you to have a look at the [CONTRIBUTING.md](CONTRIBUTING.md) and the [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).
