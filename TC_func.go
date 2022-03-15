package ThreadComputation

import (
  "strings"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
  "github.com/vulogov/ThreadComputation/parser"
)

type TCFun func(l *TCExecListener, name string, q *deque.Deque) (interface{}, error)


func (l *TCExecListener) EnterFun(c *parser.FunContext) {
  var mod interface{}
  if l.TC.Errors() > 0 {
    return
  }
  if c.GetMod() != nil {
    mod = c.GetMod().GetText()
    mod = strings.ToLower(mod.(string))
  } else {
    mod = nil
  }
  if c.GetFname() == nil {
    l.SetError("Can not get function name out of context")
    return
  }
  func_name := c.GetFname().GetText()
  if mod != nil && mod == "?" {
    log.Debugf("Conditional for %v will be created", func_name)
    return
  }
  if mod != nil && mod == "`" {
    log.Debugf("Reference for %v will be created", func_name)
    return
  }
  log.Debugf("open call: %v\\%v", mod, func_name)
  l.TC.InAttr += 1
  l.TC.Attrs.Add()
}

func (l *TCExecListener) ExitFun(c *parser.FunContext) {
  var mod interface{}
  var q  *deque.Deque

  if l.TC.InAttr > 0 {
    l.TC.InAttr -= 1
  }
  if l.TC.Errors() > 0 {
    return
  }
  if c.GetMod() != nil {
    mod = c.GetMod().GetText()
    mod = strings.ToLower(mod.(string))
  } else {
    mod = nil
  }
  func_name := c.GetFname().GetText()

  if mod != nil && mod == "`" {
    log.Debugf("Reference to the function %v will be processed", func_name)
    return
  }

  q = l.Attrs()
  //
  // First, check if variable exists
  //
  if l.TC.HasVariable(func_name) || q.Len() > 0 {
    if q.Len() > 0 {
      e := q.PopFront()
      l.TC.SetVariable(func_name, e)
    }
    data, _ := l.TC.GetVariable(func_name)
    ReturnFromFunction(l, func_name, data)
  } else if l.TC.HaveCommand(func_name) {
    //
    // Next, if command exists
    //
    log.Debugf("Command %v will be applied only to attributes", func_name)
    fun := l.TC.GetCommand(func_name)
    res, err := fun(l, func_name, q)
    if err != nil {
      l.SetError("Command %v returned error: %v", func_name, err)
    } else {
      if res != nil {
        ReturnFromFunction(l, func_name, res)
      }
    }
  } else {
    //
    // Otherwise assiming it is a function or values
    //
    if l.TC.HaveFunction(func_name) {
      if l.TC.Ready() || l.TC.HaveAttrs() {
        if mod != nil && mod == "~" {
          log.Debugf("Function %v will be applied to all data in stack", func_name)
          q = AllValuesFromStackAndAttr(l, q)
        } else {
          log.Debugf("Function %v will be applied to a single element in stack", func_name)
          q = SingleValueFromStackAndAttr(l, q)
        }
        fun := l.TC.GetFunction(func_name)
        res, err := fun(l, func_name, q)
        if err != nil {
          l.SetError("Function %v returned error: %v", func_name, err)
        } else {
          if res != nil {
            ReturnFromFunction(l, func_name, res)
          }
        }
      } else {
        l.SetError("Stack is too shallow to execute function: %v", func_name)
        return
      }
    } else {
      if IsSimpleData(func_name) {
        e := GetSimpleData(func_name)
        e = ApplyToFunction(l, e, q)
        ReturnFromFunction(l, func_name, e)
      } else {
        l.SetError("%v is not recognozed as function", func_name)
      }
    }
  }
  if mod != nil && mod == "?" {
    log.Debugf("Conditional for %v will be processed", func_name)
    return
  }
}
