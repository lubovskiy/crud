package intersect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersectStr(t *testing.T) {
	s := IntersectStr([]string{"a"}, []string{"b"})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []string{})

	s = IntersectStr([]string{"a", "b"}, []string{"b"})
	assert.Equal(t, s, []string{"b"})
}