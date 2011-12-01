package gtk3

/*
#include <gtk/gtk.h>

static GtkButton* to_GtkButton(void* obj) {
	return GTK_BUTTON(obj);
}

*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type ButtonLike interface {
	B() *Button
}

type Button struct {
	object *C.GtkButton
	*Container
}

//Create and return new button Structure
func NewButton() *Button {
	b := &Button{}
	o := C.gtk_button_new()
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Container = NewContainer(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

func NewButtonWithLabel(label string) *Button {
	b := &Button{}
	l := gobject.GString(label)
	defer l.Free()

	o := C.gtk_button_new_with_label((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Container = NewContainer(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

func NewButtonFromStock(stockId string) *Button {
	b := NewButtonWithLabel(stockId)
	b.SetUseStock(true)
	return b
}

func NewButtonWithMnemonic(label string) *Button {
	b := &Button{}
	l := gobject.GString(label)
	defer l.Free()
	o := C.gtk_button_new_with_mnemonic((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Container = NewContainer(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

// Unref button, when Button structer goes out of reach
func buttonFinalizer(b *Button) {
	runtime.SetFinalizer(b, func(b *Button) { gobject.Unref(b) })
}

// Conversion function for gobject registration map
func newButtonFromNative(obj unsafe.Pointer) interface{} {
	b := &Button{}
	b.object = C.to_GtkButton(obj)
	
	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Container = NewContainer(obj)
	buttonFinalizer(b)

	return b
}

func nativeFromButton(b interface{}) *gobject.GValue {
	if but, ok := b.(*Button); ok {
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

func (self Button) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
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
	defer s.Free()
	C.gtk_button_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Button) GetLabel() string {
	l := C.gtk_button_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(l))
}

func (self *Button) GetUseStock() bool {
	b := C.gtk_button_get_use_stock(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetUseStock(useStock bool) {
	gobject.SetProperty(self, "use-stock", useStock)
}

func (self *Button) GetUseUnderline() bool {
	b := C.gtk_button_get_use_underline(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetUseUnderline(underline bool) {
	gobject.SetProperty(self, "use-underline", underline)
}

func (self *Button) SetFocusOnClick(focusOnClick bool) {
	gobject.SetProperty(self, "focus-on-click", focusOnClick)
}

func (self *Button) GetFocusOnClick() bool {
	b := C.gtk_button_get_focus_on_click(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetAlignment(xalign, yalign float32) {
	C.gtk_button_set_alignment(self.object, C.gfloat(xalign), C.gfloat(yalign))
}

func (self *Button) GetAlignment() (float32, float32) {
	var cx C.gfloat
	var cy C.gfloat
	C.gtk_button_get_alignment(self.object, &cx, &cy)
	return float32(cx), float32(cy)
}

func (self *Button) SetImage(w WidgetLike) {
	C.gtk_button_set_image(self.object, w.W().object)
}

func (self *Button) GetImage() WidgetLike {
	cw := C.gtk_button_get_image(self.object)
	widget, err := gobject.ConvertToGo(unsafe.Pointer(&cw))
	if err != nil {
		return widget.(WidgetLike)
	}

	return nil
}

func (self *Button) SetImagePosition(position int) {
	C.gtk_button_set_image_position(self.object, C.GtkPositionType(position))
}

func (self *Button) GetImagePosition() int {
	return int(C.gtk_button_get_image_position(self.object))
}

func (self *Button) Clicked() {
	C.gtk_button_clicked(self.object)
}
//TODO: gtk_button_get_event_window
