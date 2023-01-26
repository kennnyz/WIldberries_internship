package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	createData(w, r)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var event Event
	err = json.Unmarshal(reqBytes, &event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := event.UserID
	fmt.Println("Id ---> ", id)
	deleteID := -1
	for p := range Events {
		if Events[p].UserID == id {
			deleteID = p
		}
	}
	if deleteID == -1 {
		w.Write([]byte("Data no exist"))
		return
	}

	Events = Events[:deleteID]
	fmt.Println(Events)
}
func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var event Event
	err = json.Unmarshal(reqBytes, &event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !Check("02.01.2006", event.Date) {
		w.Write([]byte("Incorrect date format. Please use dd-mm-yyyy"))
		return
	}
	id := event.UserID
	fmt.Println("Id ---> ", id)
	updateID := -1
	for p := range Events {
		if Events[p].UserID == id {
			updateID = p
		}
	}
	if updateID == -1 {
		w.Write([]byte("Data no exist"))
		return
	}

	Events[updateID].Date = event.Date
	fmt.Println(Events)
}

func createData(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var event Event
	err = json.Unmarshal(reqBytes, &event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(event.Date)
	if Check("02.01.2006", event.Date) == false {
		w.Write([]byte("Incorrect date format. Please use dd-mm-yyyy"))
		return
	}
	Events = append(Events, event)
	fmt.Println(Events)
}

func Check(format, date string) bool {
	_, err := time.Parse(format, date)
	if err != nil {
		return false
	}
	return true
}
