package ThreadComputation

import (
  "fmt"
)



type TCExecuteFun func(*TCExecListener, interface{}) interface{}

func TCExecuteEcho(l *TCExecListener, code interface{}) interface{} {
  return code
}

func TCExecuteEval(l *TCExecListener, code interface{}) interface{} {
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
