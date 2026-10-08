package main

import "fmt"

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

func main() {

	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": {},
		"E": {},
		"F": {},
	}

	bfs(graph, "A")

	fmt.Println("DFS:")
	visited := make(map[string]bool)
	dfs(graph, "A", visited)
	fmt.Println()
}