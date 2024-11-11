package main

import (
	"math/rand"
	"time"
)

func runRandomDemo() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomInt := r.Intn(10)
	println("Random int:", randomInt)
}
