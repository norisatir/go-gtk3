GTK3 (WITH GOBJECT AND GDK) FOR GO
===================================

THIS IS:
--------

	Hopefully GTK3 Bindings complete with GObject and GDK3

	Currently implemented is GObject and Conversion between 
	C and Go types (mainly through GValues).

	Callbacks are implemented with closures, channels and gorutines
	(see gobject/callback.go)

	As of now there are almost no widgets (Window, Button), but they can
	be easily implemented (see button.go, window.go in gtk3 folder)

	Everything else soon...

INSTALL:
--------

	git clone https://github.com/norisatir/go-gtk3
	cd go-gtk3
	gomake install

or:
	goinstall github.com/norisatir/go-gtk3/gtk3


For demo:
	gomake demo
	cd demo_app
	./demo


LICENSE:
--------

	Under same terms and conditions as the Go language,
	BSD style license

AUTHOR
------

	Matej Knific
