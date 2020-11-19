package string_split

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit1(t *testing.T) {
	s := `key="value if not" and=value why=not`
	c := []string{`key="value if not"`, `and=value`, "why=not"}
	r := split(s)

	assert.Equal(t, c, r)

}
