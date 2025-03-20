package core

import (
	"time"
	
	"github.com/Zatfer17/zurg/internal/book"
	"github.com/Zatfer17/zurg/internal/note"
)

func AddBook(basePath string, bookName string) error {

	b := book.Book{
		Title:    bookName,
	}

	err := b.Add(basePath)
	if err != nil {
		return err
	}

	return nil
}

func AddNote(basePath string, bookName string, noteContent string) error {

	ts := time.Now().Local().Truncate(time.Second).Format(time.RFC3339)
	
	n := note.Note{
		Book:      bookName,
		CreatedAt: ts,
		UpdatedAt: ts,
		Content:   noteContent,
	}

	err := n.Add(basePath)
	if err != nil {
		return err
	}

	return nil
}