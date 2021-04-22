package controller

import (
	"encoding/json"
	"net/http"
	"phonebook/database"
	"phonebook/model"
)

func CountContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.CountContactsResponseDTO
	res.Status = "Success"
	res.Count = len(database.Phonebook.Contacts)
	json.NewEncoder(w).Encode(res)
}
