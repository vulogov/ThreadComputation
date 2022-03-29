package ThreadComputation

import (
  "gonum.org/v1/gonum/floats"
)

func tcNumbersMathInt(name string, x *TCNumbers, y int64) *TCNumbers {
  for i := 0; i < len(x.N); i++ {
    switch name {
    case "+":
      x.N[i] = x.N[i] + float64(y)
    case "-":
      x.N[i] = x.N[i] - float64(y)
    case "*":
      x.N[i] = x.N[i] * float64(y)
    case "/":
      x.N[i] = x.N[i] / float64(y)
    }
  }
  return x
}

func tcNumbersMathFloat(name string, x *TCNumbers, y float64) *TCNumbers {
  for i := 0; i < len(x.N); i++ {
    switch name {
    case "+":
      x.N[i] = x.N[i] + y
    case "-":
      x.N[i] = x.N[i] - y
    case "*":
      x.N[i] = x.N[i] * y
    case "/":
      x.N[i] = x.N[i] / y
    }
  }
  return x
}

func tcNumbersMathNumbers(name string, x *TCNumbers, y *TCNumbers) *TCNumbers {
  if x.Len() != y.Len() {
    return nil
  }
  switch name {
  case "+":
    floats.Add(x.N, y.N)
  case "-":
    floats.Sub(x.N, y.N)
  case "*":
    floats.Mul(x.N, y.N)
  case "/":
    floats.Div(x.N, y.N)
  }
  return x
}

func TCNumbersMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCNumbers:
    switch y.(type) {
    case int64:
      return tcNumbersMathInt(name, x.(*TCNumbers), y.(int64))
    case float64:
      return tcNumbersMathFloat(name, x.(*TCNumbers), y.(float64))
    case *TCNumbers:
      return tcNumbersMathNumbers(name, x.(*TCNumbers), y.(*TCNumbers))
    case *TCValue:
      switch y.(*TCValue).Value.(type) {
      case int64:
        return MakeValue(tcNumbersMathInt(name, x.(*TCNumbers), y.(*TCValue).Value.(int64)),y.(*TCValue).P, 0)
      case float64:
        return MakeValue(tcNumbersMathFloat(name, x.(*TCNumbers), y.(*TCValue).Value.(float64)),y.(*TCValue).P, 0)
      case *TCNumbers:
        return MakeValue(tcNumbersMathNumbers(name, x.(*TCNumbers), y.(*TCValue).Value.(*TCNumbers)),y.(*TCValue).P, 0)
      }
    }
  }
  return nil
}

func TCValueNumbersMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCValue:
    switch x.(*TCValue).Value.(type) {
    case *TCNumbers:
      switch y.(type) {
      case *TCNumbers:
        return MakeValue(tcNumbersMathNumbers(name, x.(*TCValue).Value.(*TCNumbers), y.(*TCNumbers)), x.(*TCValue).P, 0)
      }
    }
  }
  return nil
}

func init() {
  RegisterMathCallback(Numbers, Int, TCNumbersMath)
  RegisterMathCallback(Numbers, Float, TCNumbersMath)
  RegisterMathCallback(Numbers, Value, TCNumbersMath)
  RegisterMathCallback(Numbers, Numbers, TCNumbersMath)
  RegisterMathCallback(Value, Numbers, TCValueNumbersMath)
}
