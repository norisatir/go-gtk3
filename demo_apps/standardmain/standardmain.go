package main

import "github.com/norisatir/go-gtk3/gtk3"

func main() {
	// Initialize gtk3
	gtk3.Init()

	// Create windows
	w := gtk3.NewWindow(gtk3.GtkWindowType.TOPLEVEL)
	w.Connect("destroy", gtk3.MainQuit)
	// Let's set a couple of window properties
	w.Set(gtk3.P{"title": "Go-GTK3 Demo", "resizable": true})

	b2 := gtk3.NewVBox(0)
	w.Add(b2)

	// Create GtkFrame
	f := gtk3.NewFrame("Button Play")
	b2.Add(f)

	// Create GtkBox
	box := gtk3.NewBox(gtk3.GtkOrientation.VERTICAL, 5)
	// Add it to window
	f.Add(box)

	// Create First Button
	fbut := gtk3.NewButtonWithLabel("Click Me")
	box.PackStart(fbut, false, false, 0)
	fbut.Connect("clicked", my_dialog, w)

	// Another one
	fbut2 := gtk3.NewButtonWithLabel("Disable my upper brother")
	box.PackStart(fbut2, false, false, 0)
	// So, let's disable him
	fbut2.Connect("clicked", func(s bool) { fbut.SetSensitive(s) }, false)

	// And another
	fbut3 := gtk3.NewButtonWithLabel("Don't do it")
	box.PackStart(fbut3, false, false, 0)
	fbut2.Connect("clicked", but_disabled, fbut3, fbut)

	fbut3.Connect("clicked", func() {
		fbut.SetSensitive(true)
		fbut3.SetLabel("There you go")
	})

	// Run applicaton
	w.ShowAll()
	gtk3.Main()
}

func but_disabled(b3 *gtk3.Button, b1 *gtk3.Button) {
	b3.SetLabel("Damn it. Enable him")
}

func my_dialog(w *gtk3.Window) {
	// NewDialog
	d := gtk3.NewMessageDialog(w, gtk3.GtkDialogFlags.DESTROY_WITH_PARENT, gtk3.GtkMessage.INFO, gtk3.GtkButtons.CLOSE, "")
	d.SetMarkup("<span foreground=\"blue\" size=\"x-large\">Blue text</span> is <i>cool</i>!")
	d.FormatSecondaryMarkup("<span foreground=\"red\"> So is red!</span>")
	//gobject.RemoveHandler(d, han)
	//d.Connect("response", func() bool { fmt.Println("true"); return true})
	d.Run()
	d.Destroy()
}
