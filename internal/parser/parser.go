package parser

import (
	"strings"
	"fmt"
	"github.com/adrg/frontmatter"
	"github.com/Zatfer17/crush/internal/note"
)

func ParseNote(content string) (note.Note, error) {

	var matter struct {
		Id        string `yaml:"id"`
		CreatedAt string `yaml:"created"`
		UpdatedAt string `yaml:"updated"`
	}

	rest, err := frontmatter.MustParse(strings.NewReader(content), &matter)
	if err != nil {
		return note.Note{}, fmt.Errorf("invalid note format")
	}

	n := note.Note{
        Id:        matter.Id,
        CreatedAt: matter.CreatedAt,
        UpdatedAt: matter.UpdatedAt,
        Content:   string(rest),
    }

	return n, nil
}