package gtk3

/*
#include <gtk/gtk.h>
#include <stdlib.h>

// Exported funcs from our module {{{
extern void _gtk_callback(GtkWidget* widget, gpointer data);
extern gboolean _gtk_entry_completion_match_func(GtkEntryCompletion* completion,
											 const gchar* key,
											 GtkTreeIter* iter,
											 gpointer user_data);
extern gboolean _gtk_cell_callback(GtkCellRenderer* renderer, gpointer data);

// End Exported funcs }}}

static void free_id(gpointer data) {
	free(data);
}

static void _gtk_init(void* argc, void* argv) {
	gtk_init((int*)argc, (char***)argv);
}


// to_Gtk*** Funcs {{{
static inline GApplication* to_GApplication(GtkApplication* g) { return G_APPLICATION(g); }
static inline GtkWidget* to_GtkWidget(void* obj) { return GTK_WIDGET(obj); }
static inline GtkContainer* to_GtkContainer(void *obj) { return GTK_CONTAINER(obj); }
static inline GtkWindow* to_GtkWindow(void* obj) { return GTK_WINDOW(obj); }
static inline GtkAssistant* to_GtkAssistant(void* obj) { return GTK_ASSISTANT(obj); }
static inline GtkBox* to_GtkBox(void* obj) { return GTK_BOX(obj); }
static inline GtkButtonBox* to_GtkButtonBox(void* obj) { return GTK_BUTTON_BOX(obj); }
static inline GtkFrame* to_GtkFrame(void* obj) { return GTK_FRAME(obj); }
static inline GtkGrid* to_GtkGrid(void* obj) { return GTK_GRID(obj); }
static inline GtkLabel* to_GtkLabel(void* obj) { return GTK_LABEL(obj); }
static inline GtkProgressBar* to_GtkProgressBar(void* obj) { return GTK_PROGRESS_BAR(obj); }
static inline GtkImage* to_GtkImage(void* obj) { return GTK_IMAGE(obj); }
static inline GtkButton* to_GtkButton(void* obj) { return GTK_BUTTON(obj); }
static inline GtkToggleButton* to_GtkToggleButton(void* obj) { return GTK_TOGGLE_BUTTON(obj); }
static inline GtkCheckButton* to_GtkCheckButton(void* obj) { return GTK_CHECK_BUTTON(obj); }
static inline GtkEntryBuffer* to_GtkEntryBuffer(void* obj) { return GTK_ENTRY_BUFFER(obj); }
static inline GtkEntry* to_GtkEntry(void* obj) { return GTK_ENTRY(obj); }
static inline GtkEntryCompletion* to_GtkEntryCompletion(void* obj) { return GTK_ENTRY_COMPLETION(obj); }
static inline GtkDialog* to_GtkDialog(void* obj) { return GTK_DIALOG(obj); }
static inline GtkMessageDialog* to_GtkMessageDialog(void* obj) { return GTK_MESSAGE_DIALOG(obj); }
static inline GtkInvisible* to_GtkInvisible(void* obj) { return GTK_INVISIBLE(obj); }
static inline GtkSeparator* to_GtkSeparator(void* obj) { return GTK_SEPARATOR(obj); }
static inline GtkAdjustment* to_GtkAdjustment(void* obj) { return GTK_ADJUSTMENT(obj); }
static inline GtkRange* to_GtkRange(void* obj) { return GTK_RANGE(obj); }
static inline GtkScrollbar* to_GtkScrollbar(void* obj) { return GTK_SCROLLBAR(obj); }
static inline GtkScrolledWindow* to_GtkScrolledWindow(void* obj) { return GTK_SCROLLED_WINDOW(obj); }
static inline GtkTreeModel* to_GtkTreeModel(void* obj) { return GTK_TREE_MODEL(obj); }
static inline GtkListStore* to_GtkListStore(void* obj) { return GTK_LIST_STORE(obj); }
static inline GtkTreeStore* to_GtkTreeStore(void* obj) { return GTK_TREE_STORE(obj); }
static inline GtkCellArea* to_GtkCellArea(void* obj) { return GTK_CELL_AREA(obj); }
static inline GtkCellAreaContext* to_GtkCellAreaContext(void* obj) { return GTK_CELL_AREA_CONTEXT(obj); }
static inline GtkCellRenderer* to_GtkCellRenderer(void* obj) { return GTK_CELL_RENDERER(obj); }
static inline GtkCellRendererText* to_GtkCellRendererText(void* obj) { return GTK_CELL_RENDERER_TEXT(obj); }
static inline GtkCellRendererProgress* to_GtkCellRendererProgress(void* obj) { return GTK_CELL_RENDERER_PROGRESS(obj); }
static inline GtkCellRendererSpinner* to_GtkCellRendererSpinner(void* obj) { return GTK_CELL_RENDERER_SPINNER(obj); }
static inline GtkCellRendererToggle* to_GtkCellRendererToggle(void* obj) { return GTK_CELL_RENDERER_TOGGLE(obj); }
static inline GtkCellRendererPixbuf* to_GtkCellRendererPixbuf(void* obj) { return GTK_CELL_RENDERER_PIXBUF(obj); }
static inline GtkCellRendererAccel* to_GtkCellRendererAccel(void* obj) { return GTK_CELL_RENDERER_ACCEL(obj); }
static inline GtkCellRendererCombo* to_GtkCellRendererCombo(void* obj) { return GTK_CELL_RENDERER_COMBO(obj); }
static inline GtkCellRendererSpin* to_GtkCellRendererSpin(void* obj) { return GTK_CELL_RENDERER_SPIN(obj); }
static inline GtkTreeViewColumn* to_GtkTreeViewColumn(void* obj) { return GTK_TREE_VIEW_COLUMN(obj); }
static inline GtkTreeView* to_GtkTreeView(void* obj) { return GTK_TREE_VIEW(obj); }
static inline GtkTreeSelection* to_GtkTreeSelection(void* obj) { return GTK_TREE_SELECTION(obj); }
static inline GtkNotebook* to_GtkNotebook(void* obj) { return GTK_NOTEBOOK(obj); }
// End }}}

// GtkContainer funcs {{{
static void _gtk_container_foreach(GtkContainer* container, gint64* id) {
	gtk_container_foreach(container, _gtk_callback, (gpointer)id);
}

//End GtkContainer funcs }}}

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
	return FALSE;
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

// GtkEntryCompletion funcs {{{

void _gtk_entry_completion_set_match_func(GtkEntryCompletion* completion, gint64 func_data) {
	gint64* data = (gint64*)malloc(sizeof(gint64));
	*data = func_data;

	gtk_entry_completion_set_match_func(completion, _gtk_entry_completion_match_func, (gpointer)data, free_id);
}

// End GtkEntryCompletion funcs }}}

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

// GtkAdjustment funcs {{{
static gdouble _gtk_adjustment_get_minimum_increment(GtkAdjustment* adjustment) {
#if GTK_CHECK_VERSION(3,2,0)
    return gtk_adjustment_get_minimum_increment(adjustment);
#else
    return 0;
#endif
}
// End GtkAdjustment funcs }}}

// GtkTreePath funcs {{{
static inline gint _gtk_tree_path_get_indice(gint* indices, int index) {
	return *(indices + index);
}
//End GtkTreePath func }}}

// GtkCellArea funcs {{{
static void _gtk_cell_area_foreach(GtkCellArea* area, gint64 id) {
	gint64* uid = &id;
	gtk_cell_area_foreach(area, _gtk_cell_callback, (gpointer)uid);
}

//End GtkCellArea funcs }}}

// Gtk*Store funcs {{{
static inline GType* gtype_array_new(int n) {
	return g_new0(GType, n);
}

static inline void gtype_array_free(GType* types) {
	g_free(types);
}

static inline void gtype_array_set_element(GType* types, int n, GType type) {
	types[n] = type;
}

static inline GValue* gvalue_array_new(int n) {
	return g_new0(GValue, n);
}

static inline void gvalue_array_free(GValue* vals) {
	g_free(vals);
}

static inline void gvalue_array_set_element(GValue* vals, int n, GValue* val) {
	vals[n] = *val;
}

static inline gint* gint_array_new(int n) {
	return g_new0(gint, n);
}

static inline void gint_array_free(gint* vals) {
	g_free(vals);
}

static inline void gint_array_set_element(gint* vals, int n, gint val) {
	vals[n] = val;
}

//End Gtk*Store func }}}

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

// Map variable for custom closures required by some gtk funcs
var _closures map[int64]gobject.ClosureFunc

// Convenient map for properties
type P map[string]interface{}

// Convenient struct and slice for buttons and id's (See Dialog)
type ButID struct {
	Text     string
	Response int
}
type B []ButID

// Convenient map for Columns and values (See ListStore, TreeStore)
type V map[int]interface{}

// Convenient struct and slice for TreeViewColumn attributes
type CollAttr struct {
	Attribute string
	Column    int
}
type A []CollAttr

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

// TreeModel like interface must have method ITreeModel()
type TreeModelLike interface {
	ITreeModel() *TreeModel
}

// CellRendererLike interface must have method CRenderer()
type CellRendererLike interface {
	CRenderer() *CellRenderer
}
//////////////////////////////
// END Interfaces
////////////////////////////// }}}

// Base Structs {{{

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

func newWidgetFromNative(obj unsafe.Pointer) interface{} {
	w := &Widget{}
	w.object = C.to_GtkWidget(obj)

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

// Widget inteface

func (self *Widget) Destroy() {
	C.gtk_widget_destroy(self.object)
	runtime.SetFinalizer(self, nil)
}

func (self *Widget) InDestruction() bool {
	b := C.gtk_widget_in_destruction(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Show() {
	C.gtk_widget_show(self.object)
}

func (self *Widget) ShowNow() {
	C.gtk_widget_show_now(self.object)
}

func (self *Widget) Hide() {
	C.gtk_widget_hide(self.object)
}

func (self *Widget) ShowAll() {
	C.gtk_widget_show_all(self.object)
}

//TODO: gtk_widget_draw

func (self *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(self.object)
}

//TODO: gtk_widget_add_accelerator
//TODO: gtk_widget_remove_accelerator
//TODO: gtk_widget_set_accel_path
//TODO: gtk_widget_list_accel_closures
//TODO: gtk_widget_can_activate_accel

func (self *Widget) Activate() bool {
	b := C.gtk_widget_activate(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) Reparent(new_parent WidgetLike) {
	C.gtk_widget_reparent(self.object, new_parent.W().object)
}

func (self *Widget) Intersect(area gdk3.Rectangle) (bool, gdk3.Rectangle) {
	var rec C.GdkRectangle
	cArea := gobject.ConvertToC(area)
	defer cArea.Free()

	b := C.gtk_widget_intersect(self.object, (*C.GdkRectangle)(cArea.GetPtr()), &rec)
	retBool := gobject.GoBool(unsafe.Pointer(&b))

	if rectangle, err := gobject.ConvertToGo(unsafe.Pointer(&rec), gdk3.GdkType.RECTANGLE); err == nil {
		return retBool, rectangle.(gdk3.Rectangle)
	}

	return retBool, gdk3.Rectangle{}
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

func (self *Widget) SetParentWindow(parentWindow *gdk3.Window) {
	C.gtk_widget_set_parent_window(self.object, (*C.GdkWindow)(parentWindow.ToNative()))
}

func (self *Widget) GetParentWindow() *gdk3.Window {
	pw := C.gtk_widget_get_parent_window(self.object)

	if parWindow, err := gobject.ConvertToGo(unsafe.Pointer(pw)); err == nil {
		return parWindow.(*gdk3.Window)
	}
	return nil
}

func (self *Widget) SetEvents(gdk_EventMask int) {
	C.gtk_widget_set_events(self.object, C.gint(gdk_EventMask))
}

func (self *Widget) GetEvents() int {
	return int(C.gtk_widget_get_events(self.object))
}

func (self *Widget) AddEvents(gdk_EventMask int) {
	C.gtk_widget_add_events(self.object, C.gint(gdk_EventMask))
}

func (self *Widget) SetDeviceEvents(device *gdk3.Device, gdk_EventMask int) {
	C.gtk_widget_set_device_events(self.object, (*C.GdkDevice)(device.ToNative()), C.GdkEventMask(gdk_EventMask))
}

func (self *Widget) GetDeviceEvents(device *gdk3.Device) int {
	return int(C.gtk_widget_get_device_events(self.object, (*C.GdkDevice)(device.ToNative())))
}

func (self *Widget) AddDeviceEvents(device *gdk3.Device, gdk_EventMask int) {
	C.gtk_widget_add_device_events(self.object, (*C.GdkDevice)(device.ToNative()), C.GdkEventMask(gdk_EventMask))
}

func (self *Widget) SetDeviceEnabled(device *gdk3.Device, enabled bool) {
	b := gobject.GBool(enabled)
	defer b.Free()
	C.gtk_widget_set_device_enabled(self.object, (*C.GdkDevice)(device.ToNative()), *((*C.gboolean)(b.GetPtr())))
}

func (self *Widget) GetDeviceEnabled(device *gdk3.Device) bool {
	b := C.gtk_widget_get_device_enabled(self.object, (*C.GdkDevice)(device.ToNative()))
	return gobject.GoBool(unsafe.Pointer(&b))
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

//TODO: gdk_widget_get_visual
//TODO: gdk_widget_set_visual

func (self *Widget) GetPointer() (x, y int) {
	var cx, cy C.gint
	C.gtk_widget_get_pointer(self.object, &cx, &cy)

	return int(cx), int(cy)
}

func (self *Widget) IsAncestor(ancestor WidgetLike) bool {
	b := C.gtk_widget_is_ancestor(self.object, ancestor.W().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Widget) TranslateCoordinates(destWidget WidgetLike, srcX, srcY int) (b bool, destX, destY int) {
	var cx, cy C.gint
	cb := C.gtk_widget_translate_coordinates(self.object, destWidget.W().object, C.gint(srcX), C.gint(srcY), &cx, &cy)
	b = gobject.GoBool(unsafe.Pointer(&cb))
	destX = int(cx)
	destY = int(cy)
	return
}

func (self *Widget) GetHalign() int {
	return int(C.gtk_widget_get_halign(self.object))
}

func (self *Widget) SetHalign(gtk_Align int) {
	C.gtk_widget_set_halign(self.object, C.GtkAlign(gtk_Align))
}

func (self *Widget) GetValign() int {
	return int(C.gtk_widget_get_valign(self.object))
}

func (self *Widget) SetValign(gtk_align int) {
	C.gtk_widget_set_valign(self.object, C.GtkAlign(gtk_align))
}

func (self *Widget) GetPreferredSize() (minimumSize, naturalSize Requisition) {
	var min, nat C.GtkRequisition
	C.gtk_widget_get_preferred_size(self.object, &min, &nat)

	minimumSize = Requisition{int(min.width), int(min.height)}
	naturalSize = Requisition{int(nat.width), int(nat.height)}
	return
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

func (self *Container) GetResizeMode() int {
	return int(C.gtk_container_get_resize_mode(self.object))
}

func (self *Container) SetResizeMode(gtk_ResizeMode int) {
	C.gtk_container_set_resize_mode(self.object, C.GtkResizeMode(gtk_ResizeMode))
}

func (self *Container) CheckResize() {
	C.gtk_container_check_resize(self.object)
}

func (self *Container) ForEach(f interface{}, data ...interface{}) {
	forClosure, id := gobject.CreateCustomClosure(f, data...)

	if forClosure == nil {
		return
	}

	addToClosures(id, forClosure)

	var cId C.gint64 = C.gint64(id)
	C._gtk_container_foreach(self.object, &cId)

	removeFromClosures(id)
}

//TODO: gtk_container_get_children
//TODO: gtk_container_get_path_for_child

func (self *Container) SetReallocateRedraws(needRedraws bool) {
	b := gobject.GBool(needRedraws)
	defer b.Free()
	C.gtk_container_set_reallocate_redraws(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Container) GetFocusChild() WidgetLike {
	w := C.gtk_container_get_focus_child(self.object)
	i, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return i.(WidgetLike)
	}
	return nil
}

func (self *Container) GetFocusVadjustment() *Adjustment {
	a := C.gtk_container_get_focus_vadjustment(self.object)
	if a == nil {
		return nil
	}
	adj, err := gobject.ConvertToGo(unsafe.Pointer(a))
	if err == nil {
		return adj.(*Adjustment)
	}
	return nil
}

func (self *Container) SetFocusVadjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_vadjustment(self.object, adjustment.object)
}

func (self *Container) GetFocusHadjustment() *Adjustment {
	a := C.gtk_container_get_focus_hadjustment(self.object)
	if a == nil {
		return nil
	}
	adj, err := gobject.ConvertToGo(unsafe.Pointer(a))
	if err == nil {
		return adj.(*Adjustment)
	}
	return nil
}

func (self *Container) SetFocusHadjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_hadjustment(self.object, adjustment.object)
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

// End Base Structs }}}

// Windows {{{

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

func (self *Window) SetDefaultSize(width, height int) {
	C.gtk_window_set_default_size(self.object, C.gint(width), C.gint(height))
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

// GtkInvisible {{{
//////////////////////////////

// GtkInvisible type
type Invisible struct {
	object *C.GtkInvisible
	*Widget
}

func NewInvisible() *Invisible {
	i := &Invisible{}
	o := C.gtk_invisible_new()
	i.object = C.to_GtkInvisible(unsafe.Pointer(o))

	if gobject.IsObjectFloating(i) {
		gobject.RefSink(i)
	}
	i.Widget = NewWidget(unsafe.Pointer(o))
	invisibleFinalizer(i)

	return i
}

func NewInvisibleForScreen(screen *gdk3.Screen) *Invisible {
	i := &Invisible{}
	o := C.gtk_invisible_new_for_screen((*C.GdkScreen)(screen.ToNative()))
	i.object = C.to_GtkInvisible(unsafe.Pointer(o))

	if gobject.IsObjectFloating(i) {
		gobject.RefSink(i)
	}
	i.Widget = NewWidget(unsafe.Pointer(o))
	invisibleFinalizer(i)

	return i
}

// Clear Invisible struct when it goes out of reach
func invisibleFinalizer(i *Invisible) {
	runtime.SetFinalizer(i, func(i *Invisible) { gobject.Unref(i) })
}

// Conversion funcs
func newInvisibleFromNative(obj unsafe.Pointer) interface{} {
	i := &Invisible{}
	i.object = C.to_GtkInvisible(obj)

	if gobject.IsObjectFloating(i) {
		gobject.RefSink(i)
	} else {
		gobject.Ref(i)
	}
	i.Widget = NewWidget(obj)
	invisibleFinalizer(i)

	return i
}

func nativeFromInvisible(i interface{}) *gobject.GValue {
	inv, ok := i.(*Invisible)
	if ok {
		gv := gobject.CreateCGValue(GtkType.INVISIBLE, inv.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self Invisible) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Invisible) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Invisible) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Invisible) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self Invisible) W() *Widget {
	return self.Widget
}

// Invisible interface

func (self *Invisible) SetScreen(screen *gdk3.Screen) {
	C.gtk_invisible_set_screen(self.object, (*C.GdkScreen)(screen.ToNative()))
}

func (self *Invisible) GetScreen() *gdk3.Screen {
	s := C.gtk_invisible_get_screen(self.object)

	if scr, err := gobject.ConvertToGo(unsafe.Pointer(s)); err == nil {
		return scr.(*gdk3.Screen)
	}
	return nil
}
//////////////////////////////
// END GtkInvisible
////////////////////////////// }}}

// GtkAssistant {{{
//////////////////////////////

// GtkAssistant type
type Assistant struct {
	object *C.GtkAssistant
	*Window
}

// Clear Assistant struct when it goes out of reach
func assistantFinalizer(as *Assistant) {
	runtime.SetFinalizer(as, func(as *Assistant) { gobject.Unref(as) })
}

func NewAssistant() *Assistant {
	as := &Assistant{}
	o := C.gtk_assistant_new()
	as.object = C.to_GtkAssistant(unsafe.Pointer(o))

	if gobject.IsObjectFloating(as) {
		gobject.RefSink(as)
	}
	as.Window = newWindowFromNative(unsafe.Pointer(o)).(*Window)
	assistantFinalizer(as)

	return as
}

// Conversion funcs
func newAssistantFromNative(obj unsafe.Pointer) interface{} {
	as := &Assistant{}
	as.object = C.to_GtkAssistant(obj)

	if gobject.IsObjectFloating(as) {
		gobject.RefSink(as)
	} else {
		gobject.Ref(as)
	}
	as.Window = newWindowFromNative(obj).(*Window)
	assistantFinalizer(as)

	return as
}

func nativeFromAssistant(as interface{}) *gobject.GValue {
	assistant, ok := as.(*Assistant)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ASSISTANT, assistant.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self Assistant) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Assistant) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Assistant) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Assistant) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Assistant interface

func (self *Assistant) GetCurrentPage() int {
	return int(C.gtk_assistant_get_current_page(self.object))
}

func (self *Assistant) SetCurrentPage(pageNum int) {
	C.gtk_assistant_set_current_page(self.object, C.gint(pageNum))
}

func (self *Assistant) GetNPages() int {
	return int(C.gtk_assistant_get_n_pages(self.object))
}

func (self *Assistant) GetNthPage(pageNum int) WidgetLike {
	w := C.gtk_assistant_get_nth_page(self.object, C.gint(pageNum))

	if w == nil {
		return nil
	}

	wid, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *Assistant) PrependPage(page WidgetLike) int {
	return int(C.gtk_assistant_prepend_page(self.object, page.W().object))
}

func (self *Assistant) AppendPage(page WidgetLike) int {
	return int(C.gtk_assistant_append_page(self.object, page.W().object))
}

func (self *Assistant) InsertPage(page WidgetLike, position int) int {
	return int(C.gtk_assistant_insert_page(self.object, page.W().object, C.gint(position)))
}

func (self *Assistant) RemovePage(pageNum int) {
	C.gtk_assistant_remove_page(self.object, C.gint(pageNum))
}

//TODO: gtk_assistant_set_forward_page_func

func (self *Assistant) SetPageType(page WidgetLike, gtk_AssistantPage int) {
	C.gtk_assistant_set_page_type(self.object, page.W().object, C.GtkAssistantPageType(gtk_AssistantPage))
}

func (self *Assistant) GetPageType(page WidgetLike) int {
	return int(C.gtk_assistant_get_page_type(self.object, page.W().object))
}

func (self *Assistant) SetPageTitle(page WidgetLike, title string) {
	s := gobject.GString(title)
	defer s.Free()

	C.gtk_assistant_set_page_title(self.object, page.W().object, (*C.gchar)(s.GetPtr()))
}

func (self *Assistant) GetPageTitle(page WidgetLike) string {
	s := C.gtk_assistant_get_page_title(self.object, page.W().object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Assistant) SetPageComplete(page WidgetLike, complete bool) {
	b := gobject.GBool(complete)
	defer b.Free()
	C.gtk_assistant_set_page_complete(self.object, page.W().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Assistant) GetPageComplete(page WidgetLike) bool {
	b := C.gtk_assistant_get_page_complete(self.object, page.W().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Assistant) AddActionWidget(child WidgetLike) {
	C.gtk_assistant_add_action_widget(self.object, child.W().object)
}

func (self *Assistant) RemoveActionWidget(child WidgetLike) {
	C.gtk_assistant_remove_action_widget(self.object, child.W().object)
}

func (self *Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(self.object)
}

func (self *Assistant) Commit() {
	C.gtk_assistant_commit(self.object)
}

func (self *Assistant) NextPage() {
	C.gtk_assistant_next_page(self.object)
}

func (self *Assistant) PreviousPage() {
	C.gtk_assistant_previous_page(self.object)
}

//////////////////////////////
// END GtkAssistant
////////////////////////////// }}}

// End Windows }}}

// Display Widgets {{{

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

// GtkProgressBar {{{
//////////////////////////////

// GtkProgressBar type
type ProgressBar struct {
	object *C.GtkProgressBar
	*Widget
}

func NewProgressBar() *ProgressBar {
	pb := &ProgressBar{}
	o := C.gtk_progress_bar_new()
	pb.object = C.to_GtkProgressBar(unsafe.Pointer(o))

	if gobject.IsObjectFloating(pb) {
		gobject.RefSink(pb)
	}
	pb.Widget = NewWidget(unsafe.Pointer(o))
	progressBarFinalizer(pb)

	return pb
}

// Clear ProgressBar struct when it goes out of reach
func progressBarFinalizer(pb *ProgressBar) {
	runtime.SetFinalizer(pb, func(pb *ProgressBar) { gobject.Unref(pb) })
}

// Conversion funcs
func newProgressBarFromNative(obj unsafe.Pointer) interface{} {
	pb := &ProgressBar{}
	pb.object = C.to_GtkProgressBar(obj)

	if gobject.IsObjectFloating(pb) {
		gobject.RefSink(pb)
	} else {
		gobject.Ref(pb)
	}
	pb.Widget = NewWidget(obj)
	progressBarFinalizer(pb)

	return pb
}

func nativeFromProgressBar(pb interface{}) *gobject.GValue {
	progress, ok := pb.(*ProgressBar)
	if ok {
		gv := gobject.CreateCGValue(GtkType.PROGRESS_BAR, progress.ToNative())
		return gv
	}
	return nil
}

// To be object like
func (self ProgressBar) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ProgressBar) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ProgressBar) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)

}

func (self ProgressBar) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be widget-like
func (self ProgressBar) W() *Widget {
	return self.Widget
}

// ProgressBar interface

func (self *ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(self.object)
}

func (self *ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(self.object, C.gdouble(fraction))
}

func (self *ProgressBar) GetFraction() float64 {
	return float64(C.gtk_progress_bar_get_fraction(self.object))
}

func (self *ProgressBar) SetInverted(inverted bool) {
	b := gobject.GBool(inverted)
	defer b.Free()
	C.gtk_progress_bar_set_inverted(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ProgressBar) GetInverted() bool {
	b := C.gtk_progress_bar_get_inverted(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ProgressBar) SetShowText(showText bool) {
	b := gobject.GBool(showText)
	defer b.Free()
	C.gtk_progress_bar_set_show_text(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *ProgressBar) GetShowText() bool {
	b := C.gtk_progress_bar_get_show_text(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ProgressBar) SetText(text interface{}) {
	if text == nil {
		C.gtk_progress_bar_set_text(self.object, nil)
		return
	}
	if st, ok := text.(string); ok {
		s := gobject.GString(st)
		defer s.Free()
		C.gtk_progress_bar_set_text(self.object, (*C.gchar)(s.GetPtr()))
	}
}

func (self *ProgressBar) GetText() string {
	s := C.gtk_progress_bar_get_text(self.object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

//TODO: gtk_progress_bar_set_ellipsize
//TODO: gtk_progress_bar_get_ellipsize

func (self *ProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(self.object, C.gdouble(fraction))
}

func (self *ProgressBar) GetPulseStep() float64 {
	return float64(C.gtk_progress_bar_get_pulse_step(self.object))
}
//////////////////////////////
// End GtkProgressBar
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

// End Display Widgets }}}

// Buttons and Toggles {{{

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

// GtkCheckButton {{{
//////////////////////////////

// GtkCheckButton type
type CheckButton struct {
	object *C.GtkCheckButton
	*ToggleButton
}

// Clear CheckButton struct when it goes out of reach
func checkButtonFinalizer(cb *CheckButton) {
	runtime.SetFinalizer(cb, func(cb *CheckButton) { gobject.Unref(cb) })
}

func NewCheckButton() *CheckButton {
	cb := &CheckButton{}
	o := C.gtk_check_button_new()
	cb.object = C.to_GtkCheckButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.ToggleButton = newToggleButtonFromNative(unsafe.Pointer(o)).(*ToggleButton)
	checkButtonFinalizer(cb)

	return cb
}

func NewCheckButtonWithLabel(label string) *CheckButton {
	cb := &CheckButton{}
	s := gobject.GString(label)
	defer s.Free()

	o := C.gtk_check_button_new_with_label((*C.gchar)(s.GetPtr()))
	cb.object = C.to_GtkCheckButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.ToggleButton = newToggleButtonFromNative(unsafe.Pointer(o)).(*ToggleButton)
	checkButtonFinalizer(cb)

	return cb
}

func NewCheckButtonWithMnemonic(label string) *CheckButton {
	cb := &CheckButton{}
	s := gobject.GString(label)
	defer s.Free()

	o := C.gtk_check_button_new_with_mnemonic((*C.gchar)(s.GetPtr()))
	cb.object = C.to_GtkCheckButton(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	}
	cb.ToggleButton = newToggleButtonFromNative(unsafe.Pointer(o)).(*ToggleButton)
	checkButtonFinalizer(cb)

	return cb
}

// Conversion funcs
func newCheckButtonFromNative(obj unsafe.Pointer) interface{} {
	cb := &CheckButton{}
	cb.object = C.to_GtkCheckButton(obj)

	if gobject.IsObjectFloating(cb) {
		gobject.RefSink(cb)
	} else {
		gobject.Ref(cb)
	}
	cb.ToggleButton = newToggleButtonFromNative(obj).(*ToggleButton)
	checkButtonFinalizer(cb)

	return cb
}

func nativeFromCheckButton(cb interface{}) *gobject.GValue {
	check, ok := cb.(*CheckButton)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CHECK_BUTTON, check.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CheckButton) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CheckButton) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CheckButton) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CheckButton) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// END GtkCheckButton
////////////////////////////// }}}

// End Buttons and Toggles }}}

// Numeric/Text Data Entry {{{

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

func (self *Entry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(self.object, completion.object)
}
//////////////////////////////
// END GtkEntry
////////////////////////////// }}}

// GtkEntryCompletion {{{
//////////////////////////////

// GtkEntryCompletion type
type EntryCompletion struct {
	object *C.GtkEntryCompletion
}

func NewEntryCompletion() *EntryCompletion {
	ec := &EntryCompletion{}
	ec.object = C.gtk_entry_completion_new()

	if gobject.IsObjectFloating(ec) {
		gobject.RefSink(ec)
	}
	entryCompletionFinalizer(ec)

	return ec
}

//TODO: gtk_entry_completion_new_with_area

// Clear EntryCompletion struct when it goes out of reach
func entryCompletionFinalizer(ec *EntryCompletion) {
	runtime.SetFinalizer(ec, func(ec *EntryCompletion) { gobject.Unref(ec) })
}

// Conversion funcs
func newEntryCompletionFromNative(obj unsafe.Pointer) interface{} {
	ec := &EntryCompletion{}
	ec.object = C.to_GtkEntryCompletion(obj)

	if gobject.IsObjectFloating(ec) {
		gobject.RefSink(ec)
	} else {
		gobject.Ref(ec)
	}
	entryCompletionFinalizer(ec)

	return ec
}

func nativeFromEntryCompletion(ec interface{}) *gobject.GValue {
	entryCompletion, ok := ec.(*EntryCompletion)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ENTRY_COMPLETION, entryCompletion.ToNative())
		return gv
	}
	return nil
}

// To be Object-like
func (self EntryCompletion) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self EntryCompletion) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self EntryCompletion) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self EntryCompletion) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// EntryCompletion interface

func (self *EntryCompletion) GetEntry() WidgetLike {
	w := C.gtk_entry_completion_get_entry(self.object)

	if w == nil {
		return nil
	}

	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return widget.(WidgetLike)
	}
	return nil
}

func (self *EntryCompletion) SetModel(model TreeModelLike) {
	var m *C.GtkTreeModel = nil

	if model != nil {
		m = model.ITreeModel().object
	}

	C.gtk_entry_completion_set_model(self.object, m)
}

func (self *EntryCompletion) GetModel() TreeModelLike {
	m := C.gtk_entry_completion_get_model(self.object)

	if m == nil {
		return nil
	}

	model, err := gobject.ConvertToGo(unsafe.Pointer(m))

	if err == nil {
		return model.(TreeModelLike)
	}
	return nil
}

func (self *EntryCompletion) SetMatchFunc(f interface{}, data ...interface{}) {
	clFunc, id := gobject.CreateCustomClosure(f, data...)

	if clFunc == nil {
		return
	}

	// Add closure to local closure map
	addToClosures(id, clFunc)

	// Register closure to C
	C._gtk_entry_completion_set_match_func(self.object, C.gint64(id))

	// Add WeakRef, so that when object is destroyed, closure func removes
	// our closure from local closure map
	gobject.WeakRef(self, func() { removeFromClosures(id) })
}

func (self *EntryCompletion) SetMinimumKeyLength(length int) {
	C.gtk_entry_completion_set_minimum_key_length(self.object, C.gint(length))
}

func (self *EntryCompletion) GetMinimumKeyLength() int {
	return int(C.gtk_entry_completion_get_minimum_key_length(self.object))
}

func (self *EntryCompletion) Complete() {
	C.gtk_entry_completion_complete(self.object)
}

func (self *EntryCompletion) GetCompletionPrefix() string {
	s := C.gtk_entry_completion_get_completion_prefix(self.object)

	if s == nil {
		return ""
	}

	return gobject.GoString(unsafe.Pointer(s))
}

func (self *EntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix(self.object)
}

func (self *EntryCompletion) InsertActionText(index int, text string) {
	s := gobject.GString(text)
	defer s.Free()
	C.gtk_entry_completion_insert_action_text(self.object, C.gint(index), (*C.gchar)(s.GetPtr()))
}

func (self *EntryCompletion) InsertActionMarkup(index int, markup string) {
	s := gobject.GString(markup)
	defer s.Free()
	C.gtk_entry_completion_insert_action_markup(self.object, C.gint(index), (*C.gchar)(s.GetPtr()))
}

func (self *EntryCompletion) DeleteAction(index int) {
	C.gtk_entry_completion_delete_action(self.object, C.gint(index))
}

func (self *EntryCompletion) SetTextColumn(column int) {
	C.gtk_entry_completion_set_text_column(self.object, C.gint(column))
}

func (self *EntryCompletion) GetTextColumn() int {
	return int(C.gtk_entry_completion_get_text_column(self.object))
}

func (self *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	b := gobject.GBool(inlineCompletion)
	defer b.Free()
	C.gtk_entry_completion_set_inline_completion(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetInlineCompletion() bool {
	b := C.gtk_entry_completion_get_inline_completion(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	b := gobject.GBool(inlineSelection)
	defer b.Free()
	C.gtk_entry_completion_set_inline_selection(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetInlineSelection() bool {
	b := C.gtk_entry_completion_get_inline_selection(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	b := gobject.GBool(popupCompletion)
	defer b.Free()
	C.gtk_entry_completion_set_popup_completion(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetPopupCompletion() bool {
	b := C.gtk_entry_completion_get_popup_completion(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	b := gobject.GBool(popupSetWidth)
	defer b.Free()
	C.gtk_entry_completion_set_popup_set_width(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetPopupSetWidth() bool {
	b := C.gtk_entry_completion_get_popup_set_width(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	b := gobject.GBool(popupSingleMatch)
	defer b.Free()
	C.gtk_entry_completion_set_popup_single_match(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *EntryCompletion) GetPopupSingleMatch() bool {
	b := C.gtk_entry_completion_get_popup_single_match(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}
//////////////////////////////
// END GtkEntryCompletion
////////////////////////////// }}}

// End Numeric/Text Data Entry }}}

// Tree, List and Icon Grid Widgets {{{

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

// GtkTreeRowReference {{{
//////////////////////////////

// GtkTreeRowReference type
type TreeRowReference struct {
	object *C.GtkTreeRowReference
}

//func NewTreeRowReference() *TreeRowReference {

// Clear TreeRowReference when it goes out of reach
func treeRowReferenceFinalizer(trr *TreeRowReference) {
	runtime.SetFinalizer(trr, func(trr *TreeRowReference) { trr.free() })
}

// TreeRowReference interface
func (self *TreeRowReference) free() {
	C.gtk_tree_row_reference_free(self.object)
}
//////////////////////////////
// END TreeRowReference
////////////////////////////// }}}

// GtkTreeIter {{{
//////////////////////////////

// GtkTreeRowReference type
type TreeIter struct {
	object C.GtkTreeIter
}

func (self *TreeIter) Copy() *TreeIter {
	ti := C.gtk_tree_iter_copy(&self.object)
	gti := &TreeIter{*ti}
	C.gtk_tree_iter_free(ti)

	return gti
}

//////////////////////////////
// END GtkTreeIter
////////////////////////////// }}}

// GtkTreeModel {{{
//////////////////////////////

// GtkTreeModel type
type TreeModel struct {
	object *C.GtkTreeModel
}

func NewTreeModel(ctm unsafe.Pointer) *TreeModel {
	tm := &TreeModel{}
	tm.object = C.to_GtkTreeModel(ctm)

	if gobject.IsObjectFloating(tm) {
		gobject.RefSink(tm)
	} else {
		gobject.Ref(tm)
	}
	treeModelFinalizer(tm)

	return tm
}

// Clear TreeModel struct when it goes out of scope
func treeModelFinalizer(tm *TreeModel) {
	runtime.SetFinalizer(tm, func(tm *TreeModel) { gobject.Unref(tm) })
}

// To be object-like
func (self TreeModel) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeModel) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeModel) Set(properties map[string]interface{}) {
}

func (self TreeModel) Get(properties []string) map[string]interface{} {
	return nil
}

// TreeModel interface

func (self *TreeModel) GetFlags() int {
	return int(C.gtk_tree_model_get_flags(self.object))
}

func (self *TreeModel) GetNColumns() int {
	return int(C.gtk_tree_model_get_n_columns(self.object))
}

//TODO: gtk_tree_model_get_column_type

func (self *TreeModel) GetIter(iter *TreeIter, path *TreePath) bool {
	b := C.gtk_tree_model_get_iter(self.object, &iter.object, path.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) GetIterFromString(iter *TreeIter, pathString string) bool {
	s := gobject.GString(pathString)
	defer s.Free()
	b := C.gtk_tree_model_get_iter_from_string(self.object, &iter.object, (*C.gchar)(s.GetPtr()))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) GetIterFirst(iter *TreeIter) bool {
	b := C.gtk_tree_model_get_iter_first(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) GetPath(iter *TreeIter) *TreePath {
	p := C.gtk_tree_model_get_path(self.object, &iter.object)
	path := &TreePath{p}
	treePathFinalizer(path)

	return path
}

func (self *TreeModel) GetValue(iter *TreeIter, column int) interface{} {
	var val C.GValue
	C.gtk_tree_model_get_value(self.object, &iter.object, C.gint(column), &val)

	gv := gobject.NewGValueFromNative(unsafe.Pointer(&val))
	defer gv.Free()
	t, e := gobject.ConvertToGo(gv.GetPtr(), gv.GetTypeID())

	if e == nil {
		return t
	}
	return nil
}

func (self *TreeModel) IterNext(iter *TreeIter) bool {
	b := C.gtk_tree_model_iter_next(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterPrevious(iter *TreeIter) bool {
	b := C.gtk_tree_model_iter_previous(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterChildren(iter *TreeIter, parent *TreeIter) bool {
	b := C.gtk_tree_model_iter_children(self.object, &iter.object, &parent.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterHasChild(iter *TreeIter) bool {
	b := C.gtk_tree_model_iter_has_child(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterNChildren(iter *TreeIter) int {
	return int(C.gtk_tree_model_iter_n_children(self.object, &iter.object))
}

func (self *TreeModel) IterNthChild(iter *TreeIter, parent *TreeIter, n int) bool {
	b := C.gtk_tree_model_iter_nth_child(self.object, &iter.object, &parent.object, C.gint(n))
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) IterParent(iter *TreeIter, child *TreeIter) bool {
	b := C.gtk_tree_model_iter_parent(self.object, &iter.object, &child.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeModel) GetStringFromIter(iter *TreeIter) string {
	s := C.gtk_tree_model_get_string_from_iter(self.object, &iter.object)
	return gobject.GoString(unsafe.Pointer(&s))
}

//TODO: gtk_tree_model_ref_node
//TODO: gtk_tree_model_unref_node
//TODO: gtk_tree_model_get_valist
//TODO: gtk_tree_model_foreach
//...
//////////////////////////////
// END GtkTreeModel
////////////////////////////// }}}

// GtkListStore {{{
//////////////////////////////

// GtkListStore type
type ListStore struct {
	object *C.GtkListStore
	*TreeModel
}

func NewListStore(gtypeSlice []gobject.GType) *ListStore {
	count := len(gtypeSlice)
	if count == 0 {
		return nil
	}

	arr := C.gtype_array_new(C.int(count))
	defer C.gtype_array_free(arr)

	for i, gtype := range gtypeSlice {
		C.gtype_array_set_element(arr, C.int(i), C.GType(gtype))
	}

	ls := &ListStore{}
	ls.object = C.gtk_list_store_newv(C.gint(count), arr)

	if gobject.IsObjectFloating(ls) {
		gobject.RefSink(ls)
	}
	ls.TreeModel = NewTreeModel(unsafe.Pointer(ls.object))
	listStoreFinalizer(ls)

	return ls
}

// Clear ListStore struct when it goes out of reach
func listStoreFinalizer(ls *ListStore) {
	runtime.SetFinalizer(ls, func(ls *ListStore) { gobject.Unref(ls) })
}

// Conversion functions
func newListStoreFromNative(obj unsafe.Pointer) interface{} {
	ls := &ListStore{}
	ls.object = C.to_GtkListStore(obj)

	if gobject.IsObjectFloating(ls) {
		gobject.RefSink(ls)
	} else {
		gobject.Ref(ls)
	}
	ls.TreeModel = NewTreeModel(obj)
	listStoreFinalizer(ls)

	return ls
}

func nativeFromListStore(ls interface{}) *gobject.GValue {
	listStore, ok := ls.(*ListStore)
	if ok {
		gv := gobject.CreateCGValue(GtkType.LIST_STORE, listStore.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self ListStore) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self ListStore) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self ListStore) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self ListStore) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// ListStore implements TreeModel
func (self ListStore) ITreeModel() *TreeModel {
	return self.TreeModel
}

// ListStore interface

func (self *ListStore) SetValue(iter *TreeIter, column int, value interface{}) {
	cval := gobject.ConvertToC(value)

	// If value cannot be converted to fundamnetal
	// gobject value, then we don't do anything
	if cval == nil {
		return
	}
	defer cval.Free()

	C.gtk_list_store_set_value(self.object, &iter.object, C.gint(column), (*C.GValue)(cval.ToCGValue()))
}

func (self *ListStore) SetValues(iter *TreeIter, values map[int]interface{}) {
	for k, v := range values {
		self.SetValue(iter, k, v)
	}
}

func (self *ListStore) Remove(iter *TreeIter) bool {
	b := C.gtk_list_store_remove(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *ListStore) Insert(iter *TreeIter, position int) {
	C.gtk_list_store_insert(self.object, &iter.object, C.gint(position))
}

func (self *ListStore) InsertBefore(iter, sibling *TreeIter) {
	var cSibling *C.GtkTreeIter = nil

	if sibling != nil {
		cSibling = &sibling.object
	}
	C.gtk_list_store_insert_before(self.object, &iter.object, cSibling)
}

func (self *ListStore) InsertAfter(iter, sibling *TreeIter) {
	var cSibling *C.GtkTreeIter = nil

	if sibling != nil {
		cSibling = &sibling.object
	}
	C.gtk_list_store_insert_after(self.object, &iter.object, cSibling)
}

func (self *ListStore) InsertWithValues(iter *TreeIter, position int, values map[int]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	valArr := C.gvalue_array_new(C.int(len(values)))
	columns := C.gint_array_new(C.int(len(values)))
	defer C.gvalue_array_free(valArr)
	defer C.gint_array_free(columns)

	var i int = 0
	for k, v := range values {
		C.gint_array_set_element(columns, C.int(i), C.gint(k))
		cval := gobject.ConvertToC(v)
		defer cval.Free()
		C.gvalue_array_set_element(valArr, C.int(i), (*C.GValue)(cval.ToCGValue()))
		i++
	}
	var Citer *C.GtkTreeIter = nil

	if iter != nil {
		Citer = &iter.object
	}

	C.gtk_list_store_insert_with_valuesv(self.object, Citer, C.gint(position), columns, valArr, C.gint(len(values)))
}

func (self *ListStore) Prepend(iter *TreeIter) {
	C.gtk_list_store_prepend(self.object, &iter.object)
}

func (self *ListStore) Append(iter *TreeIter) {
	C.gtk_list_store_append(self.object, &iter.object)
}

func (self *ListStore) Clear() {
	C.gtk_list_store_clear(self.object)
}

//TODO: gtk_list_store_reorder

func (self *ListStore) Swap(a, b *TreeIter) {
	C.gtk_list_store_swap(self.object, &a.object, &b.object)
}

func (self *ListStore) MoveBefore(iter, position *TreeIter) {
	C.gtk_list_store_move_before(self.object, &iter.object, &position.object)
}

func (self *ListStore) MoveAfter(iter, position *TreeIter) {
	C.gtk_list_store_move_after(self.object, &iter.object, &position.object)
}
//////////////////////////////
// END GtkListStore
////////////////////////////// }}}

// GtkTreeStore {{{
//////////////////////////////

// GtkTreeStore type
type TreeStore struct {
	object *C.GtkTreeStore
	*TreeModel
}

func NewTreeStore(gtypeSlice []gobject.GType) *TreeStore {
	count := len(gtypeSlice)
	if count == 0 {
		return nil
	}

	arr := C.gtype_array_new(C.int(count))
	defer C.gtype_array_free(arr)

	for i, gtype := range gtypeSlice {
		C.gtype_array_set_element(arr, C.int(i), C.GType(gtype))
	}

	ts := &TreeStore{}
	ts.object = C.gtk_tree_store_newv(C.gint(count), arr)

	if gobject.IsObjectFloating(ts) {
		gobject.RefSink(ts)
	}
	ts.TreeModel = NewTreeModel(unsafe.Pointer(ts.object))
	treeStoreFinalizer(ts)

	return ts
}

// Clear TreeStore struct when it goes out of reach
func treeStoreFinalizer(ts *TreeStore) {
	runtime.SetFinalizer(ts, func(ts *TreeStore) { gobject.Unref(ts) })
}

// Conversion functions
func newTreeStoreFromNative(obj unsafe.Pointer) interface{} {
	ts := &TreeStore{}
	ts.object = C.to_GtkTreeStore(obj)

	if gobject.IsObjectFloating(ts) {
		gobject.RefSink(ts)
	} else {
		gobject.Ref(ts)
	}
	ts.TreeModel = NewTreeModel(obj)
	treeStoreFinalizer(ts)

	return ts
}

func nativeFromTreeStore(ts interface{}) *gobject.GValue {
	treeStore, ok := ts.(*TreeStore)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_STORE, treeStore.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeStore) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeStore) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeStore) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeStore) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TreeStore implements TreeModel
func (self TreeStore) ITreeModel() *TreeModel {
	return self.TreeModel
}

// TreeStore interface

func (self *TreeStore) SetValue(iter *TreeIter, column int, value interface{}) {
	cval := gobject.ConvertToC(value)

	// If value cannot be converted to fundamnetal
	// gobject value, then we don't do anything
	if cval == nil {
		return
	}
	defer cval.Free()

	C.gtk_tree_store_set_value(self.object, &iter.object, C.gint(column), (*C.GValue)(cval.ToCGValue()))
}

func (self *TreeStore) SetValues(iter *TreeIter, values map[int]interface{}) {
	for k, v := range values {
		self.SetValue(iter, k, v)
	}
}

func (self *TreeStore) Remove(iter *TreeIter) bool {
	b := C.gtk_tree_store_remove(self.object, &iter.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeStore) Insert(iter, parent *TreeIter, position int) {
	var cparent *C.GtkTreeIter = nil

	if parent != nil {
		cparent = &parent.object
	}
	C.gtk_tree_store_insert(self.object, &iter.object, cparent, C.gint(position))
}

func (self *TreeStore) InsertBefore(iter, parent, sibling *TreeIter) {
	var cSibling, cParent *C.GtkTreeIter = nil, nil

	if sibling != nil {
		cSibling = &sibling.object
	}

	if parent != nil {
		cParent = &parent.object
	}
	C.gtk_tree_store_insert_before(self.object, &iter.object, cParent, cSibling)
}

func (self *TreeStore) InsertAfter(iter, parent, sibling *TreeIter) {
	var cSibling, cParent *C.GtkTreeIter = nil, nil

	if sibling != nil {
		cSibling = &sibling.object
	}

	if parent != nil {
		cParent = &parent.object
	}
	C.gtk_tree_store_insert_after(self.object, &iter.object, cParent, cSibling)
}

func (self *TreeStore) InsertWithValues(iter, parent *TreeIter, position int, values map[int]interface{}) {
	if values == nil || len(values) == 0 {
		return
	}

	valArr := C.gvalue_array_new(C.int(len(values)))
	columns := C.gint_array_new(C.int(len(values)))
	defer C.gvalue_array_free(valArr)
	defer C.gint_array_free(columns)

	var i int = 0
	for k, v := range values {
		C.gint_array_set_element(columns, C.int(i), C.gint(k))
		cval := gobject.ConvertToC(v)
		defer cval.Free()
		C.gvalue_array_set_element(valArr, C.int(i), (*C.GValue)(cval.ToCGValue()))
		i++
	}
	var Citer, Cparent *C.GtkTreeIter = nil, nil

	if iter != nil {
		Citer = &iter.object
	}

	if parent != nil {
		Cparent = &parent.object
	}

	C.gtk_tree_store_insert_with_valuesv(self.object, Citer, Cparent, C.gint(position), columns, valArr, C.gint(len(values)))
}

func (self *TreeStore) Prepend(iter, parent *TreeIter) {
	var cParent *C.GtkTreeIter = nil

	if parent != nil {
		cParent = &parent.object
	}
	C.gtk_tree_store_prepend(self.object, &iter.object, cParent)
}

func (self *TreeStore) Append(iter, parent *TreeIter) {
	var cParent *C.GtkTreeIter = nil
	if parent != nil {
		cParent = &parent.object
	}
	C.gtk_tree_store_append(self.object, &iter.object, cParent)
}

func (self *TreeStore) IsAncestor(iter, descendant *TreeIter) bool {
	b := C.gtk_tree_store_is_ancestor(self.object, &iter.object, &descendant.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeStore) IterDepth(iter *TreeIter) int {
	return int(C.gtk_tree_store_iter_depth(self.object, &iter.object))
}

func (self *TreeStore) Clear() {
	C.gtk_tree_store_clear(self.object)
}

//TODO: gtk_tree_store_reorder

func (self *TreeStore) Swap(a, b *TreeIter) {
	C.gtk_tree_store_swap(self.object, &a.object, &b.object)
}

func (self *TreeStore) MoveBefore(iter, position *TreeIter) {
	C.gtk_tree_store_move_before(self.object, &iter.object, &position.object)
}

func (self *TreeStore) MoveAfter(iter, position *TreeIter) {
	C.gtk_tree_store_move_after(self.object, &iter.object, &position.object)
}
//////////////////////////////
// END GtkTreeStore
////////////////////////////// }}}

// GtkCellArea {{{
//////////////////////////////

// GtkCellArea type
type CellArea struct {
	object *C.GtkCellArea
}

// Clear CellArea struct when it goes out of reach
func cellAreaFinalizer(ca *CellArea) {
	runtime.SetFinalizer(ca, func(ca *CellArea) { gobject.Unref(ca) })
}

// Conversion funcs
func newCellAreaFromNative(obj unsafe.Pointer) interface{} {
	ca := &CellArea{}
	ca.object = C.to_GtkCellArea(obj)

	if gobject.IsObjectFloating(ca) {
		gobject.RefSink(ca)
	} else {
		gobject.Ref(ca)
	}
	cellAreaFinalizer(ca)

	return ca
}

func nativeFromCellArea(ca interface{}) *gobject.GValue {
	cell, ok := ca.(*CellArea)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_AREA, cell.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellArea) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellArea) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellArea) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellArea) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// CellArea interface

func (self *CellArea) Add(renderer CellRendererLike) {
	C.gtk_cell_area_add(self.object, renderer.CRenderer().object)
}

func (self *CellArea) Remove(renderer CellRendererLike) {
	C.gtk_cell_area_remove(self.object, renderer.CRenderer().object)
}

func (self *CellArea) HasRenderer(renderer CellRendererLike) bool {
	b := C.gtk_cell_area_has_renderer(self.object, renderer.CRenderer().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellArea) Foreach(callback interface{}, data ...interface{}) {
	f, id := gobject.CreateCustomClosure(callback, data...)
	_closures[id] = f
	C._gtk_cell_area_foreach(self.object, C.gint64(id))
	delete(_closures, id)
}
//////////////////////////////
// End GtkCellArea
////////////////////////////// }}}

// GtkCellAreaContext {{{
//////////////////////////////

// GtkCellAreaContext type
type CellAreaContext struct {
	object *C.GtkCellAreaContext
}

// Clear CellAreaContext struct when it goes out of reach
func cellAreaContextFinalizer(ca *CellAreaContext) {
	runtime.SetFinalizer(ca, func(ca *CellAreaContext) { gobject.Unref(ca) })
}

// Conversion funcs
func newCellAreaContextFromNative(obj unsafe.Pointer) interface{} {
	ca := &CellAreaContext{}
	ca.object = C.to_GtkCellAreaContext(obj)

	if gobject.IsObjectFloating(ca) {
		gobject.RefSink(ca)
	} else {
		gobject.Ref(ca)
	}
	cellAreaContextFinalizer(ca)

	return ca
}

func nativeFromCellAreaContext(ca interface{}) *gobject.GValue {
	cell, ok := ca.(*CellAreaContext)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_AREA_CONTEXT, cell.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellAreaContext) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellAreaContext) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellAreaContext) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellAreaContext) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// CellAreaContext interface

func (self *CellAreaContext) GetArea() *CellArea {
	ar := C.gtk_cell_area_context_get_area(self.object)
	if ar == nil {
		return nil
	}
	area, _ := gobject.ConvertToGo(unsafe.Pointer(ar))

	return area.(*CellArea)
}

func (self *CellAreaContext) Allocate(width, height int) {
	C.gtk_cell_area_context_allocate(self.object, C.gint(width), C.gint(height))
}

func (self *CellAreaContext) Reset() {
	C.gtk_cell_area_context_reset(self.object)
}

func (self *CellAreaContext) GetPreferredWidth() (mWidth, nWidth int) {
	var m, n C.gint
	C.gtk_cell_area_context_get_preferred_width(self.object, &m, &n)
	mWidth = int(m)
	nWidth = int(n)
	return
}

func (self *CellAreaContext) GetPreferredHeight() (mHeight, nHeight int) {
	var m, n C.gint
	C.gtk_cell_area_context_get_preferred_width(self.object, &m, &n)
	mHeight = int(m)
	nHeight = int(n)
	return
}

func (self *CellAreaContext) GetPreferredHeightForWidth(width int) (mHeight, nHeight int) {
	var m, n C.gint
	C.gtk_cell_area_context_get_preferred_height_for_width(self.object, C.gint(width), &m, &n)
	mHeight = int(m)
	nHeight = int(n)
	return
}

func (self *CellAreaContext) GetPreferredWidthForHeight(height int) (mWidth, nWidth int) {
	var m, n C.gint
	C.gtk_cell_area_context_get_preferred_width_for_height(self.object, C.gint(height), &m, &n)
	mWidth = int(m)
	nWidth = int(n)
	return
}

func (self *CellAreaContext) GetAllocation() (width, height int) {
	var w, h C.gint
	C.gtk_cell_area_context_get_allocation(self.object, &w, &h)
	width = int(w)
	height = int(h)
	return
}

func (self *CellAreaContext) PushPreferredWidth(minimumWidth, naturalWidth int) {
	C.gtk_cell_area_context_push_preferred_width(self.object, C.gint(minimumWidth), C.gint(naturalWidth))
}

func (self *CellAreaContext) PushPreferredHeight(minimumHeight, naturalHeight int) {
	C.gtk_cell_area_context_push_preferred_height(self.object, C.gint(minimumHeight), C.gint(naturalHeight))
}
//////////////////////////////
// End GtkCellAreaContext
////////////////////////////// }}}

// GtkCellRenderer {{{
//////////////////////////////

// GtkCellRenderer type
type CellRenderer struct {
	object *C.GtkCellRenderer
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererFinalizer(cl *CellRenderer) {
	runtime.SetFinalizer(cl, func(cl *CellRenderer) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRenderer{}
	cl.object = C.to_GtkCellRenderer(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cellRendererFinalizer(cl)

	return cl
}

func nativeFromCellRenderer(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRenderer)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRenderer) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRenderer) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRenderer) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRenderer) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Interface functions

func (self *CellRenderer) GetAlignedArea(w WidgetLike, gtk_CellRendererState_flags int, cellArea gdk3.Rectangle) gdk3.Rectangle {
	ca := gobject.ConvertToC(cellArea)
	defer ca.Free()

	var alignedArea C.GdkRectangle

	C.gtk_cell_renderer_get_aligned_area(self.object, w.W().object, C.GtkCellRendererState(gtk_CellRendererState_flags),
		(*C.GdkRectangle)(ca.GetPtr()), &alignedArea)

	if rec, err := gobject.ConvertToGo(unsafe.Pointer(&alignedArea), gdk3.GdkType.RECTANGLE); err == nil {
		return rec.(gdk3.Rectangle)
	}
	return gdk3.Rectangle{}
}

//TODO: gtk_cell_renderer_render
//TODO: gtk_cell_renderer_activate
//TODO: gtk_cell_renderer_start_editing
//TODO: gtk_cell_renderer_stop_editing

func (self *CellRenderer) GetFixedSize() (width, height int) {
	var w C.gint
	var h C.gint
	C.gtk_cell_renderer_get_fixed_size(self.object, &w, &h)

	width = int(w)
	height = int(h)
	return
}

func (self *CellRenderer) SetFixedSize(width, height int) {
	C.gtk_cell_renderer_set_fixed_size(self.object, C.gint(width), C.gint(height))
}

func (self *CellRenderer) GetVisible() bool {
	b := C.gtk_cell_renderer_get_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRenderer) SetVisible(visible bool) {
	b := gobject.GBool(visible)
	defer b.Free()
	C.gtk_cell_renderer_set_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellRenderer) GetSensitive() bool {
	b := C.gtk_cell_renderer_get_sensitive(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRenderer) SetSensitive(sensitive bool) {
	b := gobject.GBool(sensitive)
	defer b.Free()
	C.gtk_cell_renderer_set_sensitive(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellRenderer) GetAlignment() (xalign, yalign float32) {
	var x C.gfloat
	var y C.gfloat
	C.gtk_cell_renderer_get_alignment(self.object, &x, &y)

	xalign = float32(x)
	yalign = float32(y)
	return
}

func (self *CellRenderer) SetAlignment(xalign, yalign float32) {
	C.gtk_cell_renderer_set_alignment(self.object, C.gfloat(xalign), C.gfloat(yalign))
}

func (self *CellRenderer) GetPadding() (xpad, ypad int) {
	var x C.gint
	var y C.gint
	C.gtk_cell_renderer_get_padding(self.object, &x, &y)
	xpad = int(x)
	ypad = int(y)
	return
}

func (self *CellRenderer) SetPadding(xpad, ypad int) {
	C.gtk_cell_renderer_set_padding(self.object, C.gint(xpad), C.gint(ypad))
}

func (self *CellRenderer) GetState(w WidgetLike, gtk_CellRendererState int) int {
	return int(C.gtk_cell_renderer_get_state(self.object, w.W().object, C.GtkCellRendererState(gtk_CellRendererState)))
}

func (self *CellRenderer) IsActivatable() bool {
	b := C.gtk_cell_renderer_is_activatable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRenderer) GetPreferredHeight(w WidgetLike) (minimumSize, naturalSize int) {
	var m, n C.gint
	C.gtk_cell_renderer_get_preferred_height(self.object, w.W().object, &m, &n)
	minimumSize = int(m)
	naturalSize = int(n)
	return
}

func (self *CellRenderer) GetPreferredHeightForWidth(w WidgetLike, width int) (minimumHeight, naturalHeight int) {
	var m, n C.gint
	C.gtk_cell_renderer_get_preferred_height_for_width(self.object, w.W().object, C.gint(width), &m, &n)
	minimumHeight = int(m)
	naturalHeight = int(n)
	return
}

func (self *CellRenderer) GetPreferredSize(w WidgetLike) (minSize, natSize Requisition) {
	var m, n C.GtkRequisition
	C.gtk_cell_renderer_get_preferred_size(self.object, w.W().object, &m, &n)

	minSize.Height = int(m.height)
	minSize.Width = int(m.width)
	natSize.Height = int(n.height)
	natSize.Width = int(n.width)
	return
}

func (self *CellRenderer) GetPreferredWidth(w WidgetLike) (minSize, natSize int) {
	var m, n C.gint
	C.gtk_cell_renderer_get_preferred_width(self.object, w.W().object, &m, &n)
	minSize = int(m)
	natSize = int(n)
	return
}

func (self *CellRenderer) GetPreferredWidthForHeight(w WidgetLike, height int) (minWidth, natWidth int) {
	var m, n C.gint
	C.gtk_cell_renderer_get_preferred_width_for_height(self.object, w.W().object, C.gint(height), &m, &n)
	minWidth = int(m)
	natWidth = int(n)
	return
}

func (self *CellRenderer) GetRequestMode() int {
	return int(C.gtk_cell_renderer_get_request_mode(self.object))
}
//////////////////////////////
// END GtkCellRenderer
////////////////////////////// }}}

// GtkCellRendererText {{{
//////////////////////////////

// GtkCellRendererText type
type CellRendererText struct {
	object *C.GtkCellRendererText
	*CellRenderer
}

func NewCellRendererText() *CellRendererText {
	cl := &CellRendererText{}
	o := C.gtk_cell_renderer_text_new()
	cl.object = C.to_GtkCellRendererText(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererTextFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererTextFinalizer(cl *CellRendererText) {
	runtime.SetFinalizer(cl, func(cl *CellRendererText) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererTextFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererText{}
	cl.object = C.to_GtkCellRendererText(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererTextFinalizer(cl)

	return cl
}

func nativeFromCellRendererText(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererText)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_TEXT, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererText) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererText) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererText) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererText) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererText) CRenderer() *CellRenderer {
	return self.CellRenderer
}

// Interface functions

func (self *CellRendererText) SetFixedHeigthFromFont(numberOfRows int) {
	C.gtk_cell_renderer_text_set_fixed_height_from_font(self.object, C.gint(numberOfRows))
}

//////////////////////////////
// END GtkCellRendererText
////////////////////////////// }}}

// GtkCellRendererProgress {{{
//////////////////////////////

// GtkCellRendererProgress type
type CellRendererProgress struct {
	object *C.GtkCellRendererProgress
	*CellRenderer
}

func NewCellRendererProgress() *CellRendererProgress {
	cl := &CellRendererProgress{}
	o := C.gtk_cell_renderer_progress_new()
	cl.object = C.to_GtkCellRendererProgress(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererProgressFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererProgressFinalizer(cl *CellRendererProgress) {
	runtime.SetFinalizer(cl, func(cl *CellRendererProgress) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererProgressFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererProgress{}
	cl.object = C.to_GtkCellRendererProgress(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererProgressFinalizer(cl)

	return cl
}

func nativeFromCellRendererProgress(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererProgress)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_PROGRESS, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererProgress) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererProgress) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererProgress) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererProgress) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererProgress) CRenderer() *CellRenderer {
	return self.CellRenderer
}
//////////////////////////////
// END GtkCellRendererProgress
////////////////////////////// }}}

// GtkCellRendererSpinner {{{
//////////////////////////////

// GtkCellRendererSpinner type
type CellRendererSpinner struct {
	object *C.GtkCellRendererSpinner
	*CellRenderer
}

func NewCellRendererSpinner() *CellRendererSpinner {
	cl := &CellRendererSpinner{}
	o := C.gtk_cell_renderer_spinner_new()
	cl.object = C.to_GtkCellRendererSpinner(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererSpinnerFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererSpinnerFinalizer(cl *CellRendererSpinner) {
	runtime.SetFinalizer(cl, func(cl *CellRendererSpinner) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererSpinnerFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererSpinner{}
	cl.object = C.to_GtkCellRendererSpinner(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererSpinnerFinalizer(cl)

	return cl
}

func nativeFromCellRendererSpinner(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererSpinner)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_SPINNER, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererSpinner) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererSpinner) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererSpinner) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererSpinner) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererSpinner) CRenderer() *CellRenderer {
	return self.CellRenderer
}
//////////////////////////////
// END GtkCellRendererSpinner
////////////////////////////// }}}

// GtkCellRendererToggle {{{
//////////////////////////////

// GtkCellRendererToggle type
type CellRendererToggle struct {
	object *C.GtkCellRendererToggle
	*CellRenderer
}

func NewCellRendererToggle() *CellRendererToggle {
	cl := &CellRendererToggle{}
	o := C.gtk_cell_renderer_toggle_new()
	cl.object = C.to_GtkCellRendererToggle(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererToggleFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererToggleFinalizer(cl *CellRendererToggle) {
	runtime.SetFinalizer(cl, func(cl *CellRendererToggle) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererToggleFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererToggle{}
	cl.object = C.to_GtkCellRendererToggle(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererToggleFinalizer(cl)

	return cl
}

func nativeFromCellRendererToggle(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererToggle)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_TOGGLE, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererToggle) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererToggle) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererToggle) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererToggle) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererToggle) CRenderer() *CellRenderer {
	return self.CellRenderer
}

// CellRendererToggle Interface

func (self *CellRendererToggle) GetRadio() bool {
	b := C.gtk_cell_renderer_toggle_get_radio(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRendererToggle) SetRadio(radio bool) {
	b := gobject.GBool(radio)
	defer b.Free()
	C.gtk_cell_renderer_toggle_set_radio(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellRendererToggle) GetActive() bool {
	b := C.gtk_cell_renderer_toggle_get_active(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRendererToggle) SetActive(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_cell_renderer_toggle_set_active(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *CellRendererToggle) GetActivatable() bool {
	b := C.gtk_cell_renderer_toggle_get_activatable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *CellRendererToggle) SetActivatable(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_cell_renderer_toggle_set_activatable(self.object, *((*C.gboolean)(b.GetPtr())))
}
//////////////////////////////
// END GtkCellRendererToggle
////////////////////////////// }}}

// GtkCellRendererPixbuf {{{
//////////////////////////////

// GtkCellRendererPixbuf type
type CellRendererPixbuf struct {
	object *C.GtkCellRendererPixbuf
	*CellRenderer
}

func NewCellRendererPixbuf() *CellRendererPixbuf {
	cl := &CellRendererPixbuf{}
	o := C.gtk_cell_renderer_pixbuf_new()
	cl.object = C.to_GtkCellRendererPixbuf(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(unsafe.Pointer(o)).(*CellRenderer)
	cellRendererPixbufFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererPixbufFinalizer(cl *CellRendererPixbuf) {
	runtime.SetFinalizer(cl, func(cl *CellRendererPixbuf) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererPixbufFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererPixbuf{}
	cl.object = C.to_GtkCellRendererPixbuf(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRenderer = newCellRendererFromNative(obj).(*CellRenderer)
	cellRendererPixbufFinalizer(cl)

	return cl
}

func nativeFromCellRendererPixbuf(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererPixbuf)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_PIXBUF, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererPixbuf) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererPixbuf) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererPixbuf) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererPixbuf) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be CellRendererLike
func (self CellRendererPixbuf) CRenderer() *CellRenderer {
	return self.CellRenderer
}
//////////////////////////////
// END GtkCellRendererPixbuf
////////////////////////////// }}}

// GtkCellRendererAccel {{{
//////////////////////////////

// GtkCellRendererAccel type
type CellRendererAccel struct {
	object *C.GtkCellRendererAccel
	*CellRendererText
}

func NewCellRendererAccel() *CellRendererAccel {
	cl := &CellRendererAccel{}
	o := C.gtk_cell_renderer_accel_new()
	cl.object = C.to_GtkCellRendererAccel(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(unsafe.Pointer(o)).(*CellRendererText)
	cellRendererAccelFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererAccelFinalizer(cl *CellRendererAccel) {
	runtime.SetFinalizer(cl, func(cl *CellRendererAccel) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererAccelFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererAccel{}
	cl.object = C.to_GtkCellRendererAccel(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(obj).(*CellRendererText)
	cellRendererAccelFinalizer(cl)

	return cl
}

func nativeFromCellRendererAccel(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererAccel)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_ACCEL, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererAccel) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererAccel) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererAccel) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererAccel) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// END GtkCellRendererAccel
////////////////////////////// }}}

// GtkCellRendererCombo {{{
//////////////////////////////

// GtkCellRendererCombo type
type CellRendererCombo struct {
	object *C.GtkCellRendererCombo
	*CellRendererText
}

func NewCellRendererCombo() *CellRendererCombo {
	cl := &CellRendererCombo{}
	o := C.gtk_cell_renderer_combo_new()
	cl.object = C.to_GtkCellRendererCombo(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(unsafe.Pointer(o)).(*CellRendererText)
	cellRendererComboFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererComboFinalizer(cl *CellRendererCombo) {
	runtime.SetFinalizer(cl, func(cl *CellRendererCombo) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererComboFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererCombo{}
	cl.object = C.to_GtkCellRendererCombo(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(obj).(*CellRendererText)
	cellRendererComboFinalizer(cl)

	return cl
}

func nativeFromCellRendererCombo(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererCombo)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_COMBO, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererCombo) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererCombo) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererCombo) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererCombo) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// END GtkCellRendererCombo
////////////////////////////// }}}

// GtkCellRendererSpin {{{
//////////////////////////////

// GtkCellRendererSpin type
type CellRendererSpin struct {
	object *C.GtkCellRendererSpin
	*CellRendererText
}

func NewCellRendererSpin() *CellRendererSpin {
	cl := &CellRendererSpin{}
	o := C.gtk_cell_renderer_spin_new()
	cl.object = C.to_GtkCellRendererSpin(unsafe.Pointer(o))

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(unsafe.Pointer(o)).(*CellRendererText)
	cellRendererSpinFinalizer(cl)

	return cl
}

// Clear CellRenderer struct when it goes out of reach
func cellRendererSpinFinalizer(cl *CellRendererSpin) {
	runtime.SetFinalizer(cl, func(cl *CellRendererSpin) { gobject.Unref(cl) })
}

// Conversion functions
func newCellRendererSpinFromNative(obj unsafe.Pointer) interface{} {
	cl := &CellRendererSpin{}
	cl.object = C.to_GtkCellRendererSpin(obj)

	if gobject.IsObjectFloating(cl) {
		gobject.RefSink(cl)
	} else {
		gobject.Ref(cl)
	}
	cl.CellRendererText = newCellRendererTextFromNative(obj).(*CellRendererText)
	cellRendererSpinFinalizer(cl)

	return cl
}

func nativeFromCellRendererSpin(cl interface{}) *gobject.GValue {
	cellRend, ok := cl.(*CellRendererSpin)
	if ok {
		gv := gobject.CreateCGValue(GtkType.CELL_RENDERER_SPIN, cellRend.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self CellRendererSpin) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self CellRendererSpin) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self CellRendererSpin) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self CellRendererSpin) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// END GtkCellRendererSpin
////////////////////////////// }}}

// GtkTreeViewColumn {{{
//////////////////////////////

// GtkTreeViewColumn type
type TreeViewColumn struct {
	object *C.GtkTreeViewColumn
}

func NewTreeViewColumn() *TreeViewColumn {
	tc := &TreeViewColumn{}
	tc.object = C.gtk_tree_view_column_new()

	if gobject.IsObjectFloating(tc) {
		gobject.RefSink(tc)
	}
	treeViewColumnFinalizer(tc)

	return tc
}
// TODO: gtk_tree_column_new_with_area

func NewTreeViewColumnWithAttributes(title string, cell CellRendererLike, attr A) *TreeViewColumn {
	tc := NewTreeViewColumn()
	tc.SetTitle(title)
	tc.PackStart(cell, true)
	tc.SetAttributes(cell, attr)

	return tc
}

// Clear TreeViewColumn struct when it goes out of reach
func treeViewColumnFinalizer(tc *TreeViewColumn) {
	runtime.SetFinalizer(tc, func(tc *TreeViewColumn) { gobject.Unref(tc) })
}

// Conversion functions
func newTreeViewColumnFromNative(obj unsafe.Pointer) interface{} {
	tc := &TreeViewColumn{}
	tc.object = C.to_GtkTreeViewColumn(obj)

	if gobject.IsObjectFloating(tc) {
		gobject.RefSink(tc)
	} else {
		gobject.Ref(tc)
	}
	treeViewColumnFinalizer(tc)

	return tc
}

func nativeFromTreeViewColumn(tc interface{}) *gobject.GValue {
	tvc, ok := tc.(*TreeViewColumn)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_VIEW_COLUMN, tvc.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeViewColumn) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeViewColumn) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeViewColumn) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeViewColumn) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// TreeViewColumn interface

func (self *TreeViewColumn) PackStart(cell CellRendererLike, expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_tree_view_column_pack_start(self.object, cell.CRenderer().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) PackEnd(cell CellRendererLike, expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_tree_view_column_pack_start(self.object, cell.CRenderer().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(self.object)
}

func (self *TreeViewColumn) AddAttribute(cellRenderer CellRendererLike, attribute string, column int) {
	s := gobject.GString(attribute)
	defer s.Free()
	C.gtk_tree_view_column_add_attribute(self.object, cellRenderer.CRenderer().object, (*C.gchar)(s.GetPtr()), C.gint(column))
}

func (self *TreeViewColumn) SetAttributes(cellRenderer CellRendererLike, attrs A) {
	self.ClearAttributes(cellRenderer)
	for _, cellAttr := range attrs {
		self.AddAttribute(cellRenderer, cellAttr.Attribute, cellAttr.Column)
	}
}

//TODO: gtk_tree_view_column_set_data_func

func (self *TreeViewColumn) ClearAttributes(cellRenderer CellRendererLike) {
	C.gtk_tree_view_column_clear_attributes(self.object, cellRenderer.CRenderer().object)
}

func (self *TreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(self.object, C.gint(spacing))
}

func (self *TreeViewColumn) GetSpacing() int {
	return int(C.gtk_tree_view_column_get_spacing(self.object))
}

func (self *TreeViewColumn) SetVisible(visible bool) {
	b := gobject.GBool(visible)
	defer b.Free()
	C.gtk_tree_view_column_set_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetVisible() bool {
	b := C.gtk_tree_view_column_get_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetResizable(resizable bool) {
	b := gobject.GBool(resizable)
	defer b.Free()
	C.gtk_tree_view_column_set_resizable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetResizable() bool {
	b := C.gtk_tree_view_column_get_resizable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetSizing(gtk_TreeViewColumnSizing int) {
	C.gtk_tree_view_column_set_sizing(self.object, C.GtkTreeViewColumnSizing(gtk_TreeViewColumnSizing))
}

func (self *TreeViewColumn) GetSizing() int {
	return int(C.gtk_tree_view_column_get_sizing(self.object))
}

func (self *TreeViewColumn) GetWidth() int {
	return int(C.gtk_tree_view_column_get_width(self.object))
}

func (self *TreeViewColumn) GetFixedWidth() int {
	return int(C.gtk_tree_view_column_get_fixed_width(self.object))
}

func (self *TreeViewColumn) SetFixedWidth(fixedWidth int) {
	C.gtk_tree_view_column_set_fixed_width(self.object, C.gint(fixedWidth))
}

func (self *TreeViewColumn) SetMinWidth(minWidth int) {
	C.gtk_tree_view_column_set_min_width(self.object, C.gint(minWidth))
}

func (self *TreeViewColumn) GetMinWidth() int {
	return int(C.gtk_tree_view_column_get_min_width(self.object))
}

func (self *TreeViewColumn) SetMaxWidth(maxWidth int) {
	C.gtk_tree_view_column_set_max_width(self.object, C.gint(maxWidth))
}

func (self *TreeViewColumn) GetMaxWidth() int {
	return int(C.gtk_tree_view_column_get_max_width(self.object))
}

func (self *TreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(self.object)
}

func (self *TreeViewColumn) SetTitle(title string) {
	ctitle := gobject.GString(title)
	defer ctitle.Free()
	C.gtk_tree_view_column_set_title(self.object, (*C.gchar)(ctitle.GetPtr()))
}

func (self *TreeViewColumn) GetTitle() string {
	s := C.gtk_tree_view_column_get_title(self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *TreeViewColumn) SetExpand(expand bool) {
	b := gobject.GBool(expand)
	defer b.Free()
	C.gtk_tree_view_column_set_expand(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetExpand() bool {
	b := C.gtk_tree_view_column_get_expand(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetClickable(clickable bool) {
	b := gobject.GBool(clickable)
	defer b.Free()
	C.gtk_tree_view_column_set_clickable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetClickable() bool {
	b := C.gtk_tree_view_column_get_clickable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetWidget(w WidgetLike) {
	var cw *C.GtkWidget = nil
	if w != nil {
		cw = w.W().object
	}
	C.gtk_tree_view_column_set_widget(self.object, cw)
}

func (self *TreeViewColumn) GetWidget() WidgetLike {
	w := C.gtk_tree_view_column_get_widget(self.object)
	if w == nil {
		return nil
	}

	wid, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *TreeViewColumn) GetButton() WidgetLike {
	w := C.gtk_tree_view_column_get_button(self.object)

	wid, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return wid.(WidgetLike)
	}
	return nil
}

func (self *TreeViewColumn) SetAlignment(xalign float32) {
	C.gtk_tree_view_column_set_alignment(self.object, C.gfloat(xalign))
}

func (self *TreeViewColumn) GetAlignment() float32 {
	return float32(C.gtk_tree_view_column_get_alignment(self.object))
}

func (self *TreeViewColumn) SetReorderable(reorderable bool) {
	b := gobject.GBool(reorderable)
	defer b.Free()
	C.gtk_tree_view_column_set_reorderable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetReorderable() bool {
	b := C.gtk_tree_view_column_get_reorderable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetSortColumnId(sortColumnId int) {
	C.gtk_tree_view_column_set_sort_column_id(self.object, C.gint(sortColumnId))
}

func (self *TreeViewColumn) GetSortColumnId() int {
	return int(C.gtk_tree_view_column_get_sort_column_id(self.object))
}

func (self *TreeViewColumn) SetSortIndicator(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_tree_view_column_set_sort_indicator(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeViewColumn) GetSortIndicator() bool {
	b := C.gtk_tree_view_column_get_sort_indicator(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) SetSortOrder(gtk_SortType int) {
	C.gtk_tree_view_column_set_sort_order(self.object, C.GtkSortType(gtk_SortType))
}

func (self *TreeViewColumn) GetSortOrder() int {
	return int(C.gtk_tree_view_column_get_sort_order(self.object))
}

func (self *TreeViewColumn) CellSetCellData(model TreeModelLike, iter *TreeIter, isExpander, isExpanded bool) {
	b1 := gobject.GBool(isExpander)
	defer b1.Free()
	b2 := gobject.GBool(isExpanded)
	defer b2.Free()
	C.gtk_tree_view_column_cell_set_cell_data(self.object, model.ITreeModel().object, &iter.object, *((*C.gboolean)(b1.GetPtr())),
		*((*C.gboolean)(b2.GetPtr())))
}

func (self *TreeViewColumn) CellGetSize(rec gdk3.Rectangle) (xOffset, yOffset, width, height int) {
	rectangle := gobject.ConvertToC(rec)
	defer rectangle.Free()

	var x, y, w, h C.gint
	C.gtk_tree_view_column_cell_get_size(self.object, (*C.GdkRectangle)(rectangle.GetPtr()), &x, &y, &w, &h)
	xOffset = int(x)
	yOffset = int(y)
	width = int(w)
	height = int(h)
	return
}

func (self *TreeViewColumn) CellGetPosition(renderer CellRendererLike) (bool, int, int) {
	var o, w C.gint
	b := C.gtk_tree_view_column_cell_get_position(self.object, renderer.CRenderer().object, &o, &w)
	retBool := gobject.GoBool(unsafe.Pointer(&b))
	return retBool, int(o), int(w)
}

func (self *TreeViewColumn) CellIsVisible() bool {
	b := C.gtk_tree_view_column_cell_is_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeViewColumn) FocusCell(renderer CellRendererLike) {
	C.gtk_tree_view_column_focus_cell(self.object, renderer.CRenderer().object)
}

func (self *TreeViewColumn) QueueResize() {
	C.gtk_tree_view_column_queue_resize(self.object)
}

func (self *TreeViewColumn) GetTreeView() WidgetLike {
	w := C.gtk_tree_view_column_get_tree_view(self.object)
	if w == nil {
		return nil
	}

	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return widget.(WidgetLike)
	}
	return nil
}

func (self *TreeViewColumn) GetXOffset() int {
	return int(C.gtk_tree_view_column_get_x_offset(self.object))
}
//////////////////////////////
// End GtkTreeViewColumn
////////////////////////////// }}}

// GtkTreeSelection {{{
//////////////////////////////

// GtkTreeSelection type
type TreeSelection struct {
	object *C.GtkTreeSelection
}

// Clear TreeSelection struct when it goes out of reach
func treeSelectionFinalizer(ts *TreeSelection) {
	runtime.SetFinalizer(ts, func(ts *TreeSelection) { gobject.Unref(ts) })
}

// Conversion function
func newTreeSelectionFromNative(obj unsafe.Pointer) interface{} {
	ts := &TreeSelection{}
	ts.object = C.to_GtkTreeSelection(obj)

	if gobject.IsObjectFloating(ts) {
		gobject.RefSink(ts)
	} else {
		gobject.Ref(ts)
	}
	treeSelectionFinalizer(ts)

	return ts
}

func nativeFromTreeSelection(ts interface{}) *gobject.GValue {
	treeSelection, ok := ts.(*TreeSelection)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_SELECTION, treeSelection.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeSelection) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeSelection) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeSelection) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeSelection) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
//////////////////////////////
// End GtkTreeSelection
////////////////////////////// }}}

// GtkTreeView {{{
//////////////////////////////

// GtkTreeView type
type TreeView struct {
	object *C.GtkTreeView
	*Container
}

func NewTreeView() *TreeView {
	tw := &TreeView{}
	o := C.gtk_tree_view_new()
	tw.object = C.to_GtkTreeView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(tw) {
		gobject.RefSink(tw)
	}
	tw.Container = NewContainer(unsafe.Pointer(o))
	treeViewFinalizer(tw)

	return tw
}

func NewTreeViewWithModel(model TreeModelLike) *TreeView {
	tw := &TreeView{}
	o := C.gtk_tree_view_new_with_model(model.ITreeModel().object)
	tw.object = C.to_GtkTreeView(unsafe.Pointer(o))

	if gobject.IsObjectFloating(tw) {
		gobject.RefSink(tw)
	}
	tw.Container = NewContainer(unsafe.Pointer(o))
	treeViewFinalizer(tw)

	return tw
}

// Clear TreeView struct when it goes out of reach
func treeViewFinalizer(tw *TreeView) {
	runtime.SetFinalizer(tw, func(tw *TreeView) { gobject.Unref(tw) })
}

// Conversion functions
func newTreeViewFromNative(obj unsafe.Pointer) interface{} {
	tw := &TreeView{}
	tw.object = C.to_GtkTreeView(obj)

	if gobject.IsObjectFloating(tw) {
		gobject.RefSink(tw)
	} else {
		gobject.Ref(tw)
	}
	tw.Container = NewContainer(obj)
	treeViewFinalizer(tw)

	return tw
}

func nativeFromTreeView(tw interface{}) *gobject.GValue {
	treeV, ok := tw.(*TreeView)
	if ok {
		gv := gobject.CreateCGValue(GtkType.TREE_VIEW, treeV.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self TreeView) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self TreeView) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self TreeView) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self TreeView) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self TreeView) C() *Container {
	return self.Container
}

// TreeView interface

func (self *TreeView) GetLevelIndentation() int {
	return int(C.gtk_tree_view_get_level_indentation(self.object))
}

func (self *TreeView) GetShowExpanders() bool {
	b := C.gtk_tree_view_get_show_expanders(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetLevelIndentation(indentation int) {
	C.gtk_tree_view_set_level_indentation(self.object, C.gint(indentation))
}

func (self *TreeView) SetShowExpanders(enabled bool) {
	b := gobject.GBool(enabled)
	defer b.Free()
	C.gtk_tree_view_set_show_expanders(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetModel() TreeModelLike {
	m := C.gtk_tree_view_get_model(self.object)
	if m == nil {
		return nil
	}
	model, err := gobject.ConvertToGo(unsafe.Pointer(m))

	if err == nil {
		return model.(TreeModelLike)
	}
	return nil
}

func (self *TreeView) SetModel(model TreeModelLike) {
	C.gtk_tree_view_set_model(self.object, model.ITreeModel().object)
}

func (self *TreeView) GetSelection() *TreeSelection {
	s := C.gtk_tree_view_get_selection(self.object)

	sel, err := gobject.ConvertToGo(unsafe.Pointer(s))
	if err == nil {
		return sel.(*TreeSelection)
	}
	return nil
}

func (self *TreeView) GetHeadersVisible() bool {
	b := C.gtk_tree_view_get_headers_visible(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetHeadersVisible(headersVisible bool) {
	b := gobject.GBool(headersVisible)
	defer b.Free()
	C.gtk_tree_view_set_headers_visible(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) ColumnsAutoSize() {
	C.gtk_tree_view_columns_autosize(self.object)
}

func (self *TreeView) GetHeadersClickable() bool {
	b := C.gtk_tree_view_get_headers_clickable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) SetHeadersClickable(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_tree_view_set_headers_clickable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) SetRulesHint(setting bool) {
	b := gobject.GBool(setting)
	defer b.Free()
	C.gtk_tree_view_set_rules_hint(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) GetRulesHint() bool {
	b := C.gtk_tree_view_get_rules_hint(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *TreeView) AppendColumn(column *TreeViewColumn) int {
	return int(C.gtk_tree_view_append_column(self.object, column.object))
}

func (self *TreeView) RemoveColumn(colum *TreeViewColumn) int {
	return int(C.gtk_tree_view_remove_column(self.object, colum.object))
}

func (self *TreeView) InsertColumn(colum *TreeViewColumn, position int) int {
	return int(C.gtk_tree_view_insert_column(self.object, colum.object, C.gint(position)))
}

//TODO: gtk_tree_view_insert_column_with_attributes
//TODO: gtk_tree_view_insert_column_with_data_func

func (self *TreeView) GetColumn(n int) *TreeViewColumn {
	tc := C.gtk_tree_view_get_column(self.object, C.gint(n))
	if tc == nil {
		return nil
	}

	column, err := gobject.ConvertToGo(unsafe.Pointer(tc))
	if err == nil {
		return column.(*TreeViewColumn)
	}
	return nil
}

//TODO: gtk_tree_view_get_columns

func (self *TreeView) MoveColumnAfter(column, baseColumn *TreeViewColumn) {
	var base *C.GtkTreeViewColumn = nil
	if baseColumn != nil {
		base = baseColumn.object
	}
	C.gtk_tree_view_move_column_after(self.object, column.object, base)
}

func (self *TreeView) SetExpanderColumn(column *TreeViewColumn) {
	var col *C.GtkTreeViewColumn = nil
	if column != nil {
		col = column.object
	}
	C.gtk_tree_view_set_expander_column(self.object, col)
}

func (self *TreeView) GetExpanderColumn() *TreeViewColumn {
	ec := C.gtk_tree_view_get_expander_column(self.object)

	if ec != nil {
		expanderCol, err := gobject.ConvertToGo(unsafe.Pointer(ec))
		if err == nil {
			return expanderCol.(*TreeViewColumn)
		}
		return nil
	}
	return nil
}

//TODO: gtk_tree_view_set_column_drag_function

func (self *TreeView) ScrollToPoint(xTree, yTree int) {
	C.gtk_tree_view_scroll_to_point(self.object, C.gint(xTree), C.gint(yTree))
}

func (self *TreeView) ScrollToCell(path *TreePath, column *TreeViewColumn, useAlign bool, rowAlign, colAlign float32) {
	var p *C.GtkTreePath = nil
	var c *C.GtkTreeViewColumn = nil
	b := gobject.GBool(useAlign)
	defer b.Free()

	if path != nil {
		p = path.object
	}

	if column != nil {
		c = column.object
	}
	C.gtk_tree_view_scroll_to_cell(self.object, p, c, *((*C.gboolean)(b.GetPtr())), C.gfloat(rowAlign), C.gfloat(colAlign))
}

func (self *TreeView) SetCursor(path *TreePath, focusColumn *TreeViewColumn, startEditing bool) {
	var c *C.GtkTreeViewColumn = nil
	if focusColumn != nil {
		c = focusColumn.object
	}

	b := gobject.GBool(startEditing)
	defer b.Free()

	C.gtk_tree_view_set_cursor(self.object, path.object, c, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) SetCursorOnCell(path *TreePath, focusColumn *TreeViewColumn, focusCell CellRendererLike, startEditing bool) {
	var c *C.GtkTreeViewColumn = nil
	var r *C.GtkCellRenderer = nil

	if focusColumn != nil {
		c = focusColumn.object
	}

	if focusCell != nil {
		r = focusCell.CRenderer().object
	}

	b := gobject.GBool(startEditing)
	defer b.Free()
	C.gtk_tree_view_set_cursor_on_cell(self.object, path.object, c, r, *((*C.gboolean)(b.GetPtr())))
}

func (self *TreeView) SetSearchColumn(column int) {
	C.gtk_tree_view_set_search_column(self.object, C.gint(column))
}
//////////////////////////////
// End GtkTreeView
////////////////////////////// }}}

// End Tree, List and Icon Grid Widgets }}}

// Layout Containers {{{

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

// GtkNotebook {{{
//////////////////////////////

// GtkNotebook type
type Notebook struct {
	object *C.GtkNotebook
	*Container
}

func NewNotebook() *Notebook {
	n := &Notebook{}
	o := C.gtk_notebook_new()
	n.object = C.to_GtkNotebook(unsafe.Pointer(o))

	if gobject.IsObjectFloating(n) {
		gobject.RefSink(n)
	}
	n.Container = NewContainer(unsafe.Pointer(o))
	notebookFinalizer(n)

	return n
}

// Clear Notebook struct when it goes out of reach
func notebookFinalizer(n *Notebook) {
	runtime.SetFinalizer(n, func(n *Notebook) { gobject.Unref(n) })
}

// Conversion functions
func newNotebookFromNative(obj unsafe.Pointer) interface{} {
	n := &Notebook{}
	n.object = C.to_GtkNotebook(obj)

	if gobject.IsObjectFloating(n) {
		gobject.RefSink(n)
	} else {
		gobject.Ref(n)
	}
	n.Container = NewContainer(obj)
	notebookFinalizer(n)

	return n
}

func nativeFromNotebook(note interface{}) *gobject.GValue {
	n, ok := note.(*Notebook)
	if ok {
		gv := gobject.CreateCGValue(GtkType.NOTEBOOK, n.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Notebook) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Notebook) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Notebook) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Notebook) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// To be container-like
func (self Notebook) C() *Container {
	return self.Container
}

// Notebook interface

func (self *Notebook) AppendPage(child, tabLabel WidgetLike) int {
	var tab *C.GtkWidget = nil
	if tabLabel != nil {
		tab = tabLabel.W().object
	}
	return int(C.gtk_notebook_append_page(self.object, child.W().object, tab))
}

func (self *Notebook) AppendPageMenu(child, tabLabel, menuLabel WidgetLike) int {
	var t, m *C.GtkWidget = nil, nil

	if tabLabel != nil {
		t = tabLabel.W().object
	}

	if menuLabel != nil {
		m = menuLabel.W().object
	}
	return int(C.gtk_notebook_append_page_menu(self.object, child.W().object, t, m))
}

func (self *Notebook) PrependPage(child, tabLabel WidgetLike) int {
	var t *C.GtkWidget
	if tabLabel != nil {
		t = tabLabel.W().object
	}
	return int(C.gtk_notebook_prepend_page(self.object, child.W().object, t))
}

func (self *Notebook) PrependPageMenu(child, tabLabel, menuLabel WidgetLike) int {
	var t, m *C.GtkWidget = nil, nil

	if tabLabel != nil {
		t = tabLabel.W().object
	}

	if menuLabel != nil {
		m = menuLabel.W().object
	}
	return int(C.gtk_notebook_prepend_page_menu(self.object, child.W().object, t, m))
}

func (self *Notebook) InsertPage(child, tabLabel WidgetLike, position int) int {
	var t *C.GtkWidget = nil

	if tabLabel != nil {
		t = tabLabel.W().object
	}
	return int(C.gtk_notebook_insert_page(self.object, child.W().object, t, C.gint(position)))
}

func (self *Notebook) InsertPageMenu(child, tabLabel, menuLabel WidgetLike, position int) int {
	var t, m *C.GtkWidget = nil, nil

	if tabLabel != nil {
		t = tabLabel.W().object
	}

	if menuLabel != nil {
		m = menuLabel.W().object
	}
	return int(C.gtk_notebook_insert_page_menu(self.object, child.W().object, t, m, C.gint(position)))
}

func (self *Notebook) RemovePage(pageNum int) {
	C.gtk_notebook_remove_page(self.object, C.gint(pageNum))
}

func (self *Notebook) PageNum(child WidgetLike) int {
	return int(C.gtk_notebook_page_num(self.object, child.W().object))
}

func (self *Notebook) NextPage() {
	C.gtk_notebook_next_page(self.object)
}

func (self *Notebook) PrevPage() {
	C.gtk_notebook_prev_page(self.object)
}

func (self *Notebook) ReorderChild(child WidgetLike, position int) {
	C.gtk_notebook_reorder_child(self.object, child.W().object, C.gint(position))
}

func (self *Notebook) SetTabPos(gtk_PositionType int) {
	C.gtk_notebook_set_tab_pos(self.object, C.GtkPositionType(gtk_PositionType))
}

func (self *Notebook) SetShowTabs(showTabs bool) {
	b := gobject.GBool(showTabs)
	defer b.Free()
	C.gtk_notebook_set_show_tabs(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) SetShowBorder(showBorder bool) {
	b := gobject.GBool(showBorder)
	defer b.Free()
	C.gtk_notebook_set_show_border(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) SetScrollable(scrollable bool) {
	b := gobject.GBool(scrollable)
	defer b.Free()
	C.gtk_notebook_set_scrollable(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(self.object)
}

func (self *Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(self.object)
}

func (self *Notebook) GetCurrentPage() int {
	return int(C.gtk_notebook_get_current_page(self.object))
}

func (self *Notebook) GetMenuLabel(child WidgetLike) WidgetLike {
	ml := C.gtk_notebook_get_menu_label(self.object, child.W().object)
	if ml == nil {
		return nil
	}

	l, err := gobject.ConvertToGo(unsafe.Pointer(ml))
	if err == nil {
		return l.(WidgetLike)
	}
	return nil
}

func (self *Notebook) GetNthPage(pageNum int) WidgetLike {
	w := C.gtk_notebook_get_nth_page(self.object, C.gint(pageNum))
	if w == nil {
		return nil
	}

	child, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return child.(WidgetLike)
	}
	return nil
}

func (self *Notebook) GetNPages() int {
	return int(C.gtk_notebook_get_n_pages(self.object))
}

func (self *Notebook) GetTabLabel(child WidgetLike) WidgetLike {
	w := C.gtk_notebook_get_tab_label(self.object, child.W().object)
	if w == nil {
		return nil
	}

	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return widget.(WidgetLike)
	}
	return nil
}

func (self *Notebook) SetMenuLabel(child, menuLabel WidgetLike) {
	var m *C.GtkWidget = nil
	if menuLabel != nil {
		m = menuLabel.W().object
	}
	C.gtk_notebook_set_menu_label(self.object, child.W().object, m)
}

func (self *Notebook) SetMenuLabelText(child WidgetLike, menuText string) {
	s := gobject.GString(menuText)
	defer s.Free()
	C.gtk_notebook_set_menu_label_text(self.object, child.W().object, (*C.gchar)(s.GetPtr()))
}

func (self *Notebook) SetTabLabel(child, tabLabel WidgetLike) {
	var l *C.GtkWidget = nil
	if tabLabel != nil {
		l = tabLabel.W().object
	}
	C.gtk_notebook_set_tab_label(self.object, child.W().object, l)
}

func (self *Notebook) SetTabLabelText(child WidgetLike, tabText string) {
	s := gobject.GString(tabText)
	defer s.Free()
	C.gtk_notebook_set_tab_label_text(self.object, child.W().object, (*C.gchar)(s.GetPtr()))
}

func (self *Notebook) SetTabReorderable(child WidgetLike, reorderable bool) {
	b := gobject.GBool(reorderable)
	defer b.Free()
	C.gtk_notebook_set_tab_reorderable(self.object, child.W().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) SetTabDetachable(child WidgetLike, detachable bool) {
	b := gobject.GBool(detachable)
	defer b.Free()
	C.gtk_notebook_set_tab_detachable(self.object, child.W().object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Notebook) GetMenuLabelText(child WidgetLike) string {
	s := C.gtk_notebook_get_menu_label_text(self.object, child.W().object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Notebook) GetScrollable() bool {
	b := C.gtk_notebook_get_scrollable(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetShowBorder() bool {
	b := C.gtk_notebook_get_show_border(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetShowTabs() bool {
	b := C.gtk_notebook_get_show_tabs(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetTabLabelText(child WidgetLike) string {
	s := C.gtk_notebook_get_tab_label_text(self.object, child.W().object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Notebook) GetTabPos() (gtk_PositionType int) {
	return int(C.gtk_notebook_get_tab_pos(self.object))
}

func (self *Notebook) GetTabReorderable(child WidgetLike) bool {
	b := C.gtk_notebook_get_tab_reorderable(self.object, child.W().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetTabDetachable(child WidgetLike) bool {
	b := C.gtk_notebook_get_tab_detachable(self.object, child.W().object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Notebook) GetTabHborder() uint {
	return uint(C.gtk_notebook_get_tab_hborder(self.object))
}

func (self *Notebook) GetTabVborder() uint {
	return uint(C.gtk_notebook_get_tab_vborder(self.object))
}

func (self *Notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(self.object, C.gint(pageNum))
}

func (self *Notebook) SetGroupName(groupName string) {
	if groupName == "" {
		C.gtk_notebook_set_group_name(self.object, nil)
	}
	s := gobject.GString(groupName)
	defer s.Free()

	C.gtk_notebook_set_group_name(self.object, (*C.gchar)(s.GetPtr()))
}

func (self *Notebook) GetGroupName() string {
	s := C.gtk_notebook_get_group_name(self.object)
	if s == nil {
		return ""
	}
	return gobject.GoString(unsafe.Pointer(s))
}

func (self *Notebook) SetActionWidget(w WidgetLike, gtk_Pack int) {
	C.gtk_notebook_set_action_widget(self.object, w.W().object, C.GtkPackType(gtk_Pack))
}

func (self *Notebook) GetActionWidget(gtk_Pack int) WidgetLike {
	w := C.gtk_notebook_get_action_widget(self.object, C.GtkPackType(gtk_Pack))
	if w == nil {
		return nil
	}
	widget, err := gobject.ConvertToGo(unsafe.Pointer(w))
	if err == nil {
		return widget.(WidgetLike)
	}
	return nil
}
//////////////////////////////
// End GtkNotebook
////////////////////////////// }}}

// End Layout Containers }}}

// Ornaments {{{

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

// End Ornaments }}}

// Scrolling {{{

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

// End Scrolling }}}

// Miscellaneous {{{

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
	return float64(C._gtk_adjustment_get_minimum_increment(self.object))
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

// Closure add remove funcs to/from _closures map {{{
///////////////////////////////////////////////////////////

func addToClosures(key int64, f gobject.ClosureFunc) {
	_closures[key] = f
}

func removeFromClosures(key int64) {
	if _, ok := _closures[key]; ok {
		delete(_closures, key)
	}
}
//////////////////////////////
// End Closure add/remove
////////////////////////////// }}}

// End Miscellaneous }}}

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
//////////////////////////////
// END GtkApplication
////////////////////////////// }}}

// Exported functions {{{
//////////////////////////////

//export _gtk_callback
func _gtk_callback(widget, data unsafe.Pointer) {
	if w, err := gobject.ConvertToGo(widget); err == nil {
		id := int64(*((*C.gint64)(data)))
		_closures[id]([]interface{}{w})
	}

}

//export _gtk_entry_completion_match_func
func _gtk_entry_completion_match_func(completion, key, iter, userData unsafe.Pointer) C.gboolean {
	entryCompletion, _ := gobject.ConvertToGo(completion)
	s := gobject.GoString(key)
	inIter := &TreeIter{*((*C.GtkTreeIter)(iter))}

	closureId := *((*C.gint64)(userData))

	var res bool
	if c, ok := _closures[int64(closureId)]; ok {
		res = c([]interface{}{entryCompletion, s, inIter})
	}
	b := gobject.GBool(res)
	defer b.Free()

	return *((*C.gboolean)(b.GetPtr()))
}

//export _gtk_cell_callback
func _gtk_cell_callback(renderer, data unsafe.Pointer) C.gboolean {
	id := *((*C.gint64)(data))

	rend, _ := gobject.ConvertToGo(renderer)

	var res bool
	if c, ok := _closures[int64(id)]; ok {
		res = c([]interface{}{rend})
	}
	b := gobject.GBool(res)
	defer b.Free()

	return *((*C.gboolean)(b.GetPtr()))
}

//////////////////////////////
// END Exported functions
////////////////////////////// }}}

// GTK3 MODULE init function {{{
func init() {
	// Initialiize map for closures
	_closures = make(map[int64]gobject.ClosureFunc)

	// Register GtkWidget type
	gobject.RegisterCType(GtkType.WIDGET, newWidgetFromNative)
	// Register GtkApplicaton type
	gobject.RegisterCType(GtkType.APPLICATION, appFromNative)
	gobject.RegisterGoType(GtkType.APPLICATION, nativeFromApp)

	// Register GtkWindow type
	gobject.RegisterCType(GtkType.WINDOW, newWindowFromNative)
	gobject.RegisterGoType(GtkType.WINDOW, nativeFromWindow)

	// Register GtkInvisible type
	gobject.RegisterCType(GtkType.INVISIBLE, newInvisibleFromNative)
	gobject.RegisterGoType(GtkType.INVISIBLE, nativeFromInvisible)

	// Register GtkAssistant type
	gobject.RegisterCType(GtkType.ASSISTANT, newAssistantFromNative)
	gobject.RegisterGoType(GtkType.ASSISTANT, nativeFromAssistant)

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

	// Register GtkProgressBar type
	gobject.RegisterCType(GtkType.PROGRESS_BAR, newProgressBarFromNative)
	gobject.RegisterGoType(GtkType.PROGRESS_BAR, nativeFromProgressBar)

	// Register GtkImage type
	gobject.RegisterCType(GtkType.IMAGE, newImageFromNative)
	gobject.RegisterGoType(GtkType.IMAGE, nativeFromImage)

	// Register GtkButton type
	gobject.RegisterCType(GtkType.BUTTON, newButtonFromNative)
	gobject.RegisterGoType(GtkType.BUTTON, nativeFromButton)

	//Register GtkToggleButton type
	gobject.RegisterCType(GtkType.TOGGLE_BUTTON, newToggleButtonFromNative)
	gobject.RegisterGoType(GtkType.TOGGLE_BUTTON, nativeFromToggleButton)

	// Register GtkCheckButton type
	gobject.RegisterCType(GtkType.CHECK_BUTTON, newCheckButtonFromNative)
	gobject.RegisterGoType(GtkType.CHECK_BUTTON, nativeFromCheckButton)

	// Register GtkEntryBuffer type
	gobject.RegisterCType(GtkType.ENTRY_BUFFER, newEntryBufferFromNative)
	gobject.RegisterGoType(GtkType.ENTRY_BUFFER, nativeFromEntryBuffer)

	// Register GtkEntry type
	gobject.RegisterCType(GtkType.ENTRY, newEntryFromNative)
	gobject.RegisterGoType(GtkType.ENTRY, nativeFromEntry)

	// Register GtkEntryCompletion type
	gobject.RegisterCType(GtkType.ENTRY_COMPLETION, newEntryCompletionFromNative)
	gobject.RegisterGoType(GtkType.ENTRY_COMPLETION, nativeFromEntryCompletion)

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

	// Register GtkListStore type
	gobject.RegisterCType(GtkType.LIST_STORE, newListStoreFromNative)
	gobject.RegisterGoType(GtkType.LIST_STORE, nativeFromListStore)

	// Register GtkTreeStore type
	gobject.RegisterCType(GtkType.TREE_STORE, newTreeStoreFromNative)
	gobject.RegisterGoType(GtkType.TREE_STORE, nativeFromTreeStore)

	// Register GtkCellArea type
	gobject.RegisterCType(GtkType.CELL_AREA, newCellAreaFromNative)
	gobject.RegisterGoType(GtkType.CELL_AREA, nativeFromCellArea)

	// Register GtkCellAreaContext type
	gobject.RegisterCType(GtkType.CELL_AREA_CONTEXT, newCellAreaContextFromNative)
	gobject.RegisterGoType(GtkType.CELL_AREA_CONTEXT, nativeFromCellAreaContext)

	// Register GtkCellRenderer type
	gobject.RegisterCType(GtkType.CELL_RENDERER, newCellRendererFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER, nativeFromCellRenderer)

	// Register GtkCellRendererText type
	gobject.RegisterCType(GtkType.CELL_RENDERER_TEXT, newCellRendererTextFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_TEXT, nativeFromCellRendererText)

	// Register GtkCellRendererProgress type
	gobject.RegisterCType(GtkType.CELL_RENDERER_PROGRESS, newCellRendererProgressFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_PROGRESS, nativeFromCellRendererProgress)

	// Register GtkCellRendererSpinner type
	gobject.RegisterCType(GtkType.CELL_RENDERER_SPINNER, newCellRendererSpinnerFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_SPINNER, nativeFromCellRendererSpinner)

	// Register GtkCellRendererToggle type
	gobject.RegisterCType(GtkType.CELL_RENDERER_TOGGLE, newCellRendererToggleFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_TOGGLE, nativeFromCellRendererToggle)

	// Register GtkCellRendererPixbuf type
	gobject.RegisterCType(GtkType.CELL_RENDERER_PIXBUF, newCellRendererPixbufFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_PIXBUF, nativeFromCellRendererPixbuf)

	// Register GtkCellRendererAccel type
	gobject.RegisterCType(GtkType.CELL_RENDERER_ACCEL, newCellRendererAccelFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_ACCEL, nativeFromCellRendererAccel)

	// Register GtkCellRendererCombo type
	gobject.RegisterCType(GtkType.CELL_RENDERER_COMBO, newCellRendererComboFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_COMBO, nativeFromCellRendererCombo)

	// Register GtkCellRendererSpin type
	gobject.RegisterCType(GtkType.CELL_RENDERER_SPIN, newCellRendererSpinFromNative)
	gobject.RegisterGoType(GtkType.CELL_RENDERER_SPIN, nativeFromCellRendererSpin)

	// Register GtkTreeViewColumn type
	gobject.RegisterCType(GtkType.TREE_VIEW_COLUMN, newTreeViewColumnFromNative)
	gobject.RegisterGoType(GtkType.TREE_VIEW_COLUMN, nativeFromTreeViewColumn)

	// Register GtkTreeSelection type
	gobject.RegisterCType(GtkType.TREE_SELECTION, newTreeSelectionFromNative)
	gobject.RegisterGoType(GtkType.TREE_SELECTION, nativeFromTreeSelection)

	// Register GtkTreeView
	gobject.RegisterCType(GtkType.TREE_VIEW, newTreeViewFromNative)
	gobject.RegisterGoType(GtkType.TREE_VIEW, nativeFromTreeView)

	// Register GtkNotebook
	gobject.RegisterCType(GtkType.NOTEBOOK, newNotebookFromNative)
	gobject.RegisterGoType(GtkType.NOTEBOOK, nativeFromNotebook)
}
// End init function }}}
