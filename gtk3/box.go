package gtk3

/*
#include <gtk/gtk.h>

static inline GtkBox* to_GtkBox(void* obj) { return GTK_BOX(obj); }

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"

type BoxLike interface {
	CBox() *Box
}

type Box struct {
	object *C.GtkBox
	*Container
}

func NewBox(orientation int, spacing int) *Box {
	box := &Box{}

	o := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))

	box.Container = NewContainer(unsafe.Pointer(o))
	box.object = C.to_GtkBox(unsafe.Pointer(o))

	return box
}

func NewHBox(spacing int) *Box {
	return NewBox(GtkOrientation.HORIZONTAL, spacing)
}

func NewVBox(spacing int) *Box {
	return NewBox(GtkOrientation.VERTICAL, spacing)
}

// Conversion function for gobject registration map
func newBoxFromNative(obj unsafe.Pointer) interface{} {
	var box Box
	box.object = C.to_GtkBox(obj)
	box.Container = NewContainer(unsafe.Pointer(box.object))
	return &box
}

func nativeFromBox(b interface{}) *gobject.GValue {
	box, ok := b.(Box)
	if ok {
		gv := gobject.CreateCGValue(GtkType.BOX, box.ToNative())
		return gv
	}

	return nil
}

func init() {
	// Register GtkBox to gobject
	gobject.RegisterCType(GtkType.BOX, newBoxFromNative)
	gobject.RegisterGoType(GtkType.BOX, nativeFromBox)
}

// To be object-like
func (self Box) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Box) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Box) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Box) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self Box) C() *Container {
	return self.Container
}

// Box interface
func (self *Box) PackStart(w WidgetLike, expand bool, fill bool, padding uint) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_pack_start(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(e.GetPtr())), C.guint(padding))
}

func (self *Box) PackEnd(w WidgetLike, expand bool, fill bool, padding uint) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_pack_end(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(e.GetPtr())), C.guint(padding))
}

func (self *Box) GetHomogeneous() bool {
	h := gobject.GetProperty(self, "homogeneous")
	defer h.Free()
	return gobject.GoBool(h.GetPtr())
}

func (self *Box) SetHomogeneous(homogeneous bool) {
	gobject.SetProperty(self, "homogeneous", homogeneous)
}

func (self *Box) GetSpacing() int {
	return int(C.gtk_box_get_spacing(self.object))
}

func (self *Box) SetSpacing(spacing int) {
	gobject.SetProperty(self, "spacing", spacing)
}

func (self *Box) ReorderChild(w WidgetLike, position int) {
	C.gtk_box_reorder_child(self.object, w.W().object, C.gint(position))
}

func (self *Box) QueryChildPacking(w WidgetLike) (expand bool, fill bool, padding uint, ptype int) {
	var e C.gboolean
	var f C.gboolean
	var p C.guint
	var t C.GtkPackType
	C.gtk_box_query_child_packing(self.object, w.W().object, &e, &f, &p, &t)

	return gobject.GoBool(unsafe.Pointer(&e)), gobject.GoBool(unsafe.Pointer(&f)), uint(p), int(t)
}

func (self *Box) SetChildPacking(w WidgetLike, expand bool, fill bool, padding int, ptype int) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_set_child_packing(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(f.GetPtr())),
		C.guint(padding), C.GtkPackType(ptype))
}
