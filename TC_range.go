package ThreadComputation
import (
  "fmt"
)

type TCRange struct {
  X          float64
  Y          float64
  Step       float64
}

func (r *TCRange) String() string {
  return fmt.Sprintf("range[%v...%v (%v)]", r.X, r.Y, r.Step)
}
