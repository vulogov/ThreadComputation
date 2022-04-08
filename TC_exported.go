package ThreadComputation
import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/vmihailenco/msgpack/v5"
)

type TCExported struct {
  Data       []byte
}

func (r *TCExported) String() string {
  return fmt.Sprintf("exported[ %v byte(s) ]", r.Len())
}

func (r *TCExported) Len() int {
  return len(r.Data)
}

func (r *TCExported) Export() *TCBinary {
  if r.Len() > 0 {
    res := new(TCBinary)
    err := msgpack.Unmarshal(r.Data, res)
    if err == nil {
      return res
    }
    log.Debugf(fmt.Sprintf("export error: %v", err))
  }
  log.Debugf(fmt.Sprintf("export error: zero buffer"))
  return nil
}

func IsExported(x interface{}) bool {
  switch x.(type) {
  case *TCExported:
    return true
  }
  return false
}

func MakeExported(d []byte) *TCExported {
  res := new(TCExported)
  res.Data = d
  return res
}

func MakeExportedFromBinary(d *TCBinary) *TCExported {
  buf, err := msgpack.Marshal(d)
  if err != nil {
    log.Debugf(fmt.Sprintf("import error: %v", err))
    return nil
  }
  return MakeExported(buf)
}
