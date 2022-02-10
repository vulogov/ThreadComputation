package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)

type TCFun func(l *TCExecListener) error

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
  if _, ok := Functions.Load(func_name); ok {
		l.TC.InAttr = true
	} else {
    l.TC.errmsg = fmt.Sprintf("Function: %v not found", func_name)
    l.TC.errors += 1
    log.Errorf(l.TC.errmsg)
  }
}

func (l *TCExecListener) ExitFun(c *parser.FunContext) {
  l.TC.InAttr = false
  if l.TC.Errors() > 0 {
    return
  }
  func_name := c.GetFname().GetText()
  if lfun, ok := Functions.Load(func_name); ok {
    fun := lfun.(TCFun)
    res := fun(l)
    if res != nil {
      l.TC.errmsg = res.Error()
      l.TC.errors += 1
      log.Errorf(l.TC.errmsg)
    }
  }
}

func SetFunction(name string, fun TCFun) {
  Functions.Delete(name)
  Functions.Store(name, fun)
}

func initStdlib() {
  initStdlibGenerics()
}
