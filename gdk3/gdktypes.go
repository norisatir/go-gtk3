package gdk3

/*
#include <gdk/gdk.h>
#include <stdlib.h>

static GdkRGBA* _new_rgba(void) {
	GdkRGBA* r = (GdkRGBA*)malloc(sizeof(GdkRGBA));
	return r;
}

*/
import "C"
import "unsafe"
import g "github.com/norisatir/go-gtk3/gobject"

var GdkType gdkTypes

type gdkTypes struct {
	DEVICE    g.GType
	RECTANGLE g.GType
	RGBA      g.GType
	COLOR     g.GType
	EVENT     g.GType
	DISPLAY   g.GType
	SCREEN    g.GType
	WINDOW    g.GType
}

// Basic Types

// Rectangle {{{
type Rectangle struct {
	X, Y          int
	Width, Height int
}

//Conversion func
func newRectangleFromNative(obj unsafe.Pointer) interface{} {
	cr := *((*C.GdkRectangle)(obj))
	r := Rectangle{}
	r.X = int(cr.x)
	r.Y = int(cr.y)
	r.Width = int(cr.width)
	r.Height = int(cr.height)

	return r
}

func nativeFromRectangle(r interface{}) *g.GValue {
	if rec, ok := r.(Rectangle); ok {
		gv := g.CreateCGValue(GdkType.RECTANGLE, rec.ToNative())
		return gv
	}
	return nil
}

// To be BoxedLike
func (self Rectangle) GetBoxType() g.GType {
	return GdkType.RECTANGLE
}

// Rectangle interface

func (self Rectangle) ToNative() unsafe.Pointer {
	var cr C.GdkRectangle
	cr.x = C.int(self.X)
	cr.y = C.int(self.Y)
	cr.width = C.int(self.Width)
	cr.height = C.int(self.Height)

	return unsafe.Pointer(&cr)
}

func (self Rectangle) Intersect(r2 Rectangle) (bool, Rectangle) {
	var cr3 C.GdkRectangle
	cr1 := nativeFromRectangle(self)
	defer cr1.Free()
	cr2 := nativeFromRectangle(r2)
	defer cr2.Free()
	b := C.gdk_rectangle_intersect((*C.GdkRectangle)(cr1.GetPtr()), (*C.GdkRectangle)(cr2.GetPtr()), &cr3)
	rec := newRectangleFromNative(unsafe.Pointer(&cr3)).(Rectangle)

	return g.GoBool(unsafe.Pointer(&b)), rec
}

func (self Rectangle) Union(r2 Rectangle) Rectangle {
	var cr3 C.GdkRectangle
	cr1 := nativeFromRectangle(self)
	defer cr1.Free()
	cr2 := nativeFromRectangle(r2)
	defer cr2.Free()
	C.gdk_rectangle_union((*C.GdkRectangle)(cr1.GetPtr()), (*C.GdkRectangle)(cr2.GetPtr()), &cr3)
	rec := newRectangleFromNative(unsafe.Pointer(&cr3)).(Rectangle)

	return rec
}

// End Rectangle }}}

// GdkRGBA {{{

// GdkRGBA type
type RGBA struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

//Conversion func
func newRGBAFromNative(obj unsafe.Pointer) interface{} {
	crgba := *((*C.GdkRGBA)(obj))

	rgba := RGBA{}
	rgba.Red = float64(crgba.red)
	rgba.Green = float64(crgba.green)
	rgba.Blue = float64(crgba.blue)
	rgba.Alpha = float64(crgba.alpha)

	return rgba
}

func nativeFromRGBA(rgba interface{}) *g.GValue {
	if r, ok := rgba.(RGBA); ok {
		gv := g.CreateCGValue(GdkType.RGBA, r.ToNative())
		return gv
	}
	return nil
}

// To be BoxedLike
func (self RGBA) GetBoxType() g.GType {
	return GdkType.RGBA
}

func (self RGBA) ToNative() unsafe.Pointer {
	crgba := C._new_rgba()

	crgba.red = C.gdouble(self.Red)
	crgba.green = C.gdouble(self.Green)
	crgba.blue = C.gdouble(self.Blue)
	crgba.alpha = C.gdouble(self.Alpha)

	return unsafe.Pointer(crgba)
}
// End RGBA type }}}

// GdkColor {{{

// GdkColor type
type Color struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
}

//Conversion func
func newColorFromNative(obj unsafe.Pointer) interface{} {
	cColor := *((*C.GdkColor)(obj))

	color := Color{}
	color.Pixel = uint32(cColor.pixel)
	color.Red = uint16(cColor.red)
	color.Green = uint16(cColor.green)
	color.Blue = uint16(cColor.blue)

	return color
}

func nativeFromColor(c interface{}) *g.GValue {
	if col, ok := c.(Color); ok {
		gv := g.CreateCGValue(GdkType.COLOR, col.ToNative())
		return gv
	}
	return nil
}

// To be BoxedLike
func (self Color) GetBoxType() g.GType {
	return GdkType.COLOR
}

func (self Color) ToNative() unsafe.Pointer {
	var cColor C.GdkColor

	cColor.pixel = C.guint32(self.Pixel)
	cColor.red = C.guint16(self.Red)
	cColor.green = C.guint16(self.Green)
	cColor.blue = C.guint16(self.Blue)

	return unsafe.Pointer(&cColor)
}
// End Color type }}}

// GdkAtom {{{

type Atom struct {
	object *C.GdkAtom
}

func (self Atom) ToNative() unsafe.Pointer {
	if self.object == nil {
		return nil
	}
	return unsafe.Pointer(self.object)
}

func (self Atom) Name() string {
	if self.object == nil {
		return ""
	}

	s := C.gdk_atom_name(*self.object)
	return g.GoString(unsafe.Pointer(s))
}

func NewAtomFromNative(obj unsafe.Pointer) Atom {
	a := Atom{}
	a.object = new(C.GdkAtom)
	*a.object = *((*C.GdkAtom)(obj))

	return a
}

func AtomIntern(atomName string, onlyIfExists bool) Atom {
	s := g.GString(atomName)
	defer s.Free()
	b := g.GBool(onlyIfExists)
	defer b.Free()

	o := C.gdk_atom_intern((*C.gchar)(s.GetPtr()), *((*C.gboolean)(b.GetPtr())))

	a := Atom{}
	a.object = new(C.GdkAtom)
	*a.object = o

	return a
}

// End GdkAtom }}}

// GDK3 INIT FUNC {{{
func init() {
	C.gdk_init(nil, nil)
	GdkType.DEVICE = g.GType(C.gdk_device_get_type())
	GdkType.RECTANGLE = g.GType(C.gdk_rectangle_get_type())
	GdkType.RGBA = g.GType(C.gdk_rgba_get_type())
	GdkType.COLOR = g.GType(C.gdk_color_get_type())
	GdkType.EVENT = g.GType(C.gdk_event_get_type())
	GdkType.DISPLAY = g.GType(C.gdk_display_get_type())
	GdkType.SCREEN = g.GType(C.gdk_screen_get_type())
	GdkType.WINDOW = g.GType(C.gdk_window_get_type())

	// Register Rectangle
	g.RegisterCType(GdkType.RECTANGLE, newRectangleFromNative)
	g.RegisterGoType(GdkType.RECTANGLE, nativeFromRectangle)

	// Register RGBA type
	g.RegisterCType(GdkType.RGBA, newRGBAFromNative)
	g.RegisterGoType(GdkType.RGBA, nativeFromRGBA)
}
// End GDK3 INIT FUNC }}}
