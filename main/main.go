package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var un byte = '\x23'
	fmt.Printf("Hello werold %d \n\t %s\n", un, time.Now())
	fmt.Printf("|%-3d|%8.2f|\n", 12, 345.3)
	fmt.Fprint(os.Stderr, "error\n")
}
