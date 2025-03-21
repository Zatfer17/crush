package note

import (
	"encoding/json"
	"os"
	"fmt"
)

type Note struct {
	CreatedAt string
	UpdatedAt string
	Tags      []string
	Content   string
}

func (n *Note) Add(basePath string) error {

	nj, err := json.Marshal(n)
    if err != nil {
        return fmt.Errorf("could not marshal note")
    }

	path := fmt.Sprintf("%s/%s.json", basePath, n.CreatedAt)
	fmt.Println(path)
	if err := os.WriteFile(path, nj, 0644); err != nil {
		return fmt.Errorf("could not write note to file")
	}

	return nil

}
