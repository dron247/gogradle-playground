package main

import (
	"fmt"
	"os"
	"time"

	"./stringutil"
)

func main() {
	var un byte = '\x23'
	fmt.Printf("%s %d \n\t %s\n", stringutil.SayHello("World"), un, time.Now())
	fmt.Printf("|%-3d|%8.2f|\n", 12, 345.3)
	fmt.Fprint(os.Stderr, "error\n")
	stringutil.Echo("Echo")
}
