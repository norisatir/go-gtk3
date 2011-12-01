package gtk3

/*
#include <gtk/gtk.h>

static inline GtkToggleButton* to_GtkToggleButton(void* obj) {
	return GTK_TOGGLE_BUTTON(obj);
}

*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type ToggleButton struct {
	object *C.GtkToggleButton
	*Button
}

// Create and return new toggle button
func NewToggleButton() *ToggleButton {
	b := &ToggleButton{}
	o := C.gtk_toggle_button_new()
	b.object = C.to_GtkToggleButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	toggleButtonFinalizer(b)

	return b
}

func NewToggleButtonWithLabel(label string) *ToggleButton {
	b := &ToggleButton{}
	l := gobject.GString(label)
	defer l.Free()

	o := C.gtk_toggle_button_new_with_label((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkToggleButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	toggleButtonFinalizer(b)

	return b
}

// Clear ToggleButton when it goes out of reach
func toggleButtonFinalizer(b *ToggleButton) {
	runtime.SetFinalizer(b, func(b *ToggleButton) { gobject.Unref(b) })
}

// Conversion function for gobject registration map
func newToggleButtonFromNative(obj unsafe.Pointer) interface{} {
	b := &ToggleButton{}
	b.object = C.to_GtkToggleButton(obj)

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Button = newButtonFromNative(obj).(*Button)
	toggleButtonFinalizer(b)

	return b
}

func nativeFromToggleButton(b interface{}) *gobject.GValue {
	if but, ok := b.(*ToggleButton); ok {
		gv := gobject.CreateCGValue(GtkType.TOGGLE_BUTTON, but.ToNative())
		return gv
	}
	return nil
}

func init() {
	//Register GtkToggleButton to gobject type system (In Go)
	gobject.RegisterCType(GtkType.TOGGLE_BUTTON, newToggleButtonFromNative)
	gobject.RegisterGoType(GtkType.TOGGLE_BUTTON, nativeFromToggleButton)
}

// To be object-like
func (self ToggleButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ToggleButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ToggleButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ToggleButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be button-like
func (self ToggleButton) B() *Button {
	return self.Button
}

// ToggleButton interface

func (self *ToggleButton) SetMode(drawIndicator bool) {
	b := gobject.GBool(drawIndicator)
	defer b.Free()
	C.gtk_toggle_button_set_mode(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ToggleButton) GetMode() bool {
	cb := C.gtk_toggle_button_get_mode(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) Toggled() {
	C.gtk_toggle_button_toggled(self.object)
}

func (self *ToggleButton) GetActive() bool {
	cb := C.gtk_toggle_button_get_active(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) SetActive(isActive bool) {
	b := gobject.GBool(isActive)
	defer b.Free()
	C.gtk_toggle_button_set_active(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ToggleButton) GetInconsistent() bool {
	cb := C.gtk_toggle_button_get_inconsistent(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) SetInconsistent(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_toggle_button_set_inconsistent(self.object, *((*C.gboolean)(b.GetPtr())))
}
