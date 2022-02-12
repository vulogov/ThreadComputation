package ThreadComputation

import (
  "github.com/vulogov/ThreadComputation/parser"
)


func (l *TCExecListener) EnterDblock(c *parser.DblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  l.TC.Res.Add()
}

func (l *TCExecListener) ExitDblock(c *parser.DblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  l.TC.Res.Left()
}
