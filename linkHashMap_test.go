package linkedHashMap

import (
	"testing"
	"fmt"
)

func TestLinkhashMap(t *testing.T) {
	lhm := NewLinkedHashMap()
	lhm.Init(1, true)
	lhm.Put(1, 1)
	lhm.Put(3, 1)
	lhm.Base()
	lhm.Put(2, 1)
	//lhm.Base()
	lhm.Remove(1)
	//lhm.Base()
	//lhm.Put(1, 1)
	lhm.Put(3, 1)
	fmt.Println(lhm.Base().after.key,111111111)
	//fmt.Println(lhm.LinkSize())
	//fmt.Println(lhm.Base().after.key)
}
