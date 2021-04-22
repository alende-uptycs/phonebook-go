package controller

import (
	"encoding/json"
	"net/http"
	"phonebook/database"
	"phonebook/model"
	"phonebook/utils"
)

func AddContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	var name string = query.Get("name")
	var phoneNumber string = query.Get("phonenumber")

	database.Phonebook.Lock()
	database.Phonebook.Contacts[name] = phoneNumber
	database.Phonebook.Unlock()

	var res model.AddContactResponseDTO
	res.Status = "Success"
	res.TotalContacts = len(database.Phonebook.Contacts)
	json.NewEncoder(w).Encode(res)
	utils.StorePhonebook(&database.Phonebook)
}
