package algorithm

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	count := len(arr)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if arr[i] > arr[j] {
				swap(arr, j, i)
			}
		}
	}
}

// QuickSort 快速排序
func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	key := HoareSort(arr, left, right)
	QuickSort(arr, left, key-1)
	QuickSort(arr, key+1, right)
}

// HoareSort 快速排序
func HoareSort(arr []int, left, right int) int {
	mid := GetMidIndex(arr, left, right)
	swap(arr, mid, left)

	key := left
	for left < right {
		// 从右边比较当前值小的(右边，找小)
		for left < right && arr[key] <= arr[right] {
			right--
		}
		// 从左边比较当前值大的(左边，找大)
		for left < right && arr[key] >= arr[left] {
			left++
		}
		swap(arr, left, right)
	}
	swap(arr, key, left)
	return left
}

func swap(arr []int, left, right int) {
	arr[left], arr[right] = arr[right], arr[left]
}

// GetMidIndex 算法优化法1.三数取中
func GetMidIndex(arr []int, left, right int) int {
	mid := (right - left) / 2
	if arr[left] > arr[mid] {
		if arr[mid] > arr[right] {
			return mid
		} else if arr[left] < arr[right] {
			return left
		} else {
			return right
		}
	} else {
		if arr[mid] < arr[right] {
			return mid
		} else if arr[left] < arr[right] {
			return right
		} else {
			return left
		}
	}
}
