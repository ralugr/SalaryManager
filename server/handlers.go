package server

import (
	"SalaryManager/model"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var validPath = regexp.MustCompile("^/(next_salary|annual_schedule)")

type NextSalary struct {
	Days int
	Date time.Time
}

func Init() {
	http.HandleFunc("/next_salary", makeHandler(nextSalaryHandler))
	http.HandleFunc("/annual_schedule", makeHandler(annualScheduleHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)

		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

func getParam(w http.ResponseWriter, r *http.Request) int {
	keys, ok := r.URL.Query()["payday"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'payday' is missing")
		w.WriteHeader(400) // Bad request
		return -1
	}

	payday, err := strconv.Atoi(keys[0])
	if err != nil {
		log.Println("Conversion was not successful ", err)
		w.WriteHeader(500) // Internal server error
		return -1
	}
	return payday
}

func nextSalaryHandler(w http.ResponseWriter, r *http.Request) {
	payday := getParam(w, r)
	if payday < 0 {
		return
	}
	var response NextSalary
	response.Days = model.DaysTillNextSalary(payday, time.Now())
	response.Date = model.NextPayDate(payday, time.Now())
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println("Encoding failed ", err)
		w.WriteHeader(500) // Internal server error
	}
}

func annualScheduleHandler(w http.ResponseWriter, r *http.Request) {
	payday := getParam(w, r)
	if payday < 0 {
		return
	}

	schedule := model.AnnualPaySchedule(payday, time.Now())
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(schedule)

	if err != nil {
		log.Println("Encoding failed ", err)
		w.WriteHeader(500) // Internal server error
	}
}
