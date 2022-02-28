package action

import (
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Variables map[string]interface{}

func (v Variables) String() string {
	vars := varWithFullPath("", v)
	return fmt.Sprintf("%v", vars)
}

func varWithFullPath(prefix string, children Variables) []string {
	output := make([]string, 0)
	for k, v := range children {
		if len(v.(Variables)) == 0 {
			output = append(output, fmt.Sprintf("[%s.%s]", prefix, k))
			continue
		}
		output = append(output, varWithFullPath(k, v.(Variables))...)
	}
	return output
}

func (v Variables) ToYAML(path string) {
	v.pruneVariables()
	data, err := yaml.Marshal(&v)
	if err != nil {
		log.Warnf("it fails silently creating the params file: %s", err.Error())
	}
	if err := ioutil.WriteFile(path, data, os.ModePerm); err != nil {
		log.Warnf("it fails silently creating the params file: %s", err.Error())
	}
}

func (v Variables) pruneVariables() {
	for k, variable := range v {
		variableMap := variable.(Variables)
		if len(variableMap) == 0 {
			v[k] = nil
			continue
		}
		variableMap.pruneVariables()
	}
}
