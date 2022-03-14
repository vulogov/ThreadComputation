package ThreadComputation

import (
  "os"
  "fmt"
  "errors"
  "strings"
  "github.com/levigross/grequests"
  "github.com/c2fo/vfs/v6/vfssimple"
  "github.com/gammazero/deque"
)

func readVfsFile(uri string) (string, error) {
  var n int
  max_filesize, _ := GetVariable("tc.Maxfilesize")
  buf := make([]byte, max_filesize.(int))
  if strings.HasPrefix(uri, "@") {
    cwd, _ := os.Getwd()
    uri = strings.TrimPrefix(uri, "@")
    uri = fmt.Sprintf("file://%v/%v", cwd, uri)
  }
  if strings.HasPrefix(uri, "http") || strings.HasPrefix(uri, "https") {
    ro := &grequests.RequestOptions{InsecureSkipVerify: true}
    resp, err := grequests.Get(uri, ro)
    if err != nil {
      return "", err
    }
    if resp.Ok != true {
      return "", errors.New("http request did not return OK")
    }
    data := resp.String()
    resp.Close()
    return data, nil
  } else {
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
    n, err = file.Read(buf)
    if err != nil {
      return "", err
    }
    return string(buf[:n]), nil
  }
  return "", errors.New("read request did not detect a proper context")
}

func TCReadFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    fname := q.PopFront()
    switch fname.(type) {
    case string:
      data, err := readVfsFile(fname.(string))
      if err != nil {
        return nil, err
      }
      ReturnFromFunction(l, "read", data)
    }
  }
  return nil, nil
}

func init() {
  SetFunction("read", TCReadFunction)
}
