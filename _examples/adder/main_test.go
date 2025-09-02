//go:generate wkg wit fetch
//go:generate rm -rf ./internal/
//go:generate go tool wit-bindgen-go generate --out ./internal/ --versioned ./wit/
//go:generate rm -rf ./internal/wasi/random/v0.2.0/
//go:generate go tool jet -g "*.go" "github\\.com/jcbhmr/go-wasi-random/0\\.2\\.0/_examples/adder/internal/wasi/random/v0\\.2\\.0/" "github.com/jcbhmr/go-wasi-random/0.2.0/" ./internal/

package main_test

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"testing"
)

var wasmExe string

func TestMain(m *testing.M) {
	err := os.MkdirAll(".out", 0755)
	if err != nil {
		log.Fatalf("failed to create directory recursively %q: %v", ".out", err)
	}
	cmd := exec.Command("tinygo", "build", "-buildmode", "c-shared", "-o", ".out/wasm.test.wasm", "-wit-package", "wit", "-wit-world", "adder")
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

func TestAddRandom(t *testing.T) {
	cmd := exec.Command("wasmtime", "--invoke", "add-random(100)", wasmExe)
	t.Logf("Running %q", cmd)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		t.Logf("%s", output)
	}
	if err != nil {
		t.Logf("failed to run command %v: %v", cmd, err)
	}
	result, err := strconv.ParseInt(string(output), 10, 32)
	if err != nil {
		t.Logf("failed to parse output %q: %v", output, err)
	}
	if !(100 <= result && result <= 200) {
		t.Errorf("expected result in %s, got %d", "[100,200]", result)
	}
}

func TestCheckWIT(t *testing.T) {
	cmd := exec.Command("wasm-tools", "component", "wit", wasmExe)
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		t.Logf("%s", output)
	}
	if err != nil {
		t.Logf("failed to run command %v: %v", cmd, err)
	}
	if !bytes.Contains(output, []byte("package math:adder")) {
		t.Errorf("expected %q in output", "package math:adder")
	}
}
