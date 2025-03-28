package main

import (
	"strings"
    "fmt"
    "path"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Zatfer17/crush/internal/config"
    "github.com/Zatfer17/crush/internal/core"
)

func updateCollections(cfg config.Config, collectionsView *tview.List, searchBar *tview.InputField) {

    fullPath := path.Join(cfg.DefaultPath, ".queries")

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

func removeCollection(cfg config.Config, query string) {

    fullPath := path.Join(cfg.DefaultPath, ".queries")

    err := core.Remove(
        fullPath,
        query,
    )
    if err != nil {
        panic(err)
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

func updateNotes(cfg config.Config, searchText string, notesView *tview.List, noteView *tview.TextArea) {

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

func removeNote(cfg config.Config, id string, noteView *tview.TextArea) {

    err := core.Remove(
        cfg.DefaultPath,
        id,
    )
    if err != nil {
        panic(err)
    }

    if id == noteView.GetTitle() {
        noteView.SetTitle("")
        noteView.SetText("", true)
    } 

}

func main() {

    cfg, err := config.InitConfig()
    if err != nil {
        panic(err)
    }

    app := tview.NewApplication()

    collectionsView := tview.NewList()
    collectionsView.SetBorder(true)
    collectionsView.SetTitle("Collections")
    collectionsView.ShowSecondaryText(false)

    noteView := tview.NewTextArea()
    noteView.SetBorder(true)

    notesView := tview.NewList()
    notesView.SetBorder(true)
    notesView.SetTitle("Notes")
    notesView.ShowSecondaryText(false)

    searchBar := tview.NewInputField()
    searchBar.SetLabel("Search: ")
    searchBar.SetBorder(true)
    searchBar.SetChangedFunc(func(searchText string) {
        updateNotes(*cfg, searchText, notesView, noteView)
    })
    searchBar.SetDoneFunc(func(key tcell.Key){
        addNote(*cfg, searchBar.GetText(), app, noteView)
        updateNotes(*cfg, searchBar.GetText(), notesView, noteView)
    })

    footer := tview.NewTextView()
    footer.SetTextAlign(tview.AlignCenter)
    footer.SetBorder(true)
    footer.SetText("CTRL-N New • CTRL-S Save")

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

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
            case tcell.KeyTab:
                if searchBar.HasFocus() {
                    app.SetFocus(collectionsView)
                    footer.SetText("CTRL-K Search • CTRL-S Save • DEL Delete")
                } else if collectionsView.HasFocus() {
                    app.SetFocus(notesView)
                    footer.SetText("CTRL-N New • CTRL-K Search • DEL Delete")
                } else if notesView.HasFocus() {
                    app.SetFocus(noteView)
                    footer.SetText("CTRL-N New • CTRL-K Search • CTRL-S Save • DEL Delete")
                } else {
                    app.SetFocus(searchBar)
                    footer.SetText("CTRL-N New • CTRL-S Save")
                }
                return nil
            case tcell.KeyBacktab:
                if searchBar.HasFocus() {
                    app.SetFocus(noteView)
                    footer.SetText("CTRL-N New • CTRL-K Search • CTRL-S Save • DEL Delete")
                } else if collectionsView.HasFocus() {
                    app.SetFocus(searchBar)
                    footer.SetText("CTRL-N New • CTRL-S Save")
                } else if notesView.HasFocus() {
                    app.SetFocus(collectionsView)
                    footer.SetText("CTRL-K Search • CTRL-S Save • DEL Delete")
                } else {
                    app.SetFocus(notesView)
                    footer.SetText("CTRL-N New • CTRL-K Search • DEL Delete")
                } 
                return nil
            case tcell.KeyCtrlK:
                app.SetFocus(searchBar)
                return nil
            case tcell.KeyCtrlN:
                if app.GetFocus() != collectionsView{
                    addNote(*cfg, "", app, noteView)
                    updateNotes(*cfg, searchBar.GetText(), notesView, noteView)
                }
                return nil
            case tcell.KeyCtrlS:
                if app.GetFocus() == searchBar || app.GetFocus() == collectionsView {
                    if searchBar.GetText() != "" {
                        addCollection(*cfg, searchBar.GetText())
                        updateCollections(*cfg, collectionsView, searchBar)
                    }
                } else if app.GetFocus() == noteView {
                    if noteView.GetTitle() != "" {
                        editNote(*cfg, noteView)
                        updateNotes(*cfg, searchBar.GetText(), notesView, noteView)
                    }
                }
                return nil
            case tcell.KeyDelete:
                if app.GetFocus() == collectionsView {
                    fullPath := path.Join(cfg.DefaultPath, ".queries")
                    collections, err := core.List(fullPath, "")
                    if err != nil {
                        panic(err)
                    }
                    selectedCollection := collections[collectionsView.GetCurrentItem()]
                    removeCollection(*cfg, selectedCollection.Id)
                    updateCollections(*cfg, collectionsView, searchBar)
                } else if app.GetFocus() == notesView {
                    notes, err := core.List(cfg.DefaultPath, searchBar.GetText())
                    if err != nil {
                        panic(err)
                    }
                    selectedNote := notes[notesView.GetCurrentItem()]
                    removeNote(*cfg, selectedNote.Id, noteView)
                    updateNotes(*cfg, searchBar.GetText(), notesView, noteView)
                } else if app.GetFocus() == noteView {
                    removeNote(*cfg, noteView.GetTitle(), noteView)
                    updateNotes(*cfg, searchBar.GetText(), notesView, noteView)
                }
                return nil
        }
        return event
    })

    updateCollections(*cfg, collectionsView, searchBar)
    updateNotes(*cfg, "", notesView, noteView)

    if err := app.SetRoot(grid, true).EnableMouse(true).SetFocus(searchBar).Run(); err != nil {
        panic(err)
    }

}