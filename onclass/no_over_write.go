package main

import "fmt"

type Parent struct {

}
// 父类的方法
func (p Parent) SayHello() {
	fmt.Println("I am " + p.name())
}

func (p Parent) name() string {
	return "Parent"
}

type Son struct {
	Parent
}

func (s Son) name() string {
	return "Son"
}

func main() {
	son := Son{
		Parent{},
	}
	son.SayHello()
}