package utilities

func Sum[T Integer | Float](arr []T) T {
	var sum T
	for _, ele := range arr {
		sum += ele
	}
	return sum
}

func Mean[T Integer | Float](arr []T) T {
	if len(arr) == 0 {
		return 0
	}
	return Sum(arr) / T(len(arr))
}

func SumBy[K any, T Integer | Float](arr []K, f func(ele K) T) T {
	var sum T
	for _, ele := range arr {
		sum += f(ele)
	}
	return sum
}

func MeanBy[K any, T Integer | Float](arr []K, f func(ele K) T) T {
	if len(arr) == 0 {
		return 0
	}
	return SumBy(arr, f) / T(len(arr))
}
