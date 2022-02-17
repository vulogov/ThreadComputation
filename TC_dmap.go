package ThreadComputation

import (
  "fmt"
  "strings"
  log "github.com/sirupsen/logrus"
  "github.com/lrita/cmap"
  "github.com/vulogov/ThreadComputation/parser"
  conv "github.com/cstockton/go-conv"
)


func (l *TCExecListener) EnterDmap(c *parser.DmapContext) {
  var data cmap.Cmap
  if l.TC.Errors() > 0 {
    return
  }
  if l.TC.AddToUserFun("{") == true {
    return
  }
  l.TC.Res.Set(&data)
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
    v = l.TC.Res.Q().Front().(*cmap.Cmap)
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
    v.(*cmap.Cmap).Store(key, bool_val)
    return
  }
  float_val, err := conv.Float64(val)
  if err == nil {
    v.(*cmap.Cmap).Store(key, float_val)
    return
  }
  int_val, err := conv.Int64(val)
  if err == nil {
    v.(*cmap.Cmap).Store(key, int_val)
    return
  }
  if strings.HasPrefix(val, "\"") {
    val = strings.TrimSuffix(val, "\"")
    val = strings.TrimPrefix(val, "\"")
  }
  v.(*cmap.Cmap).Store(key, val)
}
