package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)



type TCExecuteFun func(*TCExecListener, interface{}, *deque.Deque) interface{}

func TCExecuteEcho(l *TCExecListener, code interface{}, q *deque.Deque) interface{} {
  return code
}

func TCExecuteValue(l *TCExecListener, v interface{}, q *deque.Deque) interface{} {
  fun := GetExecuteCallback(v.(*TCValue).Value)
  if fun == nil {
    return nil
  }
  res := fun(l, v.(*TCValue).Value, q)
  if res == nil {
    log.Debugf("Callback for Value(%T) returned nil", v.(*TCValue).Value)
    return nil
  }
  return MakeValue(res, v.(*TCValue).P, 0)
}

func TCExecuteEval(l *TCExecListener, code interface{}, q *deque.Deque) interface{} {
  l.TC.Eval(code.(string))
  return nil
}

func GetExecuteCallback(e interface{}) TCExecuteFun {
  fn := fmt.Sprintf("exec.%v", TCType(e))
  fun, ok := Callbacks.Load(fn)
  if ok {
    return fun.(TCExecuteFun)
  }
  return nil
}

func RegisterExecuteCallback(etype int, fun TCExecuteFun) {
  fname := fmt.Sprintf("exec.%v", etype)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func init() {
  RegisterExecuteCallback(Int, TCExecuteEcho)
  RegisterExecuteCallback(Float, TCExecuteEcho)
  RegisterExecuteCallback(Bool, TCExecuteEcho)
  RegisterExecuteCallback(String, TCExecuteEval)
  RegisterExecuteCallback(Value, TCExecuteValue)
}
