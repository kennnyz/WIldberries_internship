package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"time"
)

func HandleDay(w http.ResponseWriter, r *http.Request) {
	for i := range Events {
		if Events[i].Date == time.Now().Format("02.01.2006") {
			resp, err := json.Marshal(Events[i])
			if err != nil {
				log.Fatal(err)
				return
			}
			w.Write(resp)
		}
	}
}

func HandleWeek(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	for i := range Events {
		eventDate, _ := time.Parse("02.01.2006", Events[i].Date)
		if eventDate.After(startOfWeek) && eventDate.Before(endOfWeek) {
			resp, err := json.Marshal(Events[i])
			if err != nil {
				log.Fatal(err)
				return
			}
			w.Write(resp)
		}
	}
}

func HandleMonth(w http.ResponseWriter, r *http.Request) {
	for i := range Events {
		if Events[i].Date[3:] == time.Now().Format("02.01.2006")[3:] {
			resp, err := json.Marshal(Events[i])
			if err != nil {
				log.Fatal(err)
				return
			}
			w.Write(resp)
		}
	}
}
