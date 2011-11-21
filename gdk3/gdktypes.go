package gdk3

/*
#include <gdk/gdk.h>

*/
import "C"
import g "github.com/norisatir/go-gtk3/gobject"

var GdkType gdkTypes

type gdkTypes struct {
	DEVICE g.GType
}

func init() {
	GdkType = gdkTypes{}
	GdkType.DEVICE = g.GType(C.gdk_device_get_type())
}
