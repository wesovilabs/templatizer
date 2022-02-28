package action

import (
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

var varExcludedPatterns = map[string]*regexp.Regexp{
	"goTemplate": regexp.MustCompile(`{{(\s)*\x60(\s)*{{(\s)*.([a-zA-Z0-9_.]+)(\s)*}}(\s)*\x60(\s)*}}`),
}

var varPatterns = map[string]*regexp.Regexp{
	"goTemplate": regexp.MustCompile(`{{(\s)*\.([a-zA-Z0-9_\.]+(\s)*)}}`),
}

func extract(repoFiles []repoFile, mode string) Variables {
	reVars := varPatterns[mode]
	reExcluded := varExcludedPatterns[mode]
	vars := make([]string, 0)

	for _, repoFile := range repoFiles {
		varsInFolderName := findMatches(filepath.Base(repoFile.folder), reVars, reExcluded)
		if len(varsInFolderName) > 0 {
			log.Debugf("\tThe name of folder %s contains variables", filepath.Base(repoFile.folder))
			for _, v := range varsInFolderName {
				vars = append(vars, v[2])
			}
		}

		varsInFileName := findMatches(filepath.Base(repoFile.name), reVars, reExcluded)
		if len(varsInFileName) > 0 {
			log.Debugf("\tThe name of file %s contains variables", filepath.Base(repoFile.name))
			for _, v := range varsInFileName {
				vars = append(vars, v[2])
			}
		}
		all_vars := findMatches(repoFile.content, reVars, reExcluded)
		if len(all_vars) > 0 {
			log.Debugf("\tThe file %s contains variables", filepath.Base(repoFile.name))
			for _, v := range all_vars {
				vars = append(vars, v[2])
			}
		}
	}
	out := Variables{}
	for _, v := range vars {
		vParts := strings.Split(v, ".")
		appendKeysToMap(vParts, out)
	}
	return out
}

func findMatches(text string, re *regexp.Regexp, reExclude *regexp.Regexp) [][]string {
	excludes := reExclude.FindAllStringSubmatch(text, -1)
	if len(excludes) > 0 {
		for _, v := range excludes {
			text = strings.Replace(text, v[0], "", -1)
		}
	}
	return re.FindAllStringSubmatch(text, -1)
}

func appendKeysToMap(keys []string, dict Variables) {
	key := keys[0]
	_, ok := dict[key]
	if !ok {
		dict[key] = Variables{}
	}
	if len(keys) > 1 {
		appendKeysToMap(keys[1:], dict[key].(Variables))
	}
}
