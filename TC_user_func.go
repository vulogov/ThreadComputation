package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)



func (l *TCExecListener) EnterUfun(c *parser.UfunContext) {
  if l.TC.Errors() > 0 {
    return
  }
  if c.GetFname() == nil {
    l.TC.errmsg = "Can not get function name out of context"
    l.TC.errors += 1
    log.Errorf(l.TC.errmsg)
    return
  }
  func_name := c.GetFname().GetText()
  l.TC.UFNStack.PushFront(func_name)
  l.TC.UFStack.PushFront("")
}

func (l *TCExecListener) ExitUfun(c *parser.UfunContext) {
  if l.TC.Errors() > 0 {
    return
  }
  func_name := c.GetFname().GetText()
  l.TC.UserFun.Delete(func_name)
  code := l.TC.UFStack.PopFront()
  l.TC.UserFun.Store(func_name, code)
  l.TC.UFNStack.PopFront()
}

func (tc *TCstate) AddToUserFun(code string) bool {
  if tc.UFNStack.Len() > 0 {
    c := tc.UFStack.PopFront().(string)
    c += fmt.Sprintf(" %v ", code)
    tc.UFStack.PushFront(c)
    return true
  }
  return false
}
