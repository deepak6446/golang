package main

import (
//    "encoding/json"
   "github.com/gorilla/mux"
   "log"
   "net/http"
   Person "person"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)
        // Call the next handler, which can be another middleware in the chain, or the final handler.
        next.ServeHTTP(w, r)
    })
}

// main function to boot up everything
func main() {
	
	Person.InitPeople()
	router := mux.NewRouter()
	router.HandleFunc("/people", Person.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", Person.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", Person.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", Person.DeletePerson).Methods("DELETE")

	router.Use(loggingMiddleware)

	log.Print("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}