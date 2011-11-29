package gtk3

/*
#include <gtk/gtk.h>

static inline GtkSeparator* to_GtkSeparator(void* obj) { return GTK_SEPARATOR(obj); }

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"

type Separator struct {
	object *C.GtkSeparator
	*Widget
}

func NewSeparator(orientation int) *Separator {
	sep := &Separator{}

	o := C.gtk_separator_new(C.GtkOrientation(orientation))

	sep.Widget = NewWidget(unsafe.Pointer(o))
	sep.object = C.to_GtkSeparator(unsafe.Pointer(o))

	return sep
}

func NewHSeparator() *Separator {
	return NewSeparator(GtkOrientation.HORIZONTAL)
}

func NewVSeparator() *Separator {
	return NewSeparator(GtkOrientation.VERTICAL)
}


// Conversion function for gobject registration map
func newSeparatorFromNative(obj unsafe.Pointer) interface{} {
	sep := &Separator{}
	sep.object = C.to_GtkSeparator(obj)
	sep.Widget = NewWidget(obj)

	return sep
}

func nativeFromSeparator(s interface{}) *gobject.GValue {
	sep, ok := s.(*Separator)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SEPARATOR, sep.ToNative())
		return gv
	}

	return nil
}

func init() {
	// Register GtkSeparator
	gobject.RegisterCType(GtkType.SEPARATOR, newSeparatorFromNative)
	gobject.RegisterGoType(GtkType.SEPARATOR, nativeFromSeparator)
}

// To be widget-like
func (self Separator) W() *Widget {
	return self.Widget
}
