package sort

func MergeSort(nums []int32) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int32, low, high int) {
	if low < high {
		mid := low + (high-low)/2

		mergeSort(nums, low, mid)
		mergeSort(nums, mid+1, high)

		merge(nums, low, mid, high)
	}
}

func merge(nums []int32, low, mid, high int) {
	left := make([]int32, 0)
	right := make([]int32, 0)

	for i := low; i <= mid; i++ {
		left = append(left, nums[i])
	}
	for i := mid + 1; i <= high; i++ {
		right = append(right, nums[i])
	}

	l := 0
	r := 0
	for k := low; k <= high; k++ {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				nums[k] = left[l]
				l += 1
			} else {
				nums[k] = right[r]
				r += 1
			}
		} else if l < len(left) {
			nums[k] = left[l]
			l += 1
		} else {
			nums[k] = right[r]
			r += 1
		}
	}
}
