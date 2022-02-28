package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	UserId  int    `json:"user_id"`
	EventId int    `json:"event_id"`
	Message string `json:"message"`
	Date    string `json:"date"`
	Err     int    `json:"err"`
}

var err error
var EventMap = make(map[int][]Event)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var Evn Event
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&Evn)
	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
	}
	EventMap[Evn.UserId] = append(EventMap[Evn.UserId], Evn)
	fmt.Println(EventMap)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var Evn Event
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&Evn)
	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
	}
	for _, val := range EventMap[Evn.UserId] {
		if val.EventId == Evn.EventId {
			val.Message = Evn.Message
			val.Date = Evn.Date
		}
	}
	fmt.Println(Evn)
	fmt.Println(EventMap)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var Evn Event
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&Evn)
	if err != nil {
		http.Error(w, "BadRequest", http.StatusBadRequest)
	}
	for i, val := range EventMap[Evn.UserId] {
		if val.EventId == Evn.EventId {
			EventMap[Evn.UserId] = append(EventMap[Evn.UserId][:i], EventMap[Evn.UserId][i+1:]...)
		}
	}
	fmt.Println(EventMap)
}

func EventsForDay(w http.ResponseWriter, r *http.Request) {
	Result := make(map[string][]Event)
	var Evn Event
	w.Header().Set("Content-Type", "application/json")

	Evn.UserId, err = strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		Evn.Err = http.StatusInternalServerError
		Result["error"] = append(Result["error"], Evn)
		http.Error(w, "BadRequest", http.StatusInternalServerError)
	}
	Evn.Date = r.URL.Query().Get("date")

	dt, err := time.Parse("2006-01-02", Evn.Date)
	if err != nil {
		Evn.Err = http.StatusInternalServerError
		Result["error"] = append(Result["error"], Evn)
		http.Error(w, "BadRequest", http.StatusInternalServerError)
	}

	for _, val := range EventMap[Evn.UserId] {
		dtFromMp, err := time.Parse("2006-01-02", val.Date)
		if err != nil {
			Evn.Err = http.StatusInternalServerError
			Result["error"] = append(Result["error"], Evn)
			http.Error(w, "BadRequest", http.StatusInternalServerError)
		}
		if dt.Day() == dtFromMp.Day() {
			Result["result"] = append(Result["result"], val)
		}
	}
	json.NewEncoder(w).Encode(Result)

}

func EventsForMonth(w http.ResponseWriter, r *http.Request) {
	Result := make(map[string][]Event)
	var Evn Event
	w.Header().Set("Content-Type", "application/json")
	Evn.UserId, _ = strconv.Atoi(r.URL.Query().Get("user_id"))
	Evn.Date = r.URL.Query().Get("date")

	dt, err := time.Parse("2006-01-02", Evn.Date)
	if err != nil {
		Evn.Err = http.StatusInternalServerError
		Result["error"] = append(Result["error"], Evn)
		http.Error(w, "BadRequest", http.StatusInternalServerError)
	}
	for _, val := range EventMap[Evn.UserId] {
		dtFromMp, err := time.Parse("2006-01-02", val.Date)
		if err != nil {
			Evn.Err = http.StatusInternalServerError
			Result["error"] = append(Result["error"], Evn)
			http.Error(w, "BadRequest", http.StatusInternalServerError)
		}
		if dt.Month() == dtFromMp.Month() {
			Result["result"] = append(Result["result"], val)

		}
	}
	json.NewEncoder(w).Encode(Result)
}

func EventsForWeek(w http.ResponseWriter, r *http.Request) {
	Result := make(map[string][]Event)
	var Evn Event
	w.Header().Set("Content-Type", "application/json")
	Evn.UserId, _ = strconv.Atoi(r.URL.Query().Get("user_id"))
	Evn.Date = r.URL.Query().Get("date")

	dt, err := time.Parse("2006-01-02", Evn.Date)
	if err != nil {
		Evn.Err = http.StatusInternalServerError
		Result["error"] = append(Result["error"], Evn)
		http.Error(w, "BadRequest", http.StatusInternalServerError)
	}

	_, week := dt.ISOWeek()
	for _, val := range EventMap[Evn.UserId] {
		dtFromMp, err := time.Parse("2006-01-02", val.Date)
		if err != nil {
			Evn.Err = http.StatusInternalServerError
			Result["error"] = append(Result["error"], Evn)
			http.Error(w, "BadRequest", http.StatusInternalServerError)
		}
		_, weekMap := dtFromMp.ISOWeek()
		if week == weekMap {
			Result["result"] = append(Result["result"], val)
		}
	}
	json.NewEncoder(w).Encode(Result)
}

func MiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, req)
		log.Printf("Method: %s  SubCatalog: %s  time: %s", req.Method, req.RequestURI, start)
	})
}
