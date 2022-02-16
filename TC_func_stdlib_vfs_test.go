package ThreadComputation

import (
	"fmt"
	"os"
	"testing"
)



func TestVfs1(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error detecting CWD")
	}
	script := fmt.Sprintf("\"file://%s/examples/hello.txt\" read", cwd)
	tc := Init()
  tc1 := tc.Eval(script)
	if tc1.Errors() != 0 {
		t.Fatalf("%v parse failed", script)
	}
}

func TestVfs2(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error detecting CWD")
	}
	script := fmt.Sprintf("\"file://%s/examples/hello.txt\" read", cwd)
	tc := Init()
  tc1 := tc.Eval(script)
	if tc1.Errors() != 0 {
		t.Fatalf("%v parse failed", script)
	}
	res := tc.GetAsString()
	if res != "Hello world!\n" {
		t.Fatalf(" %v had failed: %v", script, res)
	}
}

func TestVfs3(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error detecting CWD")
	}
	script := fmt.Sprintf("\"file://%s/examples/hello.txt\" read strip", cwd)
	tc := Init()
  tc1 := tc.Eval(script)
	if tc1.Errors() != 0 {
		t.Fatalf("%v parse failed", script)
	}
	res := tc.GetAsString()
	if res != "Hello world!" {
		t.Fatalf(" %v had failed: %v", script, res)
	}
}

func TestVfs4(t *testing.T) {
	tc := Init()
  tc1 := tc.Eval("\"@/examples/hello.txt\" read strip")
	if tc1.Errors() != 0 {
		t.Fatalf("\"@/examples/hello.txt\" read strip parse failed")
	}
	res := tc.GetAsString()
	if res != "Hello world!" {
		t.Fatalf("\"@/examples/hello.txt\" read strip had failed: %v", res)
	}
}
