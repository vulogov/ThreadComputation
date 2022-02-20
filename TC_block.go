package ThreadComputation

import (
  "github.com/google/uuid"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)


func (l *TCExecListener) EnterDblock(c *parser.DblockContext) {
  var bname string
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("(") == true {
    return
  }
  b := c.GetBname()
  if b == nil {
    bname = uuid.NewString()
  } else {
    bname = b.GetText()
  }
  log.Debugf("Block name: %v", bname)
  l.TC.Res.Add()
}

func (l *TCExecListener) ExitDblock(c *parser.DblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun(")") == true {
    return
  }
  l.TC.Res.Left()
}
