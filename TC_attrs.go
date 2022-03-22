package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func (l *TCExecListener) Attrs() *deque.Deque {
  if l.TC.Attrs.GLen() > 0 {
    q := l.TC.Attrs.Q()
    l.TC.Attrs.Del()
    return q
  }
  return deque.New()
}

func AttrsToArray(q *deque.Deque) []interface{} {
  var res []interface{}
  for x := 0; x < q.Len(); x++ {
    res = append(res, q.At(x))
  }
  return res
}
