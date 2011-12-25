package main

import "github.com/norisatir/go-gtk3/gtk3"
import "github.com/norisatir/go-gtk3/glib"
import "github.com/norisatir/go-gtk3/gdk3"

func CreateSearchMenu(entry *gtk3.Entry) *gtk3.Menu {
	menu := gtk3.NewMenu()

	item := gtk3.NewImageMenuItemWithMnemonic("Search by _name")
	image := gtk3.NewImageFromStock(gtk3.GtkStock.FIND, gtk3.GtkIconSize.MENU)
	item.SetImage(image)
	item.SetAlwaysShowImage(true)
	item.Connect("activate", func(item gtk3.MenuItemLike, data ...interface{}) {
		entry.SetIconFromStock(gtk3.GtkEntryIconPosition.PRIMARY, gtk3.GtkStock.FIND)
		entry.SetIconTooltipText(gtk3.GtkEntryIconPosition.PRIMARY,
			"Search by name\nClick here to change the search type")
		entry.SetPlaceholderText("name")
	})
	menu.Append(item)

	item = gtk3.NewImageMenuItemWithMnemonic("Search by _description")
	image = gtk3.NewImageFromStock(gtk3.GtkStock.EDIT, gtk3.GtkIconSize.MENU)
	item.SetImage(image)
	item.SetAlwaysShowImage(true)
	item.Connect("activate", func(item gtk3.MenuItemLike, data ...interface{}) {
		entry.SetIconFromStock(gtk3.GtkEntryIconPosition.PRIMARY, gtk3.GtkStock.EDIT)
		entry.SetIconTooltipText(gtk3.GtkEntryIconPosition.PRIMARY,
			"Search by description\nClick here to change the search type")
		entry.SetPlaceholderText("description")
	})
	menu.Append(item)

	item = gtk3.NewImageMenuItemWithMnemonic("Search by _file name")
	image = gtk3.NewImageFromStock(gtk3.GtkStock.OPEN, gtk3.GtkIconSize.MENU)
	item.SetImage(image)
	item.SetAlwaysShowImage(true)
	item.Connect("activate", func(item gtk3.MenuItemLike, data ...interface{}) {
		entry.SetIconFromStock(gtk3.GtkEntryIconPosition.PRIMARY, gtk3.GtkStock.OPEN)
		entry.SetIconTooltipText(gtk3.GtkEntryIconPosition.PRIMARY,
			"Search by file name\nClick here to change the search type")
		entry.SetPlaceholderText("file name")
	})
	menu.Append(item)

	menu.ShowAll()

	return menu
}

func EntryPopulatePopup(entry *gtk3.Entry, menu *gtk3.Menu, data ...interface{}) {
	hasText := entry.GetTextLength() > 0

	item := gtk3.NewSeparatorMenuItem()
	item.Show()
	menu.Append(item)

	item2 := gtk3.NewMenuItemWithMnemonic("C_lear")
	item2.Show()
	item2.Connect("activate", func() { entry.SetText("") })
	menu.Append(item2)
	item2.SetSensitive(hasText)

	searchMenu := CreateSearchMenu(entry)
	item3 := gtk3.NewMenuItemWithLabel("Search by")
	item3.Show()
	item3.SetSubmenu(searchMenu)
	menu.Append(item3)
}

func iconPressCallback(menu *gtk3.Menu, entry *gtk3.Entry, position int, event gdk3.EventButton, data ...interface{}) {
	if position == gtk3.GtkEntryIconPosition.PRIMARY {
		menu.Popup(nil, nil, event.Button, event.Time, nil)
	} else {
		entry.SetText("")
	}
}

func main() {
	gtk3.Init()

	var searchProgressId, finishSearchId uint

	window := gtk3.NewDialogWithButtons("Search Entry", nil, 0, gtk3.B{{gtk3.GtkStock.CLOSE, gtk3.GtkResponse.NONE}})
	window.SetResizable(false)
	window.Connect("response", func() { window.Destroy() })
	window.Connect("destroy", func() {
		if finishSearchId != 0 {
			glib.GSourceRemove(finishSearchId)
		}
		if searchProgressId != 0 {
			glib.GSourceRemove(searchProgressId)
		}
	})

	contentArea := window.GetContentArea().(*gtk3.Box)

	vbox := gtk3.NewVBox(5)
	contentArea.PackStart(vbox, true, true, 0)
	vbox.SetBorderWidth(5)

	label := gtk3.NewLabel("")
	label.SetMarkup("Search entry demo")
	vbox.PackStart(label, false, false, 0)

	hbox := gtk3.NewHBox(10)
	vbox.PackStart(hbox, true, true, 0)
	hbox.SetBorderWidth(0)

	// Create our entry
	entry := gtk3.NewEntry()
	hbox.PackStart(entry, false, false, 0)

	// Create the find and cancel buttons
	notebook := gtk3.NewNotebook()
	notebook.SetShowTabs(false)
	notebook.SetShowBorder(false)
	hbox.PackStart(notebook, false, false, 0)

	startSearchFunc := func() { entry.ProgressPulse() }

	finishSearchFunc := func() bool {
		notebook.SetCurrentPage(0)
		glib.GSourceRemove(searchProgressId)
		searchProgressId = 0
		entry.SetProgressFraction(0.0)
		return false
	}

	findButton := gtk3.NewButtonWithLabel("Find")
	findButton.Connect("clicked", func() {
		notebook.SetCurrentPage(1)

		searchProgressId = glib.GTimeoutAddFull(glib.GPriority.DEFAULT, 100, startSearchFunc)

		finishSearchId = glib.GTimeoutAddSecondsFull(glib.GPriority.DEFAULT, 15, finishSearchFunc)
	})
	notebook.AppendPage(findButton, nil)
	findButton.Show()

	cancelButton := gtk3.NewButtonWithLabel("Cancel")
	cancelButton.Connect("clicked", func() {
		glib.GSourceRemove(finishSearchId)
		finishSearchFunc()
	})
	notebook.AppendPage(cancelButton, nil)
	cancelButton.Show()

	// Set up the search icon
	entry.SetIconFromStock(gtk3.GtkEntryIconPosition.PRIMARY, gtk3.GtkStock.FIND)
	entry.SetIconTooltipText(gtk3.GtkEntryIconPosition.PRIMARY,
		"Search by name\nClick here to change the search type")
	entry.SetPlaceholderText("name")

	// Set up the clear icon
	entry.SetIconFromStock(gtk3.GtkEntryIconPosition.SECONDARY, gtk3.GtkStock.CLEAR)

	menu := CreateSearchMenu(entry)
	menu.AttachToWidget(entry, nil)

	textChangedFunc := func() {
		hasText := entry.GetTextLength() > 0
		entry.SetIconSensitive(gtk3.GtkEntryIconPosition.SECONDARY, hasText)
		findButton.SetSensitive(hasText)
	}

	textChangedFunc()

	entry.Connect("icon-press", iconPressCallback, menu)
	entry.Connect("notify::text", textChangedFunc)

	entry.Connect("activate", func() {
		if searchProgressId != 0 {
			return
		}
		notebook.SetCurrentPage(1)

		searchProgressId = glib.GTimeoutAddFull(glib.GPriority.DEFAULT, 100, startSearchFunc)

		finishSearchId = glib.GTimeoutAddSecondsFull(glib.GPriority.DEFAULT, 15, finishSearchFunc)

	})

	entry.Connect("populate-popup", EntryPopulatePopup)

	button := window.GetWidgetForResponse(gtk3.GtkResponse.NONE)
	button.W().GrabFocus()

	window.ShowAll()
	window.Run()
}
