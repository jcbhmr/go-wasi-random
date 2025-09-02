//go:build generate

//go:generate rm -rf ./.out/bindings/ ./insecure/ ./insecure-seed/ ./random/
//go:generate go tool wit-bindgen-go generate --out ./.out/bindings/ ./wit/
//go:generate mv ./.out/bindings/wasi/random/insecure/ ./insecure/
//go:generate mv ./.out/bindings/wasi/random/insecure-seed/ ./insecure-seed/
//go:generate mv ./.out/bindings/wasi/random/random/ ./random/
//go:generate rm -rf ./insecure/empty.s ./insecure-seed/empty.s ./random/empty.s
//go:generate go tool jet -g "*.go" -e "DO NOT EDIT.\n" "DO NOT EDIT.\n\n//go:build wasip2\n" -e "\\.out/bindings/wasi/random/" "" ./insecure/ ./insecure-seed/ ./random/

package main
