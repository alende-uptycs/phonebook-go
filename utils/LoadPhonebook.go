package utils

import (
	"encoding/gob"
	"os"
	"phonebook/model"
)

func LoadPhonebook(phonebook *model.PhoneBook) {
	phonebook.Contacts = make(map[string]string)
	file, _ := os.Open("phonebook.bak")
	defer file.Close()
	decoder := gob.NewDecoder(file)
	phonebook.Lock()
	decoder.Decode(&phonebook.Contacts)
	phonebook.Unlock()
}
