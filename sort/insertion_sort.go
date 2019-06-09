package sort

func swap(nums []int32, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func InsertionSort(nums []int32) {
	for pos := 1; pos < len(nums); pos++ {
		item := nums[pos]

		for j := pos - 1; j >= 0; j-- {
			if nums[j] > item {
				swap(nums, j, j+1)
			}
		}
	}
}
