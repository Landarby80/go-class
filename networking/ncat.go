package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var port = flag.String("port", "8000", "the port number")
var tz = flag.String("tz", "8000", "time zone")

func main() {
	flag.Parse()
	os.Setenv("TZ", *tz)
	conn, err := net.Dial("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	shouldCopy(os.Stdout, conn)
}

func shouldCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
