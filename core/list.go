package core

import (
	"os"
	"fmt"
	"path/filepath"
	"encoding/json"

	"github.com/Zatfer17/zurg/internal/book"
	"github.com/Zatfer17/zurg/internal/note"
)

func ListBook(basePath string) ([]book.Book, error) {

	var books []book.Book

	files, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			books = append(books, book.Book{Title: file.Name()})
		}
	}

	return books, nil
}

func ListNote(basePath string, bookName string) ([]note.Note, error) {

	var notes []note.Note

	pattern := fmt.Sprintf("%s/%s/*.json", basePath, bookName)

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		
		f, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		var n note.Note
		err = json.Unmarshal(f, &n)
		if err != nil {
			return nil, err
		}

		notes = append(notes, n)
	}

	return notes, nil
}
