package main

import (
	"fmt"
	"net"

	"github.com/jeffcail/tcp-unpack/unpack"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Connect tcp err: %v", err)
		return
	}

	err = unpack.Encode(conn, "Hi mary!!!")
	if err != nil {
		fmt.Printf("Unpack err: %v", err)
		return
	}
}
