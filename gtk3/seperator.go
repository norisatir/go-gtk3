package gtk3

/*
#include <gtk/gtk.h>

static inline GtkSeparator* to_GtkSeparator(void* obj) { return GTK_SEPARATOR(obj); }

*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type Separator struct {
	object *C.GtkSeparator
	*Widget
}

func NewSeparator(orientation int) *Separator {
	sep := &Separator{}
	o := C.gtk_separator_new(C.GtkOrientation(orientation))
	sep.object = C.to_GtkSeparator(unsafe.Pointer(o))

	if gobject.IsObjectFloating(sep) {
		gobject.RefSink(sep)
	}
	sep.Widget = NewWidget(unsafe.Pointer(o))
	separatorFinalizer(sep)

	return sep
}

func NewHSeparator() *Separator {
	return NewSeparator(GtkOrientation.HORIZONTAL)
}

func NewVSeparator() *Separator {
	return NewSeparator(GtkOrientation.VERTICAL)
}

// Clear Separator struct when it goes out of reach
func separatorFinalizer(s *Separator) {
	runtime.SetFinalizer(s, func(s *Separator) { gobject.Unref(s) })
}

// Conversion function for gobject registration map
func newSeparatorFromNative(obj unsafe.Pointer) interface{} {
	sep := &Separator{}
	sep.object = C.to_GtkSeparator(obj)

	if gobject.IsObjectFloating(sep) {
		gobject.RefSink(sep)
	} else {
		gobject.Ref(sep)
	}
	sep.Widget = NewWidget(obj)
	separatorFinalizer(sep)

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

// To be object like
func (self Separator) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Separator) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Separator) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Separator) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
// To be widget-like
func (self Separator) W() *Widget {
	return self.Widget
}
