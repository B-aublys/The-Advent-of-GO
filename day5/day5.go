package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	connected_nodes map[string]*Node
}

func (node Node) String() string {
	return "Hello"
}

func newNode() *Node {
	return &Node{make(map[string]*Node)}
}

type Graph struct {
	nodes map[string]*Node
}

func makeGraph() *Graph {
	return &Graph{make(map[string]*Node)}
}

func (gr *Graph) get_or_create(node string) *Node {
	fmt.Println(node)
	val, ok := gr.nodes[node]

	if !ok {
		val = newNode()
		gr.nodes[node] = val
	}

	return val
}

func (gr Graph) String() string {
	result := ""

	for key, val := range gr.nodes {
		result += fmt.Sprintf("[%s] -> ", key)
		for key, _ := range val.connected_nodes {
			result += key
			result += ", "
		}
		result += "\n"
	}
	return result
}

func main() {
	file, err := os.Open("short_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := makeGraph()

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")
		connected_nodes := strings.Split(split[1], " ")

		current_node := graph.get_or_create(split[0])

		for _, node := range connected_nodes[1:] {
			val2, ok := current_node.connected_nodes[node]

			if !ok {
				val2 = graph.get_or_create(node)
				current_node.connected_nodes[node] = val2
			}

			_, ok = val2.connected_nodes[split[0]]

			if !ok {
				val2.connected_nodes[split[0]] = current_node
			}
		}
	}

	fmt.Println(graph)

}
