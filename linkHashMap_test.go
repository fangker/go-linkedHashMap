package linkedHashMap

import (
	"testing"
	"fmt"
)

func TestLinkhashMap(t *testing.T) {
	lhm := NewLinkedHashMap()
	lhm.Init(2222, true)
	lhm.Put(1, 1)
	//lhm.Put(3, 1)
	lhm.Put(2, 1)
	//lhm.Remove(2)
	//lhm.Put(3, 1)
	//fmt.Println(lhm.LinkSize())
	fmt.Println(lhm.Base().after.key)
}
