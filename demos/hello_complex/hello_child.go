package main;

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("[PID %d] Hello from child\n", os.Getpid())
}
