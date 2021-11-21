package datastruct

import "testing"

func TestGraph_BFS(t *testing.T) {
	g:=makeGraph()
	g.BFS(0,7)
}

func TestGraph_DFS(t *testing.T) {
	g:=makeGraph()
	g.DFS(0,7)
}

func makeGraph() *Graph {
	g:=NewGraph(8)
	g.adj[0]=append(g.adj[0], 1)
	g.adj[0]=append(g.adj[0], 3)
	g.adj[1]=append(g.adj[1], 0)
	g.adj[1]=append(g.adj[1], 2)
	g.adj[1]=append(g.adj[1], 4)
	g.adj[2]=append(g.adj[2], 1)
	g.adj[2]=append(g.adj[2], 5)
	g.adj[3]=append(g.adj[3], 0)
	g.adj[3]=append(g.adj[3], 4)
	g.adj[4]=append(g.adj[4], 1)
	g.adj[4]=append(g.adj[4], 3)
	g.adj[4]=append(g.adj[4], 5)
	g.adj[4]=append(g.adj[4], 6)
	g.adj[5]=append(g.adj[5], 2)
	g.adj[5]=append(g.adj[5], 4)
	g.adj[5]=append(g.adj[5], 7)
	g.adj[6]=append(g.adj[6], 4)
	g.adj[6]=append(g.adj[6], 7)
	g.adj[7]=append(g.adj[7], 5)
	g.adj[7]=append(g.adj[7], 6)
	return g
}
