package ThreadComputation

import (
  "github.com/gammazero/deque"
)



func TCLoopPValueCallback(l *TCExecListener, v interface{}, q *deque.Deque) *deque.Deque {
  switch e := v.(type) {
  case *TCValue:
    l.TC.SetContext("probability", float64(e.P))
    lpfun := GetLoopPCallback(e.Value)
    if lpfun != nil {
      return lpfun(l, e.Value, q)
    }
  }
  return q
}

func init() {
  RegisterLoopPCallback(Value, TCLoopPValueCallback)
}
