//go:build wasip2

package main

import (
	"fmt"

	"github.com/jcbhmr/go-wasi-random/0.2.0/random"
)

func main() {
	n := random.GetRandomU64()
	fmt.Printf("Here's a random u64: %d\n", n)
}
