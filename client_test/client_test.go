package clienttest

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestIntegration(t *testing.T) {
	cmd := exec.Command("go", "run", "../app/server/server.go")
	// cmd := exec.Command("go", "run", "../app/client/client.go", "identity")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
