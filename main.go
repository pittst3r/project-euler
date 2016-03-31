package main

import (
	"fmt"
	. "github.com/rdpitts/project-euler/problems"
	"os"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: project-euler <problem number>")
		os.Exit(1)
	}

	problem := os.Args[1]

	problemMap := map[string]func() string{
		"1": P1,
		"2": P2,
		"3": P3,
		"4": P4,
		"5": P5,
	}

	clock := time.Now()
	answer := problemMap[problem]()
	duration := time.Since(clock).Seconds()

	fmt.Println("Answer:", answer)
	fmt.Println("Run in", duration, "seconds")

}
