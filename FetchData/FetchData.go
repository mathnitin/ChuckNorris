package FetchData

import (
	"strconv"
	"log"
	"encoding/json"
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Jokes struct {
	Value []struct {
		Joke string `json:"joke"`
	} `json:"value"`
}

/*
  TODO: Open a channel and keep a check in this function. If the number of entries in the list reduces to less than 10,
        fetch the batch again.
 */
func FetchPersonBatch(count int, personList *[]Person) {
	// Set count to 50 as the max batch supported is 50.
	if count > 50 {
		count = 50
	}

	/*
	TODO: Have a max number of retries. Currently this can be an infinite loop.
	 */
fetch:
	url := "http://uinames.com/api/?amount=" + strconv.Itoa(count)
	log.Println("INFO: Fetching data from URL: ", url)
	body, err := FetchFromUrl(url)
	if err == nil {
		json.Unmarshal(body, personList)
	} else {
		// Keep on retrying. Sleep for 5 millisecond as there can be server loop detection.
		time.Sleep(5 * time.Millisecond)
		log.Println("INFO: Failed batch request from URL: ", url)
		goto fetch
	}
}

/*
  TODO: Open a channel and keep a check in this function. If the number of entries in the list reduces to less than 10,
        fetch the batch again.
 */
func FetchJokeBatch(count int, jokeList *Jokes) {
	// Set count to 50 as the max batch supported is 50.
	if count > 50 {
		count = 50
	}

	/*
	TODO: Have a max number of retries. Currently this can be an infinite loop.
	 */
fetch:
	url := "http://api.icndb.com/jokes/random/" + strconv.Itoa(count) + "?firstName=John&lastName=Doe&limitTo=[nerdy]"
	log.Println("INFO: Fetching data from URL: ", url)
	body, err := FetchFromUrl(url)
	if err == nil {
		json.Unmarshal(body, jokeList)
	} else {
		// Keep on retrying. Sleep for 5 millisecond as there can be server loop detection.
		time.Sleep(5 * time.Millisecond)
		log.Println("INFO: Failed batch request from URL: ", url)
		goto fetch
	}
}

/*
  TODO: Instead of returning the body, return the complete response. Based on the StatusCode in response user can
        make decision to retry or not.
*/
func FetchFromUrl(url string) (body []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println("ERROR: ", err.Error())
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()

	/*
		If status code is not 200, populate error correctly and set body to nil
	*/
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("status response: %v for URL: %v", res.Status, url)
		log.Println("ERROR: ", err.Error())
	} else {
		body, err = ioutil.ReadAll(res.Body)
	}
	return
}

