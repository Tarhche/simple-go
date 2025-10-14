package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Contact struct {
	Name  string
	Email string
	Phone string
}

func main() {
	// save contacts
	contactsToSave := []Contact{
		{Name: "John Doe", Email: "john.doe@example.com", Phone: "1234567890"},
		{Name: "Jane Doe", Email: "jane.doe@example.com", Phone: "0987654321"},
	}

	if err := save(contactsToSave, "contacts.gob"); err != nil {
		log.Fatal("save error:", err)
	}
}

func save(contacts []Contact, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	defer file.Sync()

	encoder := gob.NewEncoder(file)

	return encoder.Encode(contacts)
}
