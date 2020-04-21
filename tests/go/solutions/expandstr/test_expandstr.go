package main

import (
	"github.com/01-edu/z01"
)

func main() {
	arg1 := []string{"hello", "you"}
	arg2 := []string{"   only  it's harder   "}
	arg3 := []string{"you   see   it's   easy   to   display    the     same  thing"}
	args := [][]string{arg1, arg2, arg3}

	// adding of 15 random valid tests
	for i := 0; i < 15; i++ {
		randomArg := []string{z01.RandWords()}
		args = append(args, randomArg)
	}

	for _, v := range args {
		z01.ChallengeMain("expandstr", v...)
	}
	z01.ChallengeMain("expandstr")
}
