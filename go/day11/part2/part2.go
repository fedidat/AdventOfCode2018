package main

import (
	"fmt"
)

func main() {
	/*
	Runtime is fairly long (over a minute) but the max stabilizes within a few seconds,
	because the sum goes smaller after a certain value..
	*/

	// input := 18
	// input := 42
	input := 3214
	serialNumber := input
	var grid [][]int
	gridSize := 300
	for x := 0; x < gridSize; x++ {
		var row []int
		for y := 0; y < gridSize; y++ {
			rackID := x + 10
			fuelPower := rackID * y
			fuelPower += serialNumber
			fuelPower *= rackID
			fuelPower = (fuelPower % 1000) / 100
			fuelPower -= 5
			row = append(row, fuelPower)
		}
		grid = append(grid, row)
	}

	max, maxX, maxY, maxScopeSize := -1000000, 0, 0, 0
	for scopeSize := 1; scopeSize <= gridSize; scopeSize++ {
		// scopeSize := 100
		for xScoped := 0; xScoped < gridSize - scopeSize; xScoped++ {
			for yScoped := 0; yScoped < gridSize - scopeSize; yScoped++ {
				sum := 0
				for x := xScoped; x < xScoped + scopeSize; x++ {
					for y := yScoped; y < yScoped + scopeSize; y++ {
						sum += grid[x][y]
					}
				}
				if sum > max {
					max, maxX, maxY, maxScopeSize = sum, xScoped, yScoped, scopeSize
				}
			}
		}
		fmt.Println(max, " ", maxX, " ", maxY, " ", maxScopeSize, " ", scopeSize)
	}
}
