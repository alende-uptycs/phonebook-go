/** PROBLEM STATEMENT
 *
 * Create a web app (in go or node.js) to maintain student phone registry. It should listen on port 9876
 *
 * It should support following methods
 * GET /search
 * Parameters: name (student name)
 * Return: student phone number (if found) or error
 *
 * GET /count
 * Parameters: none
 * Return: Number of students in the registry
 *
 * POST /add
 * Parameters: name (student name), phone (phone number)
 * Return: error if student name already exists.
 *
 * Your server should dump your phone registry to the disk every 5 minutes.
 * And when you start/restart, you should first load the phone registry (if it is there) from the disk.
 */
package main

import (
	"log"
	"net/http"

	"phonebook/controller"
	"phonebook/cron"
	"phonebook/database"
	"phonebook/utils"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	utils.LoadPhonebook(&database.Phonebook)
	cron.Schedule(&database.Phonebook)

	// Route Handlers/Endpoints
	router.HandleFunc("/count", controller.CountContacts).Methods("GET")
	router.HandleFunc("/add", controller.AddContact).Methods("POST")
	router.HandleFunc("/search", controller.SearchContacts).Methods("GET")

	log.Fatal(http.ListenAndServe(":9876", router))
}
