package gtk3

/*
#include <gtk/gtk.h>

static inline GtkButtonBox* to_GtkButtonBox(void* obj) { return GTK_BUTTON_BOX(obj); }

static void _gtk_button_box_set_child_non_homogeneous(GtkButtonBox *widget, GtkWidget* child,
														gboolean non_homogeneous) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_button_box_set_child_non_homogeneous(widget, child, non_homogeneous);
#endif
}

static gboolean _gtk_button_box_get_child_non_homogeneous(GtkButtonBox *widget, GtkWidget* child) {
#if GTK_CHECK_VERSION(3,2,0)
	return gtk_button_box_get_child_non_homogeneous(widget, child);
#else
	return NULL;
#endif
}

*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type ButtonBox struct {
	object *C.GtkButtonBox
	*Box
}

func NewButtonBox(orientation int) *ButtonBox {
	bb := &ButtonBox{}
	o := C.gtk_button_box_new(C.GtkOrientation(orientation))
	bb.object = C.to_GtkButtonBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(bb) {
		gobject.RefSink(bb)
	}
	bb.Box = newBoxFromNative(unsafe.Pointer(o)).(*Box)
	buttonBoxFinalizer(bb)

	return bb
}

// Clear ButtonBox structure when it goes out of reach
func buttonBoxFinalizer(bb *ButtonBox) {
	runtime.SetFinalizer(bb, func(bb *ButtonBox) { gobject.Unref(bb) })
}

// Conversion function for gobject registration map
func newButtonBoxFromNative(obj unsafe.Pointer) interface{} {
	bb := &ButtonBox{}
	bb.object = C.to_GtkButtonBox(obj)

	if gobject.IsObjectFloating(bb) {
		gobject.RefSink(bb)
	}
	bb.Box = newBoxFromNative(obj).(*Box)
	buttonBoxFinalizer(bb)
	
	return bb
}

func nativeFromButtonBox(b interface{}) *gobject.GValue {
	bbox, ok := b.(*ButtonBox)
	if ok {
		gv := gobject.CreateCGValue(GtkType.BUTTON_BOX, bbox.ToNative())
		return gv
	}
	return nil
}

func init() {
	// REgister GtkButtonBox to gobject system
	gobject.RegisterCType(GtkType.BUTTON_BOX, newButtonBoxFromNative)
	gobject.RegisterGoType(GtkType.BUTTON_BOX, nativeFromButtonBox)
}

// To be object-like
func (self ButtonBox) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ButtonBox) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ButtonBox) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ButtonBox) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Box-like
func (self ButtonBox) CBox() *Box {
	return self.Box
}

// ButtonBox Interface
func (self *ButtonBox) GetLayout() int {
	l := C.gtk_button_box_get_layout(self.object)
	return int(l)
}

func (self *ButtonBox) GetChildSecondary(w WidgetLike) bool {
	cb := C.gtk_button_box_get_child_secondary(self.object, w.W().object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ButtonBox) GetChildNonHomogeneous(w WidgetLike) bool {
	cb := C._gtk_button_box_get_child_non_homogeneous(self.object, w.W().object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ButtonBox) SetLayout(layoutStyle int) {
	C.gtk_button_box_set_layout(self.object, C.GtkButtonBoxStyle(layoutStyle))
}

func (self *ButtonBox) SetChildSecondary(w WidgetLike, isSecondary bool) {
	cb := gobject.GBool(isSecondary)
	defer cb.Free()
	C.gtk_button_box_set_child_secondary(self.object, w.W().object, *((*C.gboolean)(cb.GetPtr())))
}

func (self *ButtonBox) SetChildNonHomogeneous(w WidgetLike, nonHomogeneous bool) {
	cb := gobject.GBool(nonHomogeneous)
	defer cb.Free()
	C._gtk_button_box_set_child_non_homogeneous(self.object, w.W().object, *((*C.gboolean)(cb.GetPtr())))
}
