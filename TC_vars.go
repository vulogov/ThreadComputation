package ThreadComputation

import (
  "os"
  "fmt"
  "errors"
  "github.com/vulogov/ThreadComputation/parser"
)

func (l *TCExecListener) EnterVars(c *parser.VarsContext) {
  if l.TC.Errors() > 0 {
    return
  }
}

func (tc *TCstate) SetVariable(name string, data interface{}) {
  tc.Vars.Delete(name)
  tc.Vars.Store(name, data)
}

func SetVariable(name string, data interface{}) {
  Vars.Delete(name)
  Vars.Store(name, data)
}

func GetVariable(name string) (interface{}, error) {
  if data, ok := Vars.Load(name); ok {
    return data, nil
  }
  return nil, errors.New(fmt.Sprintf("Variable %v not found", name))
}



func initStdVars() {
  SetVariable("tc.Version", VERSION)
  SetVariable("tc.Maxfilesize", 16777216)
  SetVariable("tc.Chancapacity", 102400)
  SetVariable("tc.Poolsize", 25)
  SetVariable("tc.Debuglevel", "info")
  SetVariable("tc.Logoutput", os.Stderr)
}
