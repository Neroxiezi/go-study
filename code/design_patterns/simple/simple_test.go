package simple

import "testing"

func TestType1(t *testing.T)  {
	api := NewAPI(1)
	s := api.Say("PFinal")
	if s != "Hi, PFinal" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("PFinal")
	if s != "Hello, PFinal" {
		t.Fatal("Type2 test fail")
	}
}
