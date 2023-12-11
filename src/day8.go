package main

import (
	"fmt"
	"strings"
)

func Day8RunA() {
	fmt.Println("Running Day 8 Part A")

	lines := ReadLines("../Data/day8.txt")

	directions := lines[0]
	graph := NewTreeNodeDescriptionGraph(lines[2:])

	count := graph.TraverseToEnd(directions)

	fmt.Printf("Steps To Reach Graph End = %v\n", count)
}

func Day8RunB() {
	fmt.Println("Running Day 8 Part B")

	lines := ReadLines("../Data/day8.txt")

	directions := lines[0]
	graph := NewTreeNodeDescriptionGraph(lines[2:])

	count := graph.AdvancedMultiTraverseToEnd(directions)

	fmt.Printf("Steps To Reach Graph End = %v\n", count)
}

type TreeNodeDescription struct {
	value, left, right string
}

type TreeNodeDescriptionGraph struct {
	nodes map[string]TreeNodeDescription
}

// AAA = (BBB, CCC)
func NewTreeNodeDescription(node string) TreeNodeDescription {

	parts := strings.Split(node, " = ")

	value := parts[0]

	nodeParts := strings.Split(strings.TrimSuffix(strings.TrimPrefix(parts[1], "("), ")"), ", ")

	return TreeNodeDescription{value: value, left: nodeParts[0], right: nodeParts[1]}
}

func NewTreeNodeDescriptionGraph(nodes []string) *TreeNodeDescriptionGraph {

	g := &TreeNodeDescriptionGraph{nodes: map[string]TreeNodeDescription{}}

	for _, node := range nodes {
		desc := NewTreeNodeDescription(node)

		g.nodes[desc.value] = desc
	}

	return g
}

func (g *TreeNodeDescriptionGraph) TraverseToEnd(directions string) int {

	atEndNode := func(node string) bool {
		return node == "ZZZ"
	}

	return g.TraverseToEndFunc(directions, "AAA", atEndNode)
}

func (g *TreeNodeDescriptionGraph) TraverseToEndFunc(directions string, startNode string, atEndNode func(string) bool) int {
	current := startNode
	count := 0

	for !atEndNode(current) {

		switch directions[count%len(directions)] {

		case 'L':
			current = g.nodes[current].left
		case 'R':
			current = g.nodes[current].right
		}

		count++
	}

	return count
}

func (g *TreeNodeDescriptionGraph) AdvancedMultiTraverseToEnd(directions string) int {

	startNodes, _ := g.FindStartAndEndNodes()

	counts := []int{}

	atEndNode := func(node string) bool {
		return strings.HasSuffix(node, "Z")
	}

	for _, node := range startNodes {
		counts = append(counts, g.TraverseToEndFunc(directions, node, atEndNode))
	}

	return LCMM(counts)
}

func (g *TreeNodeDescriptionGraph) BasicMultiTraverseToEnd(directions string) int {

	startNodes, _ := g.FindStartAndEndNodes()

	count := 0

	currentNodes := startNodes

	for !AreAllEndNodes(currentNodes) {

		nextNodes := []string{}

		for _, current := range currentNodes {

			switch directions[count%len(directions)] {

			case 'L':
				nextNodes = append(nextNodes, g.nodes[current].left)
			case 'R':
				nextNodes = append(nextNodes, g.nodes[current].right)
			}
		}

		count++

		currentNodes = nextNodes
	}

	return count
}

func AreAllEndNodes(currentNodes []string) bool {

	for _, node := range currentNodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}

	return true
}

func (g *TreeNodeDescriptionGraph) FindStartAndEndNodes() ([]string, []string) {

	keys := Keys(g.nodes)

	startNodes := []string{}
	endNodes := []string{}

	for _, k := range keys {
		if strings.HasSuffix(k, "A") {
			startNodes = append(startNodes, k)
		}

		if strings.HasSuffix(k, "Z") {
			endNodes = append(endNodes, k)
		}
	}

	return startNodes, endNodes
}

func GCD(a, b int) int {

	for b > 0 {

		a, b = b, a%b
	}

	return a
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

func LCMM(xs []int) int {

	m := xs[0]

	for i := 1; i < len(xs); i++ {
		m = LCM(m, xs[i])
	}

	return m
}
