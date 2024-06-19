package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "fmt", "./...") //
	cmd.Env = append(cmd.Environ(), "GOTOOLCHAIN=auto")
	// Redirect output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
