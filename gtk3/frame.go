package gtk3

/*
#include <gtk/gtk.h>

static inline GtkFrame* to_GtkFrame(void* obj) { return GTK_FRAME(obj); }

*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type Frame struct {
	object *C.GtkFrame
	*Container
}

func NewFrame(label string) *Frame {
	f := &Frame{}

	s := gobject.GString(label)
	defer s.Free()

	o := C.gtk_frame_new((*C.gchar)(s.GetPtr()))
	f.object = C.to_GtkFrame(unsafe.Pointer(o))

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	}
	f.Container = NewContainer(unsafe.Pointer(o))
	frameFinalizer(f)

	return f
}

// Clear Frame struct when it goes out of reach
func frameFinalizer(f *Frame) {
	runtime.SetFinalizer(f, func(f *Frame) { gobject.Unref(f) })
}

// Conversion function for gobject registration map
func newFrameFromNative(obj unsafe.Pointer) interface{} {
	f := &Frame{}
	f.object = C.to_GtkFrame(obj)

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	} else {
		gobject.Ref(f)
	}
	f.Container = NewContainer(unsafe.Pointer(f.object))
	frameFinalizer(f)

	return f
}

func nativeFromFrame(frame interface{}) *gobject.GValue {
	f, ok := frame.(*Frame)
	if ok {
		gv := gobject.CreateCGValue(GtkType.FRAME, f.ToNative())
		return gv
	}

	return nil
}

func init() {
	// Register GtkFrame
	gobject.RegisterCType(GtkType.FRAME, newFrameFromNative)
	gobject.RegisterGoType(GtkType.FRAME, nativeFromFrame)
}

// To be object like
func (self Frame) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Frame) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Frame) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Frame) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-lie
func (self Frame) C() *Container {
	return self.Container
}

// Frame interface
func (self *Frame) SetLabel(label string) {
	s := gobject.GString(label)
	defer s.Free()
	C.gtk_frame_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Frame) SetLabelWidget(w WidgetLike) {
	C.gtk_frame_set_label_widget(self.object, w.W().object)
}

func (self *Frame) SetLabelAlign(xalign, yalign float32) {
	C.gtk_frame_set_label_align(self.object, C.gfloat(xalign), C.gfloat(yalign))
}

func (self *Frame) SetShadowType(typ GtkShadowType) {
	C.gtk_frame_set_shadow_type(self.object, C.GtkShadowType(typ))
}

func (self *Frame) GetLabel() string {
	l := C.gtk_frame_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(l))
}

func (self *Frame) GetLabelAlign() (xalign, yalign float32) {
	var cax C.gfloat
	var cay C.gfloat
	C.gtk_frame_get_label_align(self.object, &cax, &cay)
	return float32(cax), float32(cay)
}

func (self *Frame) GetLabelWidget() WidgetLike {
	cw := C.gtk_frame_get_label_widget(self.object)
	w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
	if err != nil {
		return w.(WidgetLike)
	}
	return nil
}

func (self *Frame) GetShadowType() GtkShadowType {
	return GtkShadowType(C.gtk_frame_get_shadow_type(self.object))
}
