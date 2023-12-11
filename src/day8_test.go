package main

import "testing"

func TestNewTreeDescription(t *testing.T) {

	desc := []string{"AAA = (BBB, CCC)", "DEZ = (ZYX, CBB)", "BBB = (GGG, FFF)"}

	expected := []TreeNodeDescription{{value: "AAA", left: "BBB", right: "CCC"}, {value: "DEZ", left: "ZYX", right: "CBB"},
		{value: "BBB", left: "GGG", right: "FFF"}}

	for i := range desc {
		actual := NewTreeNodeDescription(desc[i])

		if expected[i] != actual {
			t.Errorf("Expected = %v :: Got = %v", expected[i], actual)
			return
		}
	}
}

func TestCreateTreeDescriptionMap(t *testing.T) {

	nodes := []string{"AAA = (BBB, CCC)", "BBB = (DDD, EEE)", "CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)", "EEE = (EEE, EEE)", "GGG = (GGG, GGG)", "ZZZ = (ZZZ, ZZZ)"}

	actual := NewTreeNodeDescriptionGraph(nodes)

	if len(actual.nodes) != len(nodes) {
		t.Errorf("Size of map does not match size of input node list")
		return
	}

	if actual.nodes["AAA"] != NewTreeNodeDescription("AAA = (BBB, CCC)") {
		t.Errorf("Expected AAA = BBB, CCC :: Got %v", actual.nodes["AAA"])
		return
	}

	if actual.nodes["BBB"] != NewTreeNodeDescription("BBB = (DDD, EEE)") {
		t.Errorf("Expected BBB = DDD, EEE :: Got %v", actual.nodes["BBB"])
		return
	}
}

func TestBasicTraverseToEnd(t *testing.T) {

	expected := 2
	directions := "RL"
	nodes := []string{"AAA = (BBB, CCC)", "BBB = (DDD, EEE)", "CCC = (ZZZ, GGG)",
		"DDD = (DDD, DDD)", "EEE = (EEE, EEE)", "GGG = (GGG, GGG)", "ZZZ = (ZZZ, ZZZ)"}

	graph := NewTreeNodeDescriptionGraph(nodes)

	actual := graph.TraverseToEnd(directions)

	if actual != expected {
		t.Errorf("Wrong Number of Steps: Expected %v :: Got %v", expected, actual)
		return
	}
}

func TestLongerTraverseToEnd(t *testing.T) {

	expected := 6
	directions := "LLR"
	nodes := []string{"AAA = (BBB, BBB)", "BBB = (AAA, ZZZ)", "ZZZ = (ZZZ, ZZZ)"}

	graph := NewTreeNodeDescriptionGraph(nodes)

	actual := graph.TraverseToEnd(directions)

	if actual != expected {
		t.Errorf("Wrong Number of Steps: Expected %v :: Got %v", expected, actual)
		return
	}
}

func TestMultiTraverseToEnd(t *testing.T) {
	directions := "LR"
	nodes := []string{"11A = (11B, XXX)", "11B = (XXX, 11Z)", "11Z = (11B, XXX)", "22A = (22B, XXX)",
		"22B = (22C, 22C)", "22C = (22Z, 22Z)", "22Z = (22B, 22B)", "XXX = (XXX, XXX)"}
	expected := 6

	graph := NewTreeNodeDescriptionGraph(nodes)

	actual := graph.BasicMultiTraverseToEnd(directions)

	if actual != expected {
		t.Errorf("Wrong Number of Steps: Expected %v :: Got %v", expected, actual)
		return
	}
}

func TestGCD(t *testing.T) {

	numbers := [][]int{{10, 15, 5}, {143, 227, 1}, {40, 32, 8}, {42, 56, 14}}

	for _, ar := range numbers {

		got := GCD(ar[0], ar[1])
		if ar[2] != got {
			t.Errorf("Incorrect Calculation: Expected GCD(%v, %v) = %v :: Got %v", ar[0], ar[1], ar[2], got)
			return
		}
	}
}

func TestLCM(t *testing.T) {
	numbers := [][]int{{4, 6, 12}, {18, 15, 90}} //, {143, 227, 1}, {40, 32, 8}, {42, 56, 14}}

	for _, ar := range numbers {

		got := LCM(ar[0], ar[1])
		if ar[2] != got {
			t.Errorf("Incorrect Calculation: Expected LCM(%v, %v) = %v :: Got %v", ar[0], ar[1], ar[2], got)
			return
		}
	}
}

type LcmmTestCase struct {
	xs       []int
	expected int
}

func TestLCMM(t *testing.T) {

	tests := []LcmmTestCase{{xs: []int{100, 23, 98}, expected: 112700}}

	for _, tt := range tests {

		got := LCMM(tt.xs)
		if tt.expected != got {
			t.Errorf("Incorrect Calculation: Expected %v :: Got %v", tt.expected, got)
			return
		}
	}
}
