// Package graphdeppicker implements a weight-based probability picker,
// which chooses dependencies on a graph.
//
// First the algorithm picks K nodes, then it picks N dependencies on those nodes.
// Heavy nodes are more likely to be picked.
package graphdeppicker
