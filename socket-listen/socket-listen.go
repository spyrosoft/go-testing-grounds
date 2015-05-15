package main

import (
	"net"
	"os"
)

func main() {
	port_handle, error := net.Listen("tcp", ":8080")
	if error != nil {
		os.Exit( 1 )
	}
	for {
		port_connection, error := port_handle.Accept()
		if error != nil {
			os.Exit( 1 )
		}
		go handleConnection(port_connection)
	}
}