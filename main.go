package main

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/globalsign/mgo"
	"github.com/kidstuff/mongostore"
)

func main() {
	log.Println("start...")

	// Fetch new store.
	dbsess, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer dbsess.Close()
	log.Println("client created")

	store := mongostore.NewMongoStore(dbsess.DB("test").C("test_session"), 3600, true, []byte("secret-key"))
	log.Println("store created")

	// Request y writer for testing
	req, _ := http.NewRequest("GET", "http://www.example.com", nil)
	w := httptest.NewRecorder()

	// Get a session.
	session, err := store.Get(req, "session-key")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("got session")

	// Add a value
	session.Values["foo"] = "bar"

	// Save.
	if err = session.Save(req, w); err != nil {
		log.Printf("Error saving session: %v", err)
	}

	log.Println("session saved")
}
