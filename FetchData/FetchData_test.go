package FetchData

import (
	"testing"
	"log"
)

func TestFetchPersonBatch(t *testing.T) {
	var perList []Person
	FetchPersonBatch(50, &perList)
	if len(perList) != 50 {
		log.Println("Person List:", perList)
		t.Errorf("Expected 50 response. Got %d reponse(s)", len(perList))
	}
}

func TestFetchJokesBatch(t *testing.T) {
	var jokesList Jokes
	FetchJokeBatch(50, &jokesList)
	if len(jokesList.Value) != 50 {
		log.Println("jokes List:", jokesList.Value)
		t.Errorf("Expected 50 response. Got %d reponse(s)", len(jokesList.Value))
	}
}
