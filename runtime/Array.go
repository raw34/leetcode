package runtime

func swap(arr []int, i, j int) []int {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp

	return arr
}