package entry_completion

import "github.com/norisatir/go-gtk3/gobject"
import "github.com/norisatir/go-gtk3/gtk3"

var window *gtk3.Dialog = nil

func createCompletionModel() gtk3.TreeModelLike {
	store := gtk3.NewListStore([]gobject.GType{gobject.G_TYPE_STRING})

	var iter gtk3.TreeIter
	// Append one word
	store.Append(&iter)
	store.SetValue(&iter, 0, "GNOME")

	// Append one word
	store.Append(&iter)
	store.SetValue(&iter, 0, "total")

	// Append one word
	store.Append(&iter)
	store.SetValue(&iter, 0, "totally")

	return store
}

func DoEntryCompletion(w gtk3.WidgetLike) gtk3.WidgetLike {
	if window == nil {
		window = gtk3.NewDialogWithButtons("GtkEntryCompletion", w.(*gtk3.Window),
			0, gtk3.B{{gtk3.GtkStock.CLOSE, gtk3.GtkResponse.NONE}})
		window.SetResizable(false)

		window.Connect("response", func() { window.Destroy() })
		window.Connect("destroy", func() { window = nil })

		contentArea := window.GetContentArea().(*gtk3.Box)

		vbox := gtk3.NewVBox(5)
		contentArea.PackStart(vbox, true, true, 0)
		vbox.SetBorderWidth(5)

		label := gtk3.NewLabel("")
		label.SetMarkup("Completion demo, try writing <b>total</b> or <b>gnome</b> for example.")
		vbox.PackStart(label, false, false, 0)

		// Create entry
		entry := gtk3.NewEntry()
		vbox.PackStart(entry, false, false, 0)

		// Create completion object
		completion := gtk3.NewEntryCompletion()

		// Assign completion to entry
		entry.SetCompletion(completion)

		// Create tree model and use it as the completion model
		completion_model := createCompletionModel()
		completion.SetModel(completion_model)
		completion.SetTextColumn(0)

	}

	if !window.GetVisible() {
		window.ShowAll()
	} else {
		window.Destroy()
		return nil
	}

	return window
}
