package tree_store

import "github.com/norisatir/go-gtk3/gobject"
import "github.com/norisatir/go-gtk3/gtk3"

var window *gtk3.Window = nil

type TreeItem struct {
	Label string
	Alex bool
	Havoc bool
	Tim bool
	Owen bool
	Dave bool
	WorldHoliday bool
}

const (
	HolidayNameCol = iota
	AlexCol
	HavocCol
	TimCol
	OwenCol
	DaveCol
	VisibleCol
	WorldCol
)
var months []string

var toplevel map[string][]TreeItem

func createModel() gtk3.TreeModelLike {
	model := gtk3.NewTreeStore([]gobject.GType{gobject.G_TYPE_STRING,
		gobject.G_TYPE_BOOLEAN,
		gobject.G_TYPE_BOOLEAN,
		gobject.G_TYPE_BOOLEAN,
		gobject.G_TYPE_BOOLEAN,
		gobject.G_TYPE_BOOLEAN,
		gobject.G_TYPE_BOOLEAN,
		gobject.G_TYPE_BOOLEAN})

	// Add data to store
	var iter gtk3.TreeIter

	for _, month := range months {
		model.Append(&iter, nil)
		model.SetValues(&iter, gtk3.V{HolidayNameCol:month,
			AlexCol:false,
			HavocCol:false,
			TimCol:false,
			OwenCol:false,
			DaveCol:false,
			VisibleCol:false,
			WorldCol:false})

		items := toplevel[month]

		// Add children
		for _, it := range items {
			var childIter gtk3.TreeIter
			model.Append(&childIter, &iter)
			model.SetValues(&childIter, gtk3.V{HolidayNameCol:it.Label,
				AlexCol:it.Alex,
				HavocCol:it.Havoc,
				TimCol:it.Tim,
				OwenCol:it.Owen,
				DaveCol:it.Dave,
				VisibleCol:true,
				WorldCol:it.WorldHoliday})
		}
	}

	return model
}

func itemToggled(model *gtk3.TreeStore, column int, data ...interface{}) {
	var iter gtk3.TreeIter
	pathStr := data[1].(string)
	path := gtk3.NewTreePathFromString(pathStr)

	model.GetIter(&iter, path)
	toggleItem := model.GetValue(&iter, column).(bool)

	if toggleItem {
		toggleItem = false
	} else {
		toggleItem = true
	}

	model.SetValue(&iter, column, toggleItem)
}

func addColumns(treeview *gtk3.TreeView) {
	var renderer gtk3.CellRendererLike
	model := treeview.GetModel().(*gtk3.TreeStore)

	// Column for holiday names
	renderer = gtk3.NewCellRendererText()
	renderer.CRenderer().Set(gtk3.P{"xalign":0.0})

	col := gtk3.NewTreeViewColumnWithAttributes("Holiday", renderer, gtk3.A{{"text", HolidayNameCol}})
	treeview.AppendColumn(col)
	col.SetClickable(true)

	// Alex column
	renderer = gtk3.NewCellRendererToggle()
	renderer.CRenderer().Set(gtk3.P{"xalign":0.0})
	renderer.CRenderer().Connect("toggled", itemToggled, model, AlexCol)

	col = gtk3.NewTreeViewColumnWithAttributes("Alex", renderer, 
		gtk3.A{{"active",AlexCol}, {"visible",VisibleCol}, {"activatable",WorldCol}})
	treeview.AppendColumn(col)
	col.SetSizing(gtk3.GtkTreeViewColumnSizing.FIXED)
	col.SetFixedWidth(60)
	col.SetClickable(true)

	// Havoc column
	renderer = gtk3.NewCellRendererToggle()
	renderer.CRenderer().Set(gtk3.P{"xalign":0.0})
	renderer.CRenderer().Connect("toggled", itemToggled, model, HavocCol)

	col = gtk3.NewTreeViewColumnWithAttributes("Havoc", renderer, 
		gtk3.A{{"active",HavocCol}, {"visible",VisibleCol}})
	treeview.AppendColumn(col)
	col.SetSizing(gtk3.GtkTreeViewColumnSizing.FIXED)
	col.SetFixedWidth(60)
	col.SetClickable(true)

	// Tim column
	renderer = gtk3.NewCellRendererToggle()
	renderer.CRenderer().Set(gtk3.P{"xalign":0.0})
	renderer.CRenderer().Connect("toggled", itemToggled, model, TimCol)

	col = gtk3.NewTreeViewColumnWithAttributes("Tim", renderer, 
		gtk3.A{{"active",TimCol}, {"visible",VisibleCol}, {"activatable",WorldCol}})
	treeview.AppendColumn(col)
	col.SetSizing(gtk3.GtkTreeViewColumnSizing.FIXED)
	col.SetFixedWidth(60)
	col.SetClickable(true)

	// Owen column
	renderer = gtk3.NewCellRendererToggle()
	renderer.CRenderer().Set(gtk3.P{"xalign":0.0})
	renderer.CRenderer().Connect("toggled", itemToggled, model, OwenCol)

	col = gtk3.NewTreeViewColumnWithAttributes("Owen", renderer, 
		gtk3.A{{"active",OwenCol}, {"visible",VisibleCol}})
	treeview.AppendColumn(col)
	col.SetSizing(gtk3.GtkTreeViewColumnSizing.FIXED)
	col.SetFixedWidth(60)
	col.SetClickable(true)

	// Dave column
	renderer = gtk3.NewCellRendererToggle()
	renderer.CRenderer().Set(gtk3.P{"xalign":0.0})
	renderer.CRenderer().Connect("toggled", itemToggled, model, DaveCol)

	col = gtk3.NewTreeViewColumnWithAttributes("Dave", renderer, 
		gtk3.A{{"active",DaveCol}, {"visible",VisibleCol}})
	treeview.AppendColumn(col)
	col.SetSizing(gtk3.GtkTreeViewColumnSizing.FIXED)
	col.SetFixedWidth(60)
	col.SetClickable(true)

}

func DoTreeStore(w gtk3.WidgetLike) gtk3.WidgetLike {
	if window == nil {
		window = gtk3.NewWindow(gtk3.GtkWindowType.TOPLEVEL)
		window.SetScreen(w.W().GetScreen())
		window.SetTitle("Card planning sheet")
		window.Connect("destroy", func() { window = nil })

		vbox := gtk3.NewVBox(8)
		vbox.SetBorderWidth(8)
		window.Add(vbox)

		vbox.PackStart(gtk3.NewLabel("Jonathan's Holiday Card Planning Sheet"), false, false, 0)

		sw := gtk3.NewScrolledWindow(nil, nil)
		sw.SetShadowType(gtk3.GtkShadow.ETCHED_IN)
		sw.SetPolicy(gtk3.GtkPolicy.AUTOMATIC, gtk3.GtkPolicy.AUTOMATIC)

		vbox.PackStart(sw, true, true, 0)

		// Create model
		model := createModel()

		// Create tree view
		treeview := gtk3.NewTreeViewWithModel(model)
		treeview.SetRulesHint(true)
		treeview.GetSelection().SetMode(gtk3.GtkSelectionMode.MULTIPLE)
		treeview.Connect("realize", func() { treeview.ExpandAll()})

		sw.Add(treeview)

		addColumns(treeview)

		window.SetDefaultSize(650, 400)
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

func init() {
	months = []string{"January", "February", "March", "April", "May", "June", "July",
		"September", "October", "November", "December"}

	
	toplevel = make(map[string][]TreeItem)

	// January
	toplevel["January"] = []TreeItem{
	  {"New Years Day", true, true, true, true, false, true},
	  {"Presidential Inauguration", false, true, false, true, false, false},
	  {"Martin Luther King Jr. day", false, true, false, true, false, false}}

	// February
	toplevel["February"] = []TreeItem{
	  { "Presidents' Day", false, true, false, true, false, false},
	  { "Groundhog Day", false, false, false, false, false, false},
	  { "Valentine's Day", false, false, false, false, true, true}}

	//March
	toplevel["March"] = []TreeItem{ 
	  { "National Tree Planting Day", false, false, false, false, false, false},
	  { "St Patrick's Day", false, false, false, false, false, true}}

	//April
	toplevel["April"] = []TreeItem{
	  { "April Fools' Day", false, false, false, false, false, true},
	  { "Army Day", false, false, false, false, false, false},
	  { "Earth Day", false, false, false, false, false, true},
	  { "Administrative Professionals' Day", false, false, false, false, false, false}}

	//May
	toplevel["May"] = []TreeItem{
	  { "Nurses' Day", false, false, false, false, false, false},
	  { "National Day of Prayer", false, false, false, false, false, false},
	  { "Mothers' Day", false, false, false, false, false, true},
	  { "Armed Forces Day", false, false, false, false, false, false},
	  { "Memorial Day", true, true, true, true, false, true}}

	// June
	toplevel["June"] = []TreeItem{
	  { "June Fathers' Day", false, false, false, false, false, true},
	  { "Juneteenth (Liberation of Slaves)", false, false, false, false, false, false},
	  { "Flag Day", false, true, false, true, false, false}}

	// July
	toplevel["July"] = []TreeItem{
	  { "Parents' Day", false, false, false, false, false, true},
	  { "Independence Day", false, true, false, true, false, false}}
	
	// August
	toplevel["August"] = []TreeItem{
	  { "Air Force Day", false, false, false, false, false, false},
	  { "Coast Guard Day", false, false, false, false, false, false},
	  { "Friendship Day", false, false, false, false, false, false}}
	
	// September
	toplevel["September"] = []TreeItem{
	  { "Grandparents' Day", false, false, false, false, false, true},
	  { "Citizenship Day or Constitution Day", false, false, false, false, false, false},
	  { "Labor Day", true, true, true, true, false, true}}

	// October
	toplevel["October"] = []TreeItem{
	  { "National Children's Day", false, false, false, false, false, false},
	  { "Bosses' Day", false, false, false, false, false, false},
	  { "Sweetest Day", false, false, false, false, false, false},
	  { "Mother-in-Law's Day", false, false, false, false, false, false},
	  { "Navy Day", false, false, false, false, false, false},
	  { "Columbus Day", false, true, false, true, false, false},
	  { "Halloween", false, false, false, false, false, true}}

	// November
	toplevel["November"] = []TreeItem{
	  { "Marine Corps Day", false, false, false, false, false, false},
	  { "Veterans' Day", true, true, true, true, false, true},
	  { "Thanksgiving", false, true, false, true, false, false}}

	// December
	toplevel["December"] = []TreeItem{
	  { "Pearl Harbor Remembrance Day", false, false, false, false, false, false},
	  { "Christmas", true, true, true, true, false, true},
	  { "Kwanzaa", false, false, false, false, false, false}}
}
