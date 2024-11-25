package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// ANSI color codes
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
)

var colors = []string{red, green, yellow, blue, purple, cyan}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		printGradientLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func printGradientLine(line string) {
	length := len(line)
	for i, char := range line {
		colorIndex := int(math.Floor(float64(i) * float64(len(colors)) / float64(length)))
		color := colors[colorIndex%len(colors)]
		fmt.Printf("%s%c%s", color, char, reset)
	}
	fmt.Println()
}
