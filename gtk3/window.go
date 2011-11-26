package gtk3

/*
#include <gtk/gtk.h>

static GtkWindow* to_GtkWindow(void* obj) { return GTK_WINDOW(obj); }

*/
import "C"
import "github.com/norisatir/go-gtk3/gobject"
import "unsafe"

// GtkWindowType
type GtkWindowType int

const (
	GTK_WINDOW_TOPLEVEL GtkWindowType = 0
	GTK_WINDOW_POPUP    GtkWindowType = 1
)

type WindowLike interface {
	Wnd() *Window
}

type Window struct {
	object *C.GtkWindow
	*Container
}

// Create new window
func NewWindow(wtype GtkWindowType, properties map[string]interface{}) (w *Window) {
	w = &Window{}

	o := C.gtk_window_new(C.GtkWindowType(wtype))

	w.Container = NewContainer(unsafe.Pointer(o))

	w.object = C.to_GtkWindow(unsafe.Pointer(o))
	return w
}

// Conversion function for gobject registration map
func newWindowFromNative(obj unsafe.Pointer) interface{} {
	var window Window
	window.object = C.to_GtkWindow(obj)
	window.Container = NewContainer(unsafe.Pointer(window.object))
	return &window
}

func nativeFromWindow(w interface{}) *gobject.GValue {
	win, ok := w.(Window)
	if ok {
		gv := gobject.CreateCGValue(GtkType.WINDOW, win.ToNative())
		return gv
	}
	return nil
}

func init() {
	// Register GtkWindow to gobject
	gobject.RegisterCType(GtkType.WINDOW, newWindowFromNative)
	gobject.RegisterGoType(GtkType.WINDOW, nativeFromWindow)
}

// To be object-like
func (self Window) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Window) Connect(s string, f interface{}, datas ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, s, f, datas...)
}

func (self Window) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Window) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self Window) C() *Container {
	return self.Container
}

// Window interface

func (self *Window) SetTitle(title string) {
	gobject.SetProperty(self, "title", title)
}

func (self *Window) GetTitle() string {
	t := gobject.GetProperty(self, "title")
	defer t.Free()
	return gobject.GoString(t.GetPtr())
}

func (self *Window) SetWMClass(name, class string) {
	n := gobject.GString(name)
	c := gobject.GString(class)
	defer n.Free()
	defer c.Free()
	C.gtk_window_set_wmclass(self.object, (*C.gchar)(n.GetPtr()),
		(*C.gchar)(c.GetPtr()))
}

func (self *Window) SetStartupId(id string) {
	gobject.SetProperty(self, "startup-id", id)
}

func (self *Window) ActivateFocus() bool {
	b := C.gtk_window_activate_focus(self.object)
	gb, _ := gobject.ConvertToGo(unsafe.Pointer(&b), gobject.G_TYPE_BOOLEAN)
	return gb.(bool)
}

func (self *Window) SetFocus(w WidgetLike) {
	C.gtk_window_set_focus(self.object, w.W().object)
}

func (self *Window) SetDefault(w WidgetLike) {
	C.gtk_window_set_default(self.object, w.W().object)
}

func (self *Window) ActivateDefault() bool {
	b := C.gtk_window_activate_default(self.object)
	gb, _ := gobject.ConvertToGo(unsafe.Pointer(&b), gobject.G_TYPE_BOOLEAN)
	return gb.(bool)
}

func (self *Window) SetTransientFor(parent *Window) {
	C.gtk_window_set_transient_for(self.object, parent.object)
}

func (self *Window) GetTransientFor() *Window {
	w := unsafe.Pointer(C.gtk_window_get_transient_for(self.object))
	if w == nil {
		return nil
	}
	w1 := newWindowFromNative(w).(Window)
	return &w1
}

func (self *Window) SetOpacity(opacity float64) {
	gobject.SetProperty(self, "opacity", opacity)
}

func (self *Window) GetOpacity() float64 {
	return float64(C.gtk_window_get_opacity(self.object))
}

func (self *Window) SetSkipTaskbarHint(setting bool) {
	gobject.SetProperty(self, "skip-taskbar-hint", setting)
}

func (self *Window) GetSkipTaskbarHint() bool {
	b := gobject.GetProperty(self, "skip-taskbar-hint")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) GetFocus() WidgetLike {
	widget := C.gtk_window_get_focus(self.object)
	val, err := gobject.ConvertToGo(unsafe.Pointer(widget))

	if err != nil {
		return nil
	}
	return val.(WidgetLike)
}

func (self *Window) GetDefaultWidget() WidgetLike {
	widget := C.gtk_window_get_default_widget(self.object)
	val, err := gobject.ConvertToGo(unsafe.Pointer(widget))

	if err != nil {
		return nil
	}

	return val.(WidgetLike)
}

func (self *Window) SetSkipPagerHint(setting bool) {
	gobject.SetProperty(self, "skip-pager-hint", setting)
}

func (self *Window) GetSkipPagerHint() bool {
	b := gobject.GetProperty(self, "skip-pager-hint")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetUrgencyHint(setting bool) {
	gobject.SetProperty(self, "urgency-hint", setting)
}

func (self *Window) GetUrgencyHint() bool {
	b := gobject.GetProperty(self, "urgency-hint")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetAcceptFocus(setting bool) {
	gobject.SetProperty(self, "accept-focus", setting)
}

func (self *Window) GetAcceptFocus() bool {
	b := gobject.GetProperty(self, "accept-focus")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetFocusOnMap(setting bool) {
	gobject.SetProperty(self, "focus-on-map", setting)
}

func (self *Window) GetFocusOnMap() bool {
	b := gobject.GetProperty(self, "focus-on-map")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetDestroyWithParent(setting bool) {
	gobject.SetProperty(self, "destroy-with-parent", setting)
}

func (self *Window) GetDestroyWithParent() bool {
	b := gobject.GetProperty(self, "destroy-with-parent")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetMnemonicsVisible(setting bool) {
	gobject.SetProperty(self, "mnemonics-visible", setting)
}

func (self *Window) GetMnemonicsVisible() bool {
	b := gobject.GetProperty(self, "mnemonics-visible")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetResizable(setting bool) {
	gobject.SetProperty(self, "resizable", setting)
}

func (self *Window) GetResizable() bool {
	b := gobject.GetProperty(self, "resizable")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) IsActive() bool {
	b := C.gtk_window_is_active(self.object)
	cb, _ := gobject.ConvertToGo(unsafe.Pointer(&b), gobject.G_TYPE_BOOLEAN)
	return cb.(bool)
}

func (self *Window) HasTopLevelFocus() bool {
	b := C.gtk_window_has_toplevel_focus(self.object)
	cb, _ := gobject.ConvertToGo(unsafe.Pointer(&b), gobject.G_TYPE_BOOLEAN)
	return cb.(bool)
}

func (self *Window) SetDecorated(setting bool) {
	gobject.SetProperty(self, "decorated", setting)
}

func (self *Window) GetDecorated() bool {
	b := gobject.GetProperty(self, "decorated")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetDeletable(setting bool) {
	gobject.SetProperty(self, "deletable", setting)
}

func (self *Window) GetDeletable() bool {
	b := gobject.GetProperty(self, "deletable")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) SetModal(setting bool) {
	gobject.SetProperty(self, "modal", setting)
}

func (self *Window) GetModal() bool {
	b := gobject.GetProperty(self, "modal")
	defer b.Free()
	return gobject.GoBool(b.GetPtr())
}

func (self *Window) WindowPresent() {
	C.gtk_window_present(self.object)
}

func (self *Window) Iconify() {
	C.gtk_window_iconify(self.object)
}

func (self *Window) Deiconify() {
	C.gtk_window_deiconify(self.object)
}

func (self *Window) Stick() {
	C.gtk_window_stick(self.object)
}

func (self *Window) UnStick() {
	C.gtk_window_unstick(self.object)
}

func (self *Window) Maximize() {
	C.gtk_window_maximize(self.object)
}

func (self *Window) UnMaximize() {
	C.gtk_window_unmaximize(self.object)
}

func (self *Window) FullScreen() {
	C.gtk_window_fullscreen(self.object)
}

func (self *Window) UnFullScreen() {
	C.gtk_window_unfullscreen(self.object)
}

func (self *Window) SetKeepAbove(setting bool) {
	b := gobject.ConvertToC(setting)
	defer b.Free()
	C.gtk_window_set_keep_above(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Window) SetKeepBelow(setting bool) {
	b := gobject.ConvertToC(setting)
	defer b.Free()
	C.gtk_window_set_keep_below(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Window) SetApplication(app *Application) {
	gobject.SetProperty(self, "application", app)
}

func (self *Window) GetApplication() *Application {
	a := gobject.GetProperty(self, "application")
	defer a.Free()
	app, _ := gobject.ConvertToGo(a.GetPtr())
	if app != nil {
		return app.(*Application)
	}

	return nil
}
