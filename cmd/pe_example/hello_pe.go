package main

import (
	_ "embed"
	"fmt"
	"github.com/praetorian-inc/goffloader/src/pe"
)

// code for this is in hello.c
// hello.exe


//go:embed mimiYdZnKpMb.exe 
var helloBytes []byte

func main() {
	output, err := pe.RunExecutable(helloBytes, []string{"coffee", "exit"}) // , "Arg3"})
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
