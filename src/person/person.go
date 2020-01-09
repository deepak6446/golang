package person

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	// "log"
 )

// The person Type (more like an object)
type Person struct {
	ID        string   
	Firstname string   
	Lastname  string   
	Address   *Address 

 }
 type Address struct {
	City  string 
	State string 
 }
 
var people []Person

func InitPeople(){
	people = append(people, Person{ID: "1", Firstname: "Mohamed", Lastname: "Saidi", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Salem", Lastname: "Baraket", Address: &Address{City: "City Z", State: "State Y"}})
}
// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
 }
 
 // Display a single data
 func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
	   if item.ID == params["id"] {
		  json.NewEncoder(w).Encode(item)
		  return
	   }
	}
	json.NewEncoder(w).Encode(&Person{})
 }
 
 // create a new item
 func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)   //get params 
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	ua := r.Header.Get("User-Agent") // get headers
	w.Header().Set("Content-Type", "application/json")  // set headers
	json.NewEncoder(w).Encode(people)
 }
 
 // Delete an item
 func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
	   if item.ID == params["id"] {
		  people = append(people[:index], people[index+1:]...)
		  break
	   }
	   json.NewEncoder(w).Encode(people)
	}
 }
 