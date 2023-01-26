package main

import (
	"log"
	"main/handlers"
	"net/http"
)

//POST /create_event
//POST /update_event
//POST /delete_event
//GET /events_for_day
//GET /events_for_week
//GET /events_for_month

func main() {
	http.HandleFunc("/create_event", loggerMiddleware(handlers.HandleCreate))
	http.HandleFunc("/delete_event", handlers.HandleDelete)
	http.HandleFunc("/update_event", handlers.HandleUpdate)
	http.HandleFunc("/events_for_day", loggerMiddleware(handlers.HandleDay))
	http.HandleFunc("/events_for_week", loggerMiddleware(handlers.HandleWeek))
	http.HandleFunc("/events_for_month", loggerMiddleware(handlers.HandleMonth))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
