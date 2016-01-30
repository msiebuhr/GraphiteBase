package Hashing

// Implementaiton of Graphite's consistent hashing algorithm.
//
// Should behave as https://github.com/graphite-project/carbon/blob/master/lib/carbon/hashing.py

import (
	"crypto/md5"
	"fmt"
	"sort"
)

type ringEntry struct {
	hash uint16
	key  string
}

type ringEntries []*ringEntry

func (s ringEntries) Len() int           { return len(s) }
func (s ringEntries) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ringEntries) Less(i, j int) bool { return s[i].hash < s[j].hash }

type ConsistentHashRing struct {
	ring         ringEntries
	nodes        []string
	replicaCount int
}

func NewConsistentHashRing(nodes []string, replicaCount int) *ConsistentHashRing {
	s := &ConsistentHashRing{}
	s.ring = ringEntries{}
	s.nodes = make([]string, 0)
	s.replicaCount = replicaCount

	// Add the given nodes
	for _, n := range nodes {
		s.AddNode(n)
	}

	return s
}

// Compute md5(key) and return the first two bytes as an uint16
func computeRingPosition(key string) uint16 {
	sum := md5.Sum([]byte(key))
	i := uint16(sum[0])<<8 | uint16(sum[1])
	return i
}

// Add a node to the ring
func (s *ConsistentHashRing) AddNode(node string) {
	s.nodes = append(s.nodes, node)

	// Insert node into ring
	for i := 0; i < s.replicaCount; i += 1 {
		key := fmt.Sprintf("%s:%d", node, i)
		position := computeRingPosition(key)
		s.ring = append(s.ring, &ringEntry{position, node})
	}
	sort.Sort(s.ring)
}

func (s *ConsistentHashRing) RemoveNode(node string) {
	// Copy current ring-list and empty other state
	newNodes := make([]string, 0)
	s.nodes = make([]string, 0)
	s.ring = ringEntries{}

	// Copy over nodes we want to keep
	for _, nnode := range newNodes {
		if node != nnode {
			s.AddNode(node)
		}
	}
}

func (s *ConsistentHashRing) getFirstMatchingIndex(key string) int {
	keyHash := computeRingPosition(key)

	// If the hash is greater than the last element, it means we should wrp
	// to the first element.
	if keyHash > s.ring[len(s.ring)-1].hash {
		return 0
	}

	// Do a binary search for the appropriate element
	return sort.Search(
		len(s.ring),
		func(i int) bool { return s.ring[i].hash >= keyHash },
	)
}

// Return first node hit in the ring for the given key.
func (s *ConsistentHashRing) GetNode(key string) string {
	return s.ring[s.getFirstMatchingIndex(key)].key
}

// Return all relevant nodes to send the given key to.
//
// Trivial implementation that makes no attempts at being fast...
func (s *ConsistentHashRing) GetNodes(key string) []string {
	// Array of keys to fill out
	out := make([]int, 1)

	// Get the first element
	out[0] = s.getFirstMatchingIndex(key)

	// Search forward in the hash until we have seen each server once
	for i := (out[0] + 1) % len(s.ring); i != out[0]; i = (i + 1) % len(s.ring) {
		// Check if we already have seen this node
		seen := false
		for _, o := range out {
			if s.ring[i].key == s.ring[o].key {
				seen = true
			}
		}

		if !seen {
			out = append(out, i)
		}
	}

	// Convert output to something sensible
	outStrings := make([]string, len(out))
	for i, o := range out {
		outStrings[i] = s.ring[o].key
	}

	// Can return nodes in any order
	return outStrings
}
