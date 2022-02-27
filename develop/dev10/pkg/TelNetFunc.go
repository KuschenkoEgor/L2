package pkg

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func WriteToSocket(connect net.Conn, errChan chan error) {
	for {
		read := bufio.NewReader(os.Stdin)

		txt, err := read.ReadBytes('\n')
		if err != nil {
			errChan <- err
			return
		}

		_, err = connect.Write(txt)
		if err != nil {
			errChan <- err
			return
		}
	}
}

func ReadFromSocket(connect net.Conn, errChan chan error) {
	for {
		data := make([]byte, 0)

		_, err := connect.Read(data)
		if err != nil {
			errChan <- err
			return
		}

		fmt.Println(data)
	}
}
