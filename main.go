package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"codecalc.com/m/v2/listener"
	"github.com/antlr4-go/antlr/v4"
)

var DEBUG = false

func main() {
	_, filename, _, ok := runtime.Caller(0)
	directory := path.Dir(filename)
	if !ok {
		panic("No caller information")
	}

	if len(os.Args) < 3 {
		fmt.Println("Requires input and output file commandline arguments")
		return
	}

	f, err := os.Open(os.Args[1])
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	is := antlr.NewIoStream(bufio.NewReader(f))

	module := listener.ParseStream(is)
	if DEBUG {
		fmt.Println(module.String())
	}
	err = os.WriteFile(os.Args[2]+".asm", []byte(module.Generate()), 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd := exec.Command(directory+"/fasm/fasm", os.Args[2]+".asm", os.Args[2])
	stdout, err := cmd.Output()
	if DEBUG {
		fmt.Println(string(stdout))
	}
	if err != nil {
		fmt.Println("Error Building Binary")
		fmt.Println(err)
	}
}
