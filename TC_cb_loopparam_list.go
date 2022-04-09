package ThreadComputation

import (
  "github.com/gammazero/deque"
)



func TCLoopPListCallback(l *TCExecListener, v interface{}, q *deque.Deque) *deque.Deque {
  switch e := v.(type) {
  case *TCList:
    for i := 0; i < e.Len(); i++ {
      val := e.Q.At(i)
      q.PushBack(val)
    }
  }
  return q
}

func init() {
  RegisterLoopPCallback(List, TCLoopPListCallback)
}
