package gdk3

/*
#include <gdk/gdk.h>

static void _gdk_threads_init() {
	gdk_threads_init();
}

static inline GdkDisplay* to_GdkDisplay(void* obj) { return GDK_DISPLAY(obj); }
static inline GdkDevice* to_GdkDevice(void* obj) { return GDK_DEVICE(obj); }
static inline GdkScreen* to_GdkScreen(void* obj) { return GDK_SCREEN(obj); }
static inline GdkWindow* to_GdkWindow(void* obj) { return GDK_WINDOW(obj); }

*/
// #cgo pkg-config: gdk-3.0
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

// Threads {{{
func ThreadsInit() {
	C._gdk_threads_init()
}

func ThreadsEnter() {
	C.gdk_threads_enter()
}

func ThreadsLeave() {
	C.gdk_threads_leave()
}
// End Threads }}}

// GdkDisplay {{{
//////////////////////////////

// GdkDisplay type
type Display struct {
	object *C.GdkDisplay
}

// Clear Display struct when it goes out of reach
func displayFinalizer(d *Display) {
	runtime.SetFinalizer(d, func(d *Display) { gobject.Unref(d) })
}

// Conversion funcs
func newDisplayFromNative(obj unsafe.Pointer) interface{} {
	d := &Display{}
	d.object = C.to_GdkDisplay(obj)

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	} else {
		gobject.Ref(d)
	}
	displayFinalizer(d)

	return d
}

func nativeFromDisplay(d interface{}) *gobject.GValue {
	if display, ok := d.(*Display); ok {
		gv := gobject.CreateCGValue(GdkType.DISPLAY, display.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Display) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Display) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data)
}

func (self Display) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Display) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// End GdkDisplay
////////////////////////////// }}}

// GdkDevice {{{
//////////////////////////////

// GdkDevice type
type Device struct {
	object *C.GdkDevice
}

// Clear Device struct when it goes out of reach
func deviceFinalizer(d *Device) {
	runtime.SetFinalizer(d, func(d *Device) { gobject.Unref(d) })
}

// Conversion function
func newDeviceFromNative(obj unsafe.Pointer) interface{} {
	d := &Device{}
	d.object = C.to_GdkDevice(obj)

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	} else {
		gobject.Ref(d)
	}
	deviceFinalizer(d)

	return d
}

func nativeFromDevice(d interface{}) *gobject.GValue {
	if device, ok := d.(*Device); ok {
		gv := gobject.CreateCGValue(GdkType.DEVICE, device.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Device) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Device) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data)
}

func (self Device) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Device) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// End GdkDevice
////////////////////////////// }}}

// GdkScreen {{{
//////////////////////////////

// GdkScreen type
type Screen struct {
	object *C.GdkScreen
}

// Clear Screen struct when it goes out of reach
func screenFinalizer(s *Screen) {
	runtime.SetFinalizer(s, func(s *Screen) { gobject.Unref(s) })
}

// Conversion function
func newScreenFromNative(obj unsafe.Pointer) interface{} {
	s := &Screen{}
	s.object = C.to_GdkScreen(obj)

	if gobject.IsObjectFloating(s) {
		gobject.RefSink(s)
	} else {
		gobject.Ref(s)
	}
	screenFinalizer(s)

	return s
}

func nativeFromScreen(s interface{}) *gobject.GValue {
	if screen, ok := s.(*Screen); ok {
		gv := gobject.CreateCGValue(GdkType.SCREEN, screen.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Screen) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Screen) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data)
}

func (self Screen) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Screen) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// End GdkScreen
////////////////////////////// }}}

// GdkWindow {{{
//////////////////////////////

// GdkWindow type
type Window struct {
	object *C.GdkWindow
}

// Clear Window struct when it goes out of reach
func windowFinalizer(w *Window) {
	runtime.SetFinalizer(w, func(w *Window) { gobject.Unref(w) })
}

// Conversion function
func newWindowFromNative(obj unsafe.Pointer) interface{} {
	w := &Window{}
	w.object = C.to_GdkWindow(obj)

	if gobject.IsObjectFloating(w) {
		gobject.RefSink(w)
	} else {
		gobject.Ref(w)
	}
	windowFinalizer(w)

	return w
}

func nativeFromWindow(w interface{}) *gobject.GValue {
	if window, ok := w.(*Window); ok {
		gv := gobject.CreateCGValue(GdkType.WINDOW, window.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Window) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Window) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data)
}

func (self Window) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Window) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// End GdkWindow
////////////////////////////// }}}

// GDK3 INIT FUNCS {{{
func init() {
	// Register GdkDisplay type
	gobject.RegisterCType(GdkType.DISPLAY, newDisplayFromNative)
	gobject.RegisterGoType(GdkType.DISPLAY, nativeFromDisplay)

	// Register GdkDevice type
	gobject.RegisterCType(GdkType.DEVICE, newDeviceFromNative)
	gobject.RegisterGoType(GdkType.DEVICE, nativeFromDevice)

	// Register GdkScreen type
	gobject.RegisterCType(GdkType.SCREEN, newScreenFromNative)
	gobject.RegisterGoType(GdkType.SCREEN, nativeFromScreen)

	//Register GdkWindow type
	gobject.RegisterCType(GdkType.WINDOW, newWindowFromNative)
	gobject.RegisterGoType(GdkType.WINDOW, nativeFromWindow)
}
// End GDK3 INIT FUNCS }}}
