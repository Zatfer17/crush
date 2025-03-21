package main

import (
    "fmt"
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "github.com/Zatfer17/zurg/pkg/core"
    "github.com/Zatfer17/zurg/internal/config"
)

func main() {
    app := tview.NewApplication()
    
    cfg, err := config.InitConfig()
    if err != nil {
        panic(err)
    }

    searchBar := tview.NewInputField()
    searchBar.SetLabel("Search: ").
        SetFieldWidth(0).
        SetBorder(true)

    listView := tview.NewList().
        ShowSecondaryText(false)
    listView.SetTitle("Notes").
        SetBorder(true)

    textArea := tview.NewTextArea().
        SetText("", false)
    textArea.SetTitle("Note").
        SetBorder(true)

    updateNotes := func(searchText string) {
        notes, err := core.List(cfg.DefaultPath, searchText)
        if err != nil {
            panic(err)
        }

        listView.Clear()
        for i, note := range notes {
            preview := note.Content
            if len(preview) > 50 {
                preview = preview[:47] + "..."
            }
            listView.AddItem(fmt.Sprintf("(%s) - %s", note.CreatedAt, preview), "", 0, func() {
                textArea.SetText(notes[i].Content, true)
            })
        }

        if len(notes) > 0 {
            listView.SetCurrentItem(0)
            textArea.SetText(notes[0].Content, true)
        } else {
            textArea.SetText("", false)
        }
    }

    updateNotes("")

    searchBar.SetChangedFunc(func(text string) {
        updateNotes(text)
    })

    searchBar.SetDoneFunc(func(key tcell.Key) {
        if key == tcell.KeyEnter {
            notes, _ := core.List(cfg.DefaultPath, searchBar.GetText())
            if len(notes) == 0 {
                textArea.SetText(searchBar.GetText(), true)
                app.SetFocus(textArea)
            } else {
                app.SetFocus(listView)
            }
        }
    })

    listView.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
        notes, _ := core.List(cfg.DefaultPath, searchBar.GetText())
        if index >= 0 && index < len(notes) {
            textArea.SetText(notes[index].Content, true)
        }
    })

    listView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyEnter {
            app.SetFocus(textArea)
            return nil
        }
        return event
    })

    footer := tview.NewTextView().
        SetText("CTRL-K Search • CTRL-S Save • SHIFT-DEL Delete").
        SetTextAlign(tview.AlignCenter)
    footer.SetBorder(true)

    grid := tview.NewGrid().
        SetRows(3, -2, -3, 3).
        SetColumns(0).
        SetBorders(false).
        AddItem(searchBar, 0, 0, 1, 1, 0, 0, true).
        AddItem(listView, 1, 0, 1, 1, 0, 0, false).
        AddItem(textArea, 2, 0, 1, 1, 0, 0, false).
        AddItem(footer, 3, 0, 1, 1, 0, 0, false)

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Key() {
        case tcell.KeyTab:
            if searchBar.HasFocus() {
                app.SetFocus(listView)
            } else if listView.HasFocus() {
                app.SetFocus(textArea)
            } else {
                app.SetFocus(searchBar)
            }
            return nil
        case tcell.KeyBacktab:
            if searchBar.HasFocus() {
                app.SetFocus(textArea)
            } else if listView.HasFocus() {
                app.SetFocus(searchBar)
            } else {
                app.SetFocus(listView)
            }
            return nil
        case tcell.KeyCtrlK:
            app.SetFocus(searchBar)
            return nil
        case tcell.KeyDelete:
            if event.Modifiers()&tcell.ModShift != 0 {
                currentIndex := listView.GetCurrentItem()
                notes, _ := core.List(cfg.DefaultPath, searchBar.GetText())
                
                if currentIndex >= 0 && currentIndex < len(notes) {
                    existingNote := notes[currentIndex]
                    err := core.Remove(
                        cfg.DefaultPath,
                        existingNote.CreatedAt,
                    )
                    if err != nil {
                        panic(err)
                    }
                    updateNotes(searchBar.GetText())
                }
                return nil
            }
            return event
        case tcell.KeyCtrlS:
            content := textArea.GetText()
            currentIndex := listView.GetCurrentItem()
            notes, _ := core.List(cfg.DefaultPath, searchBar.GetText())
            
            var err error
            if currentIndex >= 0 && currentIndex < len(notes) {
                existingNote := notes[currentIndex]
                err = core.Edit(
                    cfg.DefaultPath,
                    existingNote.CreatedAt,
                    content,
                )
            } else {
                err = core.Add(
                    cfg.DefaultPath,
                    content,
                )
            }

            if err != nil {
                panic(err)
            }

            updateNotes(searchBar.GetText())
            return nil
        }
        return event
    })

    if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
        panic(err)
    }
}