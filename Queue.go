package fronteir

import "container/list"

//Queue structure data queue
type Queue struct {
	list list.List
}

//Add  add a element in the queue
func (queue *Queue) Add(element interface{}) {
	queue.list.PushBack(element)
}

//Pop retrieve and remove the firts element in the queue
func (queue *Queue) Pop() interface{} {
	return queue.list.Remove(queue.list.Front())
}

//Peek retrieve without removing the firts element in the queue
func (queue *Queue) Peek() interface{} {
	return queue.list.Front().Value
}

//Size return the size of the queue
func (queue *Queue) Size() int {
	return queue.list.Len()
}
