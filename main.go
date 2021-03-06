package main

import (
	//"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sort"

	"github.com/GeertJohan/go.rice"
	//"github.com/gorilla/handlers"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"

	"./scheduling"
)

var (
	listen = flag.String("listen", ":8080", "The adress this service will be available on.")
	//dbPath = flag.String("path", "", "The path at which the database containing class scheduling information")
)

func main() {
	//running the below function runs a benchmark.
	//scheduling.FindBench()
	
	flag.Parse()

	var DevelopmentMode = true

	mux := http.NewServeMux()

	if !DevelopmentMode {
		staticFiles := rice.MustFindBox("frontend").HTTPBox()
		mux.Handle("/", http.FileServer(staticFiles))
	}

	mux.HandleFunc("/api/v0/schedules", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		if r.Method == "OPTIONS" {
			return
		}

		if r.Method != http.MethodPost {

			log.Println("Invalid method to post endpoint:", r.Method)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var scheduleRequest scheduling.ScheduleRequest
		err := decoder.Decode(&scheduleRequest)
		if err != nil {
			log.Println("Failed to decode json:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		scheduleRequest.ParseTime()

		schedules := scheduling.FindSchedules(scheduleRequest)
		log.Println("findSchedules finished")
		sort.Sort(sort.Reverse(scheduling.BySchedule(schedules)))
		//log.Println("sorted")
		log.Println("result: ", schedules)
		encoder := json.NewEncoder(rw)
		err = encoder.Encode(schedules)
		if err != nil {
			log.Println("Failed to encode json:", err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Printf("starting Server")
	if DevelopmentMode {
		http.ListenAndServe(":8080", mux)
	} else {
		server := &http.Server{
			Addr:    *listen,
			Handler: mux,
		}
		log.Fatalln(server.ListenAndServe())
	}
}
