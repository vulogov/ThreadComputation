package ThreadComputation

import (
  "os"
  "fmt"
  "errors"
  "strings"
  "github.com/c2fo/vfs/v6/vfssimple"
  "github.com/gammazero/deque"
)

func readVfsFile(uri string) (string, error) {
  max_filesize, _ := GetVariable("tc.Maxfilesize")
  buf := make([]byte, max_filesize.(int))
  if strings.HasPrefix(uri, "@") {
    cwd, _ := os.Getwd()
    uri = strings.TrimPrefix(uri, "@")
    uri = fmt.Sprintf("file://%v/%v", cwd, uri)
  }
  file, err := vfssimple.NewFile(uri)
  if err != nil {
    return "", err
  }
  exists, err := file.Exists()
  if err != nil {
    return "", err
  }
  if ! exists {
    return "", errors.New(fmt.Sprintf("File %v not exists", uri))
  }
  n, err := file.Read(buf)
  if err != nil {
    return "", err
  }
  return string(buf[:n]), nil
}

func ReadFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      fn := q.PopFront()
      switch fn.(type) {
      case string:
        data, err := readVfsFile(fn.(string))
        if err != nil {
          return nil, err
        }
        l.TC.Res.Set(data)
      default:
        break
      }
    }
    return nil, nil
  } else {
    if l.TC.Ready() {
      fn := l.TC.Get()
      switch fn.(type) {
      case string:
        return readVfsFile(fn.(string))
      }
    }
  }
  return nil, errors.New("read function did not discover proper context")
}



func initStdlibVfs() {
  SetFunction("read", ReadFunction)
}
