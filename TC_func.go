package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
  "github.com/vulogov/ThreadComputation/parser"
)

type TCFun func(l *TCExecListener, q *deque.Deque) (interface{}, error)

func (l *TCExecListener) EnterFun(c *parser.FunContext) {
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
  if _, ok := l.TC.Vars.Load(func_name); ok {
    return
  }
  if _, ok := Functions.Load(func_name); ok {
		l.TC.InAttr += 1
    l.TC.Attrs.Add()
    l.TC.FNStack.PushFront(func_name)
	} else {
    l.TC.errmsg = fmt.Sprintf("Function: %v not found", func_name)
    l.TC.errors += 1
    log.Errorf(l.TC.errmsg)
  }
}

func (l *TCExecListener) ExitFun(c *parser.FunContext) {
  if l.TC.InAttr > 0 {
    l.TC.InAttr -= 1
  }
  if l.TC.Errors() > 0 {
    return
  }
  func_name := c.GetFname().GetText()
  if vdata, ok := l.TC.Vars.Load(func_name); ok {
    l.TC.Res.Set(vdata)
    return
  }
  if lfun, ok := Functions.Load(func_name); ok {
    fun := lfun.(TCFun)
    q   := l.TC.Attrs.Q()
    l.TC.Attrs.Del()
    res, err := fun(l, q)
    l.TC.FNStack.PopFront()
    if err != nil {
      l.TC.errmsg = err.Error()
      l.TC.errors += 1
      log.Errorf(l.TC.errmsg)
    } else {
      if res != nil {
        if l.TC.Attrs.GLen() > 1 {
          l.TC.Attrs.Set(res)
        } else {
          l.TC.Res.Set(res)
        }
      }
    }
  }
}

func SetFunction(name string, fun TCFun) {
  Functions.Delete(name)
  Functions.Store(name, fun)
}

func initStdlib() {
  initStdlibGenerics()
  initStdlibMath()
  initStdlibStack()
}
