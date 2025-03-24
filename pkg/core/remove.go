package core

import (
    "fmt"
    "os"
)

func Remove(basePath string, createdAt string) error {

    path := fmt.Sprintf("%s/%s.json", basePath, createdAt)
    if err := os.Remove(path); err != nil {
        return fmt.Errorf("could not remove note: %v", err)
    }

    return nil
	
}