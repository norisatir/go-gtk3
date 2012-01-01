package spinner

import "github.com/norisatir/go-gtk3/gtk3"

var window *gtk3.Dialog = nil

func DoSpinner(w gtk3.WidgetLike) gtk3.WidgetLike {
	if window == nil {
		window = gtk3.NewDialogWithButtons("GtkSpinner", w.(*gtk3.Window), 0,
			gtk3.B{{gtk3.GtkStock.CLOSE, gtk3.GtkResponse.NONE}})

		window.SetResizable(false)
		window.Connect("response", func() { window.Destroy() })
		window.Connect("destroy", func() { window = nil })

		contentArea := window.GetContentArea().(*gtk3.Box)

		vbox := gtk3.NewVBox(5)
		contentArea.PackStart(vbox, true, true, 0)
		vbox.SetBorderWidth(5)

		// Sensitive
		hbox := gtk3.NewHBox(5)
		spinnerS := gtk3.NewSpinner()
		hbox.Add(spinnerS)
		hbox.Add(gtk3.NewEntry())
		vbox.Add(hbox)

		// Disabled
		hbox = gtk3.NewHBox(5)
		spinnerU := gtk3.NewSpinner()
		hbox.Add(spinnerU)
		hbox.Add(gtk3.NewEntry())
		vbox.Add(hbox)
		hbox.SetSensitive(false)

		button := gtk3.NewButtonFromStock(gtk3.GtkStock.MEDIA_PLAY)

		button.Connect("clicked", func() {
			spinnerS.Start()
			spinnerU.Start()
		})
		vbox.Add(button)

		button = gtk3.NewButtonFromStock(gtk3.GtkStock.MEDIA_STOP)
		button.Connect("clicked", func() {
			spinnerS.Stop()
			spinnerU.Stop()
		})
		vbox.Add(button)

		spinnerS.Start()
		spinnerU.Start()
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
