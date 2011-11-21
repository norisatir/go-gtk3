package gtk3

/*
#include <gtk/gtk.h>
#include <glib-object.h>

static GtkContainer* to_GtkContainer(void *obj) { return GTK_CONTAINER(obj); }

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"

// New Container from widget
func NewContainer(o unsafe.Pointer) *Container {
	var con Container
	con.Widget = NewWidget(o)
	con.object = C.to_GtkContainer(o)
	return &con
}
// Container-like interface must have method Container
type ContainerLike interface {
	C() *Container
	WidgetLike
}

// Container Type
type Container struct {
	object *C.GtkContainer
	*Widget
}

// To-be objectlike
func (self Container) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Container) Connect(s string, f interface{}, datas ...interface{}) {
	gobject.Connect(self, s, f, datas...)
}

func (self Container) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Container) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Widget-like
func (self Container) W() *Widget {
	return self.Widget
}

// interfaces
func (self Container) Add(w WidgetLike) {
	C.gtk_container_add(self.object, w.W().object)
}

func (self Container) Remove(w WidgetLike) {
	C.gtk_container_remove(self.object, w.W().object)
}

func (self Container) GetFocusChild() WidgetLike {
	w := C.gtk_container_get_focus_child(self.object)
	i, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err != nil {
		return i.(WidgetLike)
	}
	return nil
}
