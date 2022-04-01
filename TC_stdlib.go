package ThreadComputation

func IsStringElementExist(s []string, str string) bool {
  for _, v := range s {
    if v == str {
      return true
    }
  }
  return false
}

func initStdlib() {

}
