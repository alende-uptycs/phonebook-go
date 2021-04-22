package utils

import (
	"encoding/gob"
	"fmt"
	"os"
	"phonebook/model"
)

func StorePhonebook(phonebook *model.PhoneBook) {
	file, _ := os.Create("phonebook.bak")
	defer file.Close()
	encoder := gob.NewEncoder(file)
	encoder.Encode(phonebook.Contacts)
	fmt.Println("Phonebook backed up successfully.")
}
