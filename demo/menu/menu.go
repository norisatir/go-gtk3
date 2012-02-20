package menu

import "github.com/norisatir/go-gtk3/gtk3"
import "github.com/norisatir/go-gtk3/glib"
import "fmt"

var window *gtk3.Window = nil

func createMenu(depth int, tearoff bool) gtk3.WidgetLike {
	if depth < 1 {
		return nil
	}

	var group *glib.GSList
	menu := gtk3.NewMenu()

	if tearoff {
		mt := gtk3.NewTearoffMenuItem()
		menu.Append(mt)
		mt.Show()
	}

	for i, j := 0, 1; i < 5; i, j = i+1, j+1 {
		mi := gtk3.NewRadioMenuItemWithLabel(group, fmt.Sprintf("item %2d - %d", depth, j))
		group = mi.GetGroup()

		menu.Append(mi)
		mi.Show()

		if i == 3 {
			mi.SetSensitive(false)
		}
		mi.SetSubmenu(createMenu(depth-1, true))
	}

	return menu
}

func changeOrientation(menubar *gtk3.MenuBar, button *gtk3.Button, data ...interface{}) {
	parent := menubar.GetParent().(*gtk3.Box)
	orientation := parent.GetOrientation()
	parent.SetOrientation(1 - orientation)

	if orientation == gtk3.GtkOrientation.VERTICAL {
		menubar.SetPackDirection(gtk3.GtkPackDirection.TTB)
	} else {
		menubar.SetPackDirection(gtk3.GtkPackDirection.LTR)
	}
}
func DoMenu(w gtk3.WidgetLike) gtk3.WidgetLike {
	if window == nil {
		window = gtk3.NewWindow(gtk3.GtkWindowType.TOPLEVEL)
		window.SetScreen(w.W().GetScreen())
		window.SetTitle("Menus")
		window.Connect("destroy", func() { window.Destroy(); window = nil })

		accelGroup := gtk3.NewAccelGroup()
		window.AddAccelGroup(accelGroup)
		window.SetBorderWidth(0)

		box := gtk3.NewHBox(0)
		window.Add(box)
		box.Show()

		box1 := gtk3.NewVBox(0)
		box.Add(box1)
		box1.Show()

		menubar := gtk3.NewMenuBar()
		box1.PackStart(menubar, false, true, 0)
		menubar.Show()

		menu := createMenu(2, true)

		menuitem := gtk3.NewMenuItemWithLabel("test\nline2")
		menuitem.SetSubmenu(menu)
		menubar.Append(menuitem)
		menuitem.Show()

		menuitem = gtk3.NewMenuItemWithLabel("foo")
		menuitem.SetSubmenu(createMenu(3, true))
		menubar.Append(menuitem)
		menuitem.Show()

		menuitem = gtk3.NewMenuItemWithLabel("bar")
		menuitem.SetSubmenu(createMenu(4, true))
		menubar.Append(menuitem)
		menuitem.Show()

		box2 := gtk3.NewVBox(10)
		box2.SetBorderWidth(10)
		box1.PackStart(box2, false, true, 0)
		box2.Show()

		button := gtk3.NewButtonWithLabel("Flip")
		button.Connect("clicked", changeOrientation, menubar)
		box2.PackStart(button, true, true, 0)
		button.Show()

		button = gtk3.NewButtonWithLabel("Close")
		button.Connect("clicked", func() { window.Destroy() })
		box2.PackStart(button, true, true, 0)
		button.SetCanDefault(true)
		button.GrabDefault()
		button.Show()
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
