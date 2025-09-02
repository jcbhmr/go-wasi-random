//go:build wasip2

package random_test

import (
	"fmt"

	"github.com/jcbhmr/go-wasi-random/random"
)

func ExampleGetRandomU64() {
	n := random.GetRandomU64()
	fmt.Printf("Random uint64 is %d\n", n)
	// Example output: Random uint64 is 12345
}

func ExampleGetRandomBytes() {
	b := random.GetRandomBytes(100)
	fmt.Printf("100 random bytes are %x\n", b)
	// Example output: 100 random bytes are 010203040506070809...
}
