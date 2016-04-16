package main;

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Printf("[PID %d] Hello from parent\n", os.Getpid())

	cmd := exec.Command("./hello_child")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
