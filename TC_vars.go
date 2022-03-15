package ThreadComputation

import (
  "os"
  "fmt"
  "errors"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)

func (l *TCExecListener) ExitVars(c *parser.VarsContext) {
  if l.TC.Errors() > 0 {
    return
  }
  var_name := c.GetFname().GetText()
  if l.TC.HaveAttrs() {
    log.Debugf("Variable %v will be taken from attributes", var_name)
    v := l.TC.Attrs.Q().PopBack()
    l.TC.SetVariable(var_name, v)
  } else if l.TC.Ready() {
    log.Debugf("Variable %v will be taken from stack", var_name)
    v := l.TC.Get()
    l.TC.SetVariable(var_name, v)
  } else {
    log.Debugf("Variable %v will be taken from nowhere", var_name)
    l.SetError("Can not get variable %v value in context", var_name)
  }
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
