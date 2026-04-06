package main

import "strings"

func Render(input string, banner []string, color string, substring string) string {
	if input == "" {
		return ""
	}

	var result strings.Builder
	parts := strings.Split(input, "\\n")

	for i, part := range parts {

		if part == "" {
			if i < len(parts)-1 {
				result.WriteString("\n")
			}
			continue
		}

		indices := FindSubStringIndices(part, substring)

		for row := 0; row < 8; row++ {
			for idx, ch := range part {
				if ch < 32 || ch > 126 {
					continue
				}

				start := (int(ch)-32)*9 + 1
				line := banner[start+row]

				// 🎯 Apply coloring logic
				if color != "" {
					if substring == "" {
						// color everything
						line = Colorize(line, color)
					} else if isInRange(idx, indices, len(substring)) {
						// color only substring
						line = Colorize(line, color)
					}
				}

				result.WriteString(line)
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
