package solution

func (p Coord) add(other Coord) Coord {
	return Coord{p.x + other.x, p.y + other.y}
}

type Grid struct {
	data          []string
	width, height int
	start, end    Coord
}

func NewGrid(input []string) *Grid {
	g := &Grid{
		data:   input,
		height: len(input),
		width:  len(input[0]),
	}
	g.findStartEnd()
	return g
}

func (g *Grid) findStartEnd() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			switch g.data[y][x] {
			case 'S':
				g.start = Coord{x, y}
			case 'E':
				g.end = Coord{x, y}
			}
		}
	}
}

func (g *Grid) isValid(p Coord) bool {
	return p.x >= 0 && p.x < g.width &&
		p.y >= 0 && p.y < g.height &&
		g.data[p.y][p.x] != '#'
}

type QueueItem struct {
	pos   Coord
	steps int
}

func (g *Grid) findBasePath() int {
	directions := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make(map[Coord]bool)
	queue := []QueueItem{{g.start, 0}}
	visited[g.start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == g.end {
			return current.steps
		}

		for _, dir := range directions {
			next := current.pos.add(dir)
			if g.isValid(next) && !visited[next] {
				visited[next] = true
				queue = append(queue, QueueItem{next, current.steps + 1})
			}
		}
	}
	return -1
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

func (g *Grid) findPathWithCheat(cheat Cheat, baseTime int) int {
	// First find path to cheat start
	visited := make(map[Coord]bool)
	queue := []QueueItem{{g.start, 0}}
	visited[g.start] = true

	timeToCheatStart := -1
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == cheat.start {
			timeToCheatStart = current.steps
			break
		}

		directions := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range directions {
			next := current.pos.add(dir)
			if g.isValid(next) && !visited[next] {
				visited[next] = true
				queue = append(queue, QueueItem{next, current.steps + 1})
			}
		}
	}

	if timeToCheatStart == -1 {
		return baseTime // Can't reach cheat start
	}

	// Cost of the cheat itself
	cheatCost := abs(cheat.end.x-cheat.start.x) + abs(cheat.end.y-cheat.start.y)

	// Find path from cheat end to finish
	visited = make(map[Coord]bool)
	queue = []QueueItem{{cheat.end, timeToCheatStart + cheatCost}}
	visited[cheat.end] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == g.end {
			return current.steps
		}

		directions := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range directions {
			next := current.pos.add(dir)
			if g.isValid(next) && !visited[next] {
				visited[next] = true
				queue = append(queue, QueueItem{next, current.steps + 1})
			}
		}
	}

	return baseTime // Can't reach end from cheat end
}

func (g *Grid) findCheats(baseTime int, minSavings int) int {
	// Track unique cheats that save enough time
	validCheats := make(map[Cheat]bool)

	// For each possible cheat start position
	for y1 := 0; y1 < g.height; y1++ {
		for x1 := 0; x1 < g.width; x1++ {
			startPos := Coord{x1, y1}
			if !g.isValid(startPos) {
				continue
			}

			// For each possible cheat end position within 2 steps
			for y2 := max(0, y1-2); y2 < min(g.height, y1+3); y2++ {
				for x2 := max(0, x1-2); x2 < min(g.width, x1+3); x2++ {
					endPos := Coord{x2, y2}
					if !g.isValid(endPos) {
						continue
					}

					// Check if within 2 steps
					if abs(x2-x1)+abs(y2-y1) > 2 {
						continue
					}

					cheat := Cheat{startPos, endPos}
					timeWithCheat := g.findPathWithCheat(cheat, baseTime)
					if timeWithCheat < baseTime {
						savings := baseTime - timeWithCheat
						if savings >= minSavings {
							validCheats[cheat] = true
						}
					}
				}
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

func Part_One(grid []string, minSavings int) int {
	g := NewGrid(grid)
	baseTime := g.findBasePath()
	if baseTime == -1 {
		return 0
	}
	return g.findCheats(baseTime, minSavings)
}
