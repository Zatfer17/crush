package core

import (
	"fmt"
	"path/filepath"
	"os"
	"encoding/json"

	"github.com/Zatfer17/zurg/internal/note"
)

func List(basePath string) ([]note.Note, error) {

	var notes []note.Note

	pattern := fmt.Sprintf("%s/*.json", basePath)

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

	