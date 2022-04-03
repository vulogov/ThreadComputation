package ThreadComputation

import (
  "fmt"
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
  if l.TC.IsInConditional() {
    log.Debugf("We are in conditional. Skipping: %v", func_name)
    return
  }
  if mod != nil && mod == "@" {
    log.Debugf("User function %v will be created", func_name)
    l.TC.StartUserFun(func_name)
    return
  }
  //
  // See if we have a block callback for this function
  //
  bfun := GetBlockCallback(func_name)
  if bfun != nil {
    //
    // Processing code block start
    //
    log.Debugf("Block callaback detected for %v", func_name)
    l.TC.StartLambdaFun()
    return
  }
  //
  // Enter in conditional if #FALSE is on top of the stack
  // continue otherwise
  //
  if mod != nil && mod == "?" {
    if l.TC.Ready() {
      e := l.TC.Get()
      switch e.(type) {
      case bool:
        if ! e.(bool) {
          log.Debugf("Conditional for %v will be created", func_name)
          l.TC.BeginConditional(func_name)
          return
        } else {
          log.Debugf("Conditional for %v is met", func_name)
        }
      case *TCValue:
        switch e.(*TCValue).Value.(type) {
        case bool:
          if ! e.(*TCValue).Value.(bool) {
            log.Debugf("Conditional for %v will be created", func_name)
            l.TC.BeginConditional(func_name)
          } else {
            log.Debugf("Conditional for %v is met", func_name)
          }
        }
      default:
        l.TC.Set(e)
      }
    }
  }
  log.Debugf("open call: %v\\%v", mod, func_name)
  if mod == nil {
    l.TC.MakeUserFun("", func_name)
  } else {
    l.TC.MakeUserFun(mod.(string), func_name)
  }
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
  //
  // If we are in conditional, Skipping
  // If we are in conditional and name match, remove and skip
  //
  if l.TC.IfConditional(func_name) {
    log.Debugf("Removing conditional for: %v", func_name)
    l.TC.EndConditional(func_name)
    return
  } else {
    if l.TC.IsInConditional() {
      return
    }
  }
  log.Debugf("fname=%v, type=%v", func_name, c.GetFname().GetTokenType())
  //
  // First, let's check if we do have a block function
  //
  bfun := GetBlockCallback(func_name)
  if bfun != nil {
      //
      // Block function detected
      //
      l.TC.FinishUserFun()
      ufname, err := l.TC.EndUserFun()
      if err != nil {
        l.SetError(fmt.Sprintf("Block exit failure: %v", err))
        return
      }
      code, err := l.TC.GetUserFunCode(ufname)
      if err != nil {
        l.SetError(fmt.Sprintf("Block code: %v", err))
        return
      }
      l.TC.EndUserFun()
      bfun(l, ufname, code)
      return
  }
  log.Debugf("fname=%v pushes attributes to evattrs", func_name)
  q = l.Attrs()
  l.TC.EvAttrs.PushFront(q)

  if mod != nil && mod == "@" {
    log.Debugf("Function %v will be created", func_name)
    l.TC.FinishUserFun()
    l.TC.EndUserFun()
    return
  }
  //
  // Second, if we shall make a reference and do no more
  //
  if mod != nil && mod == "`" {
    log.Debugf("Reference to the function %v will be processed", func_name)
    ptr := l.TC.NewRef(func_name, q)
    ReturnFromFunction(l, "#REF", ptr)
    log.Debugf("fname=%v removes attributes from evattrs len=%v", func_name, l.TC.EvAttrs.Len())
    l.TC.EvAttrs.PopFront()
    return
  }
  //
  // If name prepend with '$' it is considered to be variable
  //
  if mod != nil && mod == "$" {
    if l.TC.AddToUserFun(fmt.Sprintf("$%v ", func_name)) {
      return
    }
    FuncSetVariable(l, func_name, q)
  }
  //
  // Shall we just finish user function ?
  //
  if l.TC.FinishUserFun() {
    return
  }
  //
  // First, check if variable exists
  //
  if l.TC.HasVariable(func_name)  {
    if q.Len() > 0 {
      e := q.PopFront()
      l.TC.SetVariable(func_name, e)
    }
    data, _ := l.TC.GetVariable(func_name)
    ReturnFromFunction(l, func_name, data)
  } else if l.TC.IsUserFunction(func_name) {
    //
    // User function exists
    //
    code, err := l.TC.GetUserFunCode(func_name)
    if err != nil {
      l.SetError("Getting user function code %v returned error: %v", func_name, err)
      return
    }
    log.Debugf("Evaluating user function: %v", func_name)
    l.TC.Eval(code)
    log.Debugf("fname=%v removes attributes from evattrs len=%v", func_name, l.TC.EvAttrs.Len())
    l.TC.EvAttrs.PopFront()
    return
  } else if l.TC.HaveCommand(func_name) {
    //
    // Next, if command exists
    //
    if mod == nil {
      mod = ""
    }
    fun := l.TC.GetCommand(func_name)
    if mod == "~" {
      log.Debugf("Command %v%v will be applied to all values from stack and all attributes", mod, func_name)
      q = AllValuesFromStackAndAttr(l, q)
    } else if mod == "∃" {
      log.Debugf("Command %v%v will be applied to one value from stack and all attributes", mod, func_name)
      q = SingleValueFromStackAndAttr(l, q)
    } else {
      log.Debugf("Command %v will be applied only to all attributes", func_name)
    }
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
        } else if mod != nil && mod == "∄" {
          log.Debugf("Function %v will be applied only to attrs", func_name)
        } else if mod != nil && mod == "∘" {
          log.Debugf("Function %v will be applied to ether data in stack or in attrs", func_name)
          q = AllValuesFromStackOrAttr(l, q)
        } else if mod != nil && mod == "∀" {
          log.Debugf("Function %v will be applied to all data in stack and in attrs", func_name)
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
        //
        // Give the function last chance through function router
        //
        res, err := TCFunctionRouter(l, func_name, q)
        if err != nil {
          l.SetError(fmt.Sprintf("Function router returned: %v", err))
          return
        }
        if res != nil {
          ReturnFromFunction(l, func_name, res)
        }
      }
    }
  }
  log.Debugf("fname=%v removes attributes from evattrs len=%v", func_name, l.TC.EvAttrs.Len())
  l.TC.EvAttrs.PopFront()
}
