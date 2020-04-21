package main

import (
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	table := append(z01.MultRandWords(), "dsfda")
	table = append(table, "")
	table = append(table, "1")
	table = append(table, "1")

	for _, s := range table {
		z01.ChallengeMain("displaya", strings.Fields(s)...)
	}

	z01.ChallengeMain("displaya", "1", "a")
}
