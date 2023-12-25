package graph

import (
	"fmt"
)

type Graph struct {
	Vertices map[string]*Vertex
}

type Vertex struct {
	Val   string
	Edges map[string]*Edge
}

type Edge struct {
	Vertex *Vertex
	val    string
	dir    rune
}

func (g *Graph) AddVertex(key string, val string) {
	g.Vertices[key] = &Vertex{Val: val, Edges: map[string]*Edge{}}
}

func (g *Graph) AddEdge(srcKey string, destKey string, dir rune) {
	// add edge src --> dest
	g.Vertices[srcKey].Edges[destKey] = &Edge{
		Vertex: g.Vertices[destKey], val: destKey, dir: dir,
	}
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
	}
}

func (g *Graph) CreateNodeMap(list map[string][]string) {
	for vertex, edges := range list {
		if _, ok := g.Vertices[vertex]; !ok {
			g.AddVertex(vertex, vertex)
		}

		for i, edge := range edges {
			if _, ok := g.Vertices[edge]; !ok {
				g.AddVertex(edge, edge)
			}
			if i == 0 {
				g.AddEdge(vertex, edge, 'L')
			} else {
				g.AddEdge(vertex, edge, 'R')
			}

		}
	}
}

func (g *Graph) PrintVertices() {
	for _, v := range g.Vertices {
		fmt.Println(v.Val, " with edges ", v.Edges)
	}
}

type IsFinalNode func(currentNodeVal string) bool

func (g *Graph) NumOfInstructions(
	instructions []rune,
	startNode *Vertex,
	finalValueFunc IsFinalNode,
) int {
	counter := 0
	currNode := startNode
	for {
		for _, instruction := range instructions {
			if finalValueFunc(currNode.Val) {
				return counter
			}
			counter++
			for _, edgeVertex := range currNode.Edges {
				if edgeVertex.dir == instruction || len(currNode.Edges) == 1 {
					currNode = edgeVertex.Vertex
					break
				}
			}
		}
	}
}
