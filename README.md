# Graphdeppicker
Graphdeppicker implements a weight-based probability picker which chooses dependencies on a graph.  

First the algorithm picks K nodes, then it picks N dependencies on those nodes.
Heavy nodes are more likely to be picked.

## Documentation

**Run** gets a graph and the amount of nodes. Return selected nodes

| Node | Weight | Dependency |
| ------------- |:-----------------:| -------------:|
| A | 4 | C |
| B | 5 | E, F |
| C | 6 | A, B, D |
| D | 7 | - |
| E | 2 | - |
| F | 1 | D, E |


```golang
g := graphll.New()
g.Add("a", 4, []string{"c"})
g.Add("b", 5, []string{"e", "f"})
g.Add("c", 6, []string{"a", "b", "d"})
g.Add("d", 7, []string{})
g.Add("e", 2, []string{})
g.Add("f", 1, []string{"d", "e"})

nodes, err := graphdeppicker.Run(g, 3) // nodes could be ["c", "b", "d"]
```