package gdk3

/*
#include <gdk/gdk.h>

static void _gdk_threads_init() {
	gdk_threads_init();
}

*/
// #cgo pkg-config: gdk-3.0
import "C"
import _ "github.com/norisatir/go-gtk3/gobject"

func ThreadsInit() {
	C._gdk_threads_init()
}

func ThreadsEnter() {
	C.gdk_threads_enter()
}

func ThreadsLeave() {
	C.gdk_threads_leave()
}
