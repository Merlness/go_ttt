package main

import "fmt"

func Display(grid []interface{}) string {
	var result string
	for i, val := range grid {
		if i > 0 && i%3 == 0 {
			result += "\n"
		} else if i%3 != 0 {
			result += " | "
		}
		result += fmt.Sprintf("%v", val)
	}
	return result
}
