package ThreadComputation

import (
  "os"
  "errors"
  "fmt"
  "github.com/vulogov/ThreadComputation/parser"
)


func (l *TCExecListener) EnterVars(c *parser.VarsContext) {
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun(fmt.Sprintf("$%v",c.GetVname().GetText())) {
    return
  }
  if l.TC.Ready() {
    vname := c.GetVname().GetText()
    l.TC.Vars.Delete(vname)
    e := l.TC.Get()
    l.TC.Vars.Store(vname, e)
  }
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
  SetVariable("tc.Debuglevel", "info")
  SetVariable("tc.Logoutput", os.Stderr)
}
