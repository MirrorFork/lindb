package selector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexSlotSelector_IndexOf(t *testing.T) {
	selector := NewIndexSlotSelector(1)
	assert.Equal(t, -1, selector.IndexOf(10, 5))
	assert.Equal(t, 0, selector.IndexOf(10, 10))
	assert.Equal(t, 5, selector.IndexOf(10, 15))
	assert.Equal(t, 95, selector.IndexOf(10, 105))

	selector = NewIndexSlotSelector(3)
	assert.Equal(t, 0, selector.IndexOf(10, 12))
	assert.Equal(t, 1, selector.IndexOf(10, 13))
	assert.Equal(t, 40, selector.IndexOf(10, 130))
}
