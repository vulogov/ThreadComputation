package ThreadComputation

import (
  "fmt"
  "github.com/gammazero/deque"
  "errors"
  "gonum.org/v1/gonum/floats"
  // "gonum.org/v1/gonum/stat"
)


//

//
// // func applyAggWeightFun(l *TCExecListener, fun func([]float64, w []float64) float64) error {
// //   var acc float64
// //   data := collectFloatData(l)
// //   if l.TC.Errors() > 0 {
// //     return errors.New(l.TC.Error())
// //   }
// //   acc = fun(data, w)
// //   l.TC.Res.PushFront(acc)
// //   return nil
// // }
//
func AddFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArray(l, q, floats.Sum)
}

func ArithmeticFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  data := collectData(l, q).([]float64)
  if l.TC.Errors() > 0 {
    return nil, errors.New(l.TC.Error())
  }
  acc := data[0]
  for x := 1; x < len(data); x++ {
    switch l.TC.CurrentFunctionName() {
    case "+":
      acc += data[x]
    case "-":
      acc -= data[x]
    case "*":
      acc *= data[x]
    case "/":
      acc = acc / data[x]
    default:
      return nil, errors.New(fmt.Sprintf("Unknown arthmetic operator: %v", l.TC.CurrentFunctionName()))
    }
  }
  return acc, nil
}



func MinFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArray(l, q, floats.Min)
}

func MaxFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArray(l, q, floats.Max)
}




func initStdlibMath() {
  SetFunction("+", AddFunction)
  SetFunction("-", ArithmeticFunction)
  SetFunction("*", ArithmeticFunction)
  SetFunction("/", ArithmeticFunction)
  SetFunction("min", MinFunction)
  SetFunction("max", MaxFunction)
}
