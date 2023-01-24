package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"
)

type Telnet struct {
	host    string
	timeout time.Duration
	port    string
}

func parse() (f Telnet) {
	flag.DurationVar(&f.timeout, "timout", time.Second*10, "timeout")
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Fprintln(os.Stderr, "expected host and port")
	}
	f.host = flag.Arg(0)
	f.port = flag.Arg(1)
	return
}

func main() {
	f := parse()

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(f.host, f.port), f.timeout)
	if err != nil {
		time.Sleep(f.timeout)
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			// Читаем из сокета
			reader := bufio.NewReader(conn)

			b, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					conn.Close()
					fmt.Println("Conection was closed by host")
				}
				return
			}
			fmt.Println(b)
		}
	}()

	writer := bufio.NewWriter(conn)
	_, err = writer.ReadFrom(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("EOF")
	conn.Close()
	wg.Wait()
	fmt.Println("connection closed")
}
