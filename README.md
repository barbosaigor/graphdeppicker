# Graphdeppicker
Graphdeppicker implements an algorithm that picks dependencies from a dependency graph, 
taking into account the probability of each node.

The algorithm is performed in two main steps. 
First, the algorithm chooses K nodes, then it picks N dependencies on those nodes.
Heavy nodes are more likely to be chosen.

## Documentation

**Pick** gets a graph with the maximum amount of nodes that can be picked. 
The algorithm randomly chooses the nodes, taking into account the weight of each node, and returns a set of nodes.

| Node | Weight | Dependency |
| ------------- |:-----------------:| -------------:|
| A | 4 | C |
| B | 5 | E, F |
| C | 6 | A, B, D |
| D | 7 | - |
| E | 2 | - |
| F | 1 | D, E |


```golang
// First we have to create a dependency graph
g := graphll.New()
g.Add("a", 4, []string{"c"})
g.Add("b", 5, []string{"e", "f"})
g.Add("c", 6, []string{"a", "b", "d"})
g.Add("d", 7, nil)
g.Add("e", 2, nil)
g.Add("f", 1, []string{"d", "e"})

// Pick will execute the algorithm choosing up to three nodes
nodes, err := graphdeppicker.Pick(g, 3) // nodes could be ["c", "b", "d"]
```