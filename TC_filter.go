package ThreadComputation

import (
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)


func (l *TCExecListener) EnterFilterblock(c *parser.FilterblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  log.Debug("Entering FILTER block")
  MakeTempFun(l)
}

func TCRunFilterOnQ(l *TCExecListener, code string) bool {
  l.TC.Eval(code)
  e :=  l.TC.Res.Q().Front()
  switch e.(type) {
  case bool:
    e = l.TC.Get()
    return e.(bool)
  default:
    return true
  }
  return false
}

func (l *TCExecListener) ExitFilterblock(c *parser.FilterblockContext) {
  var func_name string
  var wq deque.Deque
  var rq deque.Deque

  if l.TC.Errors() > 0 {
    return
  }
  func_name = l.TC.UFNStack.Front().(string)
  code := l.TC.UFStack.PopFront().(string)
  l.TC.UFNStack.PopFront()
  log.Debugf("Got filter: %v", func_name)
  if l.TC.LastArgs == nil {
    if l.TC.Ready() {
      log.Debug("Applying filter to the stack")
      for l.TC.Ready() {
        e := l.TC.Get()
        wq.PushBack(e)
      }
      for wq.Len() > 0 {
        e := wq.PopFront()
        l.TC.Res.Set(e)
        res := TCRunFilterOnQ(l, code)
        if res {
          log.Debugf("%T passed through filter", e)
          rq.PushBack(e)
        } else {
          log.Debugf("%T not passed through filter", e)
        }
      }
      for rq.Len() > 0 {
        e := rq.PopBack()
        l.TC.Res.Set(e)
      }
    }
  } else {
    if l.TC.FNStack.Len() > 0 {
      log.Debugf("Applying filter to the list of function arguments: %v", l.TC.FNStack.Front().(string))
    }
  }
  log.Debug("Exiting FILTER block")
}
