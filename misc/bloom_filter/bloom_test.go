package bloom_filter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	isInFilter := bloomFilter(9, 1, []int{1, 0, 0, 1, 0, 0, 1, 1, 1})
	if isInFilter == true {
		t.Error("Wrong Bloom filter", isInFilter)
	}
}
