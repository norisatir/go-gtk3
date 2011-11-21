package gtk3

/*
#include <gtk/gtk.h>
#include <glib-object.h>

static GtkWidget* to_GtkWidget(void* obj) { return GTK_WIDGET(obj); }

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"
import "github.com/norisatir/go-gtk3/gdk3"

type Allocation struct {
	object *C.GtkAllocation
}

type Requisition struct {
	Width, Height int
}

func (self Requisition) ToNative() unsafe.Pointer {
	var req C.GtkRequisition
	req.width = C.gint(self.Width)
	req.height = C.gint(self.Height)
	return unsafe.Pointer(&req)
}

// Widget-like interface must have method Widget
type WidgetLike interface {
	W() *Widget
	gobject.ObjectLike
}

func NewWidget(o unsafe.Pointer) *Widget {
	return &Widget{C.to_GtkWidget(o)}
}

// Widget Type
type Widget struct {
	object *C.GtkWidget
}

// To be object-like
func (self Widget) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Widget) Connect(s string, f interface{}, datas ...interface{}) {
	gobject.Connect(self, s, f, datas...)
}

func (self Widget) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Widget) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Widget standard funcs
func (self *Widget) Show() {
	C.gtk_widget_show(self.object)
}

func (self *Widget) ShowNow() {
	C.gtk_widget_show_now(self.object)
}

func (self *Widget) ShowAll() {
	C.gtk_widget_show_all(self.object)
}

func (self *Widget) Destroy() {
	C.gtk_widget_destroy(self.object)
}

func (self *Widget) InDestruction() bool {
	b := C.gtk_widget_in_destruction(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Hide() {
	C.gtk_widget_hide(self.object)
}

func (self *Widget) Activate() bool {
	b := C.gtk_widget_activate(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Reparent(new_parent WidgetLike) {
	C.gtk_widget_reparent(self.object, new_parent.W().object)
}

func (self *Widget) IsFocus() bool {
	b := C.gtk_widget_is_focus(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(self.object)
}

func (self *Widget) GrabDefault() {
	C.gtk_widget_grab_default(self.object)
}

func (self *Widget) SetName(name string) {
	gobject.SetProperty(self, "name", name)
}

func (self *Widget) GetName() string {
	n := gobject.GetProperty(self, "name")
	defer n.Free()
	return gobject.GoString(n.GetPtr())
}

func (self *Widget) SetSensitive(sensitive bool) {
	gobject.SetProperty(self, "sensitive", sensitive)
}

func (self *Widget) SetParent(parent WidgetLike) {
	C.gtk_widget_set_parent(self.object, parent.W().object)
}

func (self *Widget) GetTopLevel() WidgetLike {
	w := C.gtk_widget_get_toplevel(self.object)
	if tl, err := gobject.ConvertToGo(unsafe.Pointer(&w)); err == nil {
		return tl.(WidgetLike)
	}
	return nil
}

func (self *Widget) GetAncestor(widget_type gobject.GType) WidgetLike {
	w := C.gtk_widget_get_ancestor(self.object, C.GType(widget_type))
	if w == nil {
		return nil
	}
	if par, err := gobject.ConvertToGo(unsafe.Pointer(&w)); err == nil {
		return par.(WidgetLike)
	}
	return nil
}

func (self *Widget) GetPreferredSize() (minimumSize, naturalSize Requisition) {
	var min, nat C.GtkRequisition
	C.gtk_widget_get_preferred_size(self.object, &min, &nat)

	minimumSize = Requisition{int(min.width), int(min.height)}
	naturalSize = Requisition{int(nat.width), int(nat.height)}
	return
}

func (self *Widget) SetEvents(events gdk3.GdkEventMask) {
	C.gtk_widget_set_events(self.object, C.gint(events))
}

func (self *Widget) AddEvents(events gdk3.GdkEventMask) {
	C.gtk_widget_add_events(self.object, C.gint(events))
}

func (self *Widget) GetEvents() gdk3.GdkEventMask {
	return gdk3.GdkEventMask(C.gtk_widget_get_events(self.object))
}
