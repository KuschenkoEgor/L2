package main
import(
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)
//Создать программу печатающую точное время с использованием NTP -библиотеки.
//Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp.
//Написать программу печатающую текущее время / точное время с использованием этой библиотеки.


func GetTime() (time.Time,error){
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Printf(err.Error())
	}
	return time,nil
}

func main() {

	t,err := GetTime()
	if err != nil {
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}
	fmt.Printf("Текущее время: %v ч. %v мин.\n",t.Hour(),t.Minute())
	fmt.Printf("Точное время: %v ч. %v мин. %v сек.",t.Hour(),t.Minute(),t.Second())

}