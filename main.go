package main

import (
	//"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	// "sort"

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
	flag.Parse()
	_ = "breakpoint"
	// Initalize Database
	// db, err := gorm.Open("sqlite3", "data/test")
	// if err != nil {
	// }

	var DevelopmentMode = true
	// accessor := &DatabaseAccessor{db}

	mux := http.NewServeMux()

	if !DevelopmentMode {
		staticFiles := rice.MustFindBox("frontend").HTTPBox()
		mux.Handle("/", http.FileServer(staticFiles))
	}

	/*
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				log.Println("options hit")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				return
			}).Methods("OPTIONS")
	*/

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

		var courses []scheduling.Course
		err := decoder.Decode(&courses)
		if err != nil {
			log.Println("Failed to decode json:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("courses: ", courses)
		log.Println("class: ", courses[0].Classes)

		// var courses []string
		// props := make(map[string]scheduling.EventProperties)
		// for _, course := range courses.Courses {
		// 	//log.Println("looping through courses")
		// 	courses = append(courses, course.Course)
		// 	props[course.Course] = course.EventProperties
		// }
		// //log.Println("loop finished")
		// log.Println("courses: ", courses)
		// log.Println("props: ", props)
		// schedules := scheduling.FindSchedules(courses, props, accessor)
		// //log.Println("findSchedules finished")
		// sort.Sort(sort.Reverse(scheduling.BySchedule(schedules)))
		// //log.Println("sorted")
		// log.Println("result: ", schedules)
		// encoder := json.NewEncoder(rw)
		// err = encoder.Encode(schedules)
		// if err != nil {
		// 	log.Println("Failed to encode json:", err)
		// 	rw.WriteHeader(http.StatusInternalServerError)
		// }
	})

	mux.HandleFunc("/api/v0/testClass", func(rw http.ResponseWriter, r *http.Request) {
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

		var classes []scheduling.Class
		err := decoder.Decode(&classes)
		if err != nil {
			log.Println("Failed to decode json:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("classes: ", classes)
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

type CombinationsRequest struct {
	Courses []CoursesRequest
	// TODO: Some pre-existing calendar state
}

type CoursesRequest struct {
	Course string
	scheduling.EventProperties
}

// type DatabaseAccessor struct {
// 	*gorm.DB
// }

// func (accessor *DatabaseAccessor) GetClasses(courseIdentifier string) []scheduling.Class {
// 	classes := []scheduling.Class{}
// 	if courseIdentifier == "" {
// 		log.Println("identifier empty")
// 		return classes
// 	}
// 	accessor.DB.Where(&scheduling.Class{Course: courseIdentifier}).Find(&classes)
// 	return classes
// }

// func (accessor *DatabaseAccessor) GetCourse(courseIdentifier string) scheduling.Course {
// 	course := scheduling.Course{}
// 	accessor.DB.Where(&scheduling.Course{Identifier: courseIdentifier}).First(&course)
// 	return course
// }
