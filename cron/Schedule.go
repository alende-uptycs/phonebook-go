package cron

import (
	"phonebook/model"
	"phonebook/utils"
	"time"
)

func Schedule(phonebook *model.PhoneBook) {
	go func() {
		for {
			utils.StorePhonebook(phonebook)
			time.Sleep(time.Minute * 5)
		}
	}()
}
