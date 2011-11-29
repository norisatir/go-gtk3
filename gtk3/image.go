package gtk3

/*
#include <gtk/gtk.h>

static inline GtkImage* to_GtkImage(void* obj) { return GTK_IMAGE(obj); }

*/
import "C"
import "unsafe"
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

	im.Widget = NewWidget(unsafe.Pointer(o))
	im.object = C.to_GtkImage(unsafe.Pointer(o))

	return im
}

// Conversion function for gobject registration map
func newImageFromNative(obj unsafe.Pointer) interface{} {
	im := &Image{}
	im.object = C.to_GtkImage(obj)
	im.Widget = NewWidget(obj)

	return &im
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

// To be widget-like
func (self Image) W() *Widget {
	return self.Widget
}
