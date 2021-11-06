package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	/*
		if i, ok := p.(int); ok {
			fmt.Println("Integer", i)
			return
		}
		if i, ok := p.(string); ok {
			fmt.Println("String", i)
			return
		}
		fmt.Println("unknow type")

	*/
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("unknow type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething('c')
}
