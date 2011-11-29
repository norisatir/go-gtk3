package main

import "github.com/norisatir/go-gtk3/gtk3"


func main() {
	gtk3.Init()

	window := gtk3.NewWindow(gtk3.GTK_WINDOW_TOPLEVEL, nil)
	window.SetTitle("Dialogs")
	window.Connect("destroy", gtk3.MainQuit)
	window.SetBorderWidth(8)

	frame := gtk3.NewFrame("Dialogs")
	window.Add(frame)

	vbox := gtk3.NewVBox(8)
	vbox.SetBorderWidth(8)
	frame.Add(vbox)

	// Standard message dialog
	hbox := gtk3.NewHBox(8)
	vbox.PackStart(hbox, false, false, 0)

	button := gtk3.NewButtonWithMnemonic("_Message Dialog")
    button.Connect("clicked", DialogClickedClosure(window))
	hbox.PackStart(button, false, false, 0)

	vbox.PackStart(gtk3.NewHSeparator(), false, false, 0)

	// Interactive dialog
	hbox = gtk3.NewHBox(8)
	vbox.PackStart(hbox, false, false, 0)
	vbox2 := gtk3.NewVBox(0)
	
	button = gtk3.NewButtonWithMnemonic("_Interactive Dialog")
	hbox.PackStart(vbox2, false, false, 0)
	vbox2.PackStart(button, false, false, 0)

	table := gtk3.NewGrid()
	table.SetRowSpacing(4)
	table.SetColumnSpacing(4)
	hbox.PackStart(table, false, false, 0)

	label := gtk3.NewLabelWithMnemonic("_Entry 1")
    table.Attach(label, 0, 0, 1, 1)

	entry1 := gtk3.NewEntry()
	table.Attach(entry1, 1, 0, 1, 1)
	label.SetMnemonicWidget(entry1)

	label = gtk3.NewLabelWithMnemonic("E_ntry 2")
	table.Attach(label, 0, 1, 1, 1)

	entry2 := gtk3.NewEntry()
	table.Attach(entry2, 1, 1, 1, 1)
	label.SetMnemonicWidget(entry2)

	button.Connect("clicked", InteractiveDialog, window, entry1, entry2)

	window.ShowAll()
	gtk3.Main()
}

func DialogClickedClosure(w *gtk3.Window) func() {
	var i int = 1
	
	return func() {
		dialog := gtk3.NewMessageDialog(w,
			gtk3.GtkDialogFlags.MODAL | gtk3.GtkDialogFlags.DESTROY_WITH_PARENT,
			gtk3.GtkMessage.INFO,
			gtk3.GtkButtons.OK,
			"This message box has been popped up the following\nnumber of times:")
		dialog.FormatSecondaryText("%d", i)

		dialog.Run()
		dialog.Destroy()
		i++
	}
}

func InteractiveDialog(w *gtk3.Window, entry1, entry2 *gtk3.Entry) {
	dialog := gtk3.NewDialogWithButtons("Interactive Dialog",
		w,
		gtk3.GtkDialogFlags.MODAL | gtk3.GtkDialogFlags.DESTROY_WITH_PARENT,
		gtk3.B{{gtk3.GtkStock.OK, gtk3.GtkResponse.OK}, {"_Non-stock Button", gtk3.GtkResponse.CANCEL}})
	
	content_area := dialog.GetContentArea().(*gtk3.Box)

	hbox := gtk3.NewHBox(8)
	hbox.SetBorderWidth(uint(8))
	content_area.PackStart(hbox, false, false, 0)

	stock := gtk3.NewImageFromStock(gtk3.GtkStock.DIALOG_QUESTION, gtk3.GtkIconSize.DIALOG)
	hbox.PackStart(stock, false, false, 0)

	table := gtk3.NewGrid()
	table.SetRowSpacing(4);
	table.SetColumnSpacing(4);
	hbox.PackStart(table, true, true, 0)

	label := gtk3.NewLabelWithMnemonic("_Entry 1")
	table.Attach(label, 0, 0, 1, 1)

	local_entry1 := gtk3.NewEntry()
	local_entry1.SetText(entry1.GetText())
	table.Attach(local_entry1, 1, 0, 1, 1)
	label.SetMnemonicWidget(local_entry1)

	label = gtk3.NewLabelWithMnemonic("E_ntry 2")
	table.Attach(label, 0, 1, 1, 1)

	local_entry2 := gtk3.NewEntry()
	local_entry2.SetText(entry2.GetText())
	table.Attach(local_entry2, 1, 1, 1, 1)
	label.SetMnemonicWidget(local_entry2)

	hbox.ShowAll()
	response := dialog.Run()

	if response == gtk3.GtkResponse.OK {
		entry1.SetText(local_entry1.GetText())
		entry2.SetText(local_entry2.GetText())
	}

	dialog.Destroy()
}
