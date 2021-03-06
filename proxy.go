package main

import (
  "net"
  "log"
  "io"
)

func handle(src net.Conn) {
  dst, err := net.Dial("tcp", "ironzombi.com:80")
  if err != nil {
    log.Fatalln("Unable to connect to host")
  }
  defer dst.Close()

  go func() {
    if _, err := io.Copy(dst, src); err != nil {
      log.Fatalln(err)
    }
  }()
  if _, err := io.Copy(src, dst); err != nil {
    log.Fatalln(err)
  }
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
    go handle(conn)
  }
}
