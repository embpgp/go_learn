package main

import (
	"flag"
	"fmt"
)

var name string

func init(){
	flag.StringVar(&name, "name", "everyone", "The greeting obj")

}

func main(){
	flag.Parse()
	fmt.Printf("hello, %s!!\n", name)
}
