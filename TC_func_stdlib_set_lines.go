package ThreadComputation

import (
)

func TCLinesSet(x interface{}, y interface{}) {

  switch d := x.(type) {
  case *TCLines:
    switch v := y.(type) {
    case *TCPair:
      switch key := v.X.(type) {
      case int64:
        val := v.Y
        d.Q.Set(int(key), val)
      }
    }
  }
}

func init() {
  RegisterOperatorCmdCallback("set", Lines, Pair, TCLinesSet)
}
