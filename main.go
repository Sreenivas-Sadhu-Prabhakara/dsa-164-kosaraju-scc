package main

import (
	"fmt"
	"sort"
)

// SCCs returns the strongly connected components of a directed graph via
// Kosaraju's two passes: DFS finish order on the graph, then DFS on the
// transpose in reverse finish order — each second-pass tree is one SCC.
// Components and their members are sorted for deterministic output.
func SCCs(n int, edges [][2]int) [][]int {
	adj := make([][]int, n)
	radj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		radj[e[1]] = append(radj[e[1]], e[0])
	}

	// Pass 1: record finish order.
	visited := make([]bool, n)
	order := make([]int, 0, n)
	var fill func(v int)
	fill = func(v int) {
		visited[v] = true
		for _, w := range adj[v] {
			if !visited[w] {
				fill(w)
			}
		}
		order = append(order, v)
	}
	for v := 0; v < n; v++ {
		if !visited[v] {
			fill(v)
		}
	}

	// Pass 2: on the transpose, in reverse finish order.
	comp := make([]int, n)
	for i := range comp {
		comp[i] = -1
	}
	var gather func(v, c int, members *[]int)
	gather = func(v, c int, members *[]int) {
		comp[v] = c
		*members = append(*members, v)
		for _, w := range radj[v] {
			if comp[w] == -1 {
				gather(w, c, members)
			}
		}
	}
	out := [][]int{}
	for i := n - 1; i >= 0; i-- {
		v := order[i]
		if comp[v] == -1 {
			members := []int{}
			gather(v, len(out), &members)
			sort.Ints(members)
			out = append(out, members)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i][0] < out[j][0] })
	return out
}

func main() {
	fmt.Println(SCCs(5, [][2]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}}))
}
