package main

import "fmt"

// BFS dari kode asli
func bfs(graph map[string][]string, start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	fmt.Println("BFS:")

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true
		fmt.Print(current, " ")

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	fmt.Println()
}

// DFS dari kode asli
func dfs(graph map[string][]string, current string, visited map[string]bool) {
	if visited[current] {
		return
	}

	visited[current] = true
	fmt.Print(current, " ")

	for _, neighbor := range graph[current] {
		if !visited[neighbor] {
			dfs(graph, neighbor, visited)
		}
	}
}

// Tester BFS vs DFS
func testGraph(name string, graph map[string][]string, start string) {
	fmt.Println("================================")
	fmt.Println(name)
	fmt.Println("================================")

	fmt.Println("Graph:")
	for node, neighbors := range graph {
		fmt.Println(node, "->", neighbors)
	}

	fmt.Println()

	// BFS
	bfs(graph, start)

	// DFS
	fmt.Println("DFS:")
	visited := make(map[string]bool)
	dfs(graph, start, visited)
	fmt.Println()

	fmt.Println()
}

func main() {

	// TEST 1
	graph1 := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": {},
		"E": {},
		"F": {},
	}

	testGraph("TEST 1 - Graph Awal", graph1, "A")

	// TEST 2
	// Perubahan input sesuai eksperimen PDF
	graph2 := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": {},
		"E": {},
		"F": {"G"},
		"G": {},
	}

	testGraph("TEST 2 - Menambahkan Node G", graph2, "A")

	// TEST 3
	// Mengubah hubungan antar node
	graph3 := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {"E", "F"},
		"D": {},
		"E": {},
		"F": {},
	}

	testGraph("TEST 3 - Mengubah Struktur Graph", graph3, "A")
}
