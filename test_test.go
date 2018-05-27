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
	c:=&A{name:"666666"}
	a.haha(c)
	fmt.Println(a)
}
func (this *A)haha(a *A)  {
	*this=*a
}