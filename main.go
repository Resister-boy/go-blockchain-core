package main

import (
	"flag"
	"fmt"
)

var (
	configFlag = flag.String("environment", "./environment.toml", "environment toml file not found")
	difficulty = flag.Int("difficulty", 12, "difficulty error")
)

func main() {
	flag.Parse()

	fmt.Println(*configFlag)
	fmt.Println(*difficulty)
	fmt.Println("Hello, World")
}
