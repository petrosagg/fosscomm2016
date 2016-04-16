package main;

import (
	"os"
	"os/exec"
)

func main() {
	os.Remove("/bin/sh")
	os.Symlink("/bin/sh.real", "/bin/sh")

	cmd := exec.Command("/usr/bin/qemu-arm", os.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	os.Remove("/bin/sh")
	os.Symlink("/bin/sh.shim", "/bin/sh")
}
