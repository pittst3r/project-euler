package main

import (
    "time"
    "fmt"
    "os"
    . "github.com/rdpitts/project-euler/problems"
)

func main() {

    if len(os.Args) < 2 {
        fmt.Println("Usage: project-euler <problem number>")
        os.Exit(1)
    }

    problem := os.Args[1]

    problemMap := map[string]func() string {
        "1": P1,
    }

    clock := time.Now()
    answer := problemMap[problem]()
    duration := time.Since(clock).Seconds()

    fmt.Println("Answer:", answer)
    fmt.Println("Run in", duration, "seconds")

}