package fronteir

//Fronteir interface
type Fronteir interface {
	Add(element interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
}
