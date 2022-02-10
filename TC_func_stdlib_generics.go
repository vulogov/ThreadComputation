package ThreadComputation

import (
  "fmt"
)

func PrintFunction(l *TCExecListener) error {
  if l.TC.Ready() {
    data_out := l.TC.GetAsString()
    fmt.Println(data_out)
  }
  return nil
}

func initStdlibGenerics() {
  SetFunction("print", PrintFunction)
}
