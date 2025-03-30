## Markdown rendering
```golang
rendered, err := glamour.Render(note.Content, "dark")
if err != nil {
    panic(err)
}
writer := tview.ANSIWriter(noteView)
fmt.Fprint(writer, rendered)
```