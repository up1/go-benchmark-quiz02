package main

import (
	"strings"
)

func chonlasith(input string) bool {
	tokens := strings.Split(input, "{}")
	for i, n := 0, len(tokens); i < n; i++ {
		if strings.ContainsAny(tokens[i], "{}") {
			return false
		}
	}
	return true
}

func chonlasith2(input string) bool {
	for i, n := 0, len(input); i < n; i++ {
		if input[i] == '{' {
			if i < n-1 {
				if input[i+1] != '}' {
					return false
				}
				i++
			} else {
				return false
			}
		} else if input[i] == '}' {
			return false
		}
	}
	return true
}
