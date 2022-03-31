package ThreadComputation

import (
	"testing"
)

func TestExit1(t *testing.T) {
	// SetVariable("tc.Debuglevel", "debug")
	tc := Init()
	// SetVariable("tc.Debuglevel", "info")
	tc = tc.Eval("exit")
	if ! tc.ExitRequested() {
		t.Fatalf("Fail to request an exit")
	}
}
