package infobar

import "github.com/norisatir/go-gtk3/gtk3"

var window *gtk3.Window = nil

func DoInfoBar(w gtk3.WidgetLike) gtk3.WidgetLike {
    if window == nil {
        window = gtk3.NewWindow(gtk3.GtkWindowType.TOPLEVEL)
        window.SetScreen(w.W().GetScreen())

        window.SetTitle("Info Bars")
        window.Connect("destroy", func() { window.Destroy(); window = nil})
        window.SetBorderWidth(8)

        vbox := gtk3.NewVBox(0)
        window.Add(vbox)

        bar := gtk3.NewInfoBar()
        vbox.PackStart(bar, false, false, 0)
        bar.SetMessageType(gtk3.GtkMessage.INFO)
        label := gtk3.NewLabel("This is an info bar with message type GTK_MESSAGE_INFO")
        (bar.GetContentArea()).(*gtk3.Box).PackStart(label, false, false, 0)

        bar = gtk3.NewInfoBar()
        vbox.PackStart(bar, false, false, 0)
        bar.SetMessageType(gtk3.GtkMessage.WARNING)
        label = gtk3.NewLabel("This is an info bar with message type GTK_MESSAGE_WARNING")
        (bar.GetContentArea()).(*gtk3.Box).PackStart(label, false, false, 0)

        bar = gtk3.NewInfoBarWithButtons(gtk3.B{{gtk3.GtkStock.OK, gtk3.GtkResponse.OK}})
        bar.Connect("response", func(infoBar *gtk3.InfoBar, responseId int, data...interface{}) {
            dialog := gtk3.NewMessageDialog(window, gtk3.GtkDialogFlags.MODAL | gtk3.GtkDialogFlags.DESTROY_WITH_PARENT,
                        gtk3.GtkMessage.INFO,
                        gtk3.GtkButtons.OK,
                        "You Clicked a button on info bar")

            dialog.FormatSecondaryText("Your response has id %d", responseId)
            dialog.Run()
            dialog.Destroy()
        })
        vbox.PackStart(bar, false, false, 0)
        bar.SetMessageType(gtk3.GtkMessage.QUESTION)
        label = gtk3.NewLabel("This is an info bar with message type GTK_MESSAGE_QUESTION")
        (bar.GetContentArea()).(*gtk3.Box).PackStart(label, false, false, 0)

        bar = gtk3.NewInfoBar()
        vbox.PackStart(bar, false, false, 0)
        bar.SetMessageType(gtk3.GtkMessage.ERROR)
        label = gtk3.NewLabel("This is an info bar with message type GTK_MESSAGE_ERROR")
        (bar.GetContentArea()).(*gtk3.Box).PackStart(label, false, false, 0)

        bar = gtk3.NewInfoBar()
        vbox.PackStart(bar, false, false, 0)
        bar.SetMessageType(gtk3.GtkMessage.OTHER)
        label = gtk3.NewLabel("This is an info bar with message type GTK_MESSAGE_OTHER")
        (bar.GetContentArea()).(*gtk3.Box).PackStart(label, false, false, 0)

        frame := gtk3.NewFrame("Info bars")
        vbox.PackStart(frame, false, false, 8)

        vbox2 := gtk3.NewVBox(8)
        vbox.SetBorderWidth(8)
        frame.Add(vbox2)

        // Standard message dialog
        label = gtk3.NewLabel("An example of different info bars")
        vbox2.PackStart(label, false, false, 0)
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
