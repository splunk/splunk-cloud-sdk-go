// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var options struct {
	service    string
	structName string
	iface      string
	pkgName    string
	outputFile string
	comment    string
}

// Prints an error message and exits.
func fatal(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	fmt.Fprintf(os.Stderr, "error: %s\n", msg)
	os.Exit(1)
}

func init() {
	flag.StringVar(&options.service, "svc", "", "service name")
	flag.StringVar(&options.structName, "s", "", "struct to generate interface from")
	flag.StringVar(&options.iface, "i", "", "name of generated interface")
	flag.StringVar(&options.pkgName, "p", "", "pacakge name of generated interface")
	flag.StringVar(&options.outputFile, "o", "", "output file name. Print to stdout if not provided")
	flag.Parse()
}

func main() {
	var err error
	file := filepath.Join(options.service, "service.go")
	outFile := filepath.Join(options.service, "interface.go")
	args := []string{"-f", file, "-s", options.structName, "-i", options.iface, "-p", options.pkgName, "-c", "DO NOT EDIT", "-o", outFile}
	cmd := exec.Command("ifacemaker", args...)
	_, err = cmd.Output()
	if err != nil {
		fatal("%v", err)
	}
}
