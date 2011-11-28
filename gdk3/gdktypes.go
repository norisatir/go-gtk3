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
	X, Y int
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

func nativeRectangle(rec Rectangle) *C.GdkRectangle {
	var cr C.GdkRectangle
	cr.x = C.int(rec.X)
	cr.y = C.int(rec.Y)
	cr.width = C.int(rec.Width)
	cr.height = C.int(rec.Height)

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
