package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, error := net.Dial("tcp", "127.0.0.1:8080")
	exitOnError(error)
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, error := bufio.NewReader(conn).ReadString('\n')
	exitOnError(error)
	fmt.Println("-----")
	fmt.Println(status)
}

func exitOnError(error error) {
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}
}