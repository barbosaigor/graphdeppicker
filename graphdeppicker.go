package graphdeppicker

import (
	"math"

	"github.com/barbosaigor/graphll"
	"github.com/barbosaigor/picker"
)

func weightSum(buckets map[string]uint32) uint32 {
	s := uint32(0)
	for _, weight := range buckets {
		s += weight
	}
	return s
}

// chooseNodesByWeight chooses randomly n nodes taking into account the weight of each element
func chooseNodesByWeight(buckets map[string]uint32, size uint32) (map[string]uint32, error) {
	if uint32(len(buckets)) < size {
		size = uint32(len(buckets))
	}
	nodes := make(map[string]uint32, size)
	p := picker.New()
	for i := uint32(0); i < size; i++ {
		for bucket, weight := range buckets {
			p.Add(bucket, weight)
		}
		node, err := p.RollDice()
		if err != nil {
			return nil, err
		}
		// If has a node then repeat iteration
		if _, ok := nodes[node]; !ok {
			nodes[node] = buckets[node]
		} else {
			i--
		}
		p.Reset()
	}
	return nodes, nil
}

// toWeightedMap transforms a graph to a map
func toWeightedMap(graph graphll.GraphLL) map[string]uint32 {
	m := make(map[string]uint32, len(graph))
	for bucket := range graph {
		w, _ := graph.Weight(bucket)
		m[bucket] = w
	}
	return m
}

// mergeDeps merges dependencies of passed nodes
func mergeDeps(graph graphll.GraphLL, nodes []string) (map[string]uint32, error) {
	deps := make(map[string]uint32, len(nodes))
	for _, node := range nodes {
		dps, err := graph.Deps(node)
		if err != nil {
			return nil, err
		}
		for _, dep := range dps {
			w, err := graph.Weight(dep)
			if err != nil {
				return nil, err
			}
			deps[dep] = w
		}
	}
	return deps, nil
}

func toStrSlice(nodes map[string]uint32) []string {
	s := make([]string, 0, len(nodes))
	for bucket := range nodes {
		s = append(s, bucket)
	}
	return s
}

// Pick gets a graph with the maximum amount of nodes that can be picked.
// The algorithm randomly chooses the nodes, taking into account the weight of each node,
// and returns a set of nodes.
func Pick(graph graphll.GraphLL, size uint32) ([]string, error) {
	weightedNodes := toWeightedMap(graph)
	if size > uint32(len(graph)) {
		size = uint32(len(graph))
	}
	// Using an euristic for size, is possible to
	// increase the initial scope
	fstSize := uint32(math.Floor(float64(size) + float64(size)*float64(0.5)))
	if fstSize > uint32(len(graph)) {
		fstSize = size
	}
	// Choose nodes
	weightedNodes, err := chooseNodesByWeight(weightedNodes, fstSize)
	if err != nil {
		return nil, err
	}
	// Get all deps of picked nodes
	weightedNodes, err = mergeDeps(graph, toStrSlice(weightedNodes))
	if err != nil {
		return nil, err
	}
	// Choose nodes dependencies
	weightedNodes, err = chooseNodesByWeight(weightedNodes, size)
	if err != nil {
		return nil, err
	}
	return toStrSlice(weightedNodes), nil
}
