package ThreadComputation

import (
  "github.com/gammazero/deque"
)



func TCLoopPNumbersCallback(l *TCExecListener, v interface{}, q *deque.Deque) *deque.Deque {
  switch e := v.(type) {
  case *TCNumbers:
    for i := 0; i < e.Len(); i++ {
      val := e.N[i]
      q.PushBack(val)
    }
  }
  return q
}

func init() {
  RegisterLoopPCallback(Numbers, TCLoopPNumbersCallback)
}
