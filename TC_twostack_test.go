package ThreadComputation

import (
	"testing"
)

func TestInitTS(t *testing.T) {
	ts := InitTS()
	if ts.Status != true {
		t.Fatalf("TwoStack is failed to initialize")
	}
}

func TestSetGet(t *testing.T) {
	ts := InitTS()
	ts.Set("42")
	r, _ := ts.Get()
	if r != "42" {
		t.Error("ts.Set()/ts.Get() had failed")
	}
}

func TestTake(t *testing.T) {
	ts := InitTS()
	ts.Set("42")
	r, _ := ts.Take()
	if r != "42" {
		t.Error("#1 ts.Take() had failed")
	}
	if ts.Len() != 0 {
		t.Error("#2 ts.Take() had failed")
	}
}

func TestMode(t *testing.T) {
	ts := InitTS()
	ts.Set("42")
	ts.Set("41")
	ts.Reverse()
	r, _ := ts.Get()
	if r != "42" {
		t.Errorf("ts.Mode() had failed: %v", r)
	}
	ts.Normal()
	r, _ = ts.Get()
	if r != "41" {
		t.Errorf("ts.Mode() had failed: %v", r)
	}
}

func TestLeftRight(t *testing.T) {
	ts := InitTS()
	ts.Set("42")
	ts.Add()
	ts.Set("41")
	ts.Left()
	r, _ := ts.Get()
	if r != "42" {
		t.Errorf("#1 ts.Left() had failed, as got %s", r)
	}
	ts.Right()
	r, _ = ts.Get()
	if r != "41" {
		t.Errorf("#2 ts.Right() had failed, as got %s", r)
	}
}


func TestGLen(t *testing.T) {
	ts := InitTS()
	if ts.GLen() != 0 {
		t.Error("ts.GLen() had failed")
	}
}

func TestLen(t *testing.T) {
	ts := InitTS()
	ts.Set("42")
	if ts.Len() != 1 {
		t.Error("ts.Len() had failed")
	}
}

func TestQ(t *testing.T) {
	ts := InitTS()
	ts.Set("42")
	q := ts.Q()
	if q.Len() != 1 {
		t.Error("ts.Q() had failed")
	}
}

func TestDel(t *testing.T) {
	ts := InitTS()
	ts.Del()
	if ts.GLen() != 0 {
		t.Error("ts.Del() had failed")
	}
}


func TestUUID(t *testing.T) {
	ts := InitTS()
	if len(ts.ID) == 0 {
		t.Errorf(".ID generation is failed")
	}
}
