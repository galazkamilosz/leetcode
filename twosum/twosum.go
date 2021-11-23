package twosum

func TwoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, v2 := range nums[i+1:] {
			if v+v2 == target {
				return []int{i, i + j + 1}
			}
		}
	}
	return []int{0, 0}
}
