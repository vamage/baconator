/*
Package Main
Implements Agent code.
*/
package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {
	cmd := exec.Command("/bin/sleep", "5")
	var out strings.Builder
	cmd.Stdout = &out
	channel := make(chan error)
	go func() {
		channel <- cmd.Run()
	}()
Koop:
	for {
		select {
		case err := <-channel:
			if err != nil {
				fmt.Printf("Error %+v", err)
			}
			break Koop
		case <-time.After(3 * time.Second):
			fmt.Printf("Beep")
		}
	}

	fmt.Printf("status %+v, %+v,", cmd.Process, out)
}
