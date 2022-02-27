package main

import (
	"example.com/m/v2/L2_WB/develop/dev10/pkg"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Args struct {
	host string
	port string
	TimeOut string
}

func TelnetServer(A Args)  error {
	address := fmt.Sprintf("%v:%v",A.host,A.port)

	rn := []rune(A.TimeOut)
	T, _ := strconv.Atoi(string(rn[0]))

	TimeOut:= time.Duration(T)*time.Second

	connect,err := net.DialTimeout("tcp",address,TimeOut)
	if err != nil {
		return err
	}
	defer connect.Close()
	fmt.Println("Connection is completed!")

	signals := make(chan os.Signal,1)
	signal.Notify(signals)

	errChan := make(chan error)

	go pkg.ReadFromSocket(connect,errChan)
	go pkg.WriteToSocket(connect,errChan)

	return nil
}


func main() {
	var A Args

	flag.StringVar(&A.TimeOut,"timeout","10s","timeout to connect")
	flag.Parse()

	A.host = os.Args[2]
	A.port = os.Args[3]
	fmt.Println(A)
	err := TelnetServer(A)
	if err != nil {
		fmt.Println(err)
	}

}
