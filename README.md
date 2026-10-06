# algorithm_mini_lab_team_c





# Traveling Salesman Problem (TSP) Experiment

## Objective

The objective of this experiment is to understand how a
TSP route can be constructed and how changing distances
affects the resulting route.

## Algorithm

This program uses a nearest-neighbor greedy approach.

Starting from City 1, the program repeatedly selects the
nearest unvisited city. After visiting all cities, it
returns to City 1.

## Test 1

Input:

4 cities

Distances:

- City 1 - City 2 = 15
- City 1 - City 3 = 20
- City 1 - City 4 = 29
- City 2 - City 3 = 8
- City 2 - City 4 = 10
- City 3 - City 4 = 11

Result:

Route:
City 1 -> City 2 -> City 3 -> City 4 -> City 1

Total distance:
63

## Test 2

The distance between City 2 and City 4 was changed.

The route remained the same.

## Test 3

The distance between City 1 and City 4 was changed
from 29 to 5.

Result:

City 1 -> City 4 -> City 2 -> City 3 -> City 1

Total distance:
43

## Observation

Changing the distance between cities can change the route
selected by the nearest-neighbor algorithm. The algorithm
always makes its decision based on the nearest currently
unvisited city.
