package main

import "github.com/norisatir/go-gtk3/gtk3"


func CreateBBox(orientation int, title string, spacing int, layout int) *gtk3.Frame {

	frame := gtk3.NewFrame(title)

	bbox := gtk3.NewButtonBox(orientation)	
	bbox.SetLayout(layout)

	button := gtk3.NewButtonFromStock(gtk3.GtkStock.OK)
	bbox.Add(button)

	button = gtk3.NewButtonFromStock(gtk3.GtkStock.CANCEL)
	bbox.Add(button)

	button = gtk3.NewButtonFromStock(gtk3.GtkStock.HELP)
	bbox.Add(button)

	frame.Add(bbox)

	return frame
}


func DoButtonBox() *gtk3.Window {
	window := gtk3.NewWindow(gtk3.GtkWindowType.TOPLEVEL, nil)
	window.SetTitle("Button Boxes")
	window.Connect("destroy", gtk3.MainQuit)

	main_vbox := gtk3.NewVBox(0)
	window.Add(main_vbox)

	frame_horz := gtk3.NewFrame("Horizontal Button Boxes")
	main_vbox.PackStart(frame_horz, true, true, 10)

	vbox := gtk3.NewVBox(0)
	frame_horz.Add(vbox)

	vbox.PackStart(CreateBBox(gtk3.GtkOrientation.HORIZONTAL, "Spread", 40, gtk3.GtkButtonBoxStyle.SPREAD), true, true, 0)
	vbox.PackStart(CreateBBox(gtk3.GtkOrientation.HORIZONTAL, "Edge", 40, gtk3.GtkButtonBoxStyle.EDGE), true, true, 5)
	vbox.PackStart(CreateBBox(gtk3.GtkOrientation.HORIZONTAL, "Start", 40, gtk3.GtkButtonBoxStyle.START), true, true, 5)
	vbox.PackStart(CreateBBox(gtk3.GtkOrientation.HORIZONTAL, "End", 40, gtk3.GtkButtonBoxStyle.END), true, true, 5)


	frame_vert := gtk3.NewFrame("Vertical Button Boxes")
	main_vbox.PackStart(frame_vert, true, true, 10)


	hbox := gtk3.NewHBox(0)
	frame_vert.Add(hbox)


	vbox.PackStart(CreateBBox(gtk3.GtkOrientation.VERTICAL, "Spread", 30, gtk3.GtkButtonBoxStyle.SPREAD), true, true, 0)
	vbox.PackStart(CreateBBox(gtk3.GtkOrientation.VERTICAL, "Edge", 30, gtk3.GtkButtonBoxStyle.EDGE), true, true, 5)
	vbox.PackStart(CreateBBox(gtk3.GtkOrientation.VERTICAL, "Start", 30, gtk3.GtkButtonBoxStyle.START), true, true, 5)
	vbox.PackStart(CreateBBox(gtk3.GtkOrientation.VERTICAL, "End", 30, gtk3.GtkButtonBoxStyle.END), true, true, 5)

	return window
}

func main() {
	gtk3.Init()

	w := DoButtonBox()
	w.ShowAll()

	gtk3.Main()
}
