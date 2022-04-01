package ThreadComputation

import (
  "fmt"
  "math"
  "github.com/gammazero/deque"
)

type TCNumbers struct {
  P            float64
  N            []float64
}

func tcRotateFloat64(slice []float64, n int) (rotateSlice []float64) {
    var begin []float64
    var end []float64
    size := len(slice)
    rotateSlice = make([]float64, size)

    nAbs := math.Abs(float64(n))

    if int(nAbs) > size {
        remainder, _ := tcQuotientAndRemainderF64(float64(n), float64(size))
        n = int(remainder)
    }

    if n != 0 {
        if n > 0 {
            index := size - n
            begin = slice[index:]
            end = slice[0:index]
            copy(rotateSlice, begin)
            copy(rotateSlice[n:], end)
        } else {
            n = int(nAbs)
            index := size - n
            begin = slice[n:]
            end = slice[0:n]
            copy(rotateSlice, begin)
            copy(rotateSlice[index:], end)
        }
    } else {
        copy(rotateSlice, slice)
    }

    return rotateSlice
}

func tcQuotientAndRemainderF64(x, y float64) (Remainder, Quotient float64) {
    Quotient = float64(math.Floor(float64(x / y)))
    Remainder = x - y*Quotient
    return Remainder, Quotient
}

func MakeNumbers() *TCNumbers {
  res := new(TCNumbers)
  res.P = 0.0
  return res
}

func (tc *TCstate) Numbers() *TCNumbers {
  return MakeNumbers()
}

func (n *TCNumbers) Add(v interface{}) bool {
  switch v.(type) {
  case float64:
    n.P += 100.0
    n.N = append(n.N, v.(float64))
  case int64:
    n.P += 100.0
    n.N = append(n.N, float64(v.(int64)))
  case *TCList:
    n.P += 100.0
    for i := 0; i < v.(*TCList).Len(); i++ {
      n.Add(v.(*TCList).Q.At(i))
    }
  case *TCValue:
    switch v.(*TCValue).Value.(type) {
    case int64:
      n.P += float64(v.(*TCValue).P)
      n.N = append(n.N, float64(v.(*TCValue).Value.(int64)))
    case float64:
      n.P += float64(v.(*TCValue).P)
      n.N = append(n.N, v.(*TCValue).Value.(float64))
    }
  default:
    return false
  }
  return true
}

func (n *TCNumbers) Set(v []float64) {
  n.N = v
}

func (n *TCNumbers) Len() int {
  return len(n.N)
}

func (n *TCNumbers) RotateLeft() {
  n.N = tcRotateFloat64(n.N, -1)
}

func (n *TCNumbers) RotateRight() {
  n.N = tcRotateFloat64(n.N, 1)
}

func (n *TCNumbers) String() string {
  out := "numbers[ "
  for x := 0; x < len(n.N); x++ {
    out += fmt.Sprintf("%v ", n.N[x])
  }
  out += " ]"
  return out
}

func TCNumbersFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  n := l.TC.Numbers()
  for q.Len() > 0 {
    e := q.PopFront()
    if ! n.Add(e) {
      return nil, l.TC.MakeError("Error creating numbers[]")
    }
  }
  return n, nil
}

func TCNumbersValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  n := l.TC.Numbers()
  for q.Len() > 0 {
    e := q.PopFront()
    if ! n.Add(e) {
      return nil, l.TC.MakeError("Error creating Numbers[]")
    }
  }
  return MakeValue(n, float32(n.P/float64(n.Len())), 0), nil
}

func init() {
  SetCommand("numbers", TCNumbersFunction)
  SetCommand("JustNumbers", TCNumbersValueFunction)
}
