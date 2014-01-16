// Copyright 2014 Gyepi Sam. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* 
Package multiflag implements the flag.Value interface for handling repeated flag values.

It is useful for boolean flags where repeated use implies greater intensity or
values to be collected into an array.


Usage:


Imports, definitions, etc:

  import (
	"flag"
	"github.com/gyepisam/multiflag"
  )

A boolean variable counts flags and does not consume any arguments.

	  var verbosity = multiflag.NewBool("verbose", "false", "Verbosity. Repeat as necessary", "v")

All other variables consume and collect their arguments into a string array.

	  var trace = multiflag.New("trace", "none", "Trace program sections", "t")

After calling

	  flag.Parse()

The following command line flags


You can get the count

	  fmt.Println("Verbosity:", verbosity.Count())

which, given the flags:

	-v -v -verbose --verbose

produces the output:

	Verbosity: 4

or the arguments

	  for _, item := range trace.Args() {
		  fmt.Println("Tracing:", item)
	  }

which, given the flags:

	-t parse -trace compile

produces the output:

	Tracing: parse
	Tracing: compile

The examples above can be found in the file main/main.go, which can also be compiled and run.
It has the following usage text:

  Usage of main:
	-t=none: Alias for trace
	-trace=none: Trace program sections
	-v=false: Alias for verbose
	-verbose=false: Verbosity. Repeat as necessary

*/
package multiflag

import (
	"flag"
)

// Value counts and collects repeated uses of a flag.
type Value struct {
	args   []string // collected flag arguments
	val    string   // default value to display in help
	isBool bool     // denotes if Value represent a boolean value
}

// String produces a string representation.
// Provided for flag package.
func (v *Value) String() string {
	return v.val
}

// Set records a usage instance.
// Provided for flag package.
func (v *Value) Set(s string) error {
	v.args = append(v.args, s)
	return nil
}

// IsBoolFlag returns a value denoting whether the variable represents a boolean value.
// Provided for flag package.
func (v *Value) IsBoolFlag() bool { return v.isBool }

// New returns a multiflag instance.
// name, value, and usage are used to initial a flag.Value.
// aliases, if any, initialize aliases for name, with the usage text produced by the AliasUsage function.
func New(name string, value string, usage string, aliases ...string) *Value {
	v := &Value{val: value}

	flag.Var(v, name, usage)

	for _, alias := range aliases {
		flag.Var(v, alias, AliasUsage(name, alias))
	}

	return v
}

// NewBool returns a multiflag instance that represents a boolean value.
func NewBool(name string, value string, usage string, aliases ...string) *Value {
	v := New(name, value, usage, aliases...)
	v.isBool = true
	return v
}

// Args returns an array of collected arguments.
// A boolean Value, created with NewBool always returns an empty array. 
func (v *Value) Args() []string {
	if v.isBool {
		return []string{}
	} else {
		return v.args
	}
}

// Count returns the number of invocations
func (v *Value) Count() int {
	return len(v.args)
}

// AliasUsageFunc specifies the signature for an alias usage function.
type AliasUsageFunc func(orig, alias string) string

// AliasUsage returns the usage text for an alias.
// The function is a variable that may be changed to point to a custom function of type AliasUsageFunc.
var AliasUsage AliasUsageFunc = func (orig, alias string) string {
	return "Alias for " + orig
}
