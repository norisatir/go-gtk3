package main

import "github.com/norisatir/go-gtk3/gtk3"


func main() {
	app := gtk3.NewApplication("si.go-gtk3.demoapp", gtk3.G_APPLICATION_FLAGS_NONE)

	// Create windows
	w := gtk3.NewWindow(gtk3.GTK_WINDOW_TOPLEVEL, nil)
	// Add window to GTKApplication
	app.AddWindow(w)
	// Let's set a couple of window properties
	w.Set(gtk3.P{"title": "Go-GTK3 Demo", "resizable":false})

    //Create frame
    f := gtk3.NewFrame("Button play")
	// Add it to window
	w.Add(f)
    
	
	// Create GtkBox
	box := gtk3.NewBox(gtk3.ORIENTATION_VERTICAL, 5)
    f.Add(box)

	// Create First Button
	fbut := gtk3.NewButtonWithLabel("Click Me")
	box.PackStart(fbut, false, false, 0)
	fbut.Connect("clicked", func() { box.ReorderChild(fbut, 2)})

	// Another one
	fbut2 := gtk3.NewButtonWithLabel("Disable my upper brother")
	box.PackStart(fbut2, false, false, 0)
	// So, let's disable him
	fbut2.Connect("clicked", func(s bool) {fbut.SetSensitive(s)}, false)

	// And another
	fbut3 := gtk3.NewButtonWithLabel("Don't do it")
	box.PackStart(fbut3, false, false, 0)
	fbut2.Connect("clicked", but_disabled, fbut3, fbut)

	fbut3.Connect("clicked", func() { 
					fbut.SetSensitive(true)
				   	fbut3.SetLabel("There you go")})

	// Run applicaton
	w.ShowAll()
	app.Run()
}

func but_disabled(b3 *gtk3.Button, b1 *gtk3.Button) {
	b3.SetLabel("Damn it. Enable him")
}
