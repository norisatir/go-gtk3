package combobox

import "github.com/norisatir/go-gtk3/gobject"
import "github.com/norisatir/go-gtk3/gtk3"
import "github.com/norisatir/go-gtk3/gdk3"
import gp "github.com/norisatir/go-gtk3/gdkpixbuf"
import "strings"
import "regexp"

const (
	PixbufCol = iota
	TextCol
)

var window *gtk3.Window = nil

func createStockIconStore() gtk3.TreeModelLike {
	stockId := []string{gtk3.GtkStock.DIALOG_WARNING, gtk3.GtkStock.STOP,
		gtk3.GtkStock.NEW, gtk3.GtkStock.CLEAR, "", gtk3.GtkStock.OPEN}

	cellView := gtk3.NewCellView()
	store := gtk3.NewListStore([]gobject.GType{gp.G_TYPE_PIXBUF, gobject.G_TYPE_STRING})

	var iter gtk3.TreeIter

	for _, s := range stockId {
		if s != "" {
			pixbuf := cellView.RenderIconPixbuf(s, gtk3.GtkIconSize.BUTTON)
			_, stockItem := gtk3.StockLookup(s)

			label := strings.Map(func(in rune) rune {
				if in == rune('_') {
					return rune(-1)
				}
				return in
			}, stockItem.Label)

			store.Append(&iter)
			store.SetValues(&iter, gtk3.V{PixbufCol: pixbuf, TextCol: label})
		} else {
			store.Append(&iter)
			store.SetValues(&iter, gtk3.V{PixbufCol: nil, TextCol: "seperator"})
		}
	}
	return store
}

func setSensitive(data ...interface{}) {
	cell := data[1].(gtk3.CellRendererLike)
	treeModel := data[2].(gtk3.TreeModelLike)
	iter := data[3].(*gtk3.TreeIter)

	path := treeModel.ITreeModel().GetPath(iter)
	indices := path.GetIndices()
	sensitive := indices[0] != 1

	cell.CRenderer().Set(gtk3.P{"sensitive": sensitive})
}

func isSeparator(model gtk3.TreeModelLike, iter *gtk3.TreeIter, data ...interface{}) bool {
	path := model.ITreeModel().GetPath(iter)
	return path.GetIndices()[0] == 4
}

func createCapitalStore() gtk3.TreeModelLike {
	type Capital struct {
		Group   string
		Capital string
	}

	capitals := []Capital{
		{"A - B", ""},
		{"", "Albany"},
		{"", "Annapolis"},
		{"", "Atlanta"},
		{"", "Augusta"},
		{"", "Austin"},
		{"", "Baton Rouge"},
		{"", "Bismarck"},
		{"", "Boise"},
		{"", "Boston"},
		{"C - D", ""},
		{"", "Carson City"},
		{"", "Charleston"},
		{"", "Cheyenne"},
		{"", "Columbia"},
		{"", "Columbus"},
		{"", "Concord"},
		{"", "Denver"},
		{"", "Des Moines"},
		{"", "Dover"},
		{"E - J", ""},
		{"", "Frankfort"},
		{"", "Harrisburg"},
		{"", "Hartford"},
		{"", "Helena"},
		{"", "Honolulu"},
		{"", "Indianapolis"},
		{"", "Jackson"},
		{"", "Jefferson City"},
		{"", "Juneau"},
		{"K - O", ""},
		{"", "Lansing"},
		{"", "Lincoln"},
		{"", "Little Rock"},
		{"", "Ljubljana"},
		{"", "Madison"},
		{"", "Montgomery"},
		{"", "Montpelier"},
		{"", "Nashville"},
		{"", "Oklahoma City"},
		{"", "Olympia"},
		{"P - S", ""},
		{"", "Phoenix"},
		{"", "Pierre"},
		{"", "Providence"},
		{"", "Raleigh"},
		{"", "Richmond"},
		{"", "Sacramento"},
		{"", "Salem"},
		{"", "Salt Lake City"},
		{"", "Santa Fe"},
		{"", "Springfield"},
		{"", "St. Paul"},
		{"T - Z", ""},
		{"", "Tallahassee"},
		{"", "Topeka"},
		{"", "Trenton"},
		{"", ""}}

	var iter, iter2 gtk3.TreeIter
	store := gtk3.NewTreeStore([]gobject.GType{gobject.G_TYPE_STRING})

	for _, c := range capitals {
		switch {
		case c.Group != "":
			store.Append(&iter, nil)
			store.SetValues(&iter, gtk3.V{0: c.Group})
		case c.Capital != "":
			store.Append(&iter2, &iter)
			store.SetValues(&iter2, gtk3.V{0: c.Capital})
		}
	}
	return store
}

func isCapitalSensitive(data ...interface{}) {
	cell := data[1].(gtk3.CellRendererLike)
	model := data[2].(gtk3.TreeModelLike)
	iter := data[3].(*gtk3.TreeIter)
	sensitive := !model.ITreeModel().IterHasChild(iter)

	cell.CRenderer().Set(gtk3.P{"sensitive": sensitive})
}

func setBackground(entry *gtk3.Entry, mask string) {
	errorColor := gdk3.RGBA{1.0, 0.9, 0.9, 1.0}
	if mask != "" {
		m, _ := regexp.Compile(mask)
		if m.FindStringIndex(entry.GetText()) == nil {
			entry.OverrideColor(0, &errorColor)
			return
		}
	}
	entry.OverrideColor(0, nil)
}

func DoComboBox(w gtk3.WidgetLike) gtk3.WidgetLike {
	if window == nil {
		window = gtk3.NewWindow(gtk3.GtkWindowType.TOPLEVEL)
		window.SetScreen(w.W().GetScreen())
		window.SetTitle("Combo boxes")
		window.Connect("destroy", func() { window.Destroy(); window = nil} )
		window.SetBorderWidth(10)

		vbox := gtk3.NewVBox(2)
		window.Add(vbox)

		// A combobox demonstrating cell renderers, separators and insensitive rows
		frame := gtk3.NewFrame("Some stock icons")
		vbox.PackStart(frame, false, false, 0)

		box := gtk3.NewVBox(0)
		box.SetBorderWidth(5)
		frame.Add(box)

		model := createStockIconStore()
		combo := gtk3.NewComboBoxWithModel(model)
		box.Add(combo)

		var renderer gtk3.CellRendererLike

		renderer = gtk3.NewCellRendererPixbuf()
		combo.PackStart(renderer, false)
		combo.SetAttributes(renderer, gtk3.A{{"pixbuf", PixbufCol}})
		combo.SetCellDataFunc(renderer, setSensitive)

		renderer = gtk3.NewCellRendererText()
		combo.PackStart(renderer, true)
		combo.SetAttributes(renderer, gtk3.A{{"text", TextCol}})
		combo.SetCellDataFunc(renderer, setSensitive)

		combo.SetRowSeparatorFunc(isSeparator)
		combo.SetActive(0)

		// A ComboBox demonstrating trees.
		frame = gtk3.NewFrame("Where are we ?")
		vbox.PackStart(frame, false, false, 0)

		box = gtk3.NewVBox(0)
		box.SetBorderWidth(5)
		frame.Add(box)

		model1 := createCapitalStore()
		combo3 := gtk3.NewComboBoxWithModel(model1)
		box.Add(combo3)

		renderer1 := gtk3.NewCellRendererText()
		combo3.PackStart(renderer1, true)
		combo3.SetAttributes(renderer1, gtk3.A{{"text", 0}})

		combo3.SetCellDataFunc(renderer1, isCapitalSensitive)

		var iter gtk3.TreeIter
		path := gtk3.NewTreePathFromString("3:3")
		model1.ITreeModel().GetIter(&iter, path)
		combo3.SetActiveIter(&iter)

		// a combobox with validation
		frame = gtk3.NewFrame("Editable")
		vbox.PackStart(frame, false, false, 0)

		box = gtk3.NewVBox(0)
		box.SetBorderWidth(5)
		frame.Add(box)

		combo1 := gtk3.NewComboBoxTextWithEntry()
		box.Add(combo1)
		combo1.AppendText("One")
		combo1.AppendText("Two")
		combo1.AppendText("2\302\275")
		combo1.AppendText("Three")

		mask := "^([0-9]*|One|Two|2\302\275|Three)$"
		entry := combo1.GetChild().(*gtk3.Entry)
		entry.Connect("changed", setBackground, entry, mask)

		// A ComboBox with string IDs
		frame = gtk3.NewFrame("String IDs")
		vbox.PackStart(frame, false, false, 0)

		box = gtk3.NewVBox(0)
		vbox.SetBorderWidth(5)
		frame.Add(box)

		combo2 := gtk3.NewComboBoxText()
		combo2.Append("never", "Not visible")
		combo2.Append("when-active", "Visible when active")
		combo2.Append("always", "Always visible")
		box.Add(combo2)

		entry = gtk3.NewEntry()
		box.Add(entry)

		gobject.BindProperty(combo2, "active-id", entry, "text", gobject.GBindingFlags.BIDIRECTIONAL)
	}

	if !window.GetVisible() {
		window.ShowAll()
	} else {
		window.Destroy()
		window = nil
		return nil
	}

	return window
}
