package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/GeertJohan/go.rice"
	"github.com/mibzman/CourseCorrect/mock"
	"github.com/mibzman/CourseCorrect/scheduling"
)

var (
	listen = flag.String("listen", ":8080", "The adress this service will be available on.")
)

type test struct {
	blaa        string
	blaa2       string
	bsasdla4    string
	fdknenawekn string
}

func main() {
	flag.Parse()
	mux := http.NewServeMux()

	staticFiles := rice.MustFindBox("frontend").HTTPBox()
	mux.Handle("/", http.FileServer(staticFiles))

	mux.HandleFunc("/api/courses", func(rw http.ResponseWriter, r *http.Request) {
		encoder := json.NewEncoder(rw)
		err := encoder.Encode(mock.S4Courses)
		if err != nil {
			log.Println("Failed to encode json:", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/api/combos", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Invalid to post endpoint:", r.Method)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		// TODO: Close body?
		decoder := json.NewDecoder(r.Body)

		var constraints CombinationsRequest
		err := decoder.Decode(&constraints)
		if err != nil {
			log.Println("Failed to decode json:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		criteria := scheduling.Criteria{
			EarliestClass: scheduling.Criterion{
				Time:      *constraints.StartTime,
				Manditory: true,
				Weight:    10,
			},
			LatestClass: scheduling.Criterion{
				Time:      *constraints.EndTime,
				Manditory: true,
				Weight:    10,
			},
			Days: scheduling.Criterion{
				Other:     constraints.Days,
				Manditory: true,
				Weight:    10,
			},
		}

		combos := scheduling.GenerateCombos(constraints.Courses)
		for i := range combos {
			combo := &combos[i]
			sort.Sort(scheduling.ByStartTime(combo.Classes))
			combo.Score = scheduling.ScoreCombo(*combo, criteria)
		}
		sort.Sort(sort.Reverse(scheduling.ByScore(combos)))

		encoder := json.NewEncoder(rw)
		err = encoder.Encode(combos)
		if err != nil {
			log.Println("Failed to encode json:", err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	})

	server := &http.Server{
		Addr:    *listen,
		Handler: mux,
	}

	log.Printf("Starting server on %s\n", *listen)
	log.Fatalln(server.ListenAndServe())
}

type CombinationsRequest struct {
	Courses   []scheduling.Course
	StartTime *time.Time
	EndTime   *time.Time
	Days      string
}
