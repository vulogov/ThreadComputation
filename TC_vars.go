package ThreadComputation

import (
  "github.com/vulogov/ThreadComputation/parser"
)


func (l *TCExecListener) EnterVars(c *parser.VarsContext) {
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.Ready() {
    vname := c.GetVname().GetText()
    l.TC.Vars.Delete(vname)
    e := l.TC.Get()
    l.TC.Vars.Store(vname, e)
  }
}
