package ThreadComputation
import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func TCFillCtx(res *TCCtx, q *deque.Deque) *TCCtx {
  for q.Len() > 0 {
    e := q.PopFront()
    switch p := e.(type) {
    case *TCPair:
      switch key := p.X.(type) {
      case string:
        res.C.Delete(key)
        res.C.Store(key, p.Y)
      }
    }
  }
  return res
}

func tcSetCtx(l *TCExecListener, q *deque.Deque) {
  if (q.Len() % 2) != 0 {
    return
  }
  for q.Len() > 0 {
    k := q.PopFront()
    v := q.PopFront()
    switch k.(type) {
    case string:
      log.Debugf("local-context: %v=%v", k,v)
      l.TC.SetContext(k.(string), v)
    }
  }
}

func TCPopulateCtxFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  tcSetCtx(l, q)
  return nil, nil
}

func TCSetLocalCtxFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := MakeContext()
  res = TCFillCtx(res, q)
  l.TC.AddContext(res)
  return nil, nil
}

func TCCtxFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := MakeContext()
  res = TCFillCtx(res, q)
  return res, nil
}

func TCDropCtxFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  l.TC.DelContext()
  return nil, nil
}

func TCGetCtxFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := l.TC.LastContext()
  if res == nil {
    return nil, l.TC.MakeError("context stack is shallow")
  }
  return res, nil
}

func TCCtxValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var val interface{}

  res := MakeContext()
  if q.Len() > 0 {
    val = q.PopFront()
  }
  res = TCFillCtx(res, q)
  l.TC.AddContext(res)
  l.TC.SetContext("value", val)
  return nil, nil
}

func init() {
  SetCommand("local", TCSetLocalCtxFunction)
  SetCommand("here", TCPopulateCtxFunction)
  SetCommand("context", TCCtxFunction)
  SetCommand("dropcontext", TCDropCtxFunction)
  SetCommand("getcontext", TCGetCtxFunction)
  SetFunction("&", TCCtxValueFunction)
}
