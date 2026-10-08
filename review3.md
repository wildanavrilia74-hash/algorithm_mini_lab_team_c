## Graph Change Experiment

### Change
I changed the neighbor order of A from:

A -> B, C

to:

A -> C, B

### Prediction
I predicted that the traversal order would change because
BFS and DFS process neighbors according to their order in
the adjacency list.

### Observation
After running the program, the traversal order changed.

### Explanation
The graph connections were not fundamentally changed,
but the order in which neighbors were processed changed.
Therefore, the traversal output also changed.
