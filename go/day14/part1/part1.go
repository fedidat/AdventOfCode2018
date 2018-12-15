package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	target := 9
	// target := 5
	// target := 18
	// target := 2018
	// target := 380621

	head := ring.New(4)
	current := head
	current.Value = 3
	current = current.Next()
	current.Value = 7
	current = current.Next()
	current.Value = 1
	current = current.Next()
	current.Value = 0
	firstLink, secondLink := head, head.Next()
	recipes := 2
	for i := 0; recipes < target+8; i++ {
		//print
		for i, printNode := 0, head; i < head.Len(); i, printNode = i+1, printNode.Next() {
			if printNode == firstLink {
				fmt.Print("*")
			} else if printNode == secondLink {
				fmt.Print("^")
			}
			fmt.Print(printNode.Value, "->")
		}
		fmt.Println()

		//insert elements
		sum := firstLink.Value.(int) + secondLink.Value.(int)
		sumUnit, sumTens := sum%10, sum/10
		if sumTens > 0 {
			current.Link(ring.New(1))
			current = current.Next()
			current.Value = sumTens
			recipes++
		}
		current.Link(ring.New(1))
		current = current.Next()
		current.Value = sumUnit
		recipes++

		//advance positions
		firstScore, secondScore := firstLink.Value.(int), secondLink.Value.(int)
		for i := 0; i <= firstScore; i++ {
			firstLink = firstLink.Next()
		}
		for i := 0; i <= secondScore; i++ {
			secondLink = secondLink.Next()
		}
	}
	if recipes > target+8 {
		current = current.Prev()
	}

	res := ""
	for i := 0; i < 10; i++ {
		res = strconv.Itoa(current.Value.(int)) + res
		current = current.Prev()
	}
	fmt.Println(res)
}
