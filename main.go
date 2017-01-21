package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"sort"

	"github.com/GeertJohan/go.rice"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/mibzman/CourseCorrect-Student/scheduling"
)

var (
	listen = flag.String("listen", ":8080", "The adress this service will be available on.")
	dbPath = flag.String("path", "", "The path at which the database containing class scheduling information")
)

func main() {
	flag.Parse()

	// Initalize Database
	db, err := gorm.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatal(err)
	}

	accessor := &DatabaseAccessor{db}

	mux := http.NewServeMux()

	staticFiles := rice.MustFindBox("frontend").HTTPBox()
	mux.Handle("/", http.FileServer(staticFiles))

	// TODO: Restify this API!
	mux.HandleFunc("/api/v0/courses", func(rw http.ResponseWriter, r *http.Request) {
		// Query Database
		var courses []scheduling.Course
		db.Find(&courses)

		// Send Response
		encoder := json.NewEncoder(rw)
		err := encoder.Encode(courses)
		if err != nil {
			log.Println("Failed to encode json:", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/api/v0/schedules", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Invalid method to post endpoint:", r.Method)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var constraints CombinationsRequest
		err := decoder.Decode(&constraints)
		if err != nil {
			log.Println("Failed to decode json:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var courses []string
		props := make(map[string]scheduling.EventProperties)
		for _, course := range constraints.Courses {
			courses = append(courses, course.Course)
			props[course.Course] = course.EventProperties
		}

		schedules := scheduling.FindSchedules(courses, props, accessor)
		sort.Sort(sort.Reverse(scheduling.BySchedule(schedules)))

		encoder := json.NewEncoder(rw)
		err = encoder.Encode(schedules)
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
	Courses []CoursesRequest
	// TODO: Some pre-existing calendar state
}

type CoursesRequest struct {
	Course string
	scheduling.EventProperties
}

type DatabaseAccessor struct {
	*gorm.DB
}

func (accessor *DatabaseAccessor) GetClasses(courseIdentifier string) []scheduling.Class {
	classes := []scheduling.Class{}
	accessor.DB.Where(&scheduling.Class{Course: courseIdentifier}).Find(&classes)
	return classes
}

func (accessor *DatabaseAccessor) GetCourse(courseIdentifier string) scheduling.Course {
	course := scheduling.Course{}
	accessor.DB.Where(&scheduling.Course{Identifier: courseIdentifier}).First(&course)
	return course
}
