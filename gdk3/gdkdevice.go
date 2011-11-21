package gdk3

/*
#include <gdk/gdk.h>

static inline GdkDevice* to_GdkDevice(void* obj) {
	return GDK_DEVICE(obj);
}

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"

type Device struct {
	object *C.GdkDevice
}

// Conversion function
func newDeviceFromNative(obj unsafe.Pointer) interface{} {
	d := Device{}
	d.object = C.to_GdkDevice(obj)
	return &d
}

func init() {
	// Register GdkDevice to gobject
	gobject.RegisterCType(GdkType.DEVICE, newDeviceFromNative)
}

// To be object-like
func (self Device) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Device) Connect(name string, f interface{}, data ...interface{}) {
	gobject.Connect(self, name, f, data)
}

func (self Device) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Device) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
