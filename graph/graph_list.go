package main

import (
	"errors"
	"fmt"
)

//type GraphType string

type AdjacencyList struct {
	Vertices  int
	Edges     int
	GraphType GraphType
	AdjList   []*Node
}

type Node struct {
	Next   *Node
	Weight int
	Key    int
}

func (node Node) AddNode(value int) *Node {
	n := node.Next
	if n == nil {
		newNode := &Node{Next: &Node{}, Key: value}
		return newNode
	}
	nd := n.AddNode(value)
	node.Next = nd
	return &node
}

func (node Node) AddNodeWithWeight(value, weight int) *Node {
	n := node.Next
	if n == nil {
		newNode := &Node{Next: &Node{}, Key: value, Weight: weight}
		return newNode
	}
	nd := n.AddNodeWithWeight(value, weight)
	node.Next = nd
	return &node
}

func (node Node) FindNextNode(key int) (*Node, error) {
	n := node
	if n == (Node{}) {
		return &Node{}, errors.New("Node not found")
	}
	if n.Key == key {
		return &n, nil
	}
	nd := n.Next
	return nd.FindNextNode(key)
}

func (G *AdjacencyList) Init() {
	G.AdjList = make([]*Node, G.Vertices)
	G.Edges = 0
	for i := 0; i < G.Vertices; i++ {
		G.AdjList[i] = &Node{}
	}
}

func (G *AdjacencyList) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	node := G.AdjList[vertexOne].AddNode(vertexTwo)
	G.AdjList[vertexOne] = node
	G.Edges++
	if G.GraphType == UNDIRECTED {
		node := G.AdjList[vertexTwo].AddNode(vertexOne)
		G.AdjList[vertexTwo] = node
		G.Edges++
	}
	return nil
}

func (G *AdjacencyList) AddEdgeWithWeight(vertexOne, vertexTwo, weight int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	node := G.AdjList[vertexOne].AddNodeWithWeight(vertexTwo, weight)
	G.AdjList[vertexOne] = node
	G.Edges++
	if G.GraphType == UNDIRECTED {
		node := G.AdjList[vertexTwo].AddNodeWithWeight(vertexOne, weight)
		G.AdjList[vertexTwo] = node
		G.Edges++
	}
	return nil
}

func (G *AdjacencyList) RemoveEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	nodeAdj := G.AdjList[vertexOne]
	if nodeAdj == (&Node{}) {
		return errors.New("Node not found")
	}
	nextNode := nodeAdj
	newNodes := Node{}
	for nextNode != (&Node{}) && nextNode != nil {
		if nextNode.Key != vertexTwo {
			newNodes.Next = nextNode
			newNodes.Key = nextNode.Key
			newNodes.Weight = nextNode.Weight
			nextNode = nextNode.Next
		} else {
			newNodes.Next = nextNode.Next
			newNodes.Key = nextNode.Next.Key
			newNodes.Weight = nextNode.Next.Weight
			G.AdjList[vertexOne] = &newNodes
			G.Edges--
			return nil
		}
	}
	G.Edges--
	return nil
}

func (G *AdjacencyList) HasEdge(vertexOne, vertexTwo int) bool {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}

	nodeAdj := G.AdjList[vertexOne]
	if nodeAdj == (&Node{}) {
		return false
	}
	node, _ := nodeAdj.FindNextNode(vertexTwo)
	if node != nil && node.Key == vertexTwo {
		return true
	}
	return false
}

func (G *AdjacencyList) GetGraphType() GraphType {
	return G.GraphType
}

func (G *AdjacencyList) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	if vertex >= G.Vertices || vertex < 0 {
		return map[int]bool{}
	}
	nodeAdj := G.AdjList[vertex]
	nextNode := nodeAdj
	nodes := map[int]bool{}
	for nextNode != (&Node{}) && nextNode != nil {
		nodes[nextNode.Key] = true
		nextNode = nextNode.Next
	}
	return nodes
}

func (G *AdjacencyList) GetWeightOfEdge(vertexOne, vertexTwo int) (int, error) {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0, errors.New("Error getting weight for vertex")
	}
	nodeAdj := G.AdjList[vertexOne]
	if nodeAdj == (&Node{}) {
		return 0, errors.New("Error getting weight for vertex")
	}
	node, _ := nodeAdj.FindNextNode(vertexTwo)
	if node != nil && node.Key == vertexTwo {
		return nodeAdj.Weight, nil
	}
	return 0, errors.New("Error getting weight for vertex")
}

func (G *AdjacencyList) GetNumberOfVertices() int {
	return G.Vertices
}

func (G *AdjacencyList) GetNumberOfEdges() int {
	return G.Edges
}

func (G *AdjacencyList) GetIndegreeForVertex(vertex int) int {
	if vertex >= G.Vertices || vertex < 0 {
		return 0
	}

	nodeAdj := G.AdjList[vertex]
	nextNode := nodeAdj.Next
	length := 0
	for nextNode != (&Node{}) && nextNode != nil {
		length += 1
		nextNode = nextNode.Next
	}
	return length
}

func DFS(G *AdjacencyList) []int {
	visited := map[int]bool{}
	var dfsOrdered []int
	for i := 0; i < G.GetNumberOfVertices(); i++ {
		dfs(G, visited, i, &dfsOrdered)
	}
	return dfsOrdered
}

func dfs(G *AdjacencyList, visited map[int]bool, current int, dfsOrdered *[]int) {
	if visited[current] {
		return
	}
	visited[current] = true
	adjNodes := G.GetAdjacentNodesForVertex(current)
	for key, _ := range adjNodes {
		dfs(G, visited, key, dfsOrdered)
	}
	*dfsOrdered = append(*dfsOrdered, current)
	return
}

func BFS(G *AdjacencyList) []int {
	visited := map[int]bool{}
	var bfsArrayOrdered []int
	for i := 0; i < G.GetNumberOfVertices(); i++ {
		bfsArrayOrdered = bfs(G, visited, i, bfsArrayOrdered)
	}
	return bfsArrayOrdered
}

func bfs(G *AdjacencyList, visited map[int]bool, node int, bfsArrayOrdered []int) []int {
	queue := make([]int, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if visited[current] {
			continue
		}
		bfsArrayOrdered = append(bfsArrayOrdered, current)
		visited[current] = true
		adjNodes := G.GetAdjacentNodesForVertex(current)
		for key, _ := range adjNodes {
			queue = append(queue, key)
		}
	}
	return bfsArrayOrdered
}

func main() {
	G := &AdjacencyList{4, 0, DIRECTED, nil}
	G.Init()
	G.AddEdge(2, 1)
	G.AddEdge(3, 2)
	G.AddEdge(2, 0)
	G.AddEdge(1, 0)
	G.AddEdge(1, 1)
	G.AddEdge(2, 3)
	fmt.Printf("Directed Graph with DFS is %v", DFS(G))
	fmt.Printf("Directed Graph with BFS is %v", BFS(G))
	/*
		var testAdjListDirected = &AdjacencyList{4, 0, DIRECTED, nil}
		//var testAdjListUnDirected = &AdjacencyList{4, 0, UNDIRECTED, nil}

		testAdjListDirected.Init()
		testAdjListDirected.AddEdge(2, 1)

		err := testAdjListDirected.AddEdge(2, 3)
		if err != nil {
			fmt.Printf("Error adding edge")
		}
		if testAdjListDirected.AdjList[2].Key != 1 {
			fmt.Printf("Data not found at index")
		}
		if testAdjListDirected.AdjList[2].Next.Key != 3 {
			fmt.Printf("Data not found at index")
		}

		fmt.Println(testAdjListDirected)
	*/
}
