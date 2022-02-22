package ThreadComputation

import (
  "fmt"
  "github.com/gammazero/deque"
  "errors"
  "math"
  "gonum.org/v1/gonum/floats"
  "gonum.org/v1/gonum/stat"
)


func ArithmeticAggFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  data := collectAllData(l, q).([]float64)
  if l.TC.Errors() > 0 {
    return nil, errors.New(l.TC.Error())
  }
  if len(data) == 0 {
    return nil, errors.New("attempt of arithmetic aggregation over empty dataset")
  }
  acc := data[0]
  for x := 1; x < len(data); x++ {
    switch l.TC.CurrentFunctionName() {
    case "++":
      acc += data[x]
    case "--":
      acc -= data[x]
    case "**":
      acc *= data[x]
    case "//":
      acc = acc / data[x]
    default:
      return nil, errors.New(fmt.Sprintf("Unknown aggregator arthmetic operator: %v", l.TC.CurrentFunctionName()))
    }
  }
  return acc, nil
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

func MathFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  switch l.TC.CurrentFunctionName() {
  case "abs":
    return applyFunToFloats(l, q, math.Abs)
  case "exp":
    return applyFunToFloats(l, q, math.Exp)
  case "exp2":
    return applyFunToFloats(l, q, math.Exp2)
  case "log":
    return applyFunToFloats(l, q, math.Log)
  case "log10":
    return applyFunToFloats(l, q, math.Log10)
  case "sqrt":
    return applyFunToFloats(l, q, math.Sqrt)
  default:
    return nil, errors.New(fmt.Sprintf("Unknown math function: %v", l.TC.CurrentFunctionName()))
  }
  return nil, nil
}

func MeanFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArrayWithWeight(l, q, stat.Mean)
}

func VarianceFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArrayWithWeight(l, q, stat.Variance)
}

func SkewFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArrayWithWeight(l, q, stat.Skew)
}

func DeviationFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArrayWithWeight(l, q, stat.StdDev)
}

func MinFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArray(l, q, floats.Min)
}

func MaxFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return applyAggFunToFloatArray(l, q, floats.Max)
}




func initStdlibMath() {
  SetFunction("+", ArithmeticFunction)
  SetFunction("-", ArithmeticFunction)
  SetFunction("*", ArithmeticFunction)
  SetFunction("/", ArithmeticFunction)
  SetFunction("++", ArithmeticAggFunction)
  SetFunction("--", ArithmeticAggFunction)
  SetFunction("**", ArithmeticAggFunction)
  SetFunction("//", ArithmeticAggFunction)
  SetFunction("min", MinFunction)
  SetFunction("max", MaxFunction)
  SetFunction("mean", MeanFunction)
  SetFunction("variance", VarianceFunction)
  SetFunction("skew", SkewFunction)
  SetFunction("deviation", DeviationFunction)
  SetFunction("abs", MathFunction)
  SetFunction("exp", MathFunction)
  SetFunction("exp2", MathFunction)
  SetFunction("log", MathFunction)
  SetFunction("log10", MathFunction)
  SetFunction("sqrt", MathFunction)
}
