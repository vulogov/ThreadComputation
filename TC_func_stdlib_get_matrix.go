package ThreadComputation

import (
)

func TCMatrixGet(x interface{}, y interface{}) interface{} {
  rx := y.(*TCPair).X
  ry := y.(*TCPair).Y
  m, n := x.(*TCMatrix).M.Dims()
  if int(rx.(int64)) > m || int(ry.(int64)) > n {
    return nil
  }
  switch rx.(type) {
  case int64:
    switch ry.(type) {
    case int64:
      return x.(*TCMatrix).M.At(int(rx.(int64)), int(ry.(int64)))
    }
  }
  return nil
}

func init() {
  RegisterOperatorCallback("get", Matrix, Pair, TCMatrixGet)
}
