package core

import (
	"time"
	
	"github.com/Zatfer17/zurg/internal/note"
)

func Add(basePath string, noteContent string) error {

	ts := time.Now().Local().Truncate(time.Second).Format(time.RFC3339)
	
	n := note.Note{
		CreatedAt: ts,
		UpdatedAt: ts,
		Content  : noteContent,
	}

	err := n.Add(basePath)
	if err != nil {
		return err
	}

	return err
}