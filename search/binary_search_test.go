package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	Name             string
	Numbers          []int32
	Item             int32
	ExpectedPosition int32
}{
	{
		Name:             "even-number-of-elements-item-exists",
		Numbers:          []int32{1, 2, 3, 4, 5, 6},
		Item:             int32(3),
		ExpectedPosition: 2,
	},
	{
		Name:             "even-number-of-elements-item-not-exists",
		Numbers:          []int32{2, 3, 5, 6, 12, 45},
		Item:             int32(8),
		ExpectedPosition: -1,
	},
	{
		Name:             "odd-number-of-elements-item-exists-first",
		Numbers:          []int32{1, 2, 3, 4, 5},
		Item:             int32(1),
		ExpectedPosition: 0,
	},
	{
		Name:             "odd-number-of-elements-item-exists-last",
		Numbers:          []int32{1, 2, 3, 4, 5},
		Item:             int32(5),
		ExpectedPosition: 4,
	},
	{
		Name:             "odd-number-of-elements-item-not-exists",
		Numbers:          []int32{1, 2, 3, 4, 5},
		Item:             int32(9),
		ExpectedPosition: -1,
	},
	{
		Name:             "odd-number-of-elements-item-not-exists",
		Numbers:          []int32{3, 5, 6, 12, 45},
		Item:             int32(8),
		ExpectedPosition: -1,
	},
	{
		Name:             "one-element-item-exists",
		Numbers:          []int32{3},
		Item:             int32(3),
		ExpectedPosition: 0,
	},
	{
		Name:             "one-element-item-not-exists",
		Numbers:          []int32{3},
		Item:             int32(2),
		ExpectedPosition: -1,
	},
}

func TestBinaryIterative(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			res := BinaryIterative(tc.Numbers, tc.Item)
			assert.Equal(t, tc.ExpectedPosition, res)
		})
	}
}

func TestBinaryRecursive(t *testing.T) {
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			res := BinaryRecursive(tc.Numbers, tc.Item, 0, int32(len(tc.Numbers))-1)
			assert.Equal(t, tc.ExpectedPosition, res)
		})
	}
}
