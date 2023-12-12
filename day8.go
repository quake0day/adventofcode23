package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	ID    string
	Left  *Node
	Right *Node
}

func CreateGraphFromFile(filename string) (map[string]*Node, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodes := make(map[string]*Node)
	var instructions string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			// Process node line
			parts := strings.Split(line, "=")
			nodeID := strings.TrimSpace(parts[0])
			children := strings.Split(strings.Trim(parts[1], " ()"), ",")
			leftChild := strings.TrimSpace(children[0])
			rightChild := strings.TrimSpace(children[1])

			// Initialize nodes if not already
			if _, ok := nodes[nodeID]; !ok {
				nodes[nodeID] = &Node{ID: nodeID}
			}
			if _, ok := nodes[leftChild]; !ok {
				nodes[leftChild] = &Node{ID: leftChild}
			}
			if _, ok := nodes[rightChild]; !ok {
				nodes[rightChild] = &Node{ID: rightChild}
			}

			// Set children
			nodes[nodeID].Left = nodes[leftChild]
			nodes[nodeID].Right = nodes[rightChild]
		} else if line != "" {
			// This line should contain the instructions
			instructions = line
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, "", err
	}

	return nodes, instructions, nil
}

func CreateGraph() map[string]*Node {
	nodes := make(map[string]*Node)

	// Initialize nodes
	ids := []string{"AAA", "BBB", "ZZZ"}
	for _, id := range ids {
		nodes[id] = &Node{ID: id}
	}

	// Set children for the provided example
	nodes["AAA"].Left, nodes["AAA"].Right = nodes["BBB"], nodes["BBB"]
	nodes["BBB"].Left, nodes["BBB"].Right = nodes["AAA"], nodes["ZZZ"]
	nodes["ZZZ"].Left, nodes["ZZZ"].Right = nodes["ZZZ"], nodes["ZZZ"]

	return nodes
}

func TraverseGraph(start *Node, instructions string) int {
	currentNode := start
	steps := 0

	for currentNode.ID != "ZZZ" {
		steps++
		instruction := instructions[(steps-1)%len(instructions)]

		if instruction == 'L' {
			currentNode = currentNode.Left
		} else { // Assuming 'R'
			currentNode = currentNode.Right
		}

		// Safeguard against nil nodes
		if currentNode == nil {
			fmt.Println("Error: Encountered a nil node. Check your graph mappings.")
			return -1
		}
	}

	return steps
}

func TraverseAll(graph map[string]*Node, instructions string) int {
	// Find all starting nodes (nodes ending with 'A')
	var currentNodes []*Node
	for _, node := range graph {
		if strings.HasSuffix(node.ID, "A") {
			currentNodes = append(currentNodes, node)
		}
	}

	steps := 0
	for {
		steps++
		var nextNodes []*Node
		instruction := instructions[(steps-1)%len(instructions)]

		// Move each current node according to the instruction
		for _, node := range currentNodes {
			if instruction == 'L' {
				nextNodes = append(nextNodes, node.Left)
			} else { // Assuming 'R'
				nextNodes = append(nextNodes, node.Right)
			}
		}

		// Update current nodes
		currentNodes = nextNodes

		// Check if all current nodes end with 'Z'
		allEndWithZ := true
		for _, node := range currentNodes {
			if !strings.HasSuffix(node.ID, "Z") {
				allEndWithZ = false
				break
			}
		}
		fmt.Println(steps)
		if allEndWithZ {
			return steps
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func findCycleLength(node *Node, instructions string) int {
	seen := make(map[string]int)
	steps := 0
	currentNode := node

	for {
		state := currentNode.ID + strconv.Itoa(steps%len(instructions))
		if pos, exists := seen[state]; exists {
			return steps - pos
		}
		seen[state] = steps
		steps++
		instruction := instructions[steps%len(instructions)]
		if instruction == 'L' {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
}

func TraverseAllWithLCM(graph map[string]*Node, instructions string) int {
	// Find cycle lengths for all starting nodes
	cycleLengths := []int{}
	for _, node := range graph {
		if strings.HasSuffix(node.ID, "A") {
			cycleLength := findCycleLength(node, instructions)
			cycleLengths = append(cycleLengths, cycleLength)
		}
	}

	// Calculate LCM of cycle lengths
	overallLCM := 1
	for _, length := range cycleLengths {
		overallLCM = lcm(overallLCM, length)
	}
	fmt.Println(overallLCM)
	// Traverse for LCM steps and check if all end with 'Z'
	return checkAtLCMStep(graph, instructions, overallLCM)
}

func checkAtLCMStep(graph map[string]*Node, instructions string, steps int) int {
	currentNodes := make(map[string]*Node)
	for _, node := range graph {
		if strings.HasSuffix(node.ID, "A") {
			currentNodes[node.ID] = node
		}
	}

	for i := 0; i < steps; i++ {
		nextNodes := make(map[string]*Node)
		instruction := instructions[i%len(instructions)]
		for _, node := range currentNodes {
			if instruction == 'L' {
				nextNodes[node.ID] = node.Left
			} else { // Assuming 'R'
				nextNodes[node.ID] = node.Right
			}
		}
		currentNodes = nextNodes
	}

	for _, node := range currentNodes {
		if !strings.HasSuffix(node.ID, "Z") {
			return -1 // Not all nodes end with 'Z'
		}
	}
	return steps
}

func main() {
	graph, instructions, err := CreateGraphFromFile("node.txt")
	fmt.Println(instructions)
	fmt.Println(graph)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	steps := TraverseGraph(graph["AAA"], instructions)
	fmt.Println("Steps to reach ZZZ:", steps)

	// part 2

	//steps2 := TraverseAll(graph, instructions)
	steps2 := TraverseAllWithLCM(graph, instructions)
	fmt.Println("Steps to reach all Z nodes:", steps2)
}
