package book

import (
	"os"
	"fmt"
)

type Book struct {
	Title    string
}

func (b *Book) Add(basePath string) error {

	_, err := os.Stat(basePath)
    if err != nil {
        return fmt.Errorf("base path does not exist")
    }

	path := basePath + "/" + b.Title

	if err := os.Mkdir(path, 0755); err != nil && !os.IsExist(err) {
		return fmt.Errorf("could not create directory")
	}

	return nil

}

