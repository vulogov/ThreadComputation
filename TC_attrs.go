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
