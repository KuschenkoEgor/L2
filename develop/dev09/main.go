package main

import (
	"example.com/m/v2/L2_WB/develop/dev09/pkg"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Args struct {
	NameDir string
	Url     string
}

func main() {
	var A Args
	var err error
	MapLinks := make(map[string]bool)

	flag.StringVar(&A.NameDir, "O", "site", "Name dir for site")

	flag.Parse()

	A.Url = os.Args[3]

	err = os.Mkdir(A.NameDir, 0777)
	if err != nil {
		fmt.Println(err)
	}

	MapLinks = pkg.GetAllLinks(A.Url)
	for key, _ := range MapLinks {

		FName := strings.Replace(key, "/", "#", -1)
		FN := strings.Replace(FName, "https:", "", -1)
		FileName := strings.Replace(FN, ".", "_", -1)

		err = pkg.FileForPage(A.NameDir, FileName)

		if err != nil {
			fmt.Println(err)
		}

		resp,err := http.Get(key)
		if err != nil {
			fmt.Println(err)
		}

		Way := fmt.Sprintf("%v/%v",A.NameDir,FileName)
		file,err :=  os.OpenFile(Way, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()

		_, err = io.Copy(file,resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s - '%s'saved\n\n", time.Now().Format("01/02/06 15:04:05"), FileName)
	}



}
