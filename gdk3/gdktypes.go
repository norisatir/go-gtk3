package gdk3

/*
#include <gdk/gdk.h>

*/
import "C"
import "unsafe"
import g "github.com/norisatir/go-gtk3/gobject"

var GdkType gdkTypes

type gdkTypes struct {
	DEVICE g.GType
	RECTANGLE g.GType
	EVENT g.GType
}

func init() {
	GdkType = gdkTypes{}
	GdkType.DEVICE = g.GType(C.gdk_device_get_type())
	GdkType.RECTANGLE = g.GType(C.gdk_rectangle_get_type())
	GdkType.EVENT = g.GType(C.gdk_event_get_type())

	// Register Rectangle
	g.RegisterCType(GdkType.RECTANGLE, newRectangleFromNative)
}


// Basic Types

// Rectangle
type Rectangle struct {
	x, y int
	width, height int
}

func CreateRectangle(obj unsafe.Pointer) Rectangle {
	cr := *((*C.GdkRectangle)(obj))

	r := Rectangle{}
	r.x = int(cr.x)
	r.y = int(cr.y)
	r.width = int(cr.width)
	r.height = int(cr.height)

	return r
}

//Conversion func
func newRectangleFromNative(obj unsafe.Pointer) interface{} {
	r := CreateRectangle(obj)
	return r
}

func nativeRectangle(rec Rectangle) *C.GdkRectangle {
	var cr C.GdkRectangle
	cr.x = C.int(rec.x)
	cr.y = C.int(rec.y)
	cr.width = C.int(rec.width)
	cr.height = C.int(rec.height)

	return &cr
}

func RectangleIntersect(r1, r2 Rectangle) Rectangle {
	cr1 := nativeRectangle(r1)
	cr2 := nativeRectangle(r2)
	var cr3 C.GdkRectangle
	C.gdk_rectangle_intersect(cr1, cr2, &cr3)

	return CreateRectangle(unsafe.Pointer(&cr3))
}

func RectangleUnion(r1, r2 Rectangle) Rectangle {
	cr1 := nativeRectangle(r1)
	cr2 := nativeRectangle(r2)
	var cr3 C.GdkRectangle
	C.gdk_rectangle_union(cr1, cr2, &cr3)
	return CreateRectangle(unsafe.Pointer(&cr3))
}

// End Rectangle
