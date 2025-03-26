package main

import (
    "fmt"
    "path/filepath"
    "strings"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"

    "github.com/Zatfer17/crush/internal/config"
    "github.com/Zatfer17/crush/pkg/core"
)

func getCollections(cfg config.Config, collectionsView *tview.List, searchBar *tview.InputField) {

    fullPath := filepath.Join(cfg.DefaultPath, ".queries")

    collectionsView.Clear()

    notes, err := core.List(fullPath, "")
    if err != nil {
        panic(err)
    }

    for _, note := range notes {
        collectionsView.AddItem(note.Content, "", 0, func(){
            searchBar.SetText(note.Content)
        })
    }
}

func addCollection(cfg config.Config, query string) {

    _, err := core.Save(
        cfg.DefaultPath,
        query,
    )
    if err != nil {
        panic(err)
    }

}

func getNotes(cfg config.Config, searchText string, notesView *tview.List, noteView *tview.TextArea) {

    notesView.Clear()

    notes, err := core.List(cfg.DefaultPath, searchText)
    if err != nil {
        panic(err)
    }

    for _, note := range notes {
        preview := note.Content
        preview = strings.Replace(preview, "\n", " ", -1)
        if len(preview) > 50 {
            preview = preview[:47] + "..."
        }
        notesView.AddItem(fmt.Sprintf("[%s] %s", note.Id, preview), "", 0, func() {
            noteView.SetTitle(note.Id)
            noteView.SetText(note.Content, true)
        })
    }
}

func addNote(cfg config.Config, noteContent string, app *tview.Application, noteView *tview.TextArea) {

    n, err := core.Add(
        cfg.DefaultPath,
        noteContent,
    )
    if err != nil {
        panic(err)
    }

    noteView.SetTitle(n.Id)
    noteView.SetText(n.Content, true)

    app.SetFocus(noteView)

}

func editNote(cfg config.Config, noteView *tview.TextArea) {

    err := core.Edit(
        cfg.DefaultPath,
        noteView.GetTitle(),
        noteView.GetText(),
    )
    if err != nil {
        panic(err)
    }

}

func removeNote(cfg config.Config, noteView *tview.TextArea) {

    if noteView.GetTitle() != "" {
        err := core.Remove(
            cfg.DefaultPath,
            noteView.GetTitle(),
        )
        if err != nil {
            panic(err)
        }
    
        noteView.SetTitle("")
        noteView.SetText("", true)
    }

}

func main() {

    cfg, err := config.InitConfig()
    if err != nil {
        panic(err)
    }

    // CREATE WIDGETS

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

    noteView := tview.NewTextArea()
    noteView.SetBorder(true)

    footer := tview.NewTextView()
    footer.SetText("CTRL-N New • CTRL-K Search • CTRL-S Save • CTRL-W Save query • SHIFT-DEL Delete")
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
    
    grid.AddItem(collectionsView, 1, 0, 5, 2, 0, 150, false)
    grid.AddItem(notesView, 1, 2, 5, 3, 0, 150, false)
    grid.AddItem(noteView, 1, 5, 5, 5, 0, 150, false)

    // ASSIGN SHORTCUTS

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
            case tcell.KeyCtrlK:
                app.SetFocus(searchBar)
                return nil
            case tcell.KeyCtrlN:
                addNote(*cfg, "", app, noteView)
                getNotes(*cfg, searchBar.GetText(), notesView, noteView)
                return nil
            case tcell.KeyCtrlS:
                editNote(*cfg, noteView)
                getNotes(*cfg, searchBar.GetText(), notesView, noteView)
                return nil
            case tcell.KeyCtrlW:
                addCollection(*cfg, searchBar.GetText())
                getCollections(*cfg, collectionsView, searchBar)
                return nil
            case tcell.KeyDelete:
                if event.Modifiers()&tcell.ModShift != 0 {
                    removeNote(*cfg, noteView)
                    getNotes(*cfg, searchBar.GetText(), notesView, noteView)
                    return nil
                }
        }
        return event
    })

    // ASSIGN FUNCTIONALITY

    searchBar.SetChangedFunc(func(searchText string) {
        getNotes(*cfg, searchText, notesView, noteView)
    })

    searchBar.SetDoneFunc(func(key tcell.Key){
        addNote(*cfg, searchBar.GetText(), app, noteView)
        getNotes(*cfg, searchBar.GetText(), notesView, noteView)
    })

    getCollections(*cfg, collectionsView, searchBar)
    getNotes(*cfg, "", notesView, noteView)

    if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
        panic(err)
    }

}