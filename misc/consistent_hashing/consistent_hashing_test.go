package consistent_hashing

import (
	"testing"
)

func TestConsistentHashing(t *testing.T) {
	var node int = StorageNode(98765, 5)
	if node != 144 {
		t.Error("Error in StorageNode", node)
	}
}
