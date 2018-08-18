package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"github.com/mathnitin/ChuckNorris/FetchData"
)

var personList []FetchData.Person
var jokeList FetchData.Jokes

func newRouter() *http.ServeMux {
	// Declare a new router
	h := http.NewServeMux()

	// The "HandleFunc" method accepts a path and a function as arguments
	h.HandleFunc("/", handler)
	return h
}

func main() {
	go FetchData.FetchPersonBatch(50, &personList)
	go FetchData.FetchJokeBatch(50, &jokeList)

	h := newRouter()
	err := http.ListenAndServe(":5000", h)
	log.Fatal(err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Validate the Method. Basically only GET is allowed.
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Println("ERROR: Method Not Allowed.")
		return
	}
	// Validate the URL. only localhost:5000 is allowed.
	if r.URL.Path != "/" {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
		log.Println("ERROR: Not Implemented.")
		return
	}

	var wg sync.WaitGroup
	var p FetchData.Person
	var resp string
	var errPerson error
	var errJoke error

fetchAgain:
	if len(personList) == 0 {
		// Currently there are no Persons in buffer. Pull 1 entry quickly to return back.
		wg.Add(1)

		go func() {
			defer wg.Done()
			var body []byte
			log.Println("INFO: Fetching data from URL: http://uinames.com/api/")
			body, errPerson = FetchData.FetchFromUrl("http://uinames.com/api/")
			if errPerson == nil {
				json.Unmarshal(body, &p)
				log.Println("INFO: Unmarshalled data: ", p)
			}
		}()
	} else {
		// Already have entries in buffer.
		p = personList[0]
		personList = personList[1:]
	}

	if len(jokeList.Value) == 0 {
		// Currently there are no jokes in buffer. Pull 1 entry quickly to return back.
		wg.Add(1)

		go func() {
			defer wg.Done()
			type SingleJoke struct {
				Value struct {
					Joke       string   `json:"joke"`
				} `json:"value"`
			}
			var body []byte
			var singleJoke SingleJoke
			log.Println("INFO: Fetching data from URL: http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]")
			body, errJoke = FetchData.FetchFromUrl("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]")
			if errJoke == nil {
				json.Unmarshal(body, &singleJoke)
				log.Println("INFO: Unmarshalled data: ", singleJoke)
			}
			resp = singleJoke.Value.Joke
		}()
	} else {
		// Already have entries in buffer.
		resp = jokeList.Value[0].Joke
		jokeList.Value = jokeList.Value[1:]
	}
	wg.Wait()

	// Make sure error did not occur in any of the get requests and prepare output.

	if errPerson == nil && errJoke == nil {
		log.Println("INFO: No errors found. Preparing and writing response.")
		temp := strings.NewReplacer("John", p.Name, "Doe", p.Surname)
		resp = temp.Replace(resp)
	} else {
		if errPerson != nil {
			resp = errPerson.Error()
		}
		if errJoke != nil {
			resp += errJoke.Error()
		}
	}

	/*
	  TODO: In a loaded system, some times the response is empty. Investigate why and fix it accordingly.
	 */
	log.Println("INFO: Response: ", resp)
	if len(resp) == 0 {
		// for any reason if the response is empty. Try again. This scenario is only hit in case of load.
		goto fetchAgain
	}
	// Send response back.
	fmt.Fprintf(w, resp)

	// Pull the entries using batch, if we don't have anything in buffer for next pull.
	if len(personList) == 0 {
		go FetchData.FetchPersonBatch(50, &personList)
	}
	if len(jokeList.Value) == 0 {
		go FetchData.FetchJokeBatch(50, &jokeList)
	}
}