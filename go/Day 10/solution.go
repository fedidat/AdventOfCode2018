package main

import (
	"AdventOfCode2018/go/utils"
	"fmt"
	"time"
)

type velocity struct {
	X, Y int
}

type point struct {
	X, Y     int
	velocity velocity
}

func main() {

	var positionX, positionY, velocityX, velocityY int

	var points []*point

	for _, line := range utils.GetLines("input.txt") {
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &positionX, &positionY, &velocityX, &velocityY)

		points = append(points, &point{
			X: positionX,
			Y: positionY,
			velocity: velocity{
				X: velocityX,
				Y: velocityY,
			},
		})
	}

	var seconds int

	for {
		for _, point := range points {
			point.X += point.velocity.X
			point.Y += point.velocity.Y
		}

		minX, maxX, minY, maxY := getGridSize(points)

		if minX < 0 {
			minX *= -1
		}

		if minY < 0 {
			minY *= -1
		}

		seconds++

		gridX := minX + maxX + 1
		gridY := minY + maxY + 1

		if gridY < 300 && gridX < 200 {
			showGrid(points, minX, minY, gridX, gridY)
			fmt.Println("after", seconds, "seconds")

			time.Sleep(time.Second)
		}
	}
}

func getGridSize(points []*point) (minX, maxX, minY, maxY int) {
	for _, point := range points {
		if point.X > maxX {
			maxX = point.X
		}

		if point.X < minX {
			minX = point.X
		}

		if point.Y > maxY {
			maxY = point.Y
		}

		if point.Y < minY {
			minY = point.Y
		}
	}

	return
}

func showGrid(points []*point, minX, minY, gridX, gridY int) {
	grid := make([][]bool, gridY)

	for i := range grid {
		grid[i] = make([]bool, gridX)
	}

	for _, point := range points {
		grid[point.Y+minY][point.X+minX] = true
	}

	for _, line := range grid {
		var s string

		for _, point := range line {
			if point {
				s += "#"
			} else {
				s += "."
			}
		}

		fmt.Println(s)
	}
}
