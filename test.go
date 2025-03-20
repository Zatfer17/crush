package main

import (
	"fmt"
	"log"

	"github.com/Zatfer17/zurg/core"
	"github.com/Zatfer17/zurg/internal/config"
)

func run() {

	config, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = core.AddBook(config.DefaultPath, config.DefaultBookName)
	if err != nil {
		log.Fatal(err)
	}

	err = core.AddNote(config.DefaultPath, config.DefaultBookName, "This is a note")
	if err != nil {
		log.Fatal(err)
	}

	books, err := core.ListBook(config.DefaultPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)

	note, err := core.ListNote(config.DefaultPath, config.DefaultBookName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(note)

}

	