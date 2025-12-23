package main
import (
	"fmt"
	"time"
)

type Graph map[int][]int

// DFS REKURSIF
func DFSRekursif(g Graph, node int, visited map[int]bool) {
	visited[node] = true
	for _, neighbor := range g[node] {
		if !visited[neighbor] {
			DFSRekursif(g, neighbor, visited)
		}
	}
}

// DFS ITERATIF
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

func benchmarkDFS(g Graph, start int) {
	const runs = 10000
	const batches = 5

	{
		visited := make(map[int]bool)
		DFSRekursif(g, start, visited)
		DFSIteratif(g, start, len(g))
	}

	bestRec := time.Duration(1<<63 - 1)
	bestIter := time.Duration(1<<63 - 1)

	for b := 0; b < batches; b++ {
		t0 := time.Now()
		for i := 0; i < runs; i++ {
			visited := make(map[int]bool)
			DFSRekursif(g, start, visited)
		}
		rec := time.Since(t0)
		if rec < bestRec {
			bestRec = rec
		}

		t0 = time.Now()
		for i := 0; i < runs; i++ {
			DFSIteratif(g, start, len(g))
		}
		iter := time.Since(t0)
		if iter < bestIter {
			bestIter = iter
		}
	}

	fmt.Printf("DFS Rekursif : %d ms\n", bestRec.Microseconds())
	fmt.Printf("DFS Iteratif : %d ms\n", bestIter.Microseconds())
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah paket: ")
	fmt.Scan(&n)

	graph := make(Graph)

	for i := 0; i < n/2; i++ {
    left := 2*i + 1
    right := 2*i + 2
	
		if left < n {
        	graph[i] = append(graph[i], left)
    	}
    	if right < n {
        	graph[i] = append(graph[i], right)
    	}
	}
	fmt.Println("\nMenelusuri dependensi dari paket 0")
	benchmarkDFS(graph, 0)
}