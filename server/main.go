package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func handler1(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Handler 1")
}

func Now(w http.ResponseWriter, _ *http.Request) {
	now := time.Now()

	p := make(map[string]string)
	p["now"] = now.Format(time.ANSIC)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(p)

}

func zone(w http.ResponseWriter, _ *http.Request) {
	loc, _ := time.LoadLocation("Europe/Berlin")
	p1 := make(map[string]string)
	p1["Berlin"] = time.Now().In(loc).Format(time.ANSIC)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p1)

}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func handler3(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": ReadUserIP(r),
		// "date": time.Now().,
	})
	w.Write(resp)

}

// func convert json file to map
func convert(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	jsonFile := `{"id": "BM-1347",
	 	"title": "The underage storm", 
	 	"Content": "The creatives' careers can easily get uncreative but yet creative...",
	  	"Summary": "Seeking freedom"}`
	var result map[string]string

	json.Unmarshal([]byte(jsonFile), &result)

	json.NewEncoder(w).Encode(result)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")    // clients wants to get
	r.HandleFunc("/us", handler1).Methods("GET") // clients wants to get
	r.HandleFunc("/now", Now).Methods("GET")
	r.HandleFunc("/zone", zone).Methods("GET")
	r.HandleFunc("/ip", handler3).Methods("GET")
	r.HandleFunc("/convert", convert).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
