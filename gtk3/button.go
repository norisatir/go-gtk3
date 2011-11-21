package gtk3

/*
#include <gtk/gtk.h>

static GtkButton* to_GtkButton(void* obj) {
	return GTK_BUTTON(obj);
}

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"

type Button struct {
	object *C.GtkButton
	*Container
}

//Create and return new button Structure
func NewButton() *Button {
	b := &Button{}

	o := C.gtk_button_new()

	b.Container = NewContainer(unsafe.Pointer(o))

	b.object = C.to_GtkButton(unsafe.Pointer(o))
	return b
}

func NewButtonWithLabel(label string) *Button {
	b := &Button{}
	l := gobject.GString(label)
	o := C.gtk_button_new_with_label((*C.gchar)(l.GetPtr()))
	b.Container = NewContainer(unsafe.Pointer(o))
	b.object = C.to_GtkButton(unsafe.Pointer(o))
	return b
}

// Conversion function for gobject registration map
func newButtonFromNative(obj unsafe.Pointer) interface{} {
	var button Button
	button.object = C.to_GtkButton(obj)
	button.Container = NewContainer(unsafe.Pointer(obj))
	return &button
}

func nativeFromButton(b interface{}) *gobject.GValue {
	if but, ok := b.(Button); ok {
		gv := gobject.CreateCGValue(GtkType.BUTTON, but.ToNative())
		return gv
	}
	return nil
}

func init() {
	// Register GtkButton to gobject type system (in Go)
	gobject.RegisterCType(GtkType.BUTTON, newButtonFromNative)
	gobject.RegisterGoType(GtkType.BUTTON, nativeFromButton)
}

// To be object-like
func (self Button) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Button) Connect(s string, f interface{}, datas ...interface{}) {
	gobject.Connect(self, s, f, datas...)
}

func (self Button) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Button) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container like
func (self Button) C() *Container {
	return self.Container
}

// Button interface
func (self *Button) SetLabel(label string) {
	s := gobject.GString(label)
	//defer s.Free()
	C.gtk_button_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Button) GetLabel() string {
	l := C.gtk_button_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(l))
}
