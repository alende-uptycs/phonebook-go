package model

import "sync"

type PhoneBook struct {
	sync.Mutex
	Contacts map[string]string
}
