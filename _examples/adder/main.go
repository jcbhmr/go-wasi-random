//go:build wasip2

package main

import (
	"math"

	"github.com/jcbhmr/go-wasi-random/_examples/adder/internal/math/adder/add"
	"github.com/jcbhmr/go-wasi-random/random"
)

func init() {
	add.Exports.AddRandom = func(a int32) (result int32) {
		n := random.GetRandomU64()
		n100 := float64(n) / math.MaxUint64 * 100
		return a + int32(n100)
	}
}

func main() {}
