package ThreadComputation

import (
  "reflect"
  "runtime"
  log "github.com/sirupsen/logrus"
)

func ReturnFromFunction(l *TCExecListener, func_name string, data interface{}) {
  log.Debugf("return from fun=%v() res=%T", func_name, data)
  if l.TC.Attrs.GLen() >= 1 {
    log.Debugf("returning function value to attribute: %v->%v", func_name, data)
    l.TC.Attrs.Set(data)
  } else {
    log.Debugf("returning function value to stack: %v->%v", func_name, data)
    l.TC.Res.Set(data)
  }
}

func (tc *TCstate) SetFunction(name string, fun TCFun) {
  tc.Functions.Delete(name)
  tc.Functions.Store(name, fun)
}

func (tc *TCstate) GetFunction(name string) TCFun {
  if fun, ok := tc.Functions.Load(name); ok {
    return fun.(TCFun)
  }
  if fun, ok := Functions.Load(name); ok {
    return fun.(TCFun)
  }
  return nil
}

func (tc *TCstate) HaveFunction(name string) bool {
  if _, ok := tc.Functions.Load(name); ok {
    return true
  }
  if _, ok := Functions.Load(name); ok {
    return true
  }
  return false
}

func (tc *TCstate) SetCommand(name string, fun TCFun) {
  tc.Commands.Delete(name)
  tc.Commands.Store(name, fun)
}

func (tc *TCstate) GetCommand(name string) TCFun {
  if fun, ok := tc.Commands.Load(name); ok {
    return fun.(TCFun)
  }
  if fun, ok := Commands.Load(name); ok {
    return fun.(TCFun)
  }
  return nil
}

func (tc *TCstate) HaveCommand(name string) bool {
  if _, ok := Commands.Load(name); ok {
    return true
  }
  if _, ok := tc.Commands.Load(name); ok {
    return true
  }
  return false
}

func SetFunction(name string, fun TCFun) {
  Functions.Delete(name)
  Functions.Store(name, fun)
}

func SetCommand(name string, fun TCFun) {
  Commands.Delete(name)
  Commands.Store(name, fun)
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
