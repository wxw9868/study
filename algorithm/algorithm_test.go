package algorithm

import (
	"github.com/wxw9868/study/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	utils.MethodRuntime(func() {
		sorting := []int{2, 9, 4, 7, 6, 1, 3, 5, 0}
		t.Log(sorting)
		BubbleSort(sorting)
		t.Log(sorting)
	})
}

func TestSort(t *testing.T) {
	utils.MethodRuntime(func() {
		data := []int{2, 10, 6, 1, 11, 4, 8, 3}
		t.Log(data)
		QuickSort(data, 0, len(data)-1)
		t.Log(data)
	})
}
