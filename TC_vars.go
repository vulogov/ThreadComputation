package ThreadComputation

import (
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
