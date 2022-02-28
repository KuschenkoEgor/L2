package main

import "C"
import (
	"GolandProjects/L2_WB/develop/dev11/pkg"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
)

type Conf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (c *Conf) ConfigData() *Conf {

	yamlFile, err := ioutil.ReadFile("/home/zhora/GolandProjects/L2_WB/develop/dev11/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c Conf

	r := mux.NewRouter()
	rLog := pkg.MiddleWare(r)

	r.HandleFunc("/create_event", pkg.CreateEvent).Methods("POST")
	r.HandleFunc("/update_event", pkg.UpdateEvent).Methods("POST")
	r.HandleFunc("/delete_event", pkg.DeleteEvent).Methods("POST")
	r.HandleFunc("/events_for_day", pkg.EventsForDay).Methods("GET")
	r.HandleFunc("/events_for_month", pkg.EventsForMonth).Methods("GET")
	r.HandleFunc("/events_for_week", pkg.EventsForWeek).Methods("GET")

	c.ConfigData()

	Address := fmt.Sprintf("%v:%v", c.Host, c.Port)

	log.Fatal(http.ListenAndServe(Address, rLog))
}
