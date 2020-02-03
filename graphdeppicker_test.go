package graphdeppicker

import (
	"testing"
	"fmt"
	
	"github.com/barbosaigor/graphll"
)

func TestChooseNodesByWeight(t *testing.T) {
	g := graphll.New()
	g.Add("a", 4, []string{"c"})
	g.Add("b", 5, []string{"e", "f"})
	g.Add("c", 6, []string{"a", "b", "d"})
	g.Add("d", 7, []string{})
	g.Add("e", 2, []string{})
	g.Add("f", 1, []string{"d", "e"})

	for i := uint32(0); i < 100; i++ {
		n, err := chooseNodesByWeight(toWeightedMap(g), i % 10)
		// fmt.Println(i, fmt.Sprintf("{%v, %v}", i % 10, len(n)), n, err)
		if err != nil {
			t.Error("ChooseNodesByWeight have returned an error: " + err.Error())
		}
		if i % 10 > 6 && uint32(len(n)) != 6 {
			s := i % 10
			if s > 6 {
				s = 6
			}
			t.Error(fmt.Sprintf("%v: chooseNodesByWeight should return %v elements but got %v", i, s, len(n)))
		}
	}
}

func TestWeightSum(t *testing.T) {
	nodes := map[string]uint32{"a": 10, "b": 20, "c": 7, "d": 3}
	sum := weightSum(nodes)
	if sum != 40 {
		t.Error("Weight Sum should return 40 but got " + fmt.Sprint(sum))
	}
}

func TestMergeDeps(t *testing.T) {
	g := graphll.New()
	g.Add("a", 4, []string{"c"})
	g.Add("b", 5, []string{"e", "f"})
	g.Add("c", 6, []string{"a", "b", "d"})
	g.Add("d", 7, []string{})
	g.Add("e", 2, []string{})
	g.Add("f", 1, []string{"d", "e"})
	weightedNodes, err := mergeDeps(g, []string{"a", "c"})
	if err != nil {
		t.Error("MergeDeps returned an error, " + err.Error())
	}
	if len(weightedNodes) != 4 {
		t.Error("MergeDeps should return 4 elements but got " + fmt.Sprint(len(weightedNodes)))
	}
}

func TestRun(t *testing.T) {
	g := graphll.New()
	g.Add("a", 4, []string{"c"})
	g.Add("b", 5, []string{"e", "f"})
	g.Add("c", 6, []string{"a", "b", "d"})
	g.Add("d", 1, []string{})
	g.Add("e", 1, []string{})
	g.Add("f", 4, []string{"d", "e"})

	for i := uint32(0); i < 100; i++ {
		result, err := Run(g, i % 10)
		if err != nil {
			t.Error("Run returned an error, " + err.Error())
		}
		if i % 10 > 6 && uint32(len(result)) != 6 {
			s := i % 10
			if s > 6 {
				s = 6
			}
			t.Error(fmt.Sprintf("%v: Run should return %v elements but got %v", i, s, len(result)))
		}
	}
}