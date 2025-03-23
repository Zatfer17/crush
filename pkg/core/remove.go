package core

import (
    "fmt"
    "os"
)

func Remove(basePath string, baseWorkspace string, createdAt string) error {

    path := fmt.Sprintf("%s/%s/%s.json", basePath, baseWorkspace, createdAt)
    if err := os.Remove(path); err != nil {
        return fmt.Errorf("could not remove note: %v", err)
    }

    return nil
	
}