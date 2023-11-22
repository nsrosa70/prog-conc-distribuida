package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 40
	height = 20
)

func initializeGrid() [][]bool {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			grid[i][j] = rand.Intn(2) == 1
		}
	}
	return grid
}

func printGrid(grid [][]bool) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if grid[i][j] {
				fmt.Print("■ ")
			} else {
				fmt.Print("□ ")
			}
		}
		fmt.Println()
	}
}

func countNeighbors(grid [][]bool, x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// Skip the cell itself
			if i == 0 && j == 0 {
				continue
			}
			// Calculate neighbor's coordinates with wrapping around the grid
			xx := (x + i + width) % width
			yy := (y + j + height) % height
			if grid[yy][xx] {
				count++
			}
		}
	}
	return count
}

func updateGrid(grid [][]bool) [][]bool {
	newGrid := make([][]bool, height)
	for i := range newGrid {
		newGrid[i] = make([]bool, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			aliveNeighbors := countNeighbors(grid, j, i)
			if grid[i][j] {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					newGrid[i][j] = false
				} else {
					newGrid[i][j] = true
				}
			} else {
				if aliveNeighbors == 3 {
					newGrid[i][j] = true
				} else {
					newGrid[i][j] = false
				}
			}
		}
	}
	return newGrid
}

func main() {
	grid := initializeGrid()

	for i := 0; i < 100; i++ { // Change the number of iterations as needed
		fmt.Printf("Generation %d:\n", i)
		printGrid(grid)
		time.Sleep(200 * time.Millisecond) // Adjust the delay between generations

		// Clear the console
		fmt.Print("\033[H\033[2J")

		grid = updateGrid(grid)
	}
}
