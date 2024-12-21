package solution

import (
	"time"
)

func (p Coord) add(other Coord) Coord {
	return Coord{p.X + other.X, p.Y + other.Y}
}

func (g *Grid) isValid(p Coord) bool {
	return p.Y >= 0 && p.Y < g.Height && p.X >= 0 && p.X < g.Width && g.Data[p.Y][p.X] != '#'
}

func (g *Grid) is_not_out_of_band(p Coord) bool {
	return p.Y >= 0 && p.Y < g.Height && p.X >= 0 && p.X < g.Width
}

type QueueItem struct {
	pos   Coord
	steps int
}

func (g *Grid) findBasePath() (int, map[Coord]int) {
	directions := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make(map[Coord]int)
	queue := []QueueItem{{g.Start, 0}}
	visited[g.Start] = 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == g.End {
			return current.steps, visited
		}

		for _, dir := range directions {
			next := current.pos.add(dir)
			if g.isValid(next) && visited[next] == 0 && next != g.Start {
				visited[next] = current.steps + 1
				queue = append(queue, QueueItem{next, current.steps + 1})
			}
		}
	}
	return -1, visited
}

type Cheat struct {
	start, end Coord
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (g *Grid) testCheats(baseTime int, wining_path map[Coord]int, minSavings int) int {
	validCheats := make(map[Cheat]bool)

	directions := []Coord{{0, 2}, {0, -2}, {2, 0}, {-2, 0}}
	for cell := range wining_path {
		x, y := cell.X, cell.Y

		// For each possible cheat end position within 2 steps
		for _, dir := range directions {
			end_pos := Coord{x + dir.X, y + dir.Y}

			if !g.isValid(end_pos) {
				continue
			}

			cheat := Cheat{cell, end_pos}
			time_saved := wining_path[end_pos] - wining_path[cell]

			if time_saved > minSavings {
				validCheats[cheat] = true
			}
		}
	}
	return len(validCheats)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (g *Grid) manhattanBFS(baseTime int, wining_path map[Coord]int, minSavings int, start_pos Coord, validCheats *map[Cheat]bool) {
	directions := []Coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	visited := make(map[Coord]bool)

	// Queue for BFS
	queue := []Coord{start_pos} // Start at origin
	visited[start_pos] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check all directions
		for _, dir := range directions {
			next := Coord{current.X + dir.X, current.Y + dir.Y}

			if !g.is_not_out_of_band(next) {
				continue
			}

			// Calculate Manhattan distance
			distance := abs(next.X-start_pos.X) + abs(next.Y-start_pos.Y)

			// If point is within distance and not visited
			if distance <= 20 && !visited[next] {
				visited[next] = true
				queue = append(queue, next)

				if g.isValid(next) {
					cheat := Cheat{start_pos, next}
					time_saved := wining_path[next] - (wining_path[start_pos] + distance)

					if time_saved >= minSavings {
						(*validCheats)[cheat] = true
					}
				}

			}
		}
	}
}

func Part_One(grid *Grid, minSavings int) int {
	defer Track(time.Now(), "Part 1")

	baseTime, wining_path := grid.findBasePath()

	winning_cheats := grid.testCheats(baseTime, wining_path, minSavings)

	if baseTime == -1 {
		return 0
	}

	return winning_cheats
}

func Part_Two(grid *Grid, minSavings int) int {
	defer Track(time.Now(), "Part 2")

	baseTime, wining_path := grid.findBasePath()

	validCheats := make(map[Cheat]bool)
	for cell := range wining_path {
		grid.manhattanBFS(baseTime, wining_path, minSavings, cell, &validCheats)
	}

	if baseTime == -1 {
		return 0
	}

	return len(validCheats)
}
