package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

type TCFunRef struct {
  Name        string
  Attrs      *deque.Deque
}

func (tc *TCstate) NewRef(name string, q *deque.Deque) *TCFunRef {
  log.Debugf("Creating reference: %v [%v]", name, q.Len())
  res := new(TCFunRef)
  res.Name = name
  res.Attrs = q
  return res
}

func (r *TCFunRef) String() string {
  if r.Attrs != nil {
    return fmt.Sprintf("`%v[ %v element(s) ]", r.Name, r.Attrs.Len())
  }
  return fmt.Sprintf("`%v[ ]", r.Name)
}

func TCRefConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *TCFunRef:
    switch to_type {
    case String:
      return e.String()
    }
  }
  return nil
}

func TCRefExecute(l *TCExecListener, code interface{}, q *deque.Deque) interface{} {
  switch code.(type) {
  case *TCFunRef:
    if code.(*TCFunRef).Attrs != nil {
      for x := 0; x < code.(*TCFunRef).Attrs.Len(); x++ {
        q.PushBack(code.(*TCFunRef).Attrs.At(x))
      }
    }
    res, err := l.ExecFunction(code.(*TCFunRef).Name, q)
    if err != nil {
      l.SetError("Error executing reference: %v", err)
      return nil
    }
    return res
  }
  return nil
}

func init() {
  RegisterConvertCallback(Ref, TCRefConvert)
  RegisterExecuteCallback(Ref, TCRefExecute)
}
