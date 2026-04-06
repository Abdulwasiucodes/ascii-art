package main

import (
	"fmt"
	"os"
	"strings"
)

const usageMsg = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""

func main() {
	args := os.Args[1:]

	var color string
	var substring string
	var input string

	switch len(args) {

	// ✅ Only string (no color)
	case 1:
		input = args[0]

	// ✅ With color
	case 3:
		if !strings.HasPrefix(args[0], "--color=") {
			fmt.Println(usageMsg)
			return
		}

		color = strings.TrimPrefix(args[0], "--color=")
		substring = args[1]
		input = args[2]

	default:
		fmt.Println(usageMsg)
		return
	}

	// Read banner
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading banner file")
		return
	}

	banner := strings.Split(string(data), "\n")

	output := Render(input, banner, color, substring)

	fmt.Print(output)
}
