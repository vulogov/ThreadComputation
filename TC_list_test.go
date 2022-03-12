package ThreadComputation

import (
	"fmt"
	"testing"
)

func TestList1(t *testing.T) {
	l := &TCList{}
	if l.Len() != 0 {
		t.Error("TCList creation had failed")
	}
}

func TestList2(t *testing.T) {
	l := &TCList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	if l.Len() != 3 {
		t.Error("TCList.Add() creation had failed")
	}
}

func TestList3(t *testing.T) {
	l := &TCList{}
	l.Add(int64(1))
	l.Add(int64(2))
	l.Add(int64(3))
	l.Add("Hello World!")
	l.Add(true)
	fmt.Println(l.String())
}
