package controller

import (
	"encoding/json"
	"net/http"
	"phonebook/database"
	"phonebook/model"
)

func SearchContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	var name string = query.Get("name")

	_, ok := database.Phonebook.Contacts[name]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		var res model.ErrorDTO
		res.Status = "Failed"
		res.Description = "Contact not found."
		json.NewEncoder(w).Encode(res)
	} else {
		var res model.SearchContactResponseDTO
		res.Status = "Success"
		res.Contacts = append(res.Contacts, model.ContactDTO{Name: name, Phonenumber: database.Phonebook.Contacts[name]})
		res.Count = 1
		json.NewEncoder(w).Encode(res)
	}
}
