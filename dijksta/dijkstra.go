package main

import (
	"container/heap"
	"fmt"
)

type Vertext int

type PriorityQueue struct {
	items []Vertext
	// value to index
	m map[Vertext]int
	// value to priority
	pr map[Vertext]int
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.pr[pq.items[i]] < pq.pr[pq.items[j]]
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.m[pq.items[i]] = i
	pq.m[pq.items[j]] = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(Vertext)
	pq.m[item] = n
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.m[item] = -1
	pq.items = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item Vertext, priority int) {
	pq.pr[item] = priority
	heap.Fix(pq, pq.m[item])
}

func (pq *PriorityQueue) addWithPriority(item Vertext, priority int) {
	heap.Push(pq, item)
	pq.update(item, priority)
}

const (
	Infinity      = int(^uint(0) >> 1)
	Uninitialized = -1
)

func Dijkstra(G Graph, source Vertext) (distance map[Vertext]int, previous map[Vertext]Vertext) {
	distance = make(map[Vertext]int)
	previous = make(map[Vertext]Vertext)

	distance[source] = 0

	q := &PriorityQueue{[]Vertext{}, make(map[Vertext]int), make(map[Vertext]int)}

	for _, v := range G.Vertices() {
		if v != source {
			distance[v] = Infinity
		}
		previous[v] = Uninitialized
		q.addWithPriority(v, distance[v])
	}

	for len(q.items) != 0 {
		u := heap.Pop(q).(Vertext)
		for _, v := range G.Neighbors(u) {
			alt := distance[u] + G.Weight(u, v)
			if alt < distance[v] {
				distance[v] = alt
				previous[v] = u
				q.update(v, alt)
			}
		}
	}

	return distance, previous
}

type Graph interface {
	Vertices() []Vertext
	Neighbors(v Vertext) []Vertext
	Weight(u, v Vertext) int
}

type AdjacencyMap struct {
	ids   map[string]Vertext
	names map[Vertext]string
	edges map[Vertext]map[Vertext]int
}

func NewGraph(ids map[string]Vertext) AdjacencyMap {
	G := AdjacencyMap{ids: ids}
	G.names = make(map[Vertext]string)
	for k, v := range ids {
		G.names[v] = k
	}
	G.edges = make(map[Vertext]map[Vertext]int)
	return G
}

func (G AdjacencyMap) AddEdge(u, v string, w int) {
	if _, ok := G.edges[G.ids[u]]; !ok {
		G.edges[G.ids[u]] = make(map[Vertext]int)
	}
	G.edges[G.ids[u]][G.ids[v]] = w
}

func (G AdjacencyMap) GetPath(v Vertext, previous map[Vertext]Vertext) (path string) {
	path = G.names[v]
	for previous[v] >= 0 {
		v = previous[v]
		path = G.names[v] + path
	}
	return path
}

func (G AdjacencyMap) Vertices() (vertices []Vertext) {
	for _, v := range G.ids {
		vertices = append(vertices, v)
	}
	return vertices
}

func (G AdjacencyMap) Neighbors(u Vertext) (vertices []Vertext) {
	for v := range G.edges[u] {
		vertices = append(vertices, v)
	}
	return vertices
}

func (G AdjacencyMap) Weight(u, v Vertext) int {
	return G.edges[u][v]
}

func main() {
	fmt.Println(2 / 11)

	/*
		G := NewGraph(map[string]Vertext{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": 4,
			"e": 5,
			"f": 6,
		})

		G.AddEdge("a", "b", 7)
		G.AddEdge("a", "c", 9)
		G.AddEdge("a", "f", 14)
		G.AddEdge("b", "c", 10)
		G.AddEdge("c", "d", 11)
		G.AddEdge("c", "f", 2)
		G.AddEdge("d", "e", 6)
		G.AddEdge("e", "f", 9)

		distance, previous := Dijkstra(G, G.ids["a"])
		fmt.Printf("Distance to %s is %d, Path: %s\n", "e", distance[G.ids["e"]], G.GetPath(G.ids["e"], previous))
		fmt.Printf("Distance to %s is %d, Path: %s\n", "e", distance[G.ids["f"]], G.GetPath(G.ids["f"], previous))
	*/
}
