package lists

type List interface {
	Insert(index int, value interface{})
	IndexOf() int
	RemoveAtIndex(index int) interface{}
	Remove(items ...interface{}) interface{}
	//RemoveContainer(c Container) interface{}
	Set(index int, value interface{})
	Get(index int) interface{}
	Contains(items ...interface{}) bool
	LastIndexOf() int
	Sublist(to, from int) List
	//Size() int
	//IsEmpty() bool
}

//type list struct {
//	List
//}

//func (self *list) IsEmpty() bool {
//	return self.Size() == 0
//}

//func validIndex(list List, index int) bool {
//	return 0 <= index && index < list.Size()
//}
