package main

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "github.com/Zatfer17/zurg/core"
    "github.com/Zatfer17/zurg/internal/config"
    "log"
)

func main() {
    // Initialize config
    cfg, err := config.InitConfig()
    if err != nil {
        log.Fatal(err)
    }

    app := tview.NewApplication()

    booksList := tview.NewList()
    booksList.ShowSecondaryText(false).
        SetBorder(true).
        SetTitle("Books")

    notesList := tview.NewList()
    notesList.ShowSecondaryText(false).
        SetBorder(true).
        SetTitle("Notes")

    noteContent := tview.NewTextArea()
    noteContent.SetBorder(true).
        SetTitle("Note")

    // Load initial books
    books, err := core.ListBook(cfg.DefaultPath)
    if err != nil {
        log.Fatal(err)
    }

    // Populate books list
    for _, book := range books {
        booksList.AddItem(book.Title, "", 0, nil)
    }

	if len(books) > 0 {
        notes, err := core.ListNote(cfg.DefaultPath, books[0].Title)
        if err != nil {
            log.Fatal(err)
        }

        for _, note := range notes {
            notesList.AddItem(note.CreatedAt, "", 0, nil)
        }

        // If there are any notes, show the first note's content
        if len(notes) > 0 {
            noteContent.SetText(notes[0].Content, true)
        }
    }

    // Handle book selection
    booksList.SetChangedFunc(func(index int, title string, secondaryText string, shortcut rune) {
        notesList.Clear()
        
        // Get notes for selected book
        notes, err := core.ListNote(cfg.DefaultPath, title)
        if err != nil {
            // In a real app, you'd want to show this error to the user
            return
        }

        // Populate notes list
        for _, note := range notes {
            notesList.AddItem(note.CreatedAt, "", 0, nil)
        }
    })

	booksList.SetSelectedFunc(func(index int, title string, secondaryText string, shortcut rune) {
        app.SetFocus(notesList)
    })

    // Handle note navigation
    notesList.SetChangedFunc(func(index int, timestamp string, secondaryText string, shortcut rune) {
        book := books[booksList.GetCurrentItem()]
        notes, err := core.ListNote(cfg.DefaultPath, book.Title)
        if err != nil {
            return
        }
        
        if index >= 0 && index < len(notes) {
            noteContent.SetText(notes[index].Content, true)
        }
    })

	notesList.SetSelectedFunc(func(index int, timestamp string, secondaryText string, shortcut rune) {
        app.SetFocus(noteContent)
    })

    // Create layout
    grid := tview.NewGrid().
        SetColumns(-1, -2, -4). // Proportional column widths
        SetBorders(false)

    // Add components to grid
    grid.AddItem(booksList, 0, 0, 1, 1, 0, 0, true).
        AddItem(notesList, 0, 1, 1, 1, 0, 0, false).
        AddItem(noteContent, 0, 2, 1, 1, 0, 0, false)

    // Handle keyboard navigation
    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
        case tcell.KeyTab:
            // Move focus right
            if booksList.HasFocus() {
                app.SetFocus(notesList)
            } else if notesList.HasFocus() {
                app.SetFocus(noteContent)
            } else {
                app.SetFocus(booksList)
            }
            return nil
        case tcell.KeyBacktab: // This is Shift+Tab
            // Move focus left
            if booksList.HasFocus() {
                app.SetFocus(noteContent)
            } else if notesList.HasFocus() {
                app.SetFocus(booksList)
            } else {
                app.SetFocus(notesList)
            }
            return nil
        }
        return event
    })

    if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
        log.Fatal(err)
    }
}