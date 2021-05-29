package main

import (
	"os"
	"flag"
	"fmt"
)

//var name string

var name = flag.String("name", "everyone", "The greeting obj")
func init(){
	fmt.Print("init...\n")

}

func main(){
	flag.Usage = func(){
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("hello, %s!!\n", *name)
}
