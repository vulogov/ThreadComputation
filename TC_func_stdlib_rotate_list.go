package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
)

func TCRotateListLeft(v interface{}) interface{} {
  log.Debugf("Rotate list left: %T", v)
  switch l := v.(type) {
  case *TCList:
    if l.Q.Len() > 0 {
      e := l.Q.PopFront()
      l.Q.PushBack(e)
      return l
    }
  }
  return v
}

func TCRotateListRight(v interface{}) interface{} {
  log.Debugf("Rotate list right: %T", v)
  switch l := v.(type) {
  case *TCList:
    if l.Q.Len() > 0 {
      e := l.Q.PopBack()
      l.Q.PushFront(e)
      return l
    }
  }
  return v
}

func init() {
  RegisterFunctionCallback("rotateright", List, TCRotateListRight)
  RegisterFunctionCallback("rotateleft", List, TCRotateListLeft)
}
