package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	start := time.Now()

	cmd := exec.Command("go", "test")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
