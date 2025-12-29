package main

import (
	"fmt"
	"os"
	"time"
)

type Graph map[int][]int

func DFSRekursif(g Graph, node int, visited map[int]bool) {
	visited[node] = true
	for _, neighbor := range g[node] {
		if !visited[neighbor] {
			DFSRekursif(g, neighbor, visited)
		}
	}
}

func DFSIteratif(g Graph, start int, n int) {
	visited := make(map[int]bool, n)
	stack := make([]int, 0, n)
	stack = append(stack, start)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[node] {
			continue
		}
		visited[node] = true

		for _, neighbor := range g[node] {
			if !visited[neighbor] {
				stack = append(stack, neighbor)
			}
		}
	}
}

func benchmarkDFS(g Graph, start int) (int64, int64) {
	const runs = 10000

	// warm-up
	{
		visited := make(map[int]bool)
		DFSRekursif(g, start, visited)
		DFSIteratif(g, start, len(g))
	}

	t0 := time.Now()
	for i := 0; i < runs; i++ {
		visited := make(map[int]bool)
		DFSRekursif(g, start, visited)
	}
	rec := time.Since(t0)

	t0 = time.Now()
	for i := 0; i < runs; i++ {
		DFSIteratif(g, start, len(g))
	}
	iter := time.Since(t0)

	return rec.Microseconds(), iter.Microseconds()
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah paket: ")
	fmt.Scan(&n)

	graph := make(Graph)
	for i := 0; i < n/2; i++ {
		l := 2*i + 1
		r := 2*i + 2
		if l < n {
			graph[i] = append(graph[i], l)
		}
		if r < n {
			graph[i] = append(graph[i], r)
		}
	}

	rec, iter := benchmarkDFS(graph, 0)

	fmt.Printf("DFS Rekursif : %d µs\n", rec)
	fmt.Printf("DFS Iteratif : %d µs\n", iter)

	// simpan ke CSV
	f, _ := os.OpenFile("data.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	fmt.Fprintf(f, "%d,%d,%d\n", n, rec, iter)
}