package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"codecalc.com/m/v2/listener"
	"github.com/antlr4-go/antlr/v4"
)

var (
	fasmOut = flag.String("fasm", "", "Output file for generated fasm assembly.")
	debug   = flag.Bool("debug", false, "Print debug messages.")
	run     = flag.Bool("run", false, "Immediatly run the binary.")
)

func main() {
	flag.Parse()
	_, filename, _, ok := runtime.Caller(0)
	directory := path.Dir(filename)
	if !ok {
		panic("No caller information")
	}
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Requires 2 and only 2 positional commandline arguments, inputfilename and outputfilename. Got:")
		fmt.Println(strings.Join(args, ", "))
		return
	}

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	is := antlr.NewIoStream(bufio.NewReader(f))

	module, err := listener.ParseStream(is)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *debug {
		fmt.Println(module.String())
	}

	fasmFilename := *fasmOut
	if fasmFilename == "" {
		fasmFile, err := os.CreateTemp(directory+"/temp/", "*.inc")
		if err != nil {
			fmt.Println("Error creating temporary FASM file")
			fmt.Println(err)
			return
		}
		defer os.Remove(fasmFile.Name())
		fasmFilename = fasmFile.Name()
		_, err = fasmFile.WriteString(module.Generate())
		if err != nil {
			fmt.Println("Error generating fasm")
			fmt.Println(err)
			return
		}
	} else {
		err = os.WriteFile(fasmFilename, []byte(module.Generate()), 0777)
		if err != nil {
			fmt.Println("Error generating fasm")
			fmt.Println(err)
			return
		}
	}

	cmd := exec.Command(directory+"/fasm/fasm", fasmFilename, args[1])
	cmd.Stderr = os.Stderr
	stdout, err := cmd.Output()
	if *debug {
		fmt.Println(string(stdout))
	}
	if err != nil {
		fmt.Println("Error Building Binary")
		fmt.Println(err)
		return
	}

	if *run {
		cmd := exec.Command(args[1])
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error Running Binary")
			fmt.Println(err)
			return
		}
	}
}
