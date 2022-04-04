package ThreadComputation
import (
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

func init() {
  SetCommand("local", TCSetLocalCtxFunction)
  SetCommand("context", TCCtxFunction)
  SetCommand("dropcontext", TCDropCtxFunction)
  SetCommand("getcontext", TCGetCtxFunction)
}