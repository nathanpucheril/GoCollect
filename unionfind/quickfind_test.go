package unionfind

import (
	"fmt"
	"testing"
)

func TestNewQuickFind(t *testing.T) {
	qu := NewQuickFind(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Connect all odds
	qu.Union(3, 1)
	qu.Union(3, 5)
	qu.Union(7, 1)
	qu.Union(9, 3)

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			connected, err := qu.IsConnected(i, j)

			if err != nil {
				t.Fail()
			}
			if connected && i%2 != j%2 {
				fmt.Print("hi")
				t.Fail()
			}
		}
	}

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			connected, err := qu.IsConnected(i, j)
			if err != nil {
				t.Fail()
			}
			if (i%2 == 0 || j%2 == 0) && connected { // events are not connected to anything
				fmt.Println(i, j)
				t.Fail()
			}
		}
	}

	qu.Add(15)
	qu.Add(16)
	fmt.Println(qu.set)
}
