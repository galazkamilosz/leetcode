package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type example struct {
		nums     []int
		target   int
		expected []int
	}
	examples := []example{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, exmp := range examples {
		if !reflect.DeepEqual(TwoSum(exmp.nums, exmp.target), exmp.expected) {
			t.Fail()
		}
	}
}
