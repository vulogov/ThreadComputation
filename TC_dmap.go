package ThreadComputation

import (
  "fmt"
  "strings"
  "errors"
  log "github.com/sirupsen/logrus"
  "github.com/srfrog/dict"
  "github.com/vulogov/ThreadComputation/parser"
  conv "github.com/cstockton/go-conv"
  "github.com/gammazero/deque"
)


func (l *TCExecListener) EnterDmap(c *parser.DmapContext) {
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("{") == true {
    return
  }
  l.TC.Res.Set(dict.New())
}

func (l *TCExecListener) ExitDmap(c *parser.DmapContext) {
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("}") == true {
    return
  }
}

func (l *TCExecListener) EnterKey_term(c *parser.Key_termContext) {
  var v interface{}
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.Ready() {
    v = l.TC.Res.Q().Front().(*dict.Dict)
  } else {
    l.TC.errmsg = "key declaration expects map on top of the stack"
    l.TC.errors += 1
    log.Errorf(l.TC.errmsg)
    return
  }
  key := c.GetKEY().GetText()
  val := c.GetVALUE().GetText()
  if l.TC.AddToUserFun(fmt.Sprintf("%v : %v", key, val)) == true {
    return
  }
  bool_val, err := conv.Bool(val[1:])
  if err == nil {
    v.(*dict.Dict).Set(key, bool_val)
    return
  }
  float_val, err := conv.Float64(val)
  if err == nil {
    v.(*dict.Dict).Set(key, float_val)
    return
  }
  int_val, err := conv.Int64(val)
  if err == nil {
    v.(*dict.Dict).Set(key, int_val)
    return
  }
  if strings.HasPrefix(val, "\"") {
    val = strings.TrimSuffix(val, "\"")
    val = strings.TrimPrefix(val, "\"")
  }
  v.(*dict.Dict).Set(key, val)
}

func TCSetVarsFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    e, err := l.TC.Res.Get()
    if err != nil {
      return nil, err
    }
    switch e.(type) {
    case *dict.Dict:
      for item := range e.(*dict.Dict).Items() {
        l.TC.SetVariable(item.Key.(string), item.Value)
      }
      return nil, nil
    default:
      return nil, errors.New(fmt.Sprintf("invalid datatype in stack, for which we do not support conversion to variables: %T", e))
    }
  }
  return nil, errors.New("invalid context for vars operator")
}

func initStdlibDmap() {
  SetFunction("vars", TCSetVarsFunction)
}
