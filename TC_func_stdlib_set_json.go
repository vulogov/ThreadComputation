package ThreadComputation

import (
)

func TCJsonSet(x interface{}, y interface{}) {

}

func init() {
  RegisterOperatorCmdCallback("set", Json, String, TCJsonSet)
}
