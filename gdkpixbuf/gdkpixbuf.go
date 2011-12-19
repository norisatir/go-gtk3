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
import "errors"
import "github.com/norisatir/go-gtk3/gobject"

var G_TYPE_PIXBUF gobject.GType

// GdkInterp type {{{

var GdkInterp gdkInterp

type gdkInterp struct {
	NEAREST  int
	TILES    int
	BILINEAR int
	HYPER    int
}
// End GdkInterp }}}

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

// Funcs

func NewFromFile(filename string) (*Pixbuf, error) {
	fn := gobject.GString(filename)
	defer fn.Free()

	pb := C.gdk_pixbuf_new_from_file((*C.char)(fn.GetPtr()), nil)

	if pb == nil {
		return nil, errors.New("Error occured")
	}

	return newPixbufFromNative(unsafe.Pointer(pb)).(*Pixbuf), nil
}

func ScaleSimple(src *Pixbuf, destWidth, destHeight, gdk_Interp int) *Pixbuf {
	pScaled := C.gdk_pixbuf_scale_simple(src.object, C.int(destWidth), C.int(destHeight), C.GdkInterpType(gdk_Interp))

	if pScaled == nil {
		return nil
	}

	if pix, err := gobject.ConvertToGo(unsafe.Pointer(pScaled)); err == nil {
		return pix.(*Pixbuf)
	}
	return nil
}

// GDK_PIXBUF init funcs {{{
func init() {
	// Initialize G_TYPE_PIXBUF
	G_TYPE_PIXBUF = gobject.GType(C.gdk_pixbuf_get_type())

	// Register Pixbuf type
	gobject.RegisterCType(G_TYPE_PIXBUF, newPixbufFromNative)
	gobject.RegisterGoType(G_TYPE_PIXBUF, nativeFromPixbuf)

	// Initialize GdkInterp
	GdkInterp.NEAREST = 0
	GdkInterp.TILES = 1
	GdkInterp.BILINEAR = 2
	GdkInterp.HYPER = 3
} // }}}
