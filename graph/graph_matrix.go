package main

import (
	"errors"
	"fmt"
)

type GraphType string

const (
	DIRECTED   GraphType = "DIRECTED"
	UNDIRECTED GraphType = "UNDIRECTED"
)

type Graph interface {
	Init()
	AddEdge(vertexOne int, vertexTwo int) error
	AddEdgeWithWeight(vertexOne int, vertexTwo int) error
	RemoveEdge(vertexOne int, vertexTwo int) error
	HasEdge(vertexOne int, vertexTwo int) bool
	GetGraphType() GraphType
	GetAdjacentNodesForVertex(vertex int) map[int]bool
	GetWeightOfEdge(vertexOne int, vertexTwo int) (int, error)
	GetNumberOfVertices() int
	GetNumberOfEdges() int
	GetIndegreeForVertex(vertex int) int
}

type AdjacencyMatrix struct {
	Vertices  int
	Edges     int
	GraphType GraphType
	AdjMatrix [][]int
}

func (G *AdjacencyMatrix) Init() {
	G.AdjMatrix = make([][]int, G.Vertices)
	G.Edges = 0
	for i := 0; i < G.Vertices; i++ {
		G.AdjMatrix[i] = make([]int, G.Vertices) // default initialization is 0
	}
}

func (G *AdjacencyMatrix) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("index out of bounds")
	}

	G.AdjMatrix[vertexOne][vertexTwo] = 1
	G.Edges++
	if G.GraphType == UNDIRECTED {
		G.AdjMatrix[vertexTwo][vertexOne] = 1
		G.Edges++
	}

	return nil
}

func (G *AdjacencyMatrix) AddEdgeWithWeight(vertexOne, vertexTwo, weight int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("index out of bounds")
	}

	G.AdjMatrix[vertexOne][vertexTwo] = weight
	G.Edges++
	if G.GraphType == UNDIRECTED {
		G.AdjMatrix[vertexTwo][vertexOne] = weight
		G.Edges++
	}
	return nil
}

func (G *AdjacencyMatrix) RemoveEdge(vertexOne, vertexTwo int) error {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("index out of bounds")
	}

	G.AdjMatrix[vertexOne][vertexTwo] = 0
	G.Edges--
	if G.GraphType == UNDIRECTED {
		G.AdjMatrix[vertexTwo][vertexOne] = 0
		G.Edges--
	}
	return nil
}
func (G *AdjacencyMatrix) HasEdge(vertexOne, vertexTwo int) bool {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}

	return G.AdjMatrix[vertexOne][vertexTwo] != 0
}

func (G *AdjacencyMatrix) GetGraphType() GraphType {
	return G.GraphType
}

func (G *AdjacencyMatrix) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	adjacencyMatrixVertices := map[int]bool{}
	if vertex > G.Vertices || vertex < 0 {
		return adjacencyMatrixVertices
	}

	for i := 0; i < G.Vertices; i++ {
		if G.AdjMatrix[vertex][i] != 0 {
			adjacencyMatrixVertices[i] = (G.AdjMatrix[vertex][i] != 0)
		}
	}

	return adjacencyMatrixVertices
}

func (G *AdjacencyMatrix) GetWeightOfEdge(vertexOne, vertexTwo int) (int, error) {
	if vertexOne >= G.Vertices || vertexTwo >= G.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0, errors.New("Error getting weight for vertex")
	}
	return G.AdjMatrix[vertexOne][vertexTwo], nil
}

func (G *AdjacencyMatrix) GetNumberOfVertices() int {
	return G.Vertices
}

func (G *AdjacencyMatrix) GetNumberOfEdges() int {
	return G.Edges
}

func (G *AdjacencyMatrix) GetIndegreeForVertex(vertex int) int {
	indegree := 0
	adjacentNodes := G.GetAdjacentNodesForVertex(vertex)
	for key := range adjacentNodes {
		if adjacentNodes[key] {
			indegree++
		}
	}
	return indegree
}

func main() {
	var test = &AdjacencyMatrix{4, 0, DIRECTED, nil}
	test.Init()
	err := test.AddEdge(2, 1)
	if err != nil {
		fmt.Printf("Error adding edge")
	}
	if test.AdjMatrix[2][1] != 1 {
		fmt.Printf("Data not found at index")
	}
	if test.AdjMatrix[1][2] != 0 {
		fmt.Printf("Data not found at index")
	}

	fmt.Println(test)
}
