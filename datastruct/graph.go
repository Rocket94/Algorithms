package datastruct

import "fmt"

type Graph struct {
	v   int     //顶点个数
	adj [][]int //邻接矩阵
}

func NewGraph(v int) *Graph {
	g := &Graph{
		v:   v,
		adj: [][]int{},
	}
	for i := 0; i < v; i++ {
		g.adj = append(g.adj, []int{})
	}
	return g
}

//无向图添加边
func (g *Graph) AddEdge(s, t int) {
	g.adj[s] = append(g.adj[s], t)
	g.adj[t] = append(g.adj[t], s)
}

//广度优先搜索
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}
	//用于记录顶点是否被访问过
	visited := make([]bool, g.v)
	visited[s] = true
	//创建一个遍历顶点的队列，遍历到了就入队，遍历完了就出队
	queue := NewQueue()
	queue.Enqueue(s)
	//用于记录路径，表示该点前一个点是什么
	pre := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		pre[i] = -1
	}
	for !queue.IsEmpty() {
		w := queue.Dequeue().(int)
		for i := 0; i < len(g.adj[w]); i++ {
			q := g.adj[w][i]
			if !visited[q] {
				//记录路径
				pre[q] = w
				//如果找到了t那就返回并且打印路径
				if q == t {
					printPath(pre, s, t)
					return
				}
				//没找到t，就把这个顶点记录，并且入队
				visited[q] = true
				queue.Enqueue(q)
			}
		}
	}
}

func printPath(pre []int, s, t int) {
	if pre[t] != -1 && t != s {
		printPath(pre, s, pre[t])
	}
	fmt.Print(t, " ")
}

//深度优先搜索
var found = false

func (g *Graph) DFS(s, t int) {
	found=false
	visited:=make([]bool,g.v)
	prev:=make([]int,g.v)
	for i:=0;i<len(prev);i++{
		prev[i]=-1
	}
	g.recurDFS(s,t,visited,prev)
	printPath(prev,s,t)
}

func (g *Graph) recurDFS(w, t int, visited []bool, prev []int) {
	if found {
		return
	}
	visited[w] = true
	if w == t {
		found = true
		return
	}
	for i:=0;i<len(g.adj[w]);i++{
		q:=g.adj[w][i]
		if !visited[q]{
			prev[q]=w
			g.recurDFS(q,t,visited,prev)
		}
	}

}
