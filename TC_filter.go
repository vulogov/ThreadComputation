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
  var e interface{}
  log.Debugf("Runing filter: %v", code)
  l.TC.Eval(code)
  if l.TC.LastArgs == nil {
    if l.TC.Ready() {
      e = l.TC.Get()
    } else {
      log.Debug("Filter did not leave any value, returning false")
      return false
    }
  } else {
    e = l.TC.LastArgs.PopFront()
  }
  switch e.(type) {
  case bool:
    log.Debugf("Filter returned: %v", e.(bool))
    return e.(bool)
  default:
    l.TC.Res.Set(e)
    return true
  }
  return false
}

func filterTCQueue(l *TCExecListener, code string, rq *deque.Deque, wq *deque.Deque) {
  for wq.Len() > 0 {
    e := wq.PopFront()
    log.Debugf("Pushing %v to stack for filtering", e)
    l.TC.Res.Set(e)
    res := TCRunFilterOnQ(l, code)
    if res {
      log.Debugf("%T passed through filter", e)
      rq.PushBack(e)
    } else {
      log.Debugf("%T not passed through filter", e)
    }
  }
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
      filterTCQueue(l, code, &rq, &wq)
      for rq.Len() > 0 {
        e := rq.PopBack()
        l.TC.Res.Set(e)
      }
    }
  } 
  log.Debug("Exiting FILTER block")
}
