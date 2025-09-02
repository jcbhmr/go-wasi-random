//go:build wasip2

package main

import (
	"fmt"

	"github.com/jcbhmr/go-wasi-random/random"
)

func main() {
	n := random.GetRandomU64()
	fmt.Printf("Here's a random u64: %d\n", n)
}
