package random_test

import (
	"os/exec"
	"testing"
)

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
