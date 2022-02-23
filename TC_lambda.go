package ThreadComputation

import (
  "github.com/google/uuid"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
  "github.com/vulogov/ThreadComputation/parser"
)


func (l *TCExecListener) EnterLblock(c *parser.LblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  func_name := uuid.NewString()
  l.TC.UFNStack.PushFront(func_name)
  l.TC.UFStack.PushFront("")

}

func (l *TCExecListener) ExitLblock(c *parser.LblockContext) {
  var func_name string
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.UFStack.Len() > 0 {
    func_name = l.TC.UFNStack.Front().(string)
    l.TC.UserFun.Delete(func_name)
    code := l.TC.UFStack.PopFront()
    l.TC.UserFun.Store(func_name, code)
    l.TC.UFNStack.PopFront()
  }
  r := &TCFunRef{
    Name: func_name,
    Attrs: deque.New(),
  }
  log.Debugf("Creating reference to a lambda function: %v", func_name)
  ReturnFromFunction(l, "lambda", r)
}
