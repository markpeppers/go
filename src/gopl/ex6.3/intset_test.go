package intset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSet(t *testing.T) {
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

	// Add 1 again to ensure idempotency
	y.AddAll(1, 2, 3, 4, 23423)
	assert.Equal(t, 8, y.Len())
	assert.Equal(t, true, y.Has(23423))
}

func TestUnion(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 3, 5, 9, 432432, 2343299)
	y.AddAll(1, 3, 432, 2342, 32)

	x.UnionWith(&y)
	assert.Equal(t, 9, x.Len())
	assert.Equal(t, true, x.Has(5))
}

func TestIntersect(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 3, 5, 9, 15, 23, 500)
	y.AddAll(1, 3, 8, 10, 500, 23423423)

	x.IntersectWith(&y)
	assert.Equal(t, 3, x.Len())
	assert.Equal(t, true, x.Has(3))
	assert.Equal(t, false, x.Has(5))
	assert.Equal(t, true, x.Has(500))
}

func TestDifferenct(t *testing.T) {
	var x, y IntSet
	x.AddAll(12, 34, 5678, 9093939, 3234)
	y.AddAll(34, 3234, 23)

	x.DifferenceWith(&y)
	assert.Equal(t, 3, x.Len())
	assert.Equal(t, true, x.Has(9093939))
	assert.Equal(t, false, x.Has(3234))
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.AddAll(12, 34, 34242, 343, 8983)
	y.AddAll(13, 35, 34242, 343, 8983)

	x.SymmetricDifference(&y)
	assert.Equal(t, 4, x.Len())
	assert.Equal(t, true, x.Has(13))
	assert.Equal(t, false, x.Has(34242))
}
