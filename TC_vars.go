package ThreadComputation

import (
  "os"
  "fmt"
  "errors"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func FuncSetVariable(l *TCExecListener, func_name string, q *deque.Deque) {
  var e interface{}
  log.Debugf("Variable for %v will be processed", func_name)
  if q.Len() > 0 {
    e = q.PopFront()
  } else if l.TC.HaveAttrs() {
    e = l.TC.Attrs.Q().PopFront()
  } else if l.TC.Ready() {
    e = l.TC.Get()
  } else {
    l.TC.SetError("Can not get the value of $%v", func_name)
    return
  }
  l.TC.SetVariable(func_name, e)
}

func (tc *TCstate) SetVariable(name string, data interface{}) {
  log.Debugf("variable %v->%v", data, name)
  tc.Vars.Delete(name)
  tc.Vars.Store(name, data)
}

func SetVariable(name string, data interface{}) {
  Vars.Delete(name)
  Vars.Store(name, data)
}

func (tc *TCstate) GetVariable(name string) (interface{}, error) {
  if data, ok := tc.Vars.Load(name); ok {
    return data, nil
  }
  return GetVariable(name)
}

func GetVariable(name string) (interface{}, error) {
  if data, ok := Vars.Load(name); ok {
    return data, nil
  }
  return nil, errors.New(fmt.Sprintf("Variable %v not found", name))
}

func (tc *TCstate) HasVariable(name string) bool {
  if _, ok := tc.Vars.Load(name); ok {
    return true
  }
  if _, ok := Vars.Load(name); ok {
    return true
  }
  return false
}



func initStdVars() {
  SetVariable("tc.Version", VERSION)
  SetVariable("tc.Maxfilesize", 16777216)
  SetVariable("tc.Chancapacity", 102400)
  SetVariable("tc.Poolsize", 25)
  SetVariable("tc.Debuglevel", "info")
  SetVariable("tc.Logoutput", os.Stderr)
}
