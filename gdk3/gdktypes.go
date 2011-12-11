package gdk3

/*
#include <gdk/gdk.h>

*/
import "C"
import "unsafe"
import g "github.com/norisatir/go-gtk3/gobject"

var GdkType gdkTypes

type gdkTypes struct {
	DEVICE    g.GType
	RECTANGLE g.GType
	EVENT     g.GType
	DISPLAY	  g.GType
	SCREEN    g.GType
}

func init() {
	GdkType.DEVICE = g.GType(C.gdk_device_get_type())
	GdkType.RECTANGLE = g.GType(C.gdk_rectangle_get_type())
	GdkType.EVENT = g.GType(C.gdk_event_get_type())
	GdkType.DISPLAY = g.GType(C.gdk_display_get_type())
	GdkType.SCREEN = g.GType(C.gdk_screen_get_type())

	// Register Rectangle
	g.RegisterCType(GdkType.RECTANGLE, newRectangleFromNative)
}

// Basic Types

// Rectangle
type Rectangle struct {
	X, Y          int
	Width, Height int
}

func CreateRectangle(obj unsafe.Pointer) Rectangle {
	cr := *((*C.GdkRectangle)(obj))

	r := Rectangle{}
	r.X = int(cr.x)
	r.Y = int(cr.y)
	r.Width = int(cr.width)
	r.Height = int(cr.height)

	return r
}

//Conversion func
func newRectangleFromNative(obj unsafe.Pointer) interface{} {
	r := CreateRectangle(obj)
	return r
}

func NativeRectangle(rec Rectangle) unsafe.Pointer {
	var cr C.GdkRectangle
	cr.x = C.int(rec.X)
	cr.y = C.int(rec.Y)
	cr.width = C.int(rec.Width)
	cr.height = C.int(rec.Height)

	return unsafe.Pointer(&cr)
}

func RectangleIntersect(r1, r2 Rectangle) Rectangle {
	cr1 := NativeRectangle(r1)
	cr2 := NativeRectangle(r2)
	var cr3 C.GdkRectangle
	C.gdk_rectangle_intersect((*C.GdkRectangle)(cr1), (*C.GdkRectangle)(cr2), &cr3)

	return CreateRectangle(unsafe.Pointer(&cr3))
}

func RectangleUnion(r1, r2 Rectangle) Rectangle {
	cr1 := NativeRectangle(r1)
	cr2 := NativeRectangle(r2)
	var cr3 C.GdkRectangle
	C.gdk_rectangle_union((*C.GdkRectangle)(cr1), (*C.GdkRectangle)(cr2), &cr3)
	return CreateRectangle(unsafe.Pointer(&cr3))
}

// End Rectangle
