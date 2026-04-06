package main

import "strings"

const (
	Black   = "\033[0;30m"
	Red     = "\033[0;31m"
	Green   = "\033[0;32m"
	Yellow  = "\033[0;33m"
	Blue    = "\033[0;34m"
	Magenta = "\033[0;35m"
	Cyan    = "\033[0;36m"
	White   = "\033[0;37m"
	Orange  = "\033[0;38;5;208m"
	Reset   = "\033[0m"
)

var ColorMap = map[string]string{
	"black":   Black,
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"magenta": Magenta,
	"cyan":    Cyan,
	"orange":  Orange,
	"white":   White,
}

func Colorize(s, color string) string {
	colorCode, ok := ColorMap[color]
	if !ok {
		return s
	}
	return colorCode + s + Reset
}

func isInRange(index int, indices []int, length int) bool {
	for _, start := range indices {
		if index >= start && index < start+length {
			return true
		}
	}
	return false
}

func FindSubStringIndices(str, substr string) []int {
	var indices []int

	if substr == "" {
		return indices
	}

	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i:], substr) {
			indices = append(indices, i)
			i += len(substr) - 1
		}
	}
	return indices
}
