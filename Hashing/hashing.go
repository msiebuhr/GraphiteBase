package Hashing

// Implementaiton of Graphite's consistent hashing algorithm.
//
// Should behave as https://github.com/graphite-project/carbon/blob/master/lib/carbon/hashing.py

import (
	"fmt"
	"crypto/md5"
)

type ConsistentHashRing struct {
	ring map[uint16]string
	nodes []string
	replicaCount int
}

func NewConsistentHashRing(nodes []string, replicaCount int) *ConsistentHashRing {
	s := &ConsistentHashRing{}
	s.ring = make(map[uint16]string)
	s.nodes = make([]string, 0)
	s.replicaCount = replicaCount
	return s
}

// Compute md5(key) and return the first two bytes as an uint16
func computeRingPosition(key string) uint16 {
	h := md5.New()
	h.Write([]byte(key))
	sum := h.Sum(nil)
	i := uint16(sum[0]) << 8 | uint16(sum[1])
	return i
}

// Add a node to the ring
func (s *ConsistentHashRing) addNode(node string) {
	s.nodes = append(s.nodes, node)

	// Insert node into ring
	for i := 0; i < s.replicaCount ; i += 1 {
		key := fmt.Sprintf("%s:%d", node, i)
		position := computeRingPosition(key)
		s.ring[position] = node
	}
}

func (s *ConsistentHashRing) removeNode(node string) {
	// Remove node from s.nodes
	newNodes := make([]string, 0)
	for _, nnode := range s.nodes {
		if node != nnode {
			newNodes = append(newNodes, nnode)
		}
	}
	s.nodes = newNodes

	// Remove node-references from the ringitself
	positionsToDelete := make([]uint16, 0)
	// Remove from ring
	for rpos, rnode := range s.ring {
		if (rnode == node) {
			positionsToDelete = append(positionsToDelete, rpos)
		}
	}

	for _, pos := range positionsToDelete {
		delete(s.ring, pos)
	}
}

// Return all relevant nodes to send the given key to.
//func (s *ConsistentHashRing) GetNodes(key string) []string {
	//return
//}
