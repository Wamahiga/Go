package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(netData)

		c.Write([]byte("Hello from server\n"))
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:10008")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}