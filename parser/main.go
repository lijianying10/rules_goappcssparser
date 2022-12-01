package main

import (
	"flag"
	"fmt"
)

type arrayInputFiles []string

func (a *arrayInputFiles) String() string {
	if len(*a) == 0 {
		return ""
	}
	var res string
	for _, v := range *a {
		res += v + ","
	}
	return res[:len(res)-1]
}

func (a *arrayInputFiles) Set(val string) error {
	*a = append(*a, val)
	return nil
}

var input arrayInputFiles
var output string

func main() {
	flag.Var(&input, "i", "input file")
	flag.StringVar(&output, "o", "", "set file name")
	flag.Parse()
	if output == "" {
		panic("empty output")
	}
	if len(input) == 0 {
		panic("empty input")
	}
	p := Parse{}
	err := p.IO(input, output)
	if err != nil {
		panic(err)
	}
	fmt.Println("debug: ", input, output)
}
