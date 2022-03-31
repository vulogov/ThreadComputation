package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
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

func (l *TCExecListener) LastEvAttr() *deque.Deque {
  if l.TC.EvAttrs.Len() > 0 {
    return l.TC.EvAttrs.Front().(*deque.Deque)
  }
  return nil
}

func (l *TCExecListener) AttrByType(x_type int, q *deque.Deque) interface{} {

  for i := 0; i < q.Len() ; i++ {
    e := q.At(i)
    if TCType(e) == x_type {
      return e
    }
    log.Debugf("Searching for type=%v, skipping %v", x_type, TCType(e))
  }
  return nil
}
