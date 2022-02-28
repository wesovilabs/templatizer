package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_executeTemplate(t *testing.T) {
	cases := []struct {
		content   string
		variables Variables
		expected  string
	}{
		{
			content: "My name is {{ .name}}",
			variables: Variables{
				"name": "John",
			},
			expected: "My name is John",
		},
		{
			content:   "Hey {{.firstname}}",
			variables: Variables{},
			expected:  "Hey <no value>",
		},
		{
			content: "Hey you",
			variables: Variables{
				"name": "John",
			},
			expected: "Hey you",
		},
	}
	for _, c := range cases {
		res := executeTemplate("", c.content, c.variables)
		assert.Equal(t, c.expected, res)
	}
}
