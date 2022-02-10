package ThreadComputation

import (
  "github.com/vulogov/ThreadComputation/parser"
  conv "github.com/cstockton/go-conv"
)

func (l *TCExecListener) EnterBoolean_term(c *parser.Boolean_termContext) {
  if l.TC.Errors() > 0 {
    return
  }
  str, err := conv.Bool(c.GetVALUE().GetText()[1:])
  if err == nil {
    if l.TC.InAttr == false {
      l.TC.Res.PushFront(str)
    } else {
      l.TC.Attrs.PushBack(str)
    }
  }
}
