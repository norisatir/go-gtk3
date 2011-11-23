include $(GOROOT)/src/Make.inc

GC=${O}g -I../gobject/_obj -I../gdk3/_obj -I$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

all:
	cd gobject && gomake
	cd gdk3 && gomake
	cd gtk3 && gomake

install:
	cd gobject && gomake install
	cd gdk3 && gomake install
	cd gtk3 && gomake install

clean:
	cd gobject && gomake clean
	cd gdk3 && gomake clean
	cd gtk3 && gomake clean
	cd demo_apps && gomake clean

fmt_all:
	gofmt -w ./gobject/*.go
	gofmt -w ./gdk3/*.go
	gofmt -w ./gtk3/*.go

demo: install
	cd demo_apps && gomake
