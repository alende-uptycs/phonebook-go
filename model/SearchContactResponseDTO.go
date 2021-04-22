package model

type SearchContactResponseDTO struct {
	Status   string       `json:"status"`
	Contacts []ContactDTO `json:"contacts"`
	Count    int          `json:"count"`
}
