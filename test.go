package test

import "fmt"

// 选择
func SeletSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		end := i
		current := arr[i]
		for end >= 0 {
			if end == 0 || current >= arr[end-1] {
				break
			}
			arr[end] = arr[end-1]
			end--
		}
		arr[end] = current
	}
}

// 希尔
func ShellSort(arr []int) {
	gap := len(arr)
	gap /= 2
	for gap > 0 {
		for i := gap; i < len(arr); i++ { // i = 0
			end := i
			current := arr[i]
			for end >= gap && arr[end-gap] > current { ///end>=0
				arr[end] = arr[end-gap]
				end -= gap
			}
			arr[end] = current
		}
		gap /= 2
	}
}

// 堆
func HeapSort(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

// 堆方法
func heapify(arr []int, len int, i int) {
	laster := i
	left := i*2 + 1
	right := i*2 + 2
	if left < len && arr[left] > arr[laster] {
		laster = left
	}
	if right < len && arr[right] > arr[laster] {
		laster = right
	}
	if laster != i {
		arr[laster], arr[i] = arr[i], arr[laster]
		heapify(arr, len, laster)
	}
}

// 归并
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

// 归并排序的方法
func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		res[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		res[k] = right[j]
		j++
		k++
	}
	return res
}

//快速排序分开方法
func partition(arr []int, left, right int) int {
	current := arr[left]
	start := left
	end := right - 1
	for start < end {
		// end++
		for start < end {
			if current < arr[start] {
				break
			}
			start++
		}
		for end < right {
			if current > arr[end] {
				break
			}
			end--
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}
	arr[end], arr[left] = current, arr[end]
	return end
}

//快速排序
func QuickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		fmt.Println(pivot)
		QuickSort(arr, low, pivot)
		QuickSort(arr, pivot+1, high)
	}
}
