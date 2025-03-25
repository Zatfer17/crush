package core

import (
    "fmt"
    "os"
)

func Remove(basePath string, Id string) error {

    path := fmt.Sprintf("%s/%s.json", basePath, Id)
    if err := os.Remove(path); err != nil {
        return fmt.Errorf("could not remove note: %v", err)
    }

    return nil
	
}