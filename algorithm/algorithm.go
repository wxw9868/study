package algorithm

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

// 快速排序
func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	key := Hsort(arr, left, right)
	QuickSort(arr, left, key-1)
	QuickSort(arr, key+1, right)
}

// 快速排序
func Hsort(arr []int, left, right int) int {
	key := left
	for left < right {
		// 从右边比较当前值小的
		for left < right && arr[key] <= arr[right] {
			right--
		}
		// 从左边比较当前值大的
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
