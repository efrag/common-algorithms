package search

func BinaryIterative(nums []int32, item int32) int32 {
	low := int32(0)
	high := int32(len(nums))-1

	for low <= high {
		mid := low + ((high-low)/2)

		if nums[mid] == item {
			return mid
		}

		if nums[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return int32(-1)
}

func BinaryRecursive(nums []int32, item, low, high int32) int32 {
	if high < low {
		return int32(-1)
	}

	mid := low + ((high-low)/2)

	if nums[mid] == item {
		return mid
	}

	if nums[mid] > item {
		return BinaryRecursive(nums, item, low, mid-1)
	}

	return BinaryRecursive(nums, item, mid+1, high)
}