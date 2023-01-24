package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		fmt.Println("connected")

		//Reader для чтения из буффера

		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

		go func(conn net.Conn) {
			defer conn.Close()
			for {
				//читаем
				b, err := bufReader.ReadString('\n')
				if err != nil {
					fmt.Fprintln(os.Stderr, "can`t read", err)
					break
				}
				fmt.Fprintln(conn, "message:", string(b[:len(b)-1]), "received")
				fmt.Print(string(b))
			}
		}(conn)
	}
}
