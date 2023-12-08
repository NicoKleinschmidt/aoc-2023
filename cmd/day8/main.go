package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Node struct {
	Id    string
	Left  *Node
	Right *Node
}

type StrNode struct {
	Left  string
	Right string
}

func main() {
	input, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	if len(lines) < 3 {
		return
	}

	sequence := lines[0]
	strNodes := parseNodes(lines[2:])
	nodes := buildTree(strNodes)
	stepCount := walkTree(nodes["AAA"], nodes["ZZZ"], sequence)
	stepCountGhosts := walkTreeGhost(nodes, sequence)

	fmt.Printf("Step count: %d\n", stepCount)
	fmt.Printf("Step count (ghosts 👻): %d", stepCountGhosts)
}

func parseNodes(lines []string) map[string]*StrNode {
	m := make(map[string]*StrNode, len(lines))
	for _, line := range lines {
		id, node := parseNode(line)
		m[id] = node
	}
	return m
}

func parseNode(line string) (id string, node *StrNode) {
	idEnd := strings.Index(line, "=")
	id = strings.TrimSpace(line[:idEnd])
	childList := strings.Trim(line[idEnd+1:], "() ")
	childSlice := strings.Split(childList, ",")

	return id, &StrNode{
		Left:  strings.TrimSpace(childSlice[0]),
		Right: strings.TrimSpace(childSlice[1]),
	}
}

// buildTree creates a tree structure from the nodes defined by strNodes.
// This functions also handles the case where strNodes defines multiple independent trees.
func buildTree(strNodes map[string]*StrNode) map[string]*Node {
	nodes := make(map[string]*Node, len(strNodes))
	var getNode func(id string) *Node

	getNode = func(id string) *Node {
		node := nodes[id]
		if node != nil {
			return node
		}

		node = &Node{Id: id}
		nodes[id] = node
		node.Left = getNode(strNodes[id].Left)
		node.Right = getNode(strNodes[id].Right)
		return node
	}

	for id := range strNodes {
		_, ok := nodes[id]
		if !ok {
			nodes[id] = getNode(id)
		}
	}

	return nodes
}

func walkTree(start, end *Node, sequence string) (steps int) {
	sequenceLen := len(sequence)

	for {
		if start == end {
			return steps
		}

		right := sequence[steps%sequenceLen] == 'R'

		if right {
			start = start.Right
		} else {
			start = start.Left
		}

		steps++
	}
}

func walkTreeGhost(nodes map[string]*Node, sequence string) (totalSteps int) {
	var startingNodes []*Node

	for id, node := range nodes {
		if strings.HasSuffix(id, "A") {
			startingNodes = append(startingNodes, node)
		}
	}

	stepsPerThread := make([]int, len(startingNodes))
	sequenceLen := len(sequence)

	for i, node := range startingNodes {
		steps := 0

		for {
			right := sequence[steps%sequenceLen] == 'R'

			if right {
				node = node.Right
			} else {
				node = node.Left
			}

			steps++

			if strings.HasSuffix(node.Id, "Z") {
				break
			}
		}

		stepsPerThread[i] = steps
	}

	return lcmS(stepsPerThread)
}

func lcmS(s []int) int {
	i := 1
	for _, j := range s {
		i = lcm(i, j)
	}
	return i
}

func lcm(i, j int) int {
	return (i * j) / gcd(i, j)
}

func gcd(i, j int) int {
	if j == 0 {
		return i
	}
	return gcd(j, i%j)
}
