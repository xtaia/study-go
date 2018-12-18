package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone   eee", "The greeting object.")

}

func main() {
	flag.Parse()
	hello(name)
}
