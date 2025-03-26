package core

import (
    "fmt"
    "os"
    "time"

    "github.com/Zatfer17/crush/internal/core/parser"
)

func Edit(basePath string, Id string, noteContent string) error {

    path := fmt.Sprintf("%s/%s.md", basePath, Id)

    oldContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}

    n, err := parser.ParseNote(string(oldContent))
    if err != nil {
		return err
	}

    n.Content = noteContent
    n.UpdatedAt = time.Now().Format(time.RFC3339)

    err = n.Write(basePath)
    if err != nil {
        return err
    }

    return nil
}