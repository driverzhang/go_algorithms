package main

import (
	"gitHub/go_algorithms/linkedList"
)

var (
	store map[string]int
	list  []string
)

type Object struct {
	Status string
}

func main() {
	linkedList.Run()
}
