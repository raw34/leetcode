package _137

func singleNumberII(nums []int) int {
	result := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}

		if total % 3 > 0 {
			result |= 1 << i
		}
	}
	return int(result)
}