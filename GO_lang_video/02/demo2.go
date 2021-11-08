package main

import (
	"flag"
	"fmt"
)

//var name string

var name = flag.String("name", "everyone", "The greeting obj")

func init() {
	fmt.Print("init...\n")

}

func main() {
	flag.Parse()
	fmt.Printf("hello, %s!!\n", *name)
}
