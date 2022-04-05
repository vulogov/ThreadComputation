package ThreadComputation

import (
  "github.com/gammazero/deque"
)



func TCLoopPPairCallback(l *TCExecListener, v interface{}, q *deque.Deque) *deque.Deque {
  switch e := v.(type) {
  case *TCPair:
    q.PushBack(e.X)
    q.PushBack(e.Y)
    switch e.X.(type) {
    case string:
      l.TC.SetContext(e.X.(string), e.Y)
    }
  }
  return q
}

func init() {
  RegisterLoopPCallback(Pair, TCLoopPPairCallback)
}
