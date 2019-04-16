package fronteir

import "container/list"

//Stack structure data stack
type Stack struct {
	list list.List
}

//Add a element in the stack
func (stack *Stack) Add(element interface{}) {
	stack.list.PushBack(element)
}

//Pop remove the element
func (stack *Stack) Pop() interface{} {
	return stack.list.Remove(stack.list.Back())
}

//Peek see the last element without remove then
func (stack *Stack) Peek() interface{} {
	return stack.list.Back()
}

//Size return the size to the list
func (stack *Stack) Size() int {
	return stack.list.Len()
}
