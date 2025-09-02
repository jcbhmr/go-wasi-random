//go:generate rm -rf ./.out/bindings/ ./insecure/ ./insecure-seed/ ./random/
//go:generate go tool wit-bindgen-go generate --out ./.out/bindings/ ./wit/
//go:generate mv ./.out/bindings/wasi/random/insecure/ ./insecure/
//go:generate mv ./.out/bindings/wasi/random/insecure-seed/ ./insecure-seed/
//go:generate mv ./.out/bindings/wasi/random/random/ ./random/
//go:generate rm -rf ./insecure/empty.s ./insecure-seed/empty.s ./random/empty.s
//go:generate go tool jet -g "*.go" -e "DO NOT EDIT.\n" "DO NOT EDIT.\n\n//go:build wasip2\n" -e "\\.out/bindings/wasi/random/" "" ./insecure/ ./insecure-seed/ ./random/

package random_test

import (
	"os/exec"
	"testing"
)

//go:generate go generate ./_examples/adder
func TestExampleAdder(t *testing.T) {
	cmd := exec.Command("go", "test", "./_examples/adder", "-v", "-count=1")
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		t.Logf("%s", output)
	}
	if err != nil {
		t.Fatalf("failed to run command %v: %v", cmd, err)
	}
}

//go:generate go generate ./_examples/hello-world
func TestExampleHelloWorld(t *testing.T) {
	cmd := exec.Command("go", "test", "./_examples/hello-world", "-v", "-count=1")
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		t.Logf("%s", output)
	}
	if err != nil {
		t.Fatalf("failed to run command %v: %v", cmd, err)
	}
}
