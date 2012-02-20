package entry_buffer

import "github.com/norisatir/go-gtk3/gtk3"

var dialog *gtk3.Dialog = nil

func DoEntryBuffer(w gtk3.WidgetLike) gtk3.WidgetLike {
	if dialog == nil {
		dialog = gtk3.NewDialogWithButtons("GtkEntryBuffer",
			w.(*gtk3.Window), 0,
			gtk3.B{{gtk3.GtkStock.CLOSE, gtk3.GtkResponse.NONE}})
		dialog.SetResizable(false)
		dialog.Connect("response", func() { dialog.Destroy() })
		dialog.Connect("destroy", func() { dialog = nil })

		content_area := dialog.GetContentArea().(*gtk3.Box)

		vbox := gtk3.NewVBox(5)
		content_area.PackStart(vbox, true, true, 0)

		label := gtk3.NewLabel("")
		label.SetMarkup("Entries share a buffer. Typing in one is reflected in the other.")
		vbox.PackStart(label, false, false, 0)

		// Create a buffer
		buffer := gtk3.NewEntryBuffer("")

		// Create entry with visibility off
		entry := gtk3.NewEntryWithBuffer(buffer)
		entry.Set(gtk3.P{"visibility": false})
		vbox.PackStart(entry, false, false, 0)

		// Create second entry
		entry = gtk3.NewEntryWithBuffer(buffer)
		vbox.PackStart(entry, false, false, 0)
	}

	if !dialog.GetVisible() {
		dialog.ShowAll()
	} else {
		dialog.Destroy()
		dialog = nil
		return nil
	}

	return dialog
}
