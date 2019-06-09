package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	Name     string
	Numbers  []int32
	Expected []int32
}{
	{
		Name:     "All the numbers in the range [1-9]",
		Numbers:  []int32{4, 1, 9, 6, 3, 8, 7, 2, 5},
		Expected: []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		Name:     "3 numbers in the range [1-9]",
		Numbers:  []int32{4, 1, 9},
		Expected: []int32{1, 4, 9},
	},
	{
		Name:     "4 numbers in the range [1-9]",
		Numbers:  []int32{4, 1, 9, 6},
		Expected: []int32{1, 4, 6, 9},
	},
	{
		Name:     "1 number",
		Numbers:  []int32{4},
		Expected: []int32{4},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			BubbleSort(tc.Numbers)
			assert.True(t, assert.ObjectsAreEqualValues(tc.Expected, tc.Numbers))
		})
	}
}

func TestBubbleSortEarlyStop(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			BubbleSortEarlyStop(tc.Numbers)
			assert.True(t, assert.ObjectsAreEqualValues(tc.Expected, tc.Numbers))
		})
	}
}
