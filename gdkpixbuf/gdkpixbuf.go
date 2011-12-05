package gdkpixbuf

/*
#include <gdk-pixbuf/gdk-pixbuf.h>

// to_GdkPixbuf
static inline GdkPixbuf* to_GdkPixbuf(void* obj) { return GDK_PIXBUF(obj); }

*/
// #cgo pkg-config: gdk-pixbuf-2.0
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

var G_TYPE_PIXBUF gobject.GType

// GdkPixbuf {{{

type Pixbuf struct {
	object *C.GdkPixbuf
}

// Clear Pixbuf struct when it goes out of reach
func pixbufFinalizer(pb *Pixbuf) {
	runtime.SetFinalizer(pb, func(pb *Pixbuf) { gobject.Unref(pb) })
}

// Conversion functions
func newPixbufFromNative(obj unsafe.Pointer) interface{} {
	pb := &Pixbuf{}
	pb.object = C.to_GdkPixbuf(obj)

	if gobject.IsObjectFloating(pb) {
		gobject.RefSink(pb)
	}
	pixbufFinalizer(pb)

	return pb
}

func nativeFromPixbuf(pb interface{}) *gobject.GValue {
	pixbuff, ok := pb.(*Pixbuf)
	if ok {
		gv := gobject.CreateCGValue(G_TYPE_PIXBUF, pixbuff.ToNative())
		return gv
	}
	return nil
}

// To be object-like

func (self Pixbuf) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Pixbuf) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Pixbuf) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Pixbuf) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// END GdkPixbuf
////////////////////////////// }}}

// GDK_PIXBUF init funcs {{{
func init() {
	// Initialize G_TYPE_PIXBUF
	G_TYPE_PIXBUF = gobject.GType(C.gdk_pixbuf_get_type())

	// Register Pixbuf type
	gobject.RegisterCType(G_TYPE_PIXBUF, newPixbufFromNative)
	gobject.RegisterGoType(G_TYPE_PIXBUF, nativeFromPixbuf)
} // }}}
