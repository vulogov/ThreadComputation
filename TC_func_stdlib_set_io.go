package ThreadComputation

import (
)

func TCIOSet(x interface{}, y interface{}) {

  switch d := x.(type) {
  case *TCIO:
    switch v := y.(type) {
    case *TCPair:
      d.Add(v)
    case string:
      d.Add(v)
    }
  }
}

func init() {
  RegisterOperatorCmdCallback("set", IO, Pair, TCIOSet)
  RegisterOperatorCmdCallback("set", IO, String, TCIOSet)
}
