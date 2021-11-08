package main

import (
	"flag"
	"lib"
)

var name string

func init() {
	//fmt.Print("init...\n")
	flag.StringVar(&name, "name", "everyone", "The greeting obj")

}

func main() {
	flag.Parse()
	//fmt.Printf("hello, %s!!\n", *name)
	lib.Hello(name)
}
