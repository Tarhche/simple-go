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
	// load contacts
	var contactsToLoad []Contact

	if err := load("contacts.gob", &contactsToLoad); err != nil {
		log.Fatal("load error:", err)
	}

	log.Println("contacts:", contactsToLoad)
}

func load(filename string, contacts *[]Contact) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	return decoder.Decode(contacts)
}
