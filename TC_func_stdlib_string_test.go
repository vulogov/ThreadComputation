package ThreadComputation

import (
	"testing"
)



func TestString1(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("\"Hello\" upper")
	if tc1.Errors() != 0 {
		t.Fatalf("\"Hello\" upper parse failed")
	}
	res := tc.GetAsString()
	if res != "HELLO" {
		t.Fatalf("\"Hello\" upper parse failed: %v", res)
	}
}

func TestString2(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("\"HelLo\" lower")
	if tc1.Errors() != 0 {
		t.Fatalf("\"HelLo\" lower parse failed")
	}
	res := tc.GetAsString()
	if res != "hello" {
		t.Fatalf("\"HelLo\" lower parse failed: %v", res)
	}
}

func TestString3(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("\"helLo\" lower title")
	if tc1.Errors() != 0 {
		t.Fatalf("\"helLo\" lower title parse failed")
	}
	res := tc.GetAsString()
	if res != "Hello" {
		t.Fatalf("\"helLo\" lower title parse failed: %v", res)
	}
}

func TestString4(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("\"kitten\" distance[\"sitting\"]")
	if tc1.Errors() != 0 {
		t.Fatalf("\"kitten\" distance[\"sitting\"] parse failed")
	}
	res := tc.GetAsString()
	if res != "3" {
		t.Fatalf("\"kitten\" distance[\"sitting\"] parse failed: %v", res)
	}
}
