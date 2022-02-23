package ThreadComputation

import (
  "fmt"
  "github.com/gammazero/deque"
  "errors"
  "math"
  "gonum.org/v1/gonum/floats"
  "gonum.org/v1/gonum/stat"
  "gonum.org/v1/gonum/mat"
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
  var acc interface{}
  var size int

  data := collectData(l, q)
  if l.TC.Errors() > 0 {
    return nil, errors.New(l.TC.Error())
  }
  switch data.(type) {
  case []int64:
    acc = data.([]int64)[0]
    size = len(data.([]int64))
  case []float64:
    acc = data.([]float64)[0]
    size = len(data.([]float64))
  case []bool:
    acc = data.([]bool)[0]
    size = len(data.([]bool))
  case []string:
    acc = data.([]string)[0]
    size = len(data.([]string))
  case []*mat.Dense:
    acc = data.([]*mat.Dense)[0]
    size = len(data.([]*mat.Dense))
  default:
    return nil, errors.New("unsupported data type for arithmetics")
  }
  for x := 1; x < size; x++ {
    switch l.TC.CurrentFunctionName() {
    case "+":
      switch acc.(type) {
      case int64:
        acc = acc.(int64) + data.([]int64)[x]
      case float64:
        acc = acc.(float64) + data.([]float64)[x]
      case bool:
        acc = acc.(bool) && data.([]bool)[x]
      case string:
        acc = acc.(string) + data.([]string)[x]
      case *mat.Dense:
        acc.(*mat.Dense).Add(acc.(*mat.Dense), data.([]*mat.Dense)[x])
      default:
        return nil, errors.New("unsupported data type for arithmetics")
      }
    case "-":
      switch acc.(type) {
      case int64:
        acc = acc.(int64) - data.([]int64)[x]
      case float64:
        acc = acc.(float64) - data.([]float64)[x]
      case bool:
        acc = acc.(bool) || data.([]bool)[x]
      case *mat.Dense:
        acc.(*mat.Dense).Sub(acc.(*mat.Dense), data.([]*mat.Dense)[x])
      default:
        return nil, errors.New("unsupported data type for arithmetics")
      }
    case "*":
      switch acc.(type) {
      case int64:
        acc = acc.(int64) * data.([]int64)[x]
      case float64:
        acc = acc.(float64) * data.([]float64)[x]
      case *mat.Dense:
        acc.(*mat.Dense).MulElem(acc.(*mat.Dense), data.([]*mat.Dense)[x])
      default:
        return nil, errors.New("unsupported data type for arithmetics")
      }
    case "/":
      switch acc.(type) {
      case int64:
        acc = acc.(int64) / data.([]int64)[x]
      case float64:
        acc = acc.(float64) / data.([]float64)[x]
      case *mat.Dense:
        acc.(*mat.Dense).DivElem(acc.(*mat.Dense), data.([]*mat.Dense)[x])
      default:
        return nil, errors.New("unsupported data type for arithmetics")
      }
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
