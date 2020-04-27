package intset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(345)
	x.Add(763)
	x.Add(123456789)
	assert.Equal(t, 4, x.Len())
	assert.Equal(t, true, x.Has(345))

	y := x.Copy()
	assert.Equal(t, 4, y.Len())

	x.Remove(345)
	assert.Equal(t, false, x.Has(345))
	x.Clear()
	assert.Equal(t, 0, x.Len())
	assert.Equal(t, false, x.Has(1))
	x.Add(5)
	assert.Equal(t, true, x.Has(5))

	// Assert y is still the same as before
	assert.Equal(t, 4, y.Len())
	assert.Equal(t, true, y.Has(123456789))
}
