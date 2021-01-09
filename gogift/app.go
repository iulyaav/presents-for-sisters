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

	personsJSON, err := json.Marshal(persons)
	if err != nil {
		fmt.Fprintf(w, "Cannot encode to JSON ")
	} else {
		fmt.Fprintf(w, "%s", personsJSON)
	}
}

func addGift(w http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		w.Write([]byte("<h1>Welcome to my web server!</h1>"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	// Serve static files from the frontend/dist directory.
	// fs := http.FileServer(http.Dir("./frontend/dist"))
	http.HandleFunc("/", getAllPersons)
	http.HandleFunc("/add-gift", addGift)

	// Start the server.
	fmt.Println("Starting server on port 3000....")
	http.ListenAndServe(":3000", nil)
}
