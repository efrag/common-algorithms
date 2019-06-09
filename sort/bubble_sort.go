package sort

// BubbleSort is a comparison algorithm that continuously swaps elements in the array until
// the elements appear sorted. At each iteration of the outer loop the largest element in the
// array will have "bubbled" up to the last element of the array - so at each iteration we can
// reduce the size of the array we need to be looking at by 1
func BubbleSort(nums []int32) {
	for n := len(nums); n > 1; n-- {
		for j := 1; j < n; j++ {
			if nums[j-1] > nums[j] {
				swap(nums, j-1, j)
			}
		}
	}
}

// BubbleSortEarlyStop uses the Bubble sort algorithm with one minor optimization where we
// check whether any elements have been swapped and if so we stop the execution
func BubbleSortEarlyStop(nums []int32) {
	for n := len(nums); n > 1; n-- {
		swapped := false
		for j := 1; j < n; j++ {
			if nums[j-1] > nums[j] {
				swap(nums, j-1, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
