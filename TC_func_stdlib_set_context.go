package ThreadComputation

import (
)

func TCContextSet(x interface{}, y interface{}) {

  switch d := x.(type) {
  case *TCCtx:
    switch v := y.(type) {
    case *TCPair:
      switch key := v.X.(type) {
      case string:
        val := v.Y
        d.C.Store(key, val)
      }
    }
  }
}

func init() {
  RegisterOperatorCmdCallback("set", Context, Pair, TCContextSet)
}
