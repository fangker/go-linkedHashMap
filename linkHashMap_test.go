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
	//lhm.Put(2, 1)
	//lhm.Remove(2)
	//lhm.Put(3, 1)
	fmt.Println(lhm.base.after.key)
}
