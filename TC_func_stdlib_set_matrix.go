package ThreadComputation

import (
)

func TCMatrixSet(x interface{}, y interface{}) interface{} {
  var rval float64

  m := x.(*TCMatrix).M
  addr := y.(*TCPair).X
  val := y.(*TCPair).Y
  switch addr.(type) {
  case *TCPair:
    ax := addr.(*TCPair).X
    ay := addr.(*TCPair).Y
    switch ax.(type) {
    case int64:
      switch ay.(type) {
      case int64:
        switch val.(type) {
        case float64:
          rval = val.(float64)
        case int64:
          rval = float64(val.(int64))
        default:
          return nil
        }
        mx, mn := m.Dims()
        if int(ax.(int64)) > mx || int(ay.(int64)) > mn {
          return nil
        }
        m.Set(int(ax.(int64)), int(ay.(int64)), rval)
        return rval
      }
    }
  }
  return nil
}

func init() {
  RegisterOperatorCallback("set", Matrix, Pair, TCMatrixSet)
}
