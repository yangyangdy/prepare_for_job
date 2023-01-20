package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func allocate(){
	for n := 1; n<100; n++ {
		_ = make([]byte, 1<<20)
	}

}
func printGCStats() {
	t := time.NewTicker(time.Millisecond)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}
func main() {
	go printGCStats()
	allocate()
}
