package ThreadComputation

import (
)

func TCDictSet(x interface{}, y interface{}) {

  switch d := x.(type) {
  case *TCDict:
    switch v := y.(type) {
    case *TCPair:
      switch key := v.X.(type) {
      case string:
        val := v.Y
        d.D.Set(key, val)
      }
    }
  }
}

func init() {
  RegisterOperatorCmdCallback("set", Dict, Pair, TCDictSet)
}
