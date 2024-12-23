package solution

import (
	"sort"
	"strings"
	"time"
)

type Graph map[string]map[string]bool

type Node struct {
	name        string
	connections map[string]bool
	degree      int
}

func buildGraph(connections [][]string) Graph {
	graph := make(Graph)

	for _, conn := range connections {
		a, b := conn[0], conn[1]

		// Initialize maps if they don't exist
		if graph[a] == nil {
			graph[a] = make(map[string]bool)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]bool)
		}

		// Add bidirectional connection
		graph[a][b] = true
		graph[b][a] = true
	}

	return graph
}

func findTriplets(graph Graph) [][]string {
	seen := make(map[string]bool)
	var triplets [][]string

	// For each node in the graph
	for node := range graph {
		// For each neighbor of the node
		for neighbor := range graph[node] {
			// For each neighbor of the neighbor
			for thirdNode := range graph[neighbor] {
				// Check if third node connects back to first node
				if graph[thirdNode][node] {
					// Create sorted triplet to avoid duplicates
					triplet := []string{node, neighbor, thirdNode}
					sort.Strings(triplet)

					// Create key for deduplication
					key := strings.Join(triplet, ",")
					if !seen[key] {
						seen[key] = true
						triplets = append(triplets, triplet)
					}
				}
			}
		}
	}

	return triplets
}

func findMaximumClique(graph Graph) []string {
	// Convert graph to nodes with degrees for better pruning
	nodes := make([]Node, 0, len(graph))
	for name, connections := range graph {
		nodes = append(nodes, Node{
			name:        name,
			connections: connections,
			degree:      len(connections),
		})
	}

	// Sort nodes by degree in descending order for better pruning
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].degree > nodes[j].degree
	})

	// Initialize variables for branch and bound
	var maxClique []string
	currentClique := make([]string, 0)

	// Helper function to check if a node can join the current clique
	canJoinClique := func(node Node, clique []string) bool {
		for _, member := range clique {
			if !node.connections[member] {
				return false
			}
		}
		return true
	}

	// Recursive function to find maximum clique using branch and bound
	var findClique func(candidates []Node, start int)
	findClique = func(candidates []Node, start int) {
		// If remaining candidates can't beat our best, prune this branch
		if len(currentClique)+len(candidates) <= len(maxClique) {
			return
		}

		// If we've run out of candidates, check if we have a new maximum
		if start >= len(candidates) {
			if len(currentClique) > len(maxClique) {
				maxClique = make([]string, len(currentClique))
				copy(maxClique, currentClique)
			}
			return
		}

		node := candidates[start]

		// If this node can join our clique, try including it
		if canJoinClique(node, currentClique) {
			// Create new candidates list with only nodes connected to current node
			newCandidates := make([]Node, 0)
			for _, candidate := range candidates[start+1:] {
				if node.connections[candidate.name] {
					newCandidates = append(newCandidates, candidate)
				}
			}

			// Add node to clique and recurse
			currentClique = append(currentClique, node.name)
			findClique(newCandidates, 0)
			currentClique = currentClique[:len(currentClique)-1]
		}

		// Try without this node
		findClique(candidates, start+1)
	}

	// Start the search with all nodes
	findClique(nodes, 0)
	return maxClique
}

func getPassword(nodes []string) string {
	// Sort nodes alphabetically
	sort.Strings(nodes)
	return strings.Join(nodes, ",")
}

func countTripletsWithT(triplets [][]string) int {
	count := 0
	for _, triplet := range triplets {
		// Check if any node in triplet starts with 't'
		for _, node := range triplet {
			if strings.HasPrefix(node, "t") {
				count++
				break
			}
		}
	}
	return count
}

func Part_One(computer_list [][]string) int {
	defer Track(time.Now(), "Part 1")

	// Build graph and find solution
	graph := buildGraph(computer_list)
	triplets := findTriplets(graph)
	result := countTripletsWithT(triplets)

	return result
}

func Part_Two(computer_list [][]string) string {
	defer Track(time.Now(), "Part 2")

	graph := buildGraph(computer_list)
	clique := findMaximumClique(graph)
	password := getPassword(clique)

	return password
}
