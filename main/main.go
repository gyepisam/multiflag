/*
Package main shows a usage example for the multiflag package.

After building this program with

	$ go build -o multiflag

Run it with these arguments
	
	$ ./multiflag -v -v -v -v -t parse -t compile

Which should produce the following output:

	Verbosity: 4
	Tracing: parse
	Tracing: compile
*/

package main

import (
	"flag"
	"fmt"
	"github.com/gyepisam/multiflag"
)

func main() {

	var verbosity = multiflag.Bool("verbose", "false", "Verbosity. Repeat as necessary", "v")
	var trace = multiflag.String("trace", "none", "Trace program sections", "t")

	flag.Parse()

	fmt.Println("Verbosity:", verbosity.NArg())

	for _, item := range trace.Args() {
		fmt.Println("Tracing:", item)
	}
}
