package ThreadComputation

import (
  "github.com/gammazero/deque"
)



func TCLoopPDictCallback(l *TCExecListener, v interface{}, q *deque.Deque) *deque.Deque {
  switch e := v.(type) {
  case *TCDict:
    for item := range e.D.Items() {
      switch item.Key.(type) {
      case string:
        l.TC.SetContext(item.Key.(string), item.Value)
      }
    }
  }
  return q
}

func init() {
  RegisterLoopPCallback(Dict, TCLoopPDictCallback)
}
