//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	// Create a process
	proc := MockProcess{}

	quit := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		signal.Notify(quit, os.Interrupt)
		go func() {
			<-quit
			os.Exit(0)
		}()
		proc.Stop()
		done <- struct{}{}
	}()
	// Run the process (blocking)
	proc.Run()

	<-done
	log.Println("Process stoped")
}
