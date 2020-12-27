package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


func getAllPersons(w http.ResponseWriter, req *http.Request) {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened file...")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var persons Persons
	json.Unmarshal(byteValue, &persons)


	// var persons = Persons{newPerson("Mery"), newPerson("Daia"), newPerson("Iulia")}
	personsJson, err := json.Marshal(persons)
	if err != nil {
        fmt.Fprintf(w, "Cannot encode to JSON ")
    } else {
		fmt.Fprintf(w, "%s", personsJson)
	}
}

func main() {
	http.HandleFunc("/", getAllPersons)
	fmt.Println("Starting server on port 8000....")
	http.ListenAndServe(":8000", nil)
}
