package consistent_hashing

// Simple consistent Hashing implementation
// Consists of 5 nodes, a simple hashing function and a  digit input
// Returns the node number who which it adheres
func StorageNode(TicketID int, nodeCount int) int {
	var nodeList []int

	// We assume, that all nodes are arranged equally
	nodeDistance := 360 / nodeCount
	// Fill node list
	for i := 72; i <= 360; i += nodeDistance {
		nodeList = append(nodeList, i)
	}

	nodeMembership := h(TicketID) % 360

	return nextHigher(nodeMembership, nodeList)

}

// Recursive search for the next higher value in a slice
func nextHigher(nodeMembership int, nodeList []int) int {
	if nodeMembership > nodeList[len(nodeList)-1] || nodeMembership < nodeList[0] {
		// Between the highest and lowest -> Lowest
		return nodeList[0]
	} else if nodeList[0] < nodeMembership && nodeList[1] > nodeMembership {
		return nodeList[1]
	}
	return nextHigher(nodeMembership, nodeList[1:])
}

// Veeeerry simple hashing function
// Lol, actually only a mock
func h(x int) int {
	return x
}
