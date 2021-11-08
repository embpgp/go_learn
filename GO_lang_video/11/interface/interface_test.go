package interface_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

type GoProgrammer1 struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}

func (g *GoProgrammer1) WriteHelloWorld() string {
	return "fmt.Println(\"hello world111\")"
}
func TestClient(t *testing.T) {
	var p Programmer
	var p1 Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	p1 = new(GoProgrammer1)
	t.Log(p1.WriteHelloWorld())
}
