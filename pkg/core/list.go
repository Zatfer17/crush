package core

import (
	"fmt"
	"path/filepath"
	"os"
	"os/exec"
	"strings"
	"encoding/json"
	"sort"

	"github.com/Zatfer17/zurg/internal/note"
)

func List(basePath string, baseWorkspace string, content string) ([]note.Note, error) {
    var notes []note.Note
    var files []string

    if content != "" {

        path := fmt.Sprintf("%s/%s", basePath, baseWorkspace)

        cmd := exec.Command("grep", "-ril", content, "--include=*.json", path)
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
        pattern := fmt.Sprintf("%s/%s/*.json", basePath, baseWorkspace)
        var err error
        files, err = filepath.Glob(pattern)
        if err != nil {
            return nil, err
        }
    }

    sort.Sort(sort.Reverse(sort.StringSlice(files)))

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

	