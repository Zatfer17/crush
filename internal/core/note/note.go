package note

import (
	"path/filepath"
	"os"
	"fmt"
)

var TEMPLATE = `---
id: %s
created: %s
updated: %s
---
%s`

type Note struct {
	Id        string
	CreatedAt string
	UpdatedAt string
	Content   string
}

func (note Note) GetName() string {
	return fmt.Sprintf("%s.md", note.Id)
}

func (note Note) Format() string {
	return fmt.Sprintf(TEMPLATE, note.Id, note.CreatedAt, note.UpdatedAt, note.Content)
}

func (note Note) Write(basePath string) error {

	filePath := filepath.Join(basePath, note.GetName())

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(note.Format())
	return err

}
