package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

//Str2Int converts a string to an int and handles errors
func Str2Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

//gets the sum of all values in a slice of ints
func GetSum(list []int) int {
	sum := 0
	for _, v := range list {
		sum = sum + v
	}
	return sum
}

//GetLines gets the lines from a file
func GetLines(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

//RegSplit A simple regex splitter
func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}

//Max gets the max value of a slice of ints
func Max(x []int) int {
	max := 0
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

//Abs absolute value
func Abs(n int) int {
	if n < 0 {
		return n * (-1)
	} else {
		return n
	}
}

//Check check errors
func Check(err error) {
	if err != nil {
		panic(err)
	}
}