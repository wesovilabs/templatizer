package action

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMatches(t *testing.T) {
	// So far only goTemplate mode is supported
	templateMode := "goTemplate"
	reVars := varPatterns[templateMode]
	reExcluded := varExcludedPatterns[templateMode]

	res := findMatches("Hey my {{.friend}}", reVars, reExcluded)
	assert.Len(t, res, 1)
	assert.Len(t, res[0], 4)
	assert.Equal(t, res[0][2], "friend")

	res = findMatches("Hey my {{`{{.friend}}`}}", reVars, reExcluded)
	assert.Len(t, res, 0)

	res = findMatches("Hey my {{.friend}} {{`{{.friend2}}`}}", reVars, reExcluded)
	assert.Len(t, res, 1)
	assert.Len(t, res[0], 4)
	assert.Equal(t, res[0][2], "friend")
}
