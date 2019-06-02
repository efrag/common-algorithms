package sort

func swap(nums []int32, i, j int32) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func InsertionSort(nums []int32) {
	for pos := int32(1); pos < int32(len(nums)); pos++ {
		item := nums[pos]

		for j := pos - 1; j >= 0; j-- {
			if nums[j] > item {
				swap(nums, j, j+1)
			}
		}
	}
}
