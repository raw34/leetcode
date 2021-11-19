package _190

func reverseBits(num uint32) uint32 {
	rev := uint32(0)
	for i := 0; i < 32; i++ {
		rev |= num & 1 << (31 - i)
		num >>= 1
	}

	return rev
}