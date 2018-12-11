package main

import (
	"fmt"
)

func main() {
	input := 18
	// input := 42
	// input := 3214
	serialNumber := input
	var grid [][]int
	w, h, xScope, yScope := 300, 300, 3, 3
	for x := 0; x < w; x++ {
		var row []int
		for y := 0; y < h; y++ {
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

	var scopedGrid [][]int
	for xScoped := 0; xScoped < w - xScope; xScoped++ {
		var row []int
		for yScoped := 0; yScoped < h - yScope; yScoped++ {
			sum := 0
			for x := xScoped; x < xScoped + xScope; x++ {
				for y := yScoped; y < yScoped + yScope; y++ {
					sum += grid[x][y]
				}
			}
			row = append(row, sum)
		}
		scopedGrid = append(scopedGrid, row)
	}
	
	max, maxX, maxY := -100, 0, 0
	for y := 0; y < h - yScope; y++ {
		for x := 0; x < w - xScope; x++ {
			if scopedGrid[x][y] > max {
				max, maxX, maxY = scopedGrid[x][y], x, y
			}
		}
	}
	fmt.Println(max, " ", maxX, " ", maxY)

	
}
