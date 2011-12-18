include $(GOROOT)/src/Make.inc

GC=${O}g -I../gobject/_obj -I ../glib/_obj -I../gdk3/_obj -I../gdkpixbuf/_obj -I../pango/_obj -I$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

all:
	cd gobject && gomake
	cd glib && gomake
	cd gdk3 && gomake
	cd gdkpixbuf && gomake
	cd pango && gomake
	cd gtk3 && gomake

install:
	cd gobject && gomake install
	cd glib && gomake install
	cd gdk3 && gomake install
	cd gdkpixbuf && gomake install
	cd pango && gomake install
	cd gtk3 && gomake install

clean:
	cd gobject && gomake clean
	cd glib && gomake clean
	cd gdk3 && gomake clean
	cd gdkpixbuf && gomake clean
	cd pango && gomake clean
	cd gtk3 && gomake clean
	cd demo_apps && gomake clean

fmt_all:
	gofmt -w ./gobject/*.go
	gofmt -w ./glib/*.go
	gofmt -w ./gdk3/*.go
	gofmt -w ./gdkpixbuf/*.go
	gofmt -w ./pango/*.go
	gofmt -w ./gtk3/*.go

demo: install
	cd demo_apps && gomake
