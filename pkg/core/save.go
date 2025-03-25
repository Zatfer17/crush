package core

import (
	"time"
	"path/filepath"
	"os"

	"github.com/Zatfer17/zurg/internal/note"
)

func Save(basePath string, query string) error {

	fullPath := filepath.Join(basePath, "queries")
	
    err := os.MkdirAll(fullPath, 0755)
    if err != nil {
        return err
    }

	ts        := time.Now()
    timestamp := ts.Format(time.RFC3339)
    
    n := note.Note{
        Id:        query,
        CreatedAt: timestamp,
        UpdatedAt: timestamp,
        Content:   query,
    }

    err = n.Add(fullPath)
    if err != nil {
        return err
    }

    return nil

}