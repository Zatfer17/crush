package main

import (
    "fmt"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"

    "github.com/Zatfer17/zurg/internal/config"
    "github.com/Zatfer17/zurg/pkg/core"
)

func main() {

    cfg, err := config.InitConfig()
    if err != nil {
        panic(err)
    }

    app := tview.NewApplication()

    searchBar := tview.NewInputField()
    searchBar.SetLabel("Search: ")
    searchBar.SetBorder(true)

    collectionsView := tview.NewList()
    collectionsView.SetBorder(true)
    collectionsView.SetTitle("Collections")
    collectionsView.ShowSecondaryText(false)

    notesView := tview.NewList()
    notesView.SetBorder(true)
    notesView.SetTitle("Notes")
    notesView.ShowSecondaryText(false)

    noteView := tview.NewTextView()
    noteView.SetBorder(true)

    footer := tview.NewTextView()
    footer.SetText("CTRL-N New • CTRL-K Search • CTRL-S Save • CTRL-DEL Delete")
    footer.SetTextAlign(tview.AlignCenter)
    footer.SetBorder(true)

    grid := tview.NewGrid()
    grid.SetRows(3, 0, 0, 0, 0, 0, 3)
    grid.SetColumns(-1, -1, -1, -1, -1, -1, -1, -1, -1, -1)
    grid.AddItem(searchBar, 0, 0, 1, 10, 0, 0, false)
    grid.AddItem(footer, 6, 0, 1, 10, 0, 0, false)
        
    grid.AddItem(collectionsView, 1, 0, 2, 4, 0, 0, false)
    grid.AddItem(notesView, 1, 4, 2, 6, 0, 0, false)
    grid.AddItem(noteView, 3, 0, 3, 10, 0, 0, false)
    
    grid.AddItem(collectionsView, 1, 0, 5, 2, 0, 100, false)
    grid.AddItem(notesView, 1, 2, 5, 3, 0, 100, false)
    grid.AddItem(noteView, 1, 5, 5, 5, 0, 100, false)

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
            case tcell.KeyTab:
                if searchBar.HasFocus() {
                    app.SetFocus(collectionsView)
                } else if collectionsView.HasFocus() {
                    app.SetFocus(notesView)
                } else if notesView.HasFocus() {
                    app.SetFocus(noteView)
                } else {
                    app.SetFocus(searchBar)
                }
                return nil
            case tcell.KeyBacktab:
                if searchBar.HasFocus() {
                    app.SetFocus(noteView)
                } else if collectionsView.HasFocus() {
                    app.SetFocus(searchBar)
                } else if notesView.HasFocus() {
                    app.SetFocus(collectionsView)
                } else {
                    app.SetFocus(notesView)
                } 
                return nil
            }
        return event
    })

    updateNotes := func(searchText string) {

        notes, err := core.List(cfg.DefaultPath, searchText)
        if err != nil {
            panic(err)
        }

        noteView.Clear()
        for i, note := range notes {
            preview := note.Content
            if len(preview) > 50 {
                preview = preview[:47] + "..."
            }
            notesView.AddItem(fmt.Sprintf("[%s] %s", note.Id, preview), "", 0, func() {noteView.SetText(notes[i].Content)})
        }

        if len(notes) > 0 {
            notesView.SetCurrentItem(0)
            noteView.SetText(notes[0].Content)
        } else {
            noteView.Clear()
        }
    }

    updateNotes(searchBar.GetText())

    if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
        panic(err)
    }
}