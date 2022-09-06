package main

import "fmt"

//Iprocess interface
type IProcess interface {
	process()
}

//adapter struct
type Adapter struct {
	adaptee Adaptee
}

//the Adapter class has a process method that invokes the convert method on adaptee
//adapter class method process
func (adapter Adapter) process() {
	fmt.Println("adpater process")
	adapter.adaptee.convert()
}

//Adaptee struct
type Adaptee struct {
	adapterType int
}

//Adaptee class method convert
func (adaptee Adaptee) convert() {
	fmt.Println("adaptee convert method")
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}
