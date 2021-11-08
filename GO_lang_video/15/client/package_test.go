package client

import (
	"fmt"
	"go_learn/GO_lang_video/15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	if v, err := series.GetFibonacci(-10); err != nil {
		fmt.Println("something wrong", err)
	} else {
		fmt.Println(v)
	}
}
