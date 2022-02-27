package main

import (
	"GolandProjects/L2_WB/develop/dev10/pkg"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Args struct {
	host    string
	port    string
	TimeOut time.Duration
}

func TelnetServer(A Args) error {
	address := fmt.Sprintf("%v:%v", A.host, A.port)

	fmt.Println("TimeOut", A.TimeOut)
	connect, err := net.DialTimeout("tcp", address, A.TimeOut)
	if err != nil {
		return err
	}
	defer connect.Close()
	fmt.Println("Connection is completed!")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	errChan := make(chan error)

	go pkg.ReadFromSocket(connect, errChan)
	go pkg.WriteToSocket(connect, errChan)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-signals:
				println("program trying to exit....")
				return
			case err = <-errChan:
				if err != nil {
					return
				}
			default:
				continue
			}
		}
	}()

	wg.Wait()

	return nil
}

func main() {
	var A Args

	flag.DurationVar(&A.TimeOut, "timeout", 10, "timeout to connect")
	flag.Parse()

	A.host = os.Args[2]
	A.port = os.Args[3]
	fmt.Println(A)
	err := TelnetServer(A)
	if err != nil {
		fmt.Println(err)
	}

}
