package core

import (
    "encoding/json"
    "fmt"
    "os"
    "time"

    "github.com/Zatfer17/zurg/internal/note"
)

func Edit(basePath string, Id string, noteContent string) error {

    path := fmt.Sprintf("%s/%s.json", basePath, Id)

    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }

    var n note.Note
    if err := json.Unmarshal(data, &n); err != nil {
        return err
    }

    n.Content = noteContent
    n.UpdatedAt = time.Now().Format(time.RFC3339)

    updatedData, err := json.Marshal(n)
    if err != nil {
        return err
    }

    if err := os.WriteFile(path, updatedData, 0644); err != nil {
        return err
	}

    return nil
}