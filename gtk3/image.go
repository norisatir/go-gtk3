package gtk3

/*
#include <gtk/gtk.h>

static inline GtkImage* to_GtkImage(void* obj) { return GTK_IMAGE(obj); }

*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type Image struct {
	object *C.GtkImage
	*Widget
}

func NewImageFromStock(stockId string, iconSize int) *Image {
	im := &Image{}

	sId := gobject.GString(stockId)
	defer sId.Free()

	o := C.gtk_image_new_from_stock((*C.gchar)(sId.GetPtr()), C.GtkIconSize(iconSize))
	im.object = C.to_GtkImage(unsafe.Pointer(o))

	if gobject.IsObjectFloating(im) {
		gobject.RefSink(im)
	}
	im.Widget = NewWidget(unsafe.Pointer(o))
	imageFinalizer(im)

	return im
}

// Clear Image struct when it goes out of reach
func imageFinalizer(im *Image) {
	runtime.SetFinalizer(im, func(im *Image) { gobject.Unref(im) })
}

// Conversion function for gobject registration map
func newImageFromNative(obj unsafe.Pointer) interface{} {
	im := &Image{}
	im.object = C.to_GtkImage(obj)

	if gobject.IsObjectFloating(im) {
		gobject.RefSink(im)
	} else {
		gobject.Ref(im)
	}
	im.Widget = NewWidget(obj)
	imageFinalizer(im)

	return im
}

func nativeFromImage(im interface{}) *gobject.GValue {
	image, ok := im.(*Image)
	if ok {
		gv := gobject.CreateCGValue(GtkType.IMAGE, image.ToNative())
		return gv
	}

	return nil
}

func init() {
	// Register GtkImage to gobject
	gobject.RegisterCType(GtkType.IMAGE, newImageFromNative)
	gobject.RegisterGoType(GtkType.IMAGE, nativeFromImage)
}
 
// To be object like
func (self Image) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Image) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Image) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Image) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Image) W() *Widget {
	return self.Widget
}
