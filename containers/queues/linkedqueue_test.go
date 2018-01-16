package queues

import "testing"

var testQueue LinkedQueue = NewLinkedQueue()

func TestLinkedQueue_Add(t *testing.T) {
	testQueue.Add("hi")
	if value, ok := testQueue.Peek(); !(value == "hi" && ok) {
		t.Fail()
	}
	testQueue.Add("hi")
}
