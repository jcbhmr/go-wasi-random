//go:build generate

//go:generate wkg wit fetch
//go:generate rm -rf ./internal/
//go:generate go tool wit-bindgen-go generate --out ./internal/ --versioned ./wit/
//go:generate rm -rf ./internal/wasi/random/v0.2.0/
//go:generate go tool jet -g "*.go" "github\\.com/jcbhmr/go-wasi-random/0\\.2\\.0/_examples/hello-world/internal/wasi/random/v0\\.2\\.0/" "github.com/jcbhmr/go-wasi-random/0.2.0/" ./internal/

package main
