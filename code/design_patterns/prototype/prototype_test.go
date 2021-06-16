package prototype

import (
	"fmt"
	"testing"
)

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	fmt.Println(t1)
	t2 := t1.Clone()
	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}
