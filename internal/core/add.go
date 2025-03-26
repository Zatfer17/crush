package core

import (
	"time"
	"hash/fnv"
    "fmt"

	"github.com/Zatfer17/crush/internal/core/note"
)

func createHash(s string) string {
    h := fnv.New32a()
    h.Write([]byte(s))
    return fmt.Sprintf("%04x", h.Sum32())[0:4]
}

func Add(basePath string, noteContent string) (note.Note, error) {

    ts        := time.Now()
    dateStr   := ts.Format("20060102")
    timestamp := ts.Format(time.RFC3339)
    hash      := createHash(timestamp)
    noteId    := fmt.Sprintf("%s-%s", dateStr, hash)
    
    n := note.Note{
        Id:        noteId,
        CreatedAt: timestamp,
        UpdatedAt: timestamp,
        Content:   noteContent,
    }

    err := n.Write(basePath)
    if err != nil {
        return note.Note{}, err
    }

    return n, nil
}