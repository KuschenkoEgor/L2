package pkg

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func WriteToSocket(connect net.Conn, errChan chan error) {
	read := bufio.NewReader(os.Stdin)

	for {
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

	input := make([]byte, 1024)
	for {
		_, err := connect.Read(input)
		if err != nil {
			errChan <- err
			return
		}
		fmt.Println("Данные из сокета:", string(input))
	}
}
