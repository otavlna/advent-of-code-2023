package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	y int
	x int
}

func (p1 Point) Add(p2 Point) Point {
	return Point{p1.y + p2.y, p1.x + p2.x}
}

func (p1 Point) Subtract(p2 Point) Point {
	return Point{p1.y - p2.y, p1.x - p2.x}
}

func (p1 Point) IsZero() bool {
	return p1.y == 0 && p1.x == 0
}

func (p1 Point) Equals(p2 Point) bool {
	return p1.y == p2.y && p1.x == p2.x
}

func (p Point) TileRune() rune {
	if p.y >= 0 && p.x >= 0 && p.y < len(grid) && p.x < len(grid[0]) {
		return grid[p.y][p.x]
	} else {
		return 'x'
	}
}

func (p1 Point) ConnectsTo(p2 Point) bool {
	tile1 := p1.TileRune()
	tile2 := p2.TileRune()

	connections1, exists1 := tiles[tile1]
	connections2, exists2 := tiles[tile2]

	if exists1 == false || exists2 == false {
		return false
	}

	diff := p2.Subtract(p1)

	for _, con1 := range connections1 {
		for _, con2 := range connections2 {
			if con1.Add(con2).IsZero() && (diff.Equals(con1)) {
				return true
			}
		}
	}
	return false
}

var (
	DOWN  = Point{1, 0}
	RIGHT = Point{0, 1}
	UP    = Point{-1, 0}
	LEFT  = Point{0, -1}
)

var offsets []Point = []Point{DOWN, RIGHT, UP, LEFT}

var tiles map[rune][]Point = map[rune][]Point{
	'|': {UP, DOWN},
	'-': {LEFT, RIGHT},
	'L': {UP, RIGHT},
	'J': {UP, LEFT},
	'7': {DOWN, LEFT},
	'F': {DOWN, RIGHT},
	'S': {DOWN, RIGHT, UP, LEFT},
	'.': nil,
	'x': nil,
}

var grid [][]rune = make([][]rune, 0)

func stepFromTile(tile Point, lastTile Point) (Point, Point) {
	for _, offset := range offsets {
		possiblePoint := tile.Add(offset)
		if !lastTile.Equals(possiblePoint) && possiblePoint.ConnectsTo(tile) {
			return possiblePoint, tile
		}
	}
	fmt.Println("impossible")
	return tile, tile
}

func main() {
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)

	var start Point

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		grid = append(grid, nil)
		y := len(grid) - 1
		for x, tile := range line {
			grid[y] = append(grid[y], tile)
			if tile == 'S' {
				start = Point{y, x}
			}
		}
	}

	// for _, row := range grid {
	// 	fmt.Printf("%c\n", row)
	// }

	// fmt.Println("start", start)

	steps := 0
	tile1, tile2 := Point{}, Point{}
	last1, last2 := start, start
	assigned1, assigned2 := false, false

	// first step - find 2 ways
	for _, offset := range offsets {
		possiblePoint := start.Add(offset)
		if possiblePoint.ConnectsTo(start) {
			if assigned1 == false {
				tile1 = possiblePoint
				assigned1 = true
			} else if assigned2 == false {
				tile2 = possiblePoint
				assigned2 = true
			}
		}
	}
	steps++

	// next steps
	for {
		if tile1.Equals(tile2) {
			fmt.Println(steps)
			break
		}
		tile1, last1 = stepFromTile(tile1, last1)
		tile2, last2 = stepFromTile(tile2, last2)
		steps++

	}

}
