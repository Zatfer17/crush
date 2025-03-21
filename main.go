package main

import (
	"fmt"
	"log"

	"github.com/Zatfer17/zurg/internal/config"
    "github.com/Zatfer17/zurg/pkg/core"
)

func main() {

    cfg, err := config.InitConfig()
    if err != nil {
        log.Fatal(err)
    }

    var tags []string
    err = core.Add(cfg.DefaultPath, tags, "This is a test")
    if err != nil {
        log.Fatal(err)
    }

    notes, err := core.List(cfg.DefaultPath)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Println(notes)

    notes, err = core.Search(cfg.DefaultPath, "note")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(notes)

    err = core.Remove(cfg.DefaultPath, notes[0].CreatedAt)
    if err != nil {
        log.Fatal(err)
    }

    notes, err = core.Search(cfg.DefaultPath, "note")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(notes)

    err = core.Edit(cfg.DefaultPath, notes[0].CreatedAt, tags, "This is a testooooo")
    if err != nil {
        log.Fatal(err)
    }

    notes, err = core.Search(cfg.DefaultPath, "testooooo")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(notes)

}