# Templatizer
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


**Templater** is created to empower the mechanism to take advantage of template repositories. It's known that Github (nor Gitlab) doesn't provide us a handy way to take advantage of the Template repositories. When we start from a Template we usually need to do a lot of manual and inefficient replacement tasks.
## Getting started

### Installation
`templatizer` can be installed by running

####
#### Brew tap

#### From source code
## Usage: Consuming templates
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
### Create repository from template

```bash
templatizer create --from github.com/ivancorrales/seed
```

By default, a named file `templatize-params.yml` will be used. If this file
doesn't exist we will be required to pass the input file with flag `--params`

```bash
templatizer create --from github.com/ivancorrales/seed --param customize-template.yml
```

## Usage: Creating a template

By default the templatizer tool will inspect all the files in the repository template and It will create a named file `templatizer-params.yml` with the variables to be populated. Additionally we could

## Contributing

Contributions are welcome, but before doing it I hardly recommend you to have a look at the [CONTRIBUTING.md](CONTRIBUTING.md) and the [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

### Pre-requisites

if you work with MacOS you can take advantage of [homebrew](https://brew.sh/index_es) o setup your environment. You just need to run `make setup`.

In other case you would need to setup manually the following tools:

- **Hadolint**: Visit the [official repository](https://github.com/hadolint/hadolint)
- **pre-commit**: Visit the [official site])(https://pre-commit.com/).
