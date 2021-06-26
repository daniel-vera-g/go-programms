package bloom_filter

// Simple bloom filter implementation
// Returns a boolean on whether element in given filter
func bloomFilter(m int, n int, filter []int) bool {
	position := n % m

	if filter[position] == 0 {
		return false
	} else {
		return true
	}
}
