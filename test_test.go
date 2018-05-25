package linkedHashMap

import (
	"testing"
	"fmt"
)

type A struct {
	name string
}

func TestLinkhhhMap(t *testing.T) {
	a:=&A{name:"121323"}
	var b *A
	b = &A{}
	b=a
	b.name="fgfsdf"
	fmt.Println(a,b)
}
