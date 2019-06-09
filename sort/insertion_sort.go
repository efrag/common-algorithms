package sort

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
