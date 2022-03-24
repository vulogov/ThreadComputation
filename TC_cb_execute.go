package ThreadComputation

import (
  "fmt"
  "github.com/gammazero/deque"
)



type TCExecuteFun func(*TCExecListener, interface{}, *deque.Deque) interface{}

func TCExecuteEcho(l *TCExecListener, code interface{}, q *deque.Deque) interface{} {
  return code
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
}
