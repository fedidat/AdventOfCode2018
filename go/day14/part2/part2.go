package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// target := "51589"
	// target := "01245"
	// target := "92510"
	// target := "59414"
	target := "380621"

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
	for i := 0; recipes < 100000000; i++ {
		//insert elements and print when inserting
		sum := firstLink.Value.(int) + secondLink.Value.(int)
		sumUnit, sumTens := sum%10, sum/10
		if sumTens > 0 {
			current.Link(ring.New(1))
			current = current.Next()
			current.Value = sumTens
			recipes++
			// fmt.Print(i, " ", recipes, " ", printRing(head, firstLink, secondLink))
			if checkScoreDigits(target, current) {
				fmt.Println(recipes - len(target) + 2)
				return
			}
		}
		current.Link(ring.New(1))
		current = current.Next()
		current.Value = sumUnit
		recipes++
		// fmt.Print(i, " ", recipes, " ", printRing(head, firstLink, secondLink))
		if checkScoreDigits(target, current) {
			fmt.Println(recipes - len(target) + 2)
			return
		}

		//advance positions
		firstScore, secondScore := firstLink.Value.(int), secondLink.Value.(int)
		for i := 0; i <= firstScore; i++ {
			firstLink = firstLink.Next()
		}
		for i := 0; i <= secondScore; i++ {
			secondLink = secondLink.Next()
		}
	}
}

/*Not the most efficient implementation, I could have stored an additional
  node and removed the previous each time, but I'm a bit tired of Golang
  at this point.
*/
func checkScoreDigits(pattern string, r *ring.Ring) bool {
	for i := len(pattern) - 1; i >= 0; i-- {
		if int(pattern[i]-'0') != r.Value.(int) {
			return false
		}
		r = r.Prev()
	}
	return true
}

func printRing(head *ring.Ring, firstLink *ring.Ring, secondLink *ring.Ring) string {
	res := ""
	for i, printNode := 0, head; i < head.Len(); i, printNode = i+1, printNode.Next() {
		if printNode == firstLink {
			res += "*"
		} else if printNode == secondLink {
			res += "^"
		}
		res += strconv.Itoa(printNode.Value.(int)) + "->"
	}
	res += "\n"
	return res
}
