package core

import (
	"fmt"
	"path/filepath"
	"os"
	"os/exec"
	"strings"
	"sort"

	"github.com/Zatfer17/crush/internal/core/note"
    "github.com/Zatfer17/crush/internal/core/parser"
)

func List(basePath string, content string) ([]note.Note, error) {

    var notes []note.Note
    var files []string

    if content != "" {

        cmd := exec.Command("grep", "-ril", content, "--exclude-dir=.queries", "--include=*.md", basePath)
        output, err := cmd.Output()
        if err != nil {
            if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
                output = []byte{}
            } else {
                return nil, err
            }
        }
        files = strings.FieldsFunc(string(output), func(r rune) bool { return r == '\n' })

    } else {
        pattern := fmt.Sprintf("%s/*.md", basePath)
        var err error
        files, err = filepath.Glob(pattern)
        if err != nil {
            return nil, err
        }
    }

    for _, file := range files {

        fileContent, err := os.ReadFile(file)
        if err != nil {
            return nil, err
        }

        n, err := parser.ParseNote(string(fileContent))
        if err != nil {
            return nil, err
        }

        notes = append(notes, n)
    }

    sort.Slice(notes, func(i, j int) bool {
        return notes[i].CreatedAt > notes[j].CreatedAt
    })

    return notes, nil
}	