include $(GOROOT)/src/Make.inc

GC=${O}g -I../gobject/_obj -I ../glib/_obj -I../gdk3/_obj -I../gdkpixbuf/_obj -I../pango/_obj -I$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

TARG     = github.com/norisatir/go-gtk3/gdk3
CGOFILES = gdk.go gdktypes.go gdkevents.go

include $(GOROOT)/src/Make.pkg
