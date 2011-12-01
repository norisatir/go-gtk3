package gtk3

/*
#include <gtk/gtk.h>
#include <glib-object.h>
#include <stdlib.h>

static inline int run_app(GtkApplication *app) {
	int status = g_application_run(G_APPLICATION (app), 0, NULL);
	return status;
}

static inline GApplication* to_GApplication(GtkApplication* g) { return G_APPLICATION(g); }
static inline GtkApplication* to_GtkApplication(void* o) {
    return GTK_APPLICATION(o);
}


*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

// Application Type
type Application struct {
	object *C.GtkApplication
}

// Create new application
func NewApplication(id string, flags int) *Application {
	cid := gobject.GString(id)
	defer cid.Free()
	app := C.gtk_application_new((*C.gchar)(cid.GetPtr()), C.GApplicationFlags(flags))
	gtkapp := &Application{app}

	if gobject.IsObjectFloating(gtkapp) {
		gobject.RefSink(gtkapp)
	}
	C.g_application_register(C.to_GApplication(app), nil, nil)
	//C.g_application_activate(C.to_GApplication(app))
	gtkapp.Connect("activate", func() {})

	return gtkapp
}

// Release reference when this Application struct goes out of scope
func appFinalizer(a *Application) {
	runtime.SetFinalizer(a, func(a *Application) { gobject.Unref(a) })
}

// Convert from Native to Go type
func appFromNative(app unsafe.Pointer) interface{} {
	ga := C.to_GtkApplication(app)
	gtkapp := &Application{ga}

	if gobject.IsObjectFloating(gtkapp) {
		gobject.RefSink(gtkapp)
	} else {
		gobject.Ref(gtkapp)
	}
	return &Application{ga}
}

// Convert from Go to Native
func nativeFromApp(app interface{}) *gobject.GValue {
	argapp, ok := app.(*Application)
	if ok {
		gv := gobject.CreateCGValue(GtkType.APPLICATION, argapp.ToNative())
		return gv
	}

	return nil
}

func init() {
	gobject.RegisterCType(GtkType.APPLICATION, appFromNative)
	gobject.RegisterGoType(GtkType.APPLICATION, nativeFromApp)
}


// To be object-like
func (self Application) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Application) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Application) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Application) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Add Window to application object
func (self *Application) AddWindow(window *Window) {
	if window == nil {
		return
	}
	C.gtk_application_add_window(self.object, window.object)
}

// Remove Window
func (self *Application) RemoveWindow(window *Window) {
	if window == nil {
		return
	}

	C.gtk_application_remove_window(self.object, window.object)
}

// Run app
func (self *Application) Run() {
	C.run_app(self.object)
}

