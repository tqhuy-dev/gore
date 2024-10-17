package utilities

func BubbleSort[T Integer | Float](arr []T) []T {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func SelectionSort[T Integer | Float](arr []T) []T {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func InsertionSort[T Integer | Float](arr []T) []T {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func QuickSort[T Integer | Float](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := left
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return arr
}

func HeapSort[T Integer | Float](arr []T) []T {
	n := len(arr)
	// Xây dựng heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// Trích xuất phần tử từ heap
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
	return arr
}

func heapify[T Integer | Float](arr []T, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func MergeSort[T Integer | Float](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge[T Integer | Float](left, right []T) []T {
	result := []T{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
