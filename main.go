// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		cmd := exec.Command("./quadchecker", "quadA", input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			fmt.Println("error", err)
			return
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Println("error", err)
			return
		}
	}
}
