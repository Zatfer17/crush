package core

import (
	"os/exec"
	"strings"
	"os"
	"encoding/json"

	"github.com/Zatfer17/zurg/internal/note"
)

func Search(basePath string, content string) ([]note.Note, error) {

	var notes []note.Note

	cmd := exec.Command("grep", "-ril", content, "--include=*.json", basePath)
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
			output = []byte{}
		} else {
			return nil, err
		}
	}

	files := strings.FieldsFunc(string(output), func(r rune) bool { return r == '\n' })

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
