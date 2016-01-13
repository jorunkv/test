
package main

import (
    . "fmt"
    "runtime"
    "time"
)
var i int

func someGoroutine() {
   for x :=0; x<1000000; x++ {
    i++
   }
}
func someGoroutine2() {
   for x :=0; x<1000000; x++ {
    i--
   }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    go func(){
        for {}
    }()
    runtime.Gosched()
    go someGoroutine()  
    go someGoroutine2()
                        // This spawns someGoroutine() as a goroutine

    // We have no way to wait f
    //or the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    Println(i)
}
