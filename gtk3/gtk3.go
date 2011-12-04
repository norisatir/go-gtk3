package gtk3

/*
#include <gtk/gtk.h>
#include <stdlib.h>

static void _gtk_init(void* argc, void* argv) {
	gtk_init((int*)argc, (char***)argv);
}


// to_Gtk*** Funcs {{{
static inline GApplication* to_GApplication(GtkApplication* g) { return G_APPLICATION(g); }
static inline GtkWidget* to_GtkWidget(void* obj) { return GTK_WIDGET(obj); }
static inline GtkContainer* to_GtkContainer(void *obj) { return GTK_CONTAINER(obj); }
static inline GtkWindow* to_GtkWindow(void* obj) { return GTK_WINDOW(obj); }
static inline GtkBox* to_GtkBox(void* obj) { return GTK_BOX(obj); }
static inline GtkButtonBox* to_GtkButtonBox(void* obj) { return GTK_BUTTON_BOX(obj); }
static inline GtkFrame* to_GtkFrame(void* obj) { return GTK_FRAME(obj); }
static inline GtkGrid* to_GtkGrid(void* obj) { return GTK_GRID(obj); }
static inline GtkLabel* to_GtkLabel(void* obj) { return GTK_LABEL(obj); }
static inline GtkImage* to_GtkImage(void* obj) { return GTK_IMAGE(obj); }
static inline GtkButton* to_GtkButton(void* obj) { return GTK_BUTTON(obj); }
static inline GtkToggleButton* to_GtkToggleButton(void* obj) { return GTK_TOGGLE_BUTTON(obj); }
static inline GtkEntryBuffer* to_GtkEntryBuffer(void* obj) { return GTK_ENTRY_BUFFER(obj); }
static inline GtkEntry* to_GtkEntry(void* obj) { return GTK_ENTRY(obj); }
static inline GtkDialog* to_GtkDialog(void* obj) { return GTK_DIALOG(obj); }
static inline GtkMessageDialog* to_GtkMessageDialog(void* obj) { return GTK_MESSAGE_DIALOG(obj); }
static inline GtkSeparator* to_GtkSeparator(void* obj) { return GTK_SEPARATOR(obj); }
static inline GtkAdjustment* to_GtkAdjustment(void* obj) { return GTK_ADJUSTMENT(obj); }
static inline GtkRange* to_GtkRange(void* obj) { return GTK_RANGE(obj); }
static inline GtkScrollbar* to_GtkScrollbar(void* obj) { return GTK_SCROLLBAR(obj); }
static inline GtkScrolledWindow* to_GtkScrolledWindow(void* obj) { return GTK_SCROLLED_WINDOW(obj); }
// End }}}

// GtkApplication funcs {{{
static inline int run_app(GtkApplication *app) {
	int status = g_application_run(G_APPLICATION (app), 0, NULL);
	return status;
}

static inline GtkApplication* to_GtkApplication(void* o) {
    return GTK_APPLICATION(o);
}
// End }}}

// ButtonBox funcs {{{
static void _gtk_button_box_set_child_non_homogeneous(GtkButtonBox *widget, GtkWidget* child,
														gboolean non_homogeneous) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_button_box_set_child_non_homogeneous(widget, child, non_homogeneous);
#endif
}

static gboolean _gtk_button_box_get_child_non_homogeneous(GtkButtonBox *widget, GtkWidget* child) {
#if GTK_CHECK_VERSION(3,2,0)
	return gtk_button_box_get_child_non_homogeneous(widget, child);
#else
	return NULL;
#endif
}
// End ButtonBox funcs }}}

// GtkGrid funcs {{{
static GtkWidget* _gtk_grid_get_child_at(GtkGrid* grid, gint left, gint top) {
#if GTK_CHECK_VERSION(3,2,0)
	return gtk_grid_get_child_at(grid, left, top);
#else
	return NULL;
#endif
}

static void _gtk_grid_insert_row(GtkGrid* grid, gint position) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_row(grid, position);
#endif
}

static void _gtk_grid_insert_column(GtkGrid* grid, gint position) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_column(grid, position);
#endif
}

static void _gtk_grid_insert_next_to(GtkGrid* grid, GtkWidget* sibling, GtkPositionType side) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_next_to(grid, sibling, side);
#endif
}
// End GtkGrid funcs }}}

// GtkEntry funcs {{{

static void _gtk_entry_set_placeholder_text(GtkEntry* entry, const gchar* text) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_entry_set_placeholder_text(entry, text);
#endif
}

static const gchar* _gtk_entry_get_placeholder_text(GtkEntry* entry) {
#if GTK_CHECK_VERSION(3,2,0)
	return gtk_entry_get_placeholder_text(entry);
#else
	return NULL;
#endif
}
// End GtkEntry funcs }}}

// GtkDialog funcs {{{
static inline GtkDialog* _dialog_new_with_buttons(const gchar* title,
												  GtkWindow* parent,
												  GtkDialogFlags flags,
												  const gchar* firstbutt,
												  gint resid) {
	GtkWidget* w = gtk_dialog_new_with_buttons(title, parent, flags, firstbutt, resid, NULL);
	return to_GtkDialog(w);
}
// End GtkDialog funcs }}}

// GtkMessageDialog funcs {{{
static inline GtkWidget* _new_message_dialog(GtkWindow* parent,
													GtkDialogFlags flags,
													GtkMessageType typ,
													GtkButtonsType buttons,
													const gchar* message) {
	return gtk_message_dialog_new(parent, flags, typ, buttons, message, NULL);
}

static inline void _gtk_message_dialog_format_secondary_text(GtkMessageDialog* md, const gchar* message) {
	gtk_message_dialog_format_secondary_text(md, message, NULL);
}
static inline void _gtk_message_dialog_format_secondary_markup(GtkMessageDialog* md, const gchar* message) {
	gtk_message_dialog_format_secondary_markup(md, message, NULL);
}
// End GtkMessageDialog funcs }}}

// GtkTreePath funcs {{{
static inline gint _gtk_tree_path_get_indice(gint* indices, int index) {
	return *(indices + index);
}
//End GtkTreePath func }}}

*/
// #cgo pkg-config: gtk+-3.0
import "C"
import "unsafe"
import "runtime"
import "fmt"
import "github.com/norisatir/go-gtk3/gobject"
import "github.com/norisatir/go-gtk3/gdk3"

// General types and functions {{{

func Init() {
	C._gtk_init(nil, nil)
}

func Main() {
	C.gtk_main()
}

func MainQuit() {
	C.gtk_main_quit()
}

// Convinient map for properties
type P map[string]interface{}

// Convinient map for buttons and id's
type ButID struct {
	Text     string
	Response int
}
type B []ButID

////////////////////////////// }}}

// Interfaces {{{
//////////////////////////////

// Widget-like interface must have method W()
type WidgetLike interface {
	W() *Widget
	gobject.ObjectLike
}

// Container-like interface must have method C()
type ContainerLike interface {
	C() *Container
	WidgetLike
}

// Window-like interface must have method Wnd()
type WindowLike interface {
	Wnd() *Window
	ContainerLike
}

// Box-like interface must have method CBox()
type BoxLike interface {
	CBox() *Box
}

// Button-like interface must have method B()
type ButtonLike interface {
	B() *Button
}

// Dialog-like interface must have method D()
type DialogLike interface {
	D() *Dialog
}

// Range-like interface must have method R()
type RangeLike interface {
	R() *Range
}
//////////////////////////////
// END Interfaces
////////////////////////////// }}}

// GtkApplication {{{
//////////////////////////////

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
// END GtkApplication
////////////////////////////// }}}

// GtkWidget {{{
//////////////////////////////

// Requisition type
type Requisition struct {
	Width, Height int
}

func (self Requisition) ToNative() unsafe.Pointer {
	var req C.GtkRequisition
	req.width = C.gint(self.Width)
	req.height = C.gint(self.Height)
	return unsafe.Pointer(&req)
}

// Widget Type
type Widget struct {
	object *C.GtkWidget
}

func NewWidget(o unsafe.Pointer) *Widget {
	w := &Widget{C.to_GtkWidget(o)}

	if gobject.IsObjectFloating(w) {
		gobject.RefSink(w)
	} else {
		gobject.Ref(w)
	}
	widgetFinalizer(w)

	return w
}

// Clear Widget struct when it goes out of reach
func widgetFinalizer(w *Widget) {
	runtime.SetFinalizer(w, func(w *Widget) { gobject.Unref(w) })
}

// To be object-like
func (self Widget) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Widget) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Widget) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Widget) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Widget standard funcs
func (self *Widget) Show() {
	C.gtk_widget_show(self.object)
}

func (self *Widget) ShowNow() {
	C.gtk_widget_show_now(self.object)
}

func (self *Widget) ShowAll() {
	C.gtk_widget_show_all(self.object)
}

func (self *Widget) Destroy() {
	C.gtk_widget_destroy(self.object)
	runtime.SetFinalizer(self, nil)
}

func (self *Widget) InDestruction() bool {
	b := C.gtk_widget_in_destruction(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Hide() {
	C.gtk_widget_hide(self.object)
}

func (self *Widget) Activate() bool {
	b := C.gtk_widget_activate(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Reparent(new_parent WidgetLike) {
	C.gtk_widget_reparent(self.object, new_parent.W().object)
}

func (self *Widget) IsFocus() bool {
	b := C.gtk_widget_is_focus(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(self.object)
}

func (self *Widget) GrabDefault() {
	C.gtk_widget_grab_default(self.object)
}

func (self *Widget) SetName(name string) {
	gobject.SetProperty(self, "name", name)
}

func (self *Widget) GetName() string {
	n := gobject.GetProperty(self, "name")
	defer n.Free()
	return gobject.GoString(n.GetPtr())
}

func (self *Widget) SetSensitive(sensitive bool) {
	gobject.SetProperty(self, "sensitive", sensitive)
}

func (self *Widget) SetParent(parent WidgetLike) {
	C.gtk_widget_set_parent(self.object, parent.W().object)
}

func (self *Widget) GetTopLevel() WidgetLike {
	w := C.gtk_widget_get_toplevel(self.object)
	if tl, err := gobject.ConvertToGo(unsafe.Pointer(&w)); err == nil {
		return tl.(WidgetLike)
	}
	return nil
}

func (self *Widget) GetAncestor(widget_type gobject.GType) WidgetLike {
	w := C.gtk_widget_get_ancestor(self.object, C.GType(widget_type))
	if w == nil {
		return nil
	}
	if par, err := gobject.ConvertToGo(unsafe.Pointer(&w)); err == nil {
		return par.(WidgetLike)
	}
	return nil
}

func (self *Widget) GetPreferredSize() (minimumSize, naturalSize Requisition) {
	var min, nat C.GtkRequisition
	C.gtk_widget_get_preferred_size(self.object, &min, &nat)

	minimumSize = Requisition{int(min.width), int(min.height)}
	naturalSize = Requisition{int(nat.width), int(nat.height)}
	return
}

func (self *Widget) SetEvents(events gdk3.GdkEventMask) {
	C.gtk_widget_set_events(self.object, C.gint(events))
}

func (self *Widget) AddEvents(events gdk3.GdkEventMask) {
	C.gtk_widget_add_events(self.object, C.gint(events))
}

func (self *Widget) GetEvents() gdk3.GdkEventMask {
	return gdk3.GdkEventMask(C.gtk_widget_get_events(self.object))
}

func (self *Widget) Map() {
	C.gtk_widget_map(self.object)
}

func (self *Widget) Unmap() {
	C.gtk_widget_unmap(self.object)
}
//////////////////////////////
// END GtkWidget
////////////////////////////// }}}

// GtkContainer {{{
//////////////////////////////

// Container Type
type Container struct {
	object *C.GtkContainer
	*Widget
}

// New Container from widget
func NewContainer(o unsafe.Pointer) *Container {
	con := &Container{}
	con.object = C.to_GtkContainer(o)

	if gobject.IsObjectFloating(con) {
		gobject.RefSink(con)
	} else {
		gobject.Ref(con)
	}
	con.Widget = NewWidget(o)
	containerFinalizer(con)

	return con
}

// Finalizer for Container struct
func containerFinalizer(c *Container) {
	runtime.SetFinalizer(c, func(c *Container) { gobject.Unref(c) })
}

// To be Object-like
func (self Container) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Container) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Container) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Container) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Widget-like
func (self Container) W() *Widget {
	return self.Widget
}

// interfaces
func (self *Container) Add(w WidgetLike) {
	C.gtk_container_add(self.object, w.W().object)
}

func (self *Container) Remove(w WidgetLike) {
	C.gtk_container_remove(self.object, w.W().object)
}

func (self *Container) GetFocusChild() WidgetLike {
	w := C.gtk_container_get_focus_child(self.object)
	i, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return i.(WidgetLike)
	}
	return nil
}

func (self *Container) GetBorderWidth() uint {
	return uint(C.gtk_container_get_border_width(self.object))
}

func (self *Container) SetBorderWidth(width uint) {
	C.gtk_container_set_border_width(self.object, C.guint(width))
}
//////////////////////////////
// END GtkContainer
////////////////////////////// }}}

// GtkWindow {{{
//////////////////////////////

// Window type
type Window struct {
	object *C.GtkWindow
	*Container
}

// Create new window
func NewWindow(windowType int) (w *Window) {
	w = &Window{}
	o := C.gtk_window_new(C.GtkWindowType(windowType))
	w.object = C.to_GtkWindow(unsafe.Pointer(o))

	if gobject.IsObjectFloating(w) {
		gobject.RefSink(w)
	}
	w.Container = NewContainer(unsafe.Pointer(o))
	windowFinalizer(w)

	return w
}

// Clear Window struct when it goes out of reach
func windowFinalizer(w *Window) {
	runtime.SetFinalizer(w, func(w *Window) { gobject.Unref(w) })
}

// Conversion function for gobject registration map
func newWindowFromNative(obj unsafe.Pointer) interface{} {
	w := &Window{}
	w.object = C.to_GtkWindow(obj)

	if gobject.IsObjectFloating(w) {
		gobject.RefSink(w)
	} else {
		gobject.Ref(w)
	}
	w.Container = NewContainer(obj)
	windowFinalizer(w)

	return w
}

func nativeFromWindow(w interface{}) *gobject.GValue {
	win, ok := w.(*Window)
	if ok {
		gv := gobject.CreateCGValue(GtkType.WINDOW, win.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Window) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Window) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
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
//////////////////////////////
// END GtkWindow
////////////////////////////// }}}

// GtkBox {{{
//////////////////////////////

// Box type
type Box struct {
	object *C.GtkBox
	*Container
}

func NewBox(orientation int, spacing int) *Box {
	box := &Box{}
	o := C.gtk_box_new(C.GtkOrientation(orientation), C.gint(spacing))
	box.object = C.to_GtkBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(box) {
		gobject.RefSink(box)
	}
	box.Container = NewContainer(unsafe.Pointer(o))
	boxFinalizer(box)

	return box
}

func NewHBox(spacing int) *Box {
	return NewBox(GtkOrientation.HORIZONTAL, spacing)
}

func NewVBox(spacing int) *Box {
	return NewBox(GtkOrientation.VERTICAL, spacing)
}

// Clear Box struct when it goes out of reach
func boxFinalizer(box *Box) {
	runtime.SetFinalizer(box, func(box *Box) { gobject.Unref(box) })
}
// Conversion function for gobject registration map
func newBoxFromNative(obj unsafe.Pointer) interface{} {
	box := &Box{}
	box.object = C.to_GtkBox(obj)

	if gobject.IsObjectFloating(box) {
		gobject.RefSink(box)
	} else {
		gobject.Ref(box)
	}
	box.Container = NewContainer(obj)
	boxFinalizer(box)

	return box
}

func nativeFromBox(b interface{}) *gobject.GValue {
	box, ok := b.(*Box)
	if ok {
		gv := gobject.CreateCGValue(GtkType.BOX, box.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Box) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Box) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Box) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Box) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self Box) C() *Container {
	return self.Container
}

// Box interface
func (self *Box) PackStart(w WidgetLike, expand bool, fill bool, padding uint) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_pack_start(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(e.GetPtr())), C.guint(padding))
}

func (self *Box) PackEnd(w WidgetLike, expand bool, fill bool, padding uint) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_pack_end(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(e.GetPtr())), C.guint(padding))
}

func (self *Box) GetHomogeneous() bool {
	h := gobject.GetProperty(self, "homogeneous")
	defer h.Free()
	return gobject.GoBool(h.GetPtr())
}

func (self *Box) SetHomogeneous(homogeneous bool) {
	gobject.SetProperty(self, "homogeneous", homogeneous)
}

func (self *Box) GetSpacing() int {
	return int(C.gtk_box_get_spacing(self.object))
}

func (self *Box) SetSpacing(spacing int) {
	gobject.SetProperty(self, "spacing", spacing)
}

func (self *Box) ReorderChild(w WidgetLike, position int) {
	C.gtk_box_reorder_child(self.object, w.W().object, C.gint(position))
}

func (self *Box) QueryChildPacking(w WidgetLike) (expand bool, fill bool, padding uint, ptype int) {
	var e C.gboolean
	var f C.gboolean
	var p C.guint
	var t C.GtkPackType
	C.gtk_box_query_child_packing(self.object, w.W().object, &e, &f, &p, &t)

	return gobject.GoBool(unsafe.Pointer(&e)), gobject.GoBool(unsafe.Pointer(&f)), uint(p), int(t)
}

func (self *Box) SetChildPacking(w WidgetLike, expand bool, fill bool, padding int, ptype int) {
	e := gobject.GBool(expand)
	f := gobject.GBool(fill)
	defer e.Free()
	defer f.Free()
	C.gtk_box_set_child_packing(self.object, w.W().object, *((*C.gboolean)(e.GetPtr())), *((*C.gboolean)(f.GetPtr())),
		C.guint(padding), C.GtkPackType(ptype))
}
//////////////////////////////
// END GtkBox
////////////////////////////// }}}

// GtkButtonBox {{{
//////////////////////////////

// ButtonBox type
type ButtonBox struct {
	object *C.GtkButtonBox
	*Box
}

func NewButtonBox(orientation int) *ButtonBox {
	bb := &ButtonBox{}
	o := C.gtk_button_box_new(C.GtkOrientation(orientation))
	bb.object = C.to_GtkButtonBox(unsafe.Pointer(o))

	if gobject.IsObjectFloating(bb) {
		gobject.RefSink(bb)
	}
	bb.Box = newBoxFromNative(unsafe.Pointer(o)).(*Box)
	buttonBoxFinalizer(bb)

	return bb
}

// Clear ButtonBox structure when it goes out of reach
func buttonBoxFinalizer(bb *ButtonBox) {
	runtime.SetFinalizer(bb, func(bb *ButtonBox) { gobject.Unref(bb) })
}

// Conversion function for gobject registration map
func newButtonBoxFromNative(obj unsafe.Pointer) interface{} {
	bb := &ButtonBox{}
	bb.object = C.to_GtkButtonBox(obj)

	if gobject.IsObjectFloating(bb) {
		gobject.RefSink(bb)
	}
	bb.Box = newBoxFromNative(obj).(*Box)
	buttonBoxFinalizer(bb)

	return bb
}

func nativeFromButtonBox(b interface{}) *gobject.GValue {
	bbox, ok := b.(*ButtonBox)
	if ok {
		gv := gobject.CreateCGValue(GtkType.BUTTON_BOX, bbox.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ButtonBox) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ButtonBox) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ButtonBox) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ButtonBox) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Box-like
func (self ButtonBox) CBox() *Box {
	return self.Box
}

// ButtonBox Interface
func (self *ButtonBox) GetLayout() int {
	l := C.gtk_button_box_get_layout(self.object)
	return int(l)
}

func (self *ButtonBox) GetChildSecondary(w WidgetLike) bool {
	cb := C.gtk_button_box_get_child_secondary(self.object, w.W().object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ButtonBox) GetChildNonHomogeneous(w WidgetLike) bool {
	cb := C._gtk_button_box_get_child_non_homogeneous(self.object, w.W().object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ButtonBox) SetLayout(layoutStyle int) {
	C.gtk_button_box_set_layout(self.object, C.GtkButtonBoxStyle(layoutStyle))
}

func (self *ButtonBox) SetChildSecondary(w WidgetLike, isSecondary bool) {
	cb := gobject.GBool(isSecondary)
	defer cb.Free()
	C.gtk_button_box_set_child_secondary(self.object, w.W().object, *((*C.gboolean)(cb.GetPtr())))
}

func (self *ButtonBox) SetChildNonHomogeneous(w WidgetLike, nonHomogeneous bool) {
	cb := gobject.GBool(nonHomogeneous)
	defer cb.Free()
	C._gtk_button_box_set_child_non_homogeneous(self.object, w.W().object, *((*C.gboolean)(cb.GetPtr())))
}
//////////////////////////////
// END GtkButtonBox
////////////////////////////// }}}

// GtkFrame {{{
//////////////////////////////

// Frame type
type Frame struct {
	object *C.GtkFrame
	*Container
}

func NewFrame(label string) *Frame {
	f := &Frame{}

	s := gobject.GString(label)
	defer s.Free()

	o := C.gtk_frame_new((*C.gchar)(s.GetPtr()))
	f.object = C.to_GtkFrame(unsafe.Pointer(o))

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	}
	f.Container = NewContainer(unsafe.Pointer(o))
	frameFinalizer(f)

	return f
}

// Clear Frame struct when it goes out of reach
func frameFinalizer(f *Frame) {
	runtime.SetFinalizer(f, func(f *Frame) { gobject.Unref(f) })
}

// Conversion function for gobject registration map
func newFrameFromNative(obj unsafe.Pointer) interface{} {
	f := &Frame{}
	f.object = C.to_GtkFrame(obj)

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	} else {
		gobject.Ref(f)
	}
	f.Container = NewContainer(unsafe.Pointer(f.object))
	frameFinalizer(f)

	return f
}

func nativeFromFrame(frame interface{}) *gobject.GValue {
	f, ok := frame.(*Frame)
	if ok {
		gv := gobject.CreateCGValue(GtkType.FRAME, f.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self Frame) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Frame) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Frame) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Frame) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-lie
func (self Frame) C() *Container {
	return self.Container
}

// Frame interface
func (self *Frame) SetLabel(label string) {
	s := gobject.GString(label)
	defer s.Free()
	C.gtk_frame_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Frame) SetLabelWidget(w WidgetLike) {
	C.gtk_frame_set_label_widget(self.object, w.W().object)
}

func (self *Frame) SetLabelAlign(xalign, yalign float32) {
	C.gtk_frame_set_label_align(self.object, C.gfloat(xalign), C.gfloat(yalign))
}

func (self *Frame) SetShadowType(gtk_shadow int) {
	C.gtk_frame_set_shadow_type(self.object, C.GtkShadowType(gtk_shadow))
}

func (self *Frame) GetLabel() string {
	l := C.gtk_frame_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(l))
}

func (self *Frame) GetLabelAlign() (xalign, yalign float32) {
	var cax C.gfloat
	var cay C.gfloat
	C.gtk_frame_get_label_align(self.object, &cax, &cay)
	return float32(cax), float32(cay)
}

func (self *Frame) GetLabelWidget() WidgetLike {
	cw := C.gtk_frame_get_label_widget(self.object)
	w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
	if err != nil {
		return w.(WidgetLike)
	}
	return nil
}

func (self *Frame) GetShadowType() int {
	return int(C.gtk_frame_get_shadow_type(self.object))
}
//////////////////////////////
// END GtkFrame
////////////////////////////// }}}

// GtkGrid {{{
//////////////////////////////

// GtkGrid Type
type Grid struct {
	object *C.GtkGrid
	*Container
}

func NewGrid() *Grid {
	grid := &Grid{}
	o := C.gtk_grid_new()
	grid.object = C.to_GtkGrid(unsafe.Pointer(o))
	grid.Container = NewContainer(unsafe.Pointer(o))

	if gobject.IsObjectFloating(grid) {
		gobject.RefSink(grid)
	}
	gridFinalizer(grid)

	return grid
}

// Clear Grid struct when it goes out of reach
func gridFinalizer(g *Grid) {
	runtime.SetFinalizer(g, func(g *Grid) { gobject.Unref(g) })
}

// Conversion function for gobject registration map
func newGridFromNative(obj unsafe.Pointer) interface{} {
	grid := &Grid{}
	grid.object = C.to_GtkGrid(obj)

	if gobject.IsObjectFloating(grid) {
		gobject.RefSink(grid)
	} else {
		gobject.Ref(grid)
	}
	grid.Container = NewContainer(obj)
	gridFinalizer(grid)

	return grid
}

func nativeFromGrid(g interface{}) *gobject.GValue {
	grid, ok := g.(*Grid)
	if ok {
		gv := gobject.CreateCGValue(GtkType.GRID, grid.ToNative())
		return gv
	}

	return nil
}

// To be Object-like
func (self Grid) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Grid) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Grid) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Grid) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
// To be container-like
func (self Grid) C() *Container {
	return self.Container
}

// Grid interface
func (self *Grid) Attach(w WidgetLike, left, top, width, height int) {
	C.gtk_grid_attach(self.object, w.W().object, C.gint(left), C.gint(top),
		C.gint(width), C.gint(height))
}

func (self *Grid) AttachNextTo(child WidgetLike, sibling WidgetLike, side, width, height int) {
	C.gtk_grid_attach_next_to(self.object, child.W().object, sibling.W().object, C.GtkPositionType(side),
		C.gint(width), C.gint(height))
}

func (self *Grid) GetChildAt(left, top int) WidgetLike {
	c := C._gtk_grid_get_child_at(self.object, C.gint(left), C.gint(top))
	if c == nil {
		return nil
	}

	child, err := gobject.ConvertToGo(unsafe.Pointer(c))
	if err == nil {
		return child.(WidgetLike)
	}

	return nil
}

func (self *Grid) InsertRow(position int) {
	C._gtk_grid_insert_row(self.object, C.gint(position))
}

func (self *Grid) InsertColumn(position int) {
	C._gtk_grid_insert_column(self.object, C.gint(position))
}

func (self *Grid) InsertNextTo(sibling WidgetLike, side int) {
	C._gtk_grid_insert_next_to(self.object, sibling.W().object, C.GtkPositionType(side))
}

func (self *Grid) SetRowHomogeneous(homogeneous bool) {
	b := gobject.GBool(homogeneous)
	defer b.Free()
	C.gtk_grid_set_row_homogeneous(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Grid) GetRowHomogeneous() bool {
	b := C.gtk_grid_get_row_homogeneous(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(self.object, C.guint(spacing))
}

func (self *Grid) GetRowSpacing() uint {
	return uint(C.gtk_grid_get_row_spacing(self.object))
}

func (self *Grid) SetColumnHomogeneous(homogeneous bool) {
	b := gobject.GBool(homogeneous)
	defer b.Free()
	C.gtk_grid_set_column_homogeneous(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Grid) GetColumnHomogeneous() bool {
	b := C.gtk_grid_get_column_homogeneous(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(self.object, C.guint(spacing))
}

func (self *Grid) GetColumnSpacing() uint {
	return uint(C.gtk_grid_get_column_spacing(self.object))
}
//////////////////////////////
// END GtkGrid
////////////////////////////// }}}

// GtkLabel {{{
//////////////////////////////

// GtkLabel type
type Label struct {
	object *C.GtkLabel
	*Widget
}

func NewLabel(str string) *Label {
	l := &Label{}

	s := gobject.GString(str)
	defer s.Free()

	o := C.gtk_label_new((*C.gchar)(s.GetPtr()))
	l.object = C.to_GtkLabel(unsafe.Pointer(o))

	if gobject.IsObjectFloating(l) {
		gobject.RefSink(l)
	}
	l.Widget = NewWidget(unsafe.Pointer(o))
	labelFinalizer(l)

	return l
}

func NewLabelWithMnemonic(str string) *Label {
	l := &Label{}

	s := gobject.GString(str)
	defer s.Free()

	o := C.gtk_label_new_with_mnemonic((*C.gchar)(s.GetPtr()))
	l.object = C.to_GtkLabel(unsafe.Pointer(o))

	if gobject.IsObjectFloating(l) {
		gobject.RefSink(l)
	}
	l.Widget = NewWidget(unsafe.Pointer(o))
	labelFinalizer(l)

	return l
}

// Clear label struct when it goes out of reach
func labelFinalizer(l *Label) {
	runtime.SetFinalizer(l, func(l *Label) { gobject.Unref(l) })
}

// Conversion function for gobject registration map
func newLabelFromNative(obj unsafe.Pointer) interface{} {
	l := &Label{}
	l.object = C.to_GtkLabel(obj)

	if gobject.IsObjectFloating(l) {
		gobject.RefSink(l)
	} else {
		gobject.Ref(l)
	}
	l.Widget = NewWidget(unsafe.Pointer(l.object))
	labelFinalizer(l)

	return l
}

func nativeFromLabel(label interface{}) *gobject.GValue {
	l, ok := label.(*Label)
	if ok {
		gv := gobject.CreateCGValue(GtkType.LABEL, l.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self Label) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Label) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Label) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Label) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Label) W() *Widget {
	return self.Widget
}

// Label interface

func (self *Label) SetText(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_text(self.object, (*C.gchar)(s.GetPtr()))
}

//TODO: gtk_label_set_attributes

func (self *Label) SetMarkup(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_markup(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Label) SetMarkupWithMnemonic(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_markup_with_mnemonic(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Label) SetPattern(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_pattern(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Label) SetJustify(jtype int) {
	C.gtk_label_set_justify(self.object, C.GtkJustification(jtype))
}

//TODO: gtk_label_set_ellipsize

func (self *Label) SetWidthChars(n_chars int) {
	C.gtk_label_set_width_chars(self.object, C.gint(n_chars))
}

func (self *Label) SetMaxWidthChars(n_chars int) {
	C.gtk_label_set_max_width_chars(self.object, C.gint(n_chars))
}

func (self *Label) SetLineWrap(wrap bool) {
	b := gobject.GBool(wrap)
	defer b.Free()
	C.gtk_label_set_line_wrap(self.object, *((*C.gboolean)(b.GetPtr())))
}

//TODO: gtk_label_set_line_wrap_mode

func (self *Label) GetLayoutOffsets() (x, y int) {
	var cx C.gint
	var cy C.gint
	C.gtk_label_get_layout_offsets(self.object, &cx, &cy)
	return int(cx), int(cy)
}

func (self *Label) GetMnemonicKeyVal() uint {
	ckv := C.gtk_label_get_mnemonic_keyval(self.object)
	return uint(ckv)
}

func (self *Label) GetSelectable() bool {
	sel := C.gtk_label_get_selectable(self.object)
	return gobject.GoBool(unsafe.Pointer(&sel))
}

func (self *Label) GetText() string {
	txt := C.gtk_label_get_text(self.object)
	return gobject.GoString(unsafe.Pointer(txt))
}

func (self *Label) SelectRegion(start_offset, end_offset int) {
	C.gtk_label_select_region(self.object, C.gint(start_offset), C.gint(end_offset))
}

func (self *Label) SetMnemonicWidget(w WidgetLike) {
	C.gtk_label_set_mnemonic_widget(self.object, w.W().object)
}

func (self *Label) SetSelectable(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_selectable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) SetTextWithMnemonic(text string) {
	s := gobject.GString(text)
	defer s.Free()
	C.gtk_label_set_text_with_mnemonic(self.object, (*C.gchar)(s.GetPtr()))
}

//TODO: gtk_label_get_attributes

func (self *Label) GetJustify() int {
	j := C.gtk_label_get_justify(self.object)
	return int(j)
}

//TODO: gtk_label_get_ellipsize

func (self *Label) GetWidthChars() int {
	w := C.gtk_label_get_width_chars(self.object)
	return int(w)
}

func (self *Label) GetMaxWidthChars() int {
	w := C.gtk_label_get_max_width_chars(self.object)
	return int(w)
}

func (self *Label) GetLabel() string {
	l := C.gtk_label_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(&l))
}

//TODO: gtk_label_get_layout

func (self *Label) GetLineWrap() bool {
	lw := C.gtk_label_get_line_wrap(self.object)
	return gobject.GoBool(unsafe.Pointer(&lw))
}

//TODO: gtk_label_get_line_wrap_mode

func (self *Label) GetMnemonicWidget() WidgetLike {
	w := C.gtk_label_get_mnemonic_widget(self.object)
	if w == nil {
		return nil
	}

	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err != nil {
		return widget.(WidgetLike)
	}

	return nil
}

func (self *Label) GetSelectionBounds() (bool, int, int) {
	var start C.gint
	var end C.gint
	b := C.gtk_label_get_selection_bounds(self.object, &start, &end)
	isSel := gobject.GoBool(unsafe.Pointer(&b))

	return isSel, int(start), int(end)
}

func (self *Label) GetUseMarkup() bool {
	b := C.gtk_label_get_use_markup(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Label) GetUseUnderline() bool {
	b := C.gtk_label_get_use_underline(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Label) GetSingleLineMode() bool {
	b := C.gtk_label_get_single_line_mode(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Label) GetAngle() float64 {
	e := C.gtk_label_get_angle(self.object)
	return float64(e)
}

func (self *Label) SetLabel(str string) {
	s := gobject.GString(str)
	defer s.Free()
	C.gtk_label_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Label) SetUseMarkup(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_use_markup(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) SetUseUnderline(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_use_underline(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) SetSingleLineMode(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_single_line_mode(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(self.object, C.gdouble(angle))
}

func (self *Label) GetCurrentURI() string {
	uri := C.gtk_label_get_current_uri(self.object)
	return gobject.GoString(unsafe.Pointer(&uri))
}

func (self *Label) SetTrackVisitedLinks(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_label_set_track_visited_links(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Label) GetTrackVisitedLinks() bool {
	b := C.gtk_label_get_track_visited_links(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}
//////////////////////////////
// END GtkLabel
////////////////////////////// }}}

// GtkImage {{{
//////////////////////////////

// Image type
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
//////////////////////////////
// END GtkImage
////////////////////////////// }}}

// GtkButton {{{
//////////////////////////////

// GtkButton type
type Button struct {
	object *C.GtkButton
	*Container
}

//Create and return new button Structure
func NewButton() *Button {
	b := &Button{}
	o := C.gtk_button_new()
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Container = NewContainer(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

func NewButtonWithLabel(label string) *Button {
	b := &Button{}
	l := gobject.GString(label)
	defer l.Free()

	o := C.gtk_button_new_with_label((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Container = NewContainer(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

func NewButtonFromStock(stockId string) *Button {
	b := NewButtonWithLabel(stockId)
	b.SetUseStock(true)
	return b
}

func NewButtonWithMnemonic(label string) *Button {
	b := &Button{}
	l := gobject.GString(label)
	defer l.Free()
	o := C.gtk_button_new_with_mnemonic((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Container = NewContainer(unsafe.Pointer(o))
	buttonFinalizer(b)

	return b
}

// Unref button, when Button structer goes out of reach
func buttonFinalizer(b *Button) {
	runtime.SetFinalizer(b, func(b *Button) { gobject.Unref(b) })
}

// Conversion function for gobject registration map
func newButtonFromNative(obj unsafe.Pointer) interface{} {
	b := &Button{}
	b.object = C.to_GtkButton(obj)

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Container = NewContainer(obj)
	buttonFinalizer(b)

	return b
}

func nativeFromButton(b interface{}) *gobject.GValue {
	if but, ok := b.(*Button); ok {
		gv := gobject.CreateCGValue(GtkType.BUTTON, but.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Button) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Button) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Button) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Button) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container like
func (self Button) C() *Container {
	return self.Container
}

// Button interface
func (self *Button) SetLabel(label string) {
	s := gobject.GString(label)
	defer s.Free()
	C.gtk_button_set_label(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Button) GetLabel() string {
	l := C.gtk_button_get_label(self.object)
	return gobject.GoString(unsafe.Pointer(l))
}

func (self *Button) GetUseStock() bool {
	b := C.gtk_button_get_use_stock(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetUseStock(useStock bool) {
	gobject.SetProperty(self, "use-stock", useStock)
}

func (self *Button) GetUseUnderline() bool {
	b := C.gtk_button_get_use_underline(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetUseUnderline(underline bool) {
	gobject.SetProperty(self, "use-underline", underline)
}

func (self *Button) SetFocusOnClick(focusOnClick bool) {
	gobject.SetProperty(self, "focus-on-click", focusOnClick)
}

func (self *Button) GetFocusOnClick() bool {
	b := C.gtk_button_get_focus_on_click(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Button) SetAlignment(xalign, yalign float32) {
	C.gtk_button_set_alignment(self.object, C.gfloat(xalign), C.gfloat(yalign))
}

func (self *Button) GetAlignment() (float32, float32) {
	var cx C.gfloat
	var cy C.gfloat
	C.gtk_button_get_alignment(self.object, &cx, &cy)
	return float32(cx), float32(cy)
}

func (self *Button) SetImage(w WidgetLike) {
	C.gtk_button_set_image(self.object, w.W().object)
}

func (self *Button) GetImage() WidgetLike {
	cw := C.gtk_button_get_image(self.object)
	widget, err := gobject.ConvertToGo(unsafe.Pointer(&cw))
	if err != nil {
		return widget.(WidgetLike)
	}

	return nil
}

func (self *Button) SetImagePosition(position int) {
	C.gtk_button_set_image_position(self.object, C.GtkPositionType(position))
}

func (self *Button) GetImagePosition() int {
	return int(C.gtk_button_get_image_position(self.object))
}

func (self *Button) Clicked() {
	C.gtk_button_clicked(self.object)
}
//TODO: gtk_button_get_event_window
//////////////////////////////
// END GtkButton
////////////////////////////// }}}

// GtkToggleButton {{{
//////////////////////////////

// GtkToggleButton type
type ToggleButton struct {
	object *C.GtkToggleButton
	*Button
}

// Create and return new toggle button
func NewToggleButton() *ToggleButton {
	b := &ToggleButton{}
	o := C.gtk_toggle_button_new()
	b.object = C.to_GtkToggleButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	toggleButtonFinalizer(b)

	return b
}

func NewToggleButtonWithLabel(label string) *ToggleButton {
	b := &ToggleButton{}
	l := gobject.GString(label)
	defer l.Free()

	o := C.gtk_toggle_button_new_with_label((*C.gchar)(l.GetPtr()))
	b.object = C.to_GtkToggleButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	}
	b.Button = newButtonFromNative(unsafe.Pointer(o)).(*Button)
	toggleButtonFinalizer(b)

	return b
}

// Clear ToggleButton when it goes out of reach
func toggleButtonFinalizer(b *ToggleButton) {
	runtime.SetFinalizer(b, func(b *ToggleButton) { gobject.Unref(b) })
}

// Conversion function for gobject registration map
func newToggleButtonFromNative(obj unsafe.Pointer) interface{} {
	b := &ToggleButton{}
	b.object = C.to_GtkToggleButton(obj)

	if gobject.IsObjectFloating(b) {
		gobject.RefSink(b)
	} else {
		gobject.Ref(b)
	}
	b.Button = newButtonFromNative(obj).(*Button)
	toggleButtonFinalizer(b)

	return b
}

func nativeFromToggleButton(b interface{}) *gobject.GValue {
	if but, ok := b.(*ToggleButton); ok {
		gv := gobject.CreateCGValue(GtkType.TOGGLE_BUTTON, but.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ToggleButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ToggleButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ToggleButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ToggleButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be button-like
func (self ToggleButton) B() *Button {
	return self.Button
}

// ToggleButton interface

func (self *ToggleButton) SetMode(drawIndicator bool) {
	b := gobject.GBool(drawIndicator)
	defer b.Free()
	C.gtk_toggle_button_set_mode(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ToggleButton) GetMode() bool {
	cb := C.gtk_toggle_button_get_mode(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) Toggled() {
	C.gtk_toggle_button_toggled(self.object)
}

func (self *ToggleButton) GetActive() bool {
	cb := C.gtk_toggle_button_get_active(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) SetActive(isActive bool) {
	b := gobject.GBool(isActive)
	defer b.Free()
	C.gtk_toggle_button_set_active(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ToggleButton) GetInconsistent() bool {
	cb := C.gtk_toggle_button_get_inconsistent(self.object)
	return gobject.GoBool(unsafe.Pointer(&cb))
}

func (self *ToggleButton) SetInconsistent(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_toggle_button_set_inconsistent(self.object, *((*C.gboolean)(b.GetPtr())))
}
//////////////////////////////
// END GtkToggleButton
////////////////////////////// }}}

// GtkEntryBuffer {{{
//////////////////////////////

// GtkEntryBuffer type
type EntryBuffer struct {
	object *C.GtkEntryBuffer
}

func NewEntryBuffer(initial_chars string) *EntryBuffer {
	e := &EntryBuffer{}

	s := gobject.GString(initial_chars)
	defer s.Free()

	o := C.gtk_entry_buffer_new((*C.gchar)(s.GetPtr()), C.gint(len(initial_chars)))
	e.object = C.to_GtkEntryBuffer(unsafe.Pointer(o))

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	}
	entryBufferFinalizer(e)

	return e

}

// Clear EntryBuffer struct when it goes out of reach
func entryBufferFinalizer(e *EntryBuffer) {
	runtime.SetFinalizer(e, func(e *EntryBuffer) { gobject.Unref(e) })
}

// Conversion function for gobject registration map
func newEntryBufferFromNative(obj unsafe.Pointer) interface{} {
	e := &EntryBuffer{}
	e.object = C.to_GtkEntryBuffer(obj)

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	} else {
		gobject.Ref(e)
	}

	return e
}

func nativeFromEntryBuffer(eb interface{}) *gobject.GValue {
	e, ok := eb.(*EntryBuffer)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ENTRY_BUFFER, e.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self EntryBuffer) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self EntryBuffer) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self EntryBuffer) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self EntryBuffer) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// EntryBuffer interface

func (self *EntryBuffer) GetText() string {
	s := C.gtk_entry_buffer_get_text(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *EntryBuffer) SetText(chars string) {
	s := gobject.GString(chars)
	defer s.Free()

	C.gtk_entry_buffer_set_text(self.object, (*C.gchar)(s.GetPtr()), C.gint(len(chars)))
}

func (self *EntryBuffer) GetBytes() uint64 {
	return uint64(C.gtk_entry_buffer_get_bytes(self.object))
}

func (self *EntryBuffer) GetLength() uint {
	return uint(C.gtk_entry_buffer_get_length(self.object))
}

func (self *EntryBuffer) GetMaxLength() int {
	return int(C.gtk_entry_buffer_get_max_length(self.object))
}

func (self *EntryBuffer) SetMaxLength(maxLength int) {
	C.gtk_entry_buffer_set_max_length(self.object, C.gint(maxLength))
}

func (self *EntryBuffer) InsertText(position uint, chars string) {
	s := gobject.GString(chars)
	defer s.Free()

	C.gtk_entry_buffer_insert_text(self.object, C.guint(position), (*C.gchar)(s.GetPtr()), C.gint(len(chars)))
}

func (self *EntryBuffer) DeleteText(position uint, num_chars int) {
	C.gtk_entry_buffer_delete_text(self.object, C.guint(position), C.gint(num_chars))
}
//////////////////////////////
// END GtkEntryBuffer
////////////////////////////// }}}

// GtkEntry {{{
//////////////////////////////

// GtkEntry Type
type Entry struct {
	object *C.GtkEntry
	*Widget
}

func NewEntry() *Entry {
	e := &Entry{}
	o := C.gtk_entry_new()
	e.object = C.to_GtkEntry(unsafe.Pointer(o))

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	}
	e.Widget = NewWidget(unsafe.Pointer(o))
	entryFinalizer(e)

	return e
}

func NewEntryWithBuffer(eb *EntryBuffer) *Entry {
	e := &Entry{}
	o := C.gtk_entry_new_with_buffer(eb.object)
	e.object = C.to_GtkEntry(unsafe.Pointer(o))

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	}
	e.Widget = NewWidget(unsafe.Pointer(o))
	entryFinalizer(e)

	return e
}

// Clear Entry struct when it goes out of reach
func entryFinalizer(e *Entry) {
	runtime.SetFinalizer(e, func(e *Entry) { gobject.Unref(e) })
}

// Conversion function for gobject registration map
func newEntryFromNative(obj unsafe.Pointer) interface{} {
	e := &Entry{}
	e.object = C.to_GtkEntry(obj)

	if gobject.IsObjectFloating(e) {
		gobject.RefSink(e)
	} else {
		gobject.Ref(e)
	}
	e.Widget = NewWidget(obj)
	entryFinalizer(e)

	return e
}

func nativeFromEntry(entry interface{}) *gobject.GValue {
	e, ok := entry.(*Entry)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ENTRY, e.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self Entry) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Entry) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Entry) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Entry) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Entry) W() *Widget {
	return self.Widget
}

// Entry interface
func (self *Entry) GetBuffer() *EntryBuffer {
	b := C.gtk_entry_get_buffer(self.object)
	buf, err := gobject.ConvertToGo(unsafe.Pointer(b))
	if err == nil {
		return buf.(*EntryBuffer)
	}
	return nil
}

func (self *Entry) SetBuffer(buffer *EntryBuffer) {
	C.gtk_entry_set_buffer(self.object, buffer.object)
}

func (self *Entry) SetText(text string) {
	s := gobject.GString(text)
	defer s.Free()

	C.gtk_entry_set_text(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Entry) GetText() string {
	s := C.gtk_entry_get_text(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Entry) GetTextLength() uint {
	return uint(C.gtk_entry_get_text_length(self.object))
}

func (self *Entry) GetTextArea() gdk3.Rectangle {
	var crec C.GdkRectangle
	C.gtk_entry_get_text_area(self.object, &crec)

	rec, err := gobject.ConvertToGo(unsafe.Pointer(&crec), gdk3.GdkType.RECTANGLE)
	if err == nil {
		return rec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

func (self *Entry) SetVisibility(visible bool) {
	b := gobject.GBool(visible)
	defer b.Free()
	C.gtk_entry_set_visibility(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Entry) SetInvisiblechar(char rune) {
	C.gtk_entry_set_invisible_char(self.object, C.gunichar(char))
}

func (self *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(self.object)
}

func (self *Entry) SetMaxLength(max int) {
	C.gtk_entry_set_max_length(self.object, C.gint(max))
}

func (self *Entry) GetActivatesDefault() bool {
	b := C.gtk_entry_get_activates_default(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Entry) GetHasFrame() bool {
	b := C.gtk_entry_get_has_frame(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

//TODO: gtk_entry_get_inner_border

func (self *Entry) GetWidthChars() int {
	return int(C.gtk_entry_get_width_chars(self.object))
}

func (self *Entry) SetActivatesDefault(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_entry_set_activates_default(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Entry) SetHasFrame(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_entry_set_has_frame(self.object, *((*C.gboolean)(b.GetPtr())))
}

//TODO: gtk_entry_set_inner_border

func (self *Entry) SetWidthChars(nChars int) {
	C.gtk_entry_set_width_chars(self.object, C.gint(nChars))
}

func (self *Entry) GetInvisibleChar() rune {
	return rune(C.gtk_entry_get_invisible_char(self.object))
}

func (self *Entry) SetAlignment(xalign float32) {
	C.gtk_entry_set_alignment(self.object, C.gfloat(xalign))
}

func (self *Entry) GetAlignment() float32 {
	return float32(C.gtk_entry_get_alignment(self.object))
}

func (self *Entry) SetPlaceholderText(text string) {
	s := gobject.GString(text)
	defer s.Free()
	C._gtk_entry_set_placeholder_text(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Entry) GetPlaceholderText() string {
	s := C._gtk_entry_get_placeholder_text(self.object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Entry) SetOverwriteMode(overwrite bool) {
	b := gobject.GBool(overwrite)
	defer b.Free()
	C.gtk_entry_set_overwrite_mode(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Entry) GetOverwriteMode() bool {
	b := C.gtk_entry_get_overwrite_mode(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}
//////////////////////////////
// END GtkEntry
////////////////////////////// }}}

// GtkDialog {{{
//////////////////////////////

// GtkDialog type
type Dialog struct {
	object *C.GtkDialog
	*Window
}

func NewDialog() *Dialog {
	d := &Dialog{}
	o := C.gtk_dialog_new()
	d.object = C.to_GtkDialog(unsafe.Pointer(o))

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	}
	d.Window = newWindowFromNative(unsafe.Pointer(o)).(*Window)
	dialogFinalizer(d)

	return d
}

func NewDialogWithButtons(title string, parent *Window, flags int, butAndID B) *Dialog {
	// Must have at least one button
	if len(butAndID) == 0 {
		return nil
	}

	t := gobject.GString(title)
	defer t.Free()

	firstButton := butAndID[0].Text
	fb := gobject.GString(firstButton)
	defer fb.Free()

	firstId := butAndID[0].Response

	var wparent *C.GtkWindow = nil
	if parent != nil {
		wparent = parent.object
	}

	d := &Dialog{}
	o := C._dialog_new_with_buttons((*C.gchar)(t.GetPtr()), wparent, C.GtkDialogFlags(flags),
		(*C.gchar)(fb.GetPtr()), C.gint(firstId))
	d.object = C.to_GtkDialog(unsafe.Pointer(o))

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	}
	d.Window = newWindowFromNative(unsafe.Pointer(o)).(*Window)
	dialogFinalizer(d)
	d.AddButtons(butAndID[1:])

	return d
}

// Clear Dialog struct when it goes out of reach
func dialogFinalizer(d *Dialog) {
	runtime.SetFinalizer(d, func(d *Dialog) { gobject.Unref(d) })
}

// Conversion function for gobject registration map
func newDialogFromNative(obj unsafe.Pointer) interface{} {
	d := &Dialog{}
	d.object = C.to_GtkDialog(obj)

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	} else {
		gobject.Ref(d)
	}
	d.Window = newWindowFromNative(obj).(*Window)
	dialogFinalizer(d)

	return d
}

func nativeFromDialog(d interface{}) *gobject.GValue {
	dialog, ok := d.(*Dialog)
	if ok {
		gv := gobject.CreateCGValue(GtkType.DIALOG, dialog.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Dialog) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Dialog) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Dialog) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Dialog) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be window-like
func (self Dialog) Wnd() *Window {
	return self.Window
}

// Dialog interface
func (self *Dialog) Run() int {
	response := C.gtk_dialog_run(self.object)
	return int(response)
}

func (self *Dialog) Response(responseId int) {
	C.gtk_dialog_response(self.object, C.gint(responseId))
}

func (self *Dialog) AddButton(buttonText string, responseId int) *Button {
	s := gobject.GString(buttonText)
	defer s.Free()
	b := C.gtk_dialog_add_button(self.object, (*C.gchar)(s.GetPtr()), C.gint(responseId))
	btn, err := gobject.ConvertToGo(unsafe.Pointer(&b))
	if err == nil {
		return btn.(*Button)
	}
	return nil
}

func (self *Dialog) AddButtons(buttonAndId B) {
	for _, btn := range buttonAndId {
		self.AddButton(btn.Text, btn.Response)
	}
}

func (self *Dialog) AddActionWidget(w WidgetLike, responseId int) {
	C.gtk_dialog_add_action_widget(self.object, w.W().object, C.gint(responseId))
}

func (self *Dialog) SetDefaultResponse(responseId int) {
	C.gtk_dialog_set_default_response(self.object, C.gint(responseId))
}

func (self *Dialog) SetResponseSensitive(responseId int, setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_dialog_set_response_sensitive(self.object, C.gint(responseId), *((*C.gboolean)(b.GetPtr())))
}

func (self *Dialog) GetResponseForWidget(w WidgetLike) int {
	i := C.gtk_dialog_get_response_for_widget(self.object, w.W().object)
	return int(i)
}

func (self *Dialog) GetWidgetForResponse(responseId int) WidgetLike {
	cw := C.gtk_dialog_get_widget_for_response(self.object, C.gint(responseId))
	if cw != nil {
		w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
		if err == nil {
			return w.(WidgetLike)
		}
	}
	return nil
}

func (self *Dialog) GetActionArea() WidgetLike {
	cw := C.gtk_dialog_get_action_area(self.object)
	if cw != nil {
		w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
		if err == nil {
			return w.(WidgetLike)
		}
	}
	return nil
}

func (self *Dialog) GetContentArea() WidgetLike {
	cw := C.gtk_dialog_get_content_area(self.object)
	if cw != nil {
		w, err := gobject.ConvertToGo(unsafe.Pointer(cw))
		if err == nil {
			return w.(WidgetLike)
		}
	}
	return nil
}
//////////////////////////////
// END GtkDialog
////////////////////////////// }}}

// GtkMessageDialog {{{
//////////////////////////////

// GtkMessageDialog type
type MessageDialog struct {
	object *C.GtkMessageDialog
	*Dialog
}

func NewMessageDialog(parent *Window, flags int, mtype int, buttons int, messageFormat string, argsToMessage ...interface{}) *MessageDialog {
	mdiag := &MessageDialog{}
	var message string = messageFormat

	// If additional args, then call Sprintf with those args
	// and save formated message
	if len(argsToMessage) > 0 {
		message = fmt.Sprintf(messageFormat, argsToMessage...)
	}

	cmessage := gobject.GString(message)
	defer cmessage.Free()

	// If parent is not nil we will feed the parent object
	var pwin *C.GtkWindow = nil
	if parent != nil {
		pwin = parent.object
	}

	o := C._new_message_dialog(pwin, C.GtkDialogFlags(flags), C.GtkMessageType(mtype), C.GtkButtonsType(buttons),
		(*C.gchar)(cmessage.GetPtr()))
	mdiag.object = C.to_GtkMessageDialog(unsafe.Pointer(o))
	if gobject.IsObjectFloating(mdiag) {
		gobject.RefSink(mdiag)
	}
	mdiag.Dialog = newDialogFromNative(unsafe.Pointer(o)).(*Dialog)
	messageDialogFinalizer(mdiag)

	return mdiag
}

// Clear MessageDialog struct when it goes out of reach
func messageDialogFinalizer(d *MessageDialog) {
	runtime.SetFinalizer(d, func(d *MessageDialog) { gobject.Unref(d) })
}

// Conversion function for gobject registration map
func newMessageDialogFromNative(obj unsafe.Pointer) interface{} {
	d := &MessageDialog{}
	d.object = C.to_GtkMessageDialog(obj)

	if gobject.IsObjectFloating(d) {
		gobject.RefSink(d)
	} else {
		gobject.Ref(d)
	}
	d.Dialog = newDialogFromNative(obj).(*Dialog)
	messageDialogFinalizer(d)
	return d
}

func nativeFromMessageDialog(d interface{}) *gobject.GValue {
	dialog, ok := d.(*MessageDialog)
	if ok {
		gv := gobject.CreateCGValue(GtkType.MESSAGE_DIALOG, dialog.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self MessageDialog) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self MessageDialog) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self MessageDialog) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self MessageDialog) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be dialog-like
func (self MessageDialog) D() *Dialog {
	return self.Dialog
}

// Message dialog interface
func (self *MessageDialog) SetMarkup(markup string) {
	m := gobject.GString(markup)
	defer m.Free()
	C.gtk_message_dialog_set_markup(self.object, (*C.gchar)(m.GetPtr()))
}

func (self *MessageDialog) SetImage(w WidgetLike) {
	C.gtk_message_dialog_set_image(self.object, w.W().object)
}

func (self *MessageDialog) GetImage() WidgetLike {
	im := C.gtk_message_dialog_get_image(self.object)
	if im != nil {
		i, err := gobject.ConvertToGo(unsafe.Pointer(im))
		if err == nil {
			return i.(WidgetLike)
		}
	}
	return nil
}

func (self *MessageDialog) GetMessageArea() WidgetLike {
	ma := C.gtk_message_dialog_get_message_area(self.object)
	if ma != nil {
		m, err := gobject.ConvertToGo(unsafe.Pointer(ma))
		if err == nil {
			return m.(WidgetLike)
		}
	}
	return nil
}

func (self *MessageDialog) FormatSecondaryText(messageFormat string, args ...interface{}) {
	var message string = messageFormat
	if len(args) > 0 {
		message = fmt.Sprintf(messageFormat, args...)
	}

	cmessage := gobject.GString(message)
	defer cmessage.Free()

	C._gtk_message_dialog_format_secondary_text(self.object, (*C.gchar)(cmessage.GetPtr()))
}

func (self *MessageDialog) FormatSecondaryMarkup(format string, args ...interface{}) {
	var message string = format
	if len(args) > 0 {
		message = fmt.Sprintf(format, args...)
	}
	cmessage := gobject.GString(message)
	defer cmessage.Free()

	C._gtk_message_dialog_format_secondary_markup(self.object, (*C.gchar)(cmessage.GetPtr()))
}
//////////////////////////////
// END GtkMessageDialog
////////////////////////////// }}}

// GtkSeparator {{{
//////////////////////////////

// GtkSeparator type
type Separator struct {
	object *C.GtkSeparator
	*Widget
}

func NewSeparator(orientation int) *Separator {
	sep := &Separator{}
	o := C.gtk_separator_new(C.GtkOrientation(orientation))
	sep.object = C.to_GtkSeparator(unsafe.Pointer(o))

	if gobject.IsObjectFloating(sep) {
		gobject.RefSink(sep)
	}
	sep.Widget = NewWidget(unsafe.Pointer(o))
	separatorFinalizer(sep)

	return sep
}

func NewHSeparator() *Separator {
	return NewSeparator(GtkOrientation.HORIZONTAL)
}

func NewVSeparator() *Separator {
	return NewSeparator(GtkOrientation.VERTICAL)
}

// Clear Separator struct when it goes out of reach
func separatorFinalizer(s *Separator) {
	runtime.SetFinalizer(s, func(s *Separator) { gobject.Unref(s) })
}

// Conversion function for gobject registration map
func newSeparatorFromNative(obj unsafe.Pointer) interface{} {
	sep := &Separator{}
	sep.object = C.to_GtkSeparator(obj)

	if gobject.IsObjectFloating(sep) {
		gobject.RefSink(sep)
	} else {
		gobject.Ref(sep)
	}
	sep.Widget = NewWidget(obj)
	separatorFinalizer(sep)

	return sep
}

func nativeFromSeparator(s interface{}) *gobject.GValue {
	sep, ok := s.(*Separator)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SEPARATOR, sep.ToNative())
		return gv
	}

	return nil
}

// To be object like
func (self Separator) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Separator) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Separator) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self Separator) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
// To be widget-like
func (self Separator) W() *Widget {
	return self.Widget
}
//////////////////////////////
// END GtkSeparator
////////////////////////////// }}}

// GtkAdjustment {{{
//////////////////////////////

// GtkAdjustment type
type Adjustment struct {
	object *C.GtkAdjustment
}

func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) *Adjustment {
	a := &Adjustment{}
	o := C.gtk_adjustment_new(C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(stepIncrement),
		C.gdouble(pageIncrement), C.gdouble(pageSize))
	a.object = C.to_GtkAdjustment(unsafe.Pointer(o))

	if gobject.IsObjectFloating(a) {
		gobject.RefSink(a)
	}
	adjustmentFinalizer(a)

	return a
}

// Clear Adjustment struct when it goes out of reach
func adjustmentFinalizer(a *Adjustment) {
	runtime.SetFinalizer(a, func(a *Adjustment) { gobject.Unref(a) })
}

// Conversion functions for gobject registration map
func newAdjustmentFromNative(obj unsafe.Pointer) interface{} {
	a := &Adjustment{}
	a.object = C.to_GtkAdjustment(obj)

	if gobject.IsObjectFloating(a) {
		gobject.RefSink(a)
	} else {
		gobject.Ref(a)
	}
	adjustmentFinalizer(a)

	return a
}

func nativeFromAdjustment(a interface{}) *gobject.GValue {
	adj, ok := a.(*Adjustment)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ADJUSTMENT, adj.ToNative())
		return gv
	}

	return nil
}

// To be object-like
func (self Adjustment) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Adjustment) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Adjustment) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Adjustment) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Adjustment interface
func (self *Adjustment) GetValue() float64 {
	return float64(C.gtk_adjustment_get_value(self.object))
}

func (self *Adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(self.object, C.gdouble(value))
}

func (self *Adjustment) ClampPage(lower, upper float64) {
	C.gtk_adjustment_clamp_page(self.object, C.gdouble(lower), C.gdouble(upper))
}

func (self *Adjustment) Changed() {
	C.gtk_adjustment_changed(self.object)
}

func (self *Adjustment) ValueChanged() {
	C.gtk_adjustment_value_changed(self.object)
}

func (self *Adjustment) Configure(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) {
	C.gtk_adjustment_configure(self.object, C.gdouble(value), C.gdouble(lower), C.gdouble(upper),
		C.gdouble(stepIncrement), C.gdouble(pageIncrement), C.gdouble(pageSize))
}

func (self *Adjustment) GetLower() float64 {
	return float64(C.gtk_adjustment_get_lower(self.object))
}

func (self *Adjustment) GetPageIncrement() float64 {
	return float64(C.gtk_adjustment_get_page_increment(self.object))
}

func (self *Adjustment) GetPageSize() float64 {
	return float64(C.gtk_adjustment_get_page_size(self.object))
}

func (self *Adjustment) GetStepIncrement() float64 {
	return float64(C.gtk_adjustment_get_step_increment(self.object))
}

func (self *Adjustment) GetMinimumIncrement() float64 {
	return float64(C.gtk_adjustment_get_minimum_increment(self.object))
}

func (self *Adjustment) GetUpper() float64 {
	return float64(C.gtk_adjustment_get_upper(self.object))
}

func (self *Adjustment) SetLower(lower float64) {
	C.gtk_adjustment_set_lower(self.object, C.gdouble(lower))
}

func (self *Adjustment) SetPageIncrement(pageIncrement float64) {
	C.gtk_adjustment_set_page_increment(self.object, C.gdouble(pageIncrement))
}

func (self *Adjustment) SetPageSize(pageSize float64) {
	C.gtk_adjustment_set_page_size(self.object, C.gdouble(pageSize))
}

func (self *Adjustment) SetStepIncrement(stepIncrement float64) {
	C.gtk_adjustment_set_step_increment(self.object, C.gdouble(stepIncrement))
}

func (self *Adjustment) SetUpper(upper float64) {
	C.gtk_adjustment_set_upper(self.object, C.gdouble(upper))
}
//////////////////////////////
// GtkAdjustment
////////////////////////////// }}}

// GtkRange {{{
//////////////////////////////

// GtkRange type
type Range struct {
	object *C.GtkRange
	*Widget
}

// Clear Range struct when it goes out of reach
func rangeFinalizer(r *Range) {
	runtime.SetFinalizer(r, func(r *Range) { gobject.Unref(r) })
}

// Conversion functions for gobject registration map
func newRangeFromNative(obj unsafe.Pointer) interface{} {
	r := &Range{}
	r.object = C.to_GtkRange(obj)

	if gobject.IsObjectFloating(r) {
		gobject.RefSink(r)
	} else {
		gobject.Ref(r)
	}
	r.Widget = NewWidget(obj)
	rangeFinalizer(r)

	return r
}

func nativeFromRange(r interface{}) *gobject.GValue {
	ran, ok := r.(*Range)
	if ok {
		gv := gobject.CreateCGValue(GtkType.RANGE, ran.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Range) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Range) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Range) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Range) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Range) W() *Widget {
	return self.Widget
}

// Range interface
func (self *Range) GetFillLevel() float64 {
	return float64(C.gtk_range_get_fill_level(self.object))
}

func (self *Range) GetRestrictToFillLevel() bool {
	b := C.gtk_range_get_restrict_to_fill_level(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) GetShowFillLevel() bool {
	b := C.gtk_range_get_show_fill_level(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) SetFillLevel(fillLevel float64) {
	C.gtk_range_set_fill_level(self.object, C.gdouble(fillLevel))
}

func (self *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	b := gobject.GBool(restrictToFillLevel)
	defer b.Free()
	C.gtk_range_set_restrict_to_fill_level(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Range) SetShowFillLevel(showFillLevel bool) {
	b := gobject.GBool(showFillLevel)
	defer b.Free()
	C.gtk_range_set_show_fill_level(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Range) GetAdjustment() *Adjustment {
	a := C.gtk_range_get_adjustment(self.object)
	adj, err := gobject.ConvertToGo(unsafe.Pointer(&a))
	if err == nil {
		return adj.(*Adjustment)
	}
	return nil
}

func (self *Range) SetAdjustment(adjustment *Adjustment) {
	C.gtk_range_set_adjustment(self.object, adjustment.object)
}

func (self *Range) GetInverted() bool {
	b := C.gtk_range_get_inverted(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) SetInverted(inverted bool) {
	b := gobject.GBool(inverted)
	C.gtk_range_set_inverted(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Range) GetValue() float64 {
	return float64(C.gtk_range_get_value(self.object))
}

func (self *Range) SetValue(value float64) {
	C.gtk_range_set_value(self.object, C.gdouble(value))
}

func (self *Range) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(self.object, C.gdouble(step), C.gdouble(page))
}

func (self *Range) SetRange(min, max float64) {
	C.gtk_range_set_range(self.object, C.gdouble(min), C.gdouble(max))
}

func (self *Range) GetRoundDigits() int {
	return int(C.gtk_range_get_round_digits(self.object))
}

func (self *Range) SetRoundDigits(roundDigits int) {
	C.gtk_range_set_round_digits(self.object, C.gint(roundDigits))
}

func (self *Range) SetLowerStepperSensitivity(gtk_sensitivityType int) {
	C.gtk_range_set_lower_stepper_sensitivity(self.object, C.GtkSensitivityType(gtk_sensitivityType))
}

func (self *Range) GetLowerStepperSensitivity() int {
	return int(C.gtk_range_get_lower_stepper_sensitivity(self.object))
}

func (self *Range) SetUpperStepperSensitivity(gtk_sensitivityType int) {
	C.gtk_range_set_upper_stepper_sensitivity(self.object, C.GtkSensitivityType(gtk_sensitivityType))
}

func (self *Range) GetUpperStepperSensitivity() int {
	return int(C.gtk_range_get_upper_stepper_sensitivity(self.object))
}

func (self *Range) GetFlippable() bool {
	b := C.gtk_range_get_flippable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) SetFlippable(flippable bool) {
	b := gobject.GBool(flippable)
	defer b.Free()
	C.gtk_range_set_flippable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Range) GetMinSliderSize() int {
	return int(C.gtk_range_get_min_slider_size(self.object))
}

//TODO: gtk_range_get_range_rect

func (self *Range) GetSliderRange() (start, end int) {
	var s C.gint
	var e C.gint
	C.gtk_range_get_slider_range(self.object, &s, &e)
	start = int(s)
	end = int(e)
	return
}

func (self *Range) GetSliderSizeFixed() bool {
	b := C.gtk_range_get_slider_size_fixed(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Range) SetMinSliderSize(minSize int) {
	C.gtk_range_set_min_slider_size(self.object, C.gint(minSize))
}

func (self *Range) SetSliderSizeFixed(sizeFixed bool) {
	b := gobject.GBool(sizeFixed)
	defer b.Free()
	C.gtk_range_set_slider_size_fixed(self.object, *((*C.gboolean)(b.GetPtr())))
}
//////////////////////////////
// END GtkRange
////////////////////////////// }}}

// GtkScrollbar {{{
//////////////////////////////

// GtkScrollbar type
type Scrollbar struct {
	object *C.GtkScrollbar
	*Range
}

func NewScrollbar(gtk_orientation int, adjustment *Adjustment) *Scrollbar {
	sb := &Scrollbar{}
	o := C.gtk_scrollbar_new(C.GtkOrientation(gtk_orientation), adjustment.object)
	sb.object = C.to_GtkScrollbar(unsafe.Pointer(o))

	if gobject.IsObjectFloating(sb) {
		gobject.RefSink(sb)
	}
	sb.Range = newRangeFromNative(unsafe.Pointer(o)).(*Range)
	scrollbarFinalizer(sb)

	return sb
}

func NewHScrollbar(adjustment *Adjustment) *Scrollbar {
	return NewScrollbar(GtkOrientation.HORIZONTAL, adjustment)
}

func NewVScrollbar(adjustment *Adjustment) *Scrollbar {
	return NewScrollbar(GtkOrientation.VERTICAL, adjustment)
}

// Clear Scrollbar struct when it goes out of reach
func scrollbarFinalizer(sb *Scrollbar) {
	runtime.SetFinalizer(sb, func(sb *Scrollbar) { gobject.Unref(sb) })
}

// Conversions functions
func newScrollbarFromNative(sb unsafe.Pointer) interface{} {
	scrollbar := &Scrollbar{}
	scrollbar.object = C.to_GtkScrollbar(sb)

	if gobject.IsObjectFloating(scrollbar) {
		gobject.RefSink(scrollbar)
	} else {
		gobject.Ref(scrollbar)
	}
	scrollbar.Range = newRangeFromNative(sb).(*Range)
	scrollbarFinalizer(scrollbar)

	return scrollbar
}

func nativeFromScrollbar(sb interface{}) *gobject.GValue {
	scrollbar, ok := sb.(*Scrollbar)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SCROLLBAR, scrollbar.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Scrollbar) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Scrollbar) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Scrollbar) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Scrollbar) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be Range-like
func (self Scrollbar) R() *Range {
	return self.Range
}
//////////////////////////////
// END GtkScrollbar
////////////////////////////// }}}

// GtkScrolledWindow {{{
//////////////////////////////

// GtkScrolledWindow type
type ScrolledWindow struct {
	object *C.GtkScrolledWindow
	*Container
}

func NewScrolledWindow(hadjustment, vadjustment *Adjustment) *ScrolledWindow {
	sw := &ScrolledWindow{}
	var ha, va *C.GtkAdjustment
	ha, va = nil, nil

	if hadjustment != nil {
		ha = hadjustment.object
	}

	if vadjustment != nil {
		va = vadjustment.object
	}

	o := C.gtk_scrolled_window_new(ha, va)
	sw.object = C.to_GtkScrolledWindow(unsafe.Pointer(o))

	if gobject.IsObjectFloating(sw) {
		gobject.RefSink(sw)
	}
	sw.Container = NewContainer(unsafe.Pointer(o))
	scrolledWindowFinalizer(sw)

	return sw
}

// Clear ScrolledWindow struct when it goes out of reach
func scrolledWindowFinalizer(sw *ScrolledWindow) {
	runtime.SetFinalizer(sw, func(sw *ScrolledWindow) { gobject.Unref(sw) })
}

// Conversion functions
func newScrolledWindowFromNative(sw unsafe.Pointer) interface{} {
	scrolled := &ScrolledWindow{}
	scrolled.object = C.to_GtkScrolledWindow(sw)

	if gobject.IsObjectFloating(scrolled) {
		gobject.RefSink(scrolled)
	} else {
		gobject.Ref(scrolled)
	}
	scrolled.Container = NewContainer(sw)
	scrolledWindowFinalizer(scrolled)

	return scrolled
}

func nativeFromScrolledWindow(sw interface{}) *gobject.GValue {
	scrolled, ok := sw.(*ScrolledWindow)
	if ok {
		gv := gobject.CreateCGValue(GtkType.SCROLLED_WINDOW, scrolled.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ScrolledWindow) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ScrolledWindow) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ScrolledWindow) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ScrolledWindow) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// ScrolledWindow interface
func (self *ScrolledWindow) GetHadjustment() *Adjustment {
	ad := C.gtk_scrolled_window_get_hadjustment(self.object)
	a, err := gobject.ConvertToGo(unsafe.Pointer(ad))
	if err == nil {
		return a.(*Adjustment)
	}
	return nil
}

func (self *ScrolledWindow) GetVadjustment() *Adjustment {
	ad := C.gtk_scrolled_window_get_vadjustment(self.object)
	a, err := gobject.ConvertToGo(unsafe.Pointer(ad))
	if err == nil {
		return a.(*Adjustment)
	}
	return nil
}

func (self *ScrolledWindow) GetHScrollbar() WidgetLike {
	ad := C.gtk_scrolled_window_get_hscrollbar(self.object)
	a, err := gobject.ConvertToGo(unsafe.Pointer(ad))
	if err == nil {
		return a.(WidgetLike)
	}
	return nil
}

func (self *ScrolledWindow) GetVScrollbar() WidgetLike {
	ad := C.gtk_scrolled_window_get_vscrollbar(self.object)
	a, err := gobject.ConvertToGo(unsafe.Pointer(ad))
	if err == nil {
		return a.(WidgetLike)
	}
	return nil
}

func (self *ScrolledWindow) SetPolicy(gtk_policy_horizontal, gtk_policy_vertical int) {
	C.gtk_scrolled_window_set_policy(self.object, C.GtkPolicyType(gtk_policy_horizontal), C.GtkPolicyType(gtk_policy_vertical))
}

func (self *ScrolledWindow) AddWithViewport(child WidgetLike) {
	C.gtk_scrolled_window_add_with_viewport(self.object, child.W().object)
}

func (self *ScrolledWindow) SetPlacement(gtk_corner_window_placement int) {
	C.gtk_scrolled_window_set_placement(self.object, C.GtkCornerType(gtk_corner_window_placement))
}

func (self *ScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement(self.object)
}

func (self *ScrolledWindow) SetShadowType(gtk_shadow int) {
	C.gtk_scrolled_window_set_shadow_type(self.object, C.GtkShadowType(gtk_shadow))
}

func (self *ScrolledWindow) SetHadjustment(ha *Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(self.object, ha.object)
}

func (self *ScrolledWindow) SetVadjustment(va *Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(self.object, va.object)
}

func (self *ScrolledWindow) GetPlacement() int {
	return int(C.gtk_scrolled_window_get_placement(self.object))
}

func (self *ScrolledWindow) GetPolicy() (hscrollbarPolicy, vscrollbarPolicy int) {
	var hp C.GtkPolicyType
	var vp C.GtkPolicyType
	C.gtk_scrolled_window_get_policy(self.object, &hp, &vp)
	return int(hp), int(vp)
}

func (self *ScrolledWindow) GetShadowType() int {
	return int(C.gtk_scrolled_window_get_shadow_type(self.object))
}

func (self *ScrolledWindow) GetMinContentWidth() int {
	return int(C.gtk_scrolled_window_get_min_content_width(self.object))
}

func (self *ScrolledWindow) SetMinContentWidth(width int) {
	C.gtk_scrolled_window_set_min_content_width(self.object, C.gint(width))
}

func (self *ScrolledWindow) GetMinContentHeight() int {
	return int(C.gtk_scrolled_window_get_min_content_height(self.object))
}

func (self *ScrolledWindow) SetMinContentHeight(height int) {
	C.gtk_scrolled_window_set_min_content_height(self.object, C.gint(height))
}
//////////////////////////////
// END GtkScrolledWindow
////////////////////////////// }}}

// GtkTreePath {{{
//////////////////////////////

// GtkTreePath type
type TreePath struct {
	object *C.GtkTreePath
}

func NewTreePath() *TreePath {
	tp := &TreePath{}
	tp.object = C.gtk_tree_path_new()
	treePathFinalizer(tp)

	return tp
}

func NewTreePathFromString(path string) *TreePath {
	tp := &TreePath{}

	s := gobject.GString(path)
	defer s.Free()

	tp.object = C.gtk_tree_path_new_from_string((*C.gchar)(s.GetPtr()))
	treePathFinalizer(tp)

	return tp
}

func NewTreePathFirst() *TreePath {
	tp := &TreePath{}
	tp.object = C.gtk_tree_path_new_first()
	treePathFinalizer(tp)

	return tp
}

//TODO: NewTreePathFromIndices


// Clear TreePath struct when it goes out of scope
func treePathFinalizer(tp *TreePath) {
	runtime.SetFinalizer(tp, func(tp *TreePath) { tp.free() })
}

// TreePath Interface
func (self *TreePath) free() {
	C.gtk_tree_path_free(self.object)
}

func (self *TreePath) ToString() string {
	s := C.gtk_tree_path_to_string(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TreePath) AppendIndex(index int) {
	C.gtk_tree_path_append_index(self.object, C.gint(index))
}

func (self *TreePath) PrependIndex(index int) {
	C.gtk_tree_path_prepend_index(self.object, C.gint(index))
}

func (self *TreePath) GetDepth() int {
	return int(C.gtk_tree_path_get_depth(self.object))
}

func (self *TreePath) GetIndices() []int {
	cind := C.gtk_tree_path_get_indices(self.object)
	depth := self.GetDepth()
	gind := make([]int, depth)

	for i := 0; i < depth; i++ {
		indice := int(C._gtk_tree_path_get_indice(cind, C.int(i)))
		gind[i] = indice
	}

	return gind
}

func (self *TreePath) Copy() *TreePath {
	tp := &TreePath{}
	tp.object = C.gtk_tree_path_copy(self.object)
	treePathFinalizer(tp)

	return tp
}

func (self *TreePath) Compare(tp *TreePath) int {
	return int(C.gtk_tree_path_compare(self.object, tp.object))
}

func (self *TreePath) Next() {
	C.gtk_tree_path_next(self.object)
}

func (self *TreePath) Prev() bool {
	b := C.gtk_tree_path_prev(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreePath) Up() bool {
	b := C.gtk_tree_path_up(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreePath) Down() {
	C.gtk_tree_path_down(self.object)
}

func (self *TreePath) IsAncestor(descendant *TreePath) bool {
	b := C.gtk_tree_path_is_ancestor(self.object, descendant.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreePath) IsDescendant(ancestor *TreePath) bool {
	b := C.gtk_tree_path_is_descendant(self.object, ancestor.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}
//////////////////////////////
// END GtkTreePath
////////////////////////////// }}}



// GTK3 MODULE init function {{{
func init() {
	// Register GtkApplicaton type
	gobject.RegisterCType(GtkType.APPLICATION, appFromNative)
	gobject.RegisterGoType(GtkType.APPLICATION, nativeFromApp)

	// Register GtkWindow type
	gobject.RegisterCType(GtkType.WINDOW, newWindowFromNative)
	gobject.RegisterGoType(GtkType.WINDOW, nativeFromWindow)

	// Register GtkBox type
	gobject.RegisterCType(GtkType.BOX, newBoxFromNative)
	gobject.RegisterGoType(GtkType.BOX, nativeFromBox)

	// Register GtkButtonBox type
	gobject.RegisterCType(GtkType.BUTTON_BOX, newButtonBoxFromNative)
	gobject.RegisterGoType(GtkType.BUTTON_BOX, nativeFromButtonBox)

	// Register GtkFrame type
	gobject.RegisterCType(GtkType.FRAME, newFrameFromNative)
	gobject.RegisterGoType(GtkType.FRAME, nativeFromFrame)

	// Register GtkGrid type
	gobject.RegisterCType(GtkType.GRID, newGridFromNative)
	gobject.RegisterGoType(GtkType.GRID, nativeFromGrid)

	// Register GtkLabel type
	gobject.RegisterCType(GtkType.LABEL, newLabelFromNative)
	gobject.RegisterGoType(GtkType.LABEL, nativeFromLabel)

	// Register GtkImage type
	gobject.RegisterCType(GtkType.IMAGE, newImageFromNative)
	gobject.RegisterGoType(GtkType.IMAGE, nativeFromImage)

	// Register GtkButton type
	gobject.RegisterCType(GtkType.BUTTON, newButtonFromNative)
	gobject.RegisterGoType(GtkType.BUTTON, nativeFromButton)

	//Register GtkToggleButton type
	gobject.RegisterCType(GtkType.TOGGLE_BUTTON, newToggleButtonFromNative)
	gobject.RegisterGoType(GtkType.TOGGLE_BUTTON, nativeFromToggleButton)

	// Register GtkEntryBuffer type
	gobject.RegisterCType(GtkType.ENTRY_BUFFER, newEntryBufferFromNative)
	gobject.RegisterGoType(GtkType.ENTRY_BUFFER, nativeFromEntryBuffer)

	// Register GtkEntry type
	gobject.RegisterCType(GtkType.ENTRY, newEntryFromNative)
	gobject.RegisterGoType(GtkType.ENTRY, nativeFromEntry)

	// Register GtkDialog type
	gobject.RegisterCType(GtkType.DIALOG, newDialogFromNative)
	gobject.RegisterGoType(GtkType.DIALOG, nativeFromDialog)

	// Register GtkMessageDialog type
	gobject.RegisterCType(GtkType.MESSAGE_DIALOG, newMessageDialogFromNative)
	gobject.RegisterGoType(GtkType.MESSAGE_DIALOG, nativeFromMessageDialog)

	// Register GtkSeparator type
	gobject.RegisterCType(GtkType.SEPARATOR, newSeparatorFromNative)
	gobject.RegisterGoType(GtkType.SEPARATOR, nativeFromSeparator)

	// Register GtkAdjustment type
	gobject.RegisterCType(GtkType.ADJUSTMENT, newAdjustmentFromNative)
	gobject.RegisterGoType(GtkType.ADJUSTMENT, nativeFromAdjustment)

	// Register GtkRange type
	gobject.RegisterCType(GtkType.RANGE, newRangeFromNative)
	gobject.RegisterGoType(GtkType.RANGE, nativeFromRange)

	// Register GtkScrollbar type
	gobject.RegisterCType(GtkType.SCROLLBAR, newScrollbarFromNative)
	gobject.RegisterGoType(GtkType.SCROLLBAR, nativeFromScrollbar)

	// Register GtkScrolledWindow type
	gobject.RegisterCType(GtkType.SCROLLED_WINDOW, newScrolledWindowFromNative)
	gobject.RegisterGoType(GtkType.SCROLLED_WINDOW, nativeFromScrolledWindow)
}
// End init function }}}
