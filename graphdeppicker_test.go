package graphdeppicker

import (
	"testing"
	"fmt"
	
	"github.com/barbosaigor/graphll"
)

func TestChooseNodesByWeight(t *testing.T) {
	nodes := map[string]uint32{"a": 10, "b": 20, "c": 7, "d": 6}
	n, err := chooseNodesByWeight(nodes, 2)
	if err != nil {
		t.Error("ChooseNodesByWeight have returned an error: " + err.Error())
	}
	if len(n) != 2 {
		t.Error("ChooseNodesByWeight not returned correct number of elements: 2 != " + fmt.Sprint(len(n)))
	}

	for i := uint32(0); i < 500; i++ {
		n, err = chooseNodesByWeight(nodes, i % 10)
		fmt.Println(i, n, err)
		if err != nil {
			t.Error("ChooseNodesByWeight have returned an error: " + err.Error())
		}
		if (i % 10 > 4 && uint32(len(n)) != 4) || (i % 10 < 5 && uint32(len(n)) != i % 10) {
			t.Error(fmt.Sprintf("ChooseNodesByWeight not returned correct number of elements: %v != %v\n", i, len(n)))
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
	g.Add("d", 7, []string{})
	g.Add("e", 2, []string{})
	g.Add("f", 1, []string{"d", "e"})
	result, err := Run(g, 4)
	if err != nil {
		t.Error("Run returned an error, " + err.Error())
	}
	if len(result) != 4 {
		t.Error("Run should return 4 elements but got " + fmt.Sprint(len(result)))
	}
}