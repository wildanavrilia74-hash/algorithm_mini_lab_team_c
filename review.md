# BFS vs DFS Code Review

## 1. Review Purpose

This review checks the provided `BFS_VS_DFS.go` program for correctness, readability, testing, and possible improvements.

## 2. Code Being Reviewed

The program contains:
- A Breadth-First Search (BFS) function
- A Depth-First Search (DFS) function
- A graph containing A, B, C, D, E, and F
- A main function that runs both algorithms starting from A

## 3. BFS Review

The BFS function creates a visited map and a queue:

```go
visited := make(map[string]bool)
queue := []string{start}
```

It removes the first element from the queue, checks whether it was visited, marks it visited, prints it, and adds its unvisited neighbors to the queue.

### Current BFS result

For the current graph:

```text
        A
       / \
      B   C
     / \   \
    D   E   F
```

The expected BFS traversal is:

```text
A -> B -> C -> D -> E -> F
```

**Review:** The BFS implementation works for the current test graph and demonstrates queue-based, level-by-level traversal.

### Possible BFS improvement

The code marks a vertex visited when it is removed from the queue. In a graph with cycles or multiple edges to the same vertex, a vertex could be added to the queue more than once. A more robust approach is to mark a neighbor as visited when it is added to the queue.

## 4. DFS Review

The DFS function first checks whether the current vertex was visited. It then marks and prints it and recursively visits each unvisited neighbor.

### Current DFS result

Starting from A, the expected traversal is:

```text
A -> B -> D -> E -> C -> F
```

**Review:** The DFS implementation works for the current graph and demonstrates recursive depth-first traversal.

## 5. BFS vs DFS Comparison

| Feature | BFS | DFS |
|---|---|---|
| Main technique | Queue | Recursion |
| Traversal style | Level by level | Depth first |
| Current result | A B C D E F | A B D E C F |
| Starting vertex | A | A |
| Visited tracking | Yes | Yes |

The different traversal orders demonstrate the main difference between BFS and DFS.

## 6. Test Plan

### Test 1 — Current Graph

Expected:

```text
BFS: A B C D E F
DFS: A B D E C F
```

Observation: BFS visits nodes by level, while DFS follows a branch deeply before backtracking.

### Test 2 — Add G

Change the graph to:

```go
"C": {"F", "G"},
"G": {},
```

Expected:

```text
BFS: A B C D E F G
DFS: A B D E C F G
```

Record the actual output and screenshot it.

### Test 3 — Change Neighbor Order

Change:

```go
"A": {"B", "C"},
```

to:

```go
"A": {"C", "B"},
```

Observe how the traversal order changes.

### Test 4 — Disconnected Vertex

Add:

```go
"G": {},
```

without connecting G to A.

Starting from A, G will not be visited because it is not reachable from A.

## 7. Strengths

1. The code is short and beginner-friendly.
2. BFS and DFS are separated into clear functions.
3. Both algorithms use visited tracking.
4. The graph is easy to modify for experiments.
5. The program is suitable for the Team C BFS vs DFS experiment.

## 8. Improvements

1. Improve BFS visited handling by marking vertices when they enter the queue.
2. Add more test cases, including cycles and disconnected graphs.
3. Add comments explaining the queue, visited map, and recursion.
4. Allow the user to choose the starting vertex instead of always using A.

## 9. Questions for the Presentation

### What does BFS use?
A queue.

### What does DFS use in this program?
Recursion.

### Why are the BFS and DFS results different?
BFS explores level by level, while DFS explores deeply along a path before backtracking.

### What is the starting vertex?
A.

### Where is the starting vertex specified?
In:

```go
bfs(graph, "A")
```

and:

```go
dfs(graph, "A", visited)
```

### Can the traversal order change?
Yes. Changing the graph structure or the order of neighbors in the adjacency lists can change the traversal order.

### Does DFS always produce A B D E C F?
No. The order depends on the graph structure and neighbor ordering.

## 10. Overall Review

**BFS:** Working for the current graph.

**DFS:** Working for the current graph.

**Readability:** Good.

**Experiment readiness:** Good.

**Main improvement:** Test additional graph structures and improve BFS visited handling for more complex graphs.

### Final Review Comment

> The BFS and DFS implementations are clear and suitable for the Team C mini-lab. The current graph demonstrates the difference between level-order BFS traversal and depth-first DFS traversal. The next step is to change the graph and neighbor ordering, predict the new results, run the program, and document the observations.

## 11. Review Checklist

- [x] BFS function reviewed
- [x] DFS function reviewed
- [x] Graph structure reviewed
- [x] Expected BFS output identified
- [x] Expected DFS output identified
- [x] Strengths documented
- [x] Improvements documented
- [x] Test cases prepared
- [x] Presentation questions prepared
- [ ] Run the test cases
- [ ] Record actual outputs
- [ ] Add screenshots
- [ ] Submit review through GitHub
