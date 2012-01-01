package main

import (
	"github.com/norisatir/go-gtk3/gtk3"
	"github.com/norisatir/go-gtk3/pango"
	"github.com/norisatir/go-gtk3/gobject"
	"./assistant"
	"./button_box"
	"./combobox"
	"./entry_buffer"
	"./entry_completion"
	"./search_entry"
	"./dialog"
	"./menu"
	"./spinner"
	"./textview"
	"./list_store"
	"./tree_store"
)

type Demo struct {
	Title string
	Filename string
	CallBack func(gtk3.WidgetLike) gtk3.WidgetLike
	Children []Demo
}

var Demos []Demo

func init() {
	// Fill in demo apps
	Demos = []Demo{{"Assistant", "assistant.go", assistant.DoAssistant, nil},
		{"Button Boxes", "button_box.go", button_box.DoButtonBox, nil},
		{"Combo Boxes", "combobox.go", combobox.DoComboBox, nil},
		{"Entry", "", nil, []Demo{
			{"Entry Buffer", "entry_buffer.go", entry_buffer.DoEntryBuffer, nil},
			{"Entry Completion", "entry_completion.go", entry_completion.DoEntryCompletion, nil},
			{"Search Entry", "search_entry.go", search_entry.DoSearchEntry, nil}}},
		{"Dialog with Message Boxes", "dialog.go", dialog.DoDialog, nil},
		{"Menus", "menu.go", menu.DoMenu, nil},
		{"Spinner", "spinner.go", spinner.DoSpinner, nil},
		{"Text Widget", "", nil, []Demo{
			{"Multiple Views", "textview.go", textview.DoTextView, nil}}},
		{"TreeView", "", nil, []Demo{
			{"List Store", "list_store.go", list_store.DoListStore, nil},
			{"Tree Store", "tree_store.go", tree_store.DoTreeStore, nil}}}}
}

const (
	TitleColumn = iota
	StyleColumn
)


func CreateText(buffer **gtk3.TextBuffer, isSource bool) gtk3.WidgetLike {
	scrolledW := gtk3.NewScrolledWindow(nil, nil)
	scrolledW.SetPolicy(gtk3.GtkPolicy.AUTOMATIC, gtk3.GtkPolicy.AUTOMATIC)
	scrolledW.SetShadowType(gtk3.GtkShadow.IN)

	*buffer = gtk3.NewTextBuffer(nil)
	textView := gtk3.NewTextView()
	textView.SetBuffer(*buffer)
	textView.SetEditable(false)
	textView.SetCursorVisible(false)

	scrolledW.Add(textView)

	if isSource {
		fontDesc := pango.NewFontDescriptionFromString("monospace")
		textView.OverrideFont(fontDesc)
		textView.SetWrapMode(gtk3.GtkWrapMode.NONE)
	} else {
		// Make it a bit nicer for text
		textView.SetWrapMode(gtk3.GtkWrapMode.WORD)
		textView.SetPixelsAboveLines(2)
		textView.SetPixelsBelowLines(2)
	}

	return scrolledW
}

func windowClosed(model *gtk3.TreeStore, path *gtk3.TreePath) {
	var iter gtk3.TreeIter

	model.GetIter(&iter, path)
	style := model.GetValue(&iter, StyleColumn).(int)

	if style == pango.PangoStyle.ITALIC {
		model.SetValue(&iter, StyleColumn, pango.PangoStyle.NORMAL)
	}
}

func rowActivated(model *gtk3.TreeStore, 
		treeView *gtk3.TreeView, path *gtk3.TreePath, 
		column *gtk3.TreeViewColumn, data ...interface{}) {

	var iter gtk3.TreeIter
	var call func(gtk3.WidgetLike) gtk3.WidgetLike

	model.GetIter(&iter, path)
	style := model.GetValue(&iter, StyleColumn).(int)

	indices := path.GetIndices()
	if len(indices) > 1 {
		call = Demos[indices[0]].Children[indices[1]].CallBack
	} else {
		call = Demos[indices[0]].CallBack
	}

	if (call != nil) {
		if style == pango.PangoStyle.ITALIC {
			style = pango.PangoStyle.NORMAL
		} else {
			style = pango.PangoStyle.ITALIC
		}
		model.SetValue(&iter, StyleColumn, style)
		window := call(treeView.GetTopLevel())

		if window != nil {
			window.Connect("destroy", windowClosed, model, path.Copy())
		}
	}
}

func CreateTree() gtk3.WidgetLike {
	model := gtk3.NewTreeStore([]gobject.GType{gobject.G_TYPE_STRING, gobject.G_TYPE_INT})

	treeView := gtk3.NewTreeViewWithModel(model)
	selection := treeView.GetSelection()
	selection.SetMode(gtk3.GtkSelectionMode.BROWSE)

	var iter gtk3.TreeIter

	for _,demo := range Demos {
		model.Append(&iter, nil)
		model.SetValues(&iter, gtk3.V{TitleColumn:demo.Title, StyleColumn:pango.PangoStyle.NORMAL})

		if len(demo.Children) > 0 {
			var childIter gtk3.TreeIter
			for _, childDemo := range demo.Children {
				model.Append(&childIter, &iter)
				model.SetValues(&childIter, gtk3.V{TitleColumn:childDemo.Title, StyleColumn:pango.PangoStyle.NORMAL})
			}
		}
	}


	cell := gtk3.NewCellRendererText()
	column := gtk3.NewTreeViewColumnWithAttributes("Widget (double click for demo)",
		cell,
		gtk3.A{{"text", TitleColumn}, {"style", StyleColumn}})

	treeView.AppendColumn(column)

	model.GetIterFirst(&iter)
	selection.SelectIter(&iter)
	
	treeView.Connect("row_activated", rowActivated, model)

	treeView.CollapseAll()
	treeView.SetHeadersVisible(false)

	sw := gtk3.NewScrolledWindow(nil, nil)
	sw.SetPolicy(gtk3.GtkPolicy.AUTOMATIC, gtk3.GtkPolicy.AUTOMATIC)
	sw.Add(treeView)

	label := gtk3.NewLabel("Widget (double click for demo)")

	box := gtk3.NewNotebook()
	box.AppendPage(sw, label)

	treeView.GrabFocus()

	return box
}


func main() {
	gtk3.Init()

	var infoBuf, sourceBuf *gtk3.TextBuffer

	window := gtk3.NewWindow(gtk3.GtkWindowType.TOPLEVEL)
	window.SetTitle("Gtk+ Code Demos")
	window.Connect("destroy", gtk3.MainQuit)


	hbox := gtk3.NewHBox(0)
	window.Add(hbox)

	// Create tree
	tree := CreateTree()
	hbox.PackStart(tree, false, false, 0)


	notebook := gtk3.NewNotebook()
	hbox.PackStart(notebook, true ,true, 0)

	notebook.AppendPage(CreateText(&infoBuf, false), gtk3.NewLabelWithMnemonic("_Info"))

	infoBuf.CreateTag("title", gtk3.P{"font":"Sans 18"})

	notebook.AppendPage(CreateText(&sourceBuf, true), gtk3.NewLabelWithMnemonic("_Source"))

	sourceBuf.CreateTag("comment", gtk3.P{"foreground":"DodgerBlue"})
	window.SetDefaultSize(600, 400)
	window.ShowAll()

	gtk3.Main()
}
