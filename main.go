package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFileToArray(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Day 1
// Consider sums of a three-measurement sliding window.
// How many sums are larger than the previous sum?
func dayOne() {
	// Read file contents whole
	//https://gosamples.dev/read-file/

	// Read line by line into an array
	//https://stackoverflow.com/questions/61633338/how-to-split-the-lines-of-a-txt-file-and-add-them-in-a-slice-in-go?noredirect=1&lq=1

	file, err := os.Open("data/day1.txt")

	// Go error handling
	if err != nil {
		log.Fatal(err)
	}

	// The defer statement defers the execution of an expression
	// untill nearby functions have executed
	defer file.Close()

	// Buffer scanner
	sc := bufio.NewScanner(file)

	// Creates object of correct size
	readings := make([]int, 0)

	// Read through 'tokens' until an EOF is encountered.
	for sc.Scan() {
		// https://pkg.go.dev/strconv#Atoi
		// Parses String to int
		x, err := strconv.Atoi(sc.Text())

		if err != nil {
			log.Fatal(err)
		}

		readings = append(readings, x)
	}

	// Error handling
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	// Convert these to integers
	// and calculate the sliding window
	windowSums := 0

	for i := 2; i < (len(readings)); i++ {
		if readings[i-2] < readings[i] {
			windowSums += 1
		}
	}

	fmt.Println(windowSums)
}

func dayTwo() {
	// Read in the file again
	var dataFile string
	dataFile = "data/day2.txt"

	lineData := readFileToArray(dataFile)

	// Initialise positions
	var xPosition int // Horisontal Position
	var yPosition int // Depth, down increases the depth and up decreases the depth

	for i := 0; i < len(lineData); i++ {
		line := strings.Split(lineData[i], " ")
		value, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}
		switch line[0] {
		case "forward":
			xPosition += value
		case "up":
			yPosition -= value
		case "down":
			yPosition += value
		}
	}
	fmt.Println("Part One: ", xPosition*yPosition)

	// Part Two
	aim := 0
	xPosition = 0
	yPosition = 0

	for i := 0; i < len(lineData); i++ {
		line := strings.Split(lineData[i], " ")
		value, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		switch line[0] {
		case "forward":
			xPosition += value
			yPosition += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	fmt.Println("Part Two: ", xPosition*yPosition)
}

// Function to recursively filter a list
func filterArray(arr []string, i int) []string {
	// base case
	if i == 0 {
		return arr
	}

	n := len(arr[0])

	// One pass over the array to determine the most
	// popular
	numOnes := 0
	for idx := range arr {
		if arr[idx][n-i-1] == '1' {
			{
				numOnes += 1
			}
		} else {
			numOnes -= 1
		}
	}
	if numOnes < 0 {
		mostCommonValue := '0'
	} else {
		mostCommonValue := '1'
	}

	res = make([]string, 0)

	for idx := range arr {
		if arr[idx][n-i-1] == mostCommonValue {
			{
				res = append(res, arr[idx])
			}
		}
	}

	return res

}

// Day Three
func dayThree() {
	fileData := readFileToArray("data/day3.txt")

	// Get the length of the first item
	numLength := len(fileData[0])

	// Number of entries
	numLines := len(fileData)

	// Generate an empty array to get the most common bit
	mostCommonBits := make([]int, numLength)

	for i := 0; i < numLines; i++ {
		// Iterate over the string
		for j := 0; j < numLength; j++ {
			if fileData[i][j] == '1' {
				mostCommonBits[j] += 1
			}
		}
	}

	// Divide by the numLines and round down or up
	for i := 0; i < numLength; i++ {
		if (float64(mostCommonBits[i]) / float64(numLines)) > 0.5 {
			mostCommonBits[i] = 1
		} else {
			mostCommonBits[i] = 0
		}
	}

	// Calculate the gamma and epsilon rate and convert them to decimal
	gamma := 0
	epsilon := 0

	for i := 0; i < numLength; i++ {
		gamma += int(math.Pow(2, float64(i))) * mostCommonBits[numLength-1-i]
		epsilon += int(math.Pow(2, float64(i))) * int(math.Abs(float64(mostCommonBits[numLength-1-i]-1)))
	}
	fmt.Println("Part 1:", epsilon*gamma)

	// Part Two
	// Have to filter both values

}

// Day 1
func main() {
	dayThree()
}
