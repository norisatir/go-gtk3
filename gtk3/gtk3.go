package gtk3

/*
#include <gtk/gtk.h>

static void _gtk_init(void* argc, void* argv) {
	gtk_init((int*)argc, (char***)argv);
}

*/
// #cgo pkg-config: gtk+-3.0 gobject-2.0
import "C"
import _ "github.com/norisatir/go-gtk3/gdk3"

func Init() {
	C._gtk_init(nil, nil)
}

type P map[string]interface{}
