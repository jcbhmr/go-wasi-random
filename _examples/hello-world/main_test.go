//go:generate wkg wit fetch
//go:generate rm -rf ./internal/
//go:generate go tool wit-bindgen-go generate --out ./internal/ ./wit/
//go:generate rm -rf ./internal/wasi/random/
//go:generate go tool jet -g "*.go" "github\\.com/jcbhmr/go-wasi-random/_examples/hello-world/internal/wasi/random/" "github.com/jcbhmr/go-wasi-random/" ./internal/

package main_test

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var wasmExe string

func TestMain(m *testing.M) {
	err := os.MkdirAll(".out", 0755)
	if err != nil {
		log.Fatalf("failed to create directory recursively %q: %v", ".out", err)
	}
	cmd := exec.Command("tinygo", "build", "-buildmode", "default", "-o", ".out/wasm.test.wasm", "-wit-package", "wit", "-wit-world", "my-world")
	cmd.Env = append(os.Environ(), "GOOS=wasip2", "GOARCH=wasm")
	output, err := cmd.CombinedOutput()
	if err != nil {
		if len(output) > 0 {
			log.Printf("%s", output)
		}
		log.Fatalf("failed to run command %v: %v", cmd, err)
	}
	wasmExe, err = filepath.Abs(".out/wasm.test.wasm")
	if err != nil {
		log.Fatalf("failed to get absolute path of %q: %v", ".out/wasm.test.wasm", err)
	}
	os.Exit(m.Run())
}

func TestRun(t *testing.T) {
	cmd := exec.Command("wasmtime", wasmExe)
	t.Logf("Running %q", cmd)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		t.Logf("%s", output)
	}
	if err != nil {
		t.Logf("failed to run command %v: %v", cmd, err)
	}
}

func TestPrintWIT(t *testing.T) {
	cmd := exec.Command("wasm-tools", "component", "wit", wasmExe)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		t.Logf("%s", output)
	}
	if err != nil {
		t.Logf("failed to run command %v: %v", cmd, err)
	}
}
