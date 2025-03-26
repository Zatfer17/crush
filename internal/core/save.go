package core

import (
	"time"
	"path/filepath"
	"os"

	"github.com/Zatfer17/crush/internal/core/note"
)

func Save(basePath string, query string) (note.Note, error) {

	fullPath := filepath.Join(basePath, ".queries")
	
    err := os.MkdirAll(fullPath, 0755)
    if err != nil {
        return note.Note{}, err
    }

	timestamp := time.Now().Format(time.RFC3339)
    
    n := note.Note{
        Id:        query,
        CreatedAt: timestamp,
        UpdatedAt: timestamp,
        Content:   query,
    }

    err = n.Write(fullPath)
    if err != nil {
        return note.Note{}, err
    }

    return n, nil

}