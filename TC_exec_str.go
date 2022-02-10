package ThreadComputation

import (
  "github.com/vulogov/ThreadComputation/parser"
  conv "github.com/cstockton/go-conv"
)

func (l *TCExecListener) EnterString_term(c *parser.String_termContext) {
  str, err := conv.String(c.GetVALUE().GetText())
  if err == nil {
    if l.TC.InAttr == false {
      l.TC.Res.PushFront(str)
    } else {
      l.TC.Attrs.PushBack(str)
    }
  }
}
