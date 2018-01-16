package containers

type Iterator interface {
	Next() (interface{}, error)
	HasNext() bool
	//Remove() interface{}
}

type Iterable interface {
	Iterator() Iterator
}

type ChanIterable interface {
	ChanIterator() <-chan interface{}
}

func ToChanIter(iterator Iterator) <-chan interface{} {
	chn := make(chan interface{})
	go func() {
		for iterator.HasNext() {
			item, err := iterator.Next()
			if err != nil {
				panic(err) // shouldnt happen right?
			}

			chn <- item
		}
		// Ensure that at the end of the loop we close the channel!
		close(chn)
	}()
	return chn
}
