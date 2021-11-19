package _191

func hammingWeight(num uint32) int {
	result := uint32(0)
	for i := 0; i < 32; i++ {
		result += num >> i & 1
	}

	return int(result)
}