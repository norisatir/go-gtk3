package gdk3

/*
#include <gdk/gdk.h>
#include <stdlib.h>

static void _gdk_threads_init() {
	gdk_threads_init();
}

static inline GdkDisplay* to_GdkDisplay(void* obj) { return GDK_DISPLAY(obj); }
static inline GdkDevice* to_GdkDevice(void* obj) { return GDK_DEVICE(obj); }
static inline GdkScreen* to_GdkScreen(void* obj) { return GDK_SCREEN(obj); }
static inline GdkWindow* to_GdkWindow(void* obj) { return GDK_WINDOW(obj); }

static GdkRGBA* _new_rgba(void) {
	GdkRGBA* r = (GdkRGBA*)malloc(sizeof(GdkRGBA));
	return r;
}

static inline guint getModifiers(GdkEventKey* ke) {
	return ke->is_modifier;
}

*/
// #cgo pkg-config: gdk-3.0
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

// Basic Types {{{

// GdkType {{{
var GdkType gdkTypes

type gdkTypes struct {
	DEVICE        gobject.GType
	RECTANGLE     gobject.GType
	RGBA          gobject.GType
	COLOR         gobject.GType
	EVENT         gobject.GType
	DISPLAY       gobject.GType
	SCREEN        gobject.GType
	WINDOW        gobject.GType
	MODIFIER_TYPE gobject.GType
}

// End GdkType }}}

// Rectangle {{{
type Rectangle struct {
	X, Y          int
	Width, Height int
}

//Conversion func
func newRectangleFromNative(obj unsafe.Pointer) interface{} {
	cr := *((*C.GdkRectangle)(obj))
	r := Rectangle{}
	r.X = int(cr.x)
	r.Y = int(cr.y)
	r.Width = int(cr.width)
	r.Height = int(cr.height)

	return r
}

func nativeFromRectangle(r interface{}) *gobject.GValue {
	if rec, ok := r.(Rectangle); ok {
		gv := gobject.CreateCGValue(GdkType.RECTANGLE, rec.ToNative())
		return gv
	}
	return nil
}

// To be BoxedLike
func (self Rectangle) GetBoxType() gobject.GType {
	return GdkType.RECTANGLE
}

// Rectangle interface

func (self Rectangle) ToNative() unsafe.Pointer {
	var cr C.GdkRectangle
	cr.x = C.int(self.X)
	cr.y = C.int(self.Y)
	cr.width = C.int(self.Width)
	cr.height = C.int(self.Height)

	return unsafe.Pointer(&cr)
}

func (self Rectangle) Intersect(r2 Rectangle) (bool, Rectangle) {
	var cr3 C.GdkRectangle
	cr1 := nativeFromRectangle(self)
	defer cr1.Free()
	cr2 := nativeFromRectangle(r2)
	defer cr2.Free()
	b := C.gdk_rectangle_intersect((*C.GdkRectangle)(cr1.GetPtr()), (*C.GdkRectangle)(cr2.GetPtr()), &cr3)
	rec := newRectangleFromNative(unsafe.Pointer(&cr3)).(Rectangle)

	return gobject.GoBool(unsafe.Pointer(&b)), rec
}

func (self Rectangle) Union(r2 Rectangle) Rectangle {
	var cr3 C.GdkRectangle
	cr1 := nativeFromRectangle(self)
	defer cr1.Free()
	cr2 := nativeFromRectangle(r2)
	defer cr2.Free()
	C.gdk_rectangle_union((*C.GdkRectangle)(cr1.GetPtr()), (*C.GdkRectangle)(cr2.GetPtr()), &cr3)
	rec := newRectangleFromNative(unsafe.Pointer(&cr3)).(Rectangle)

	return rec
}

// End Rectangle }}}

// GdkRGBA {{{

// GdkRGBA type
type RGBA struct {
	Red   float64
	Green float64
	Blue  float64
	Alpha float64
}

//Conversion func
func newRGBAFromNative(obj unsafe.Pointer) interface{} {
	crgba := *((*C.GdkRGBA)(obj))

	rgba := RGBA{}
	rgba.Red = float64(crgba.red)
	rgba.Green = float64(crgba.green)
	rgba.Blue = float64(crgba.blue)
	rgba.Alpha = float64(crgba.alpha)

	return rgba
}

func nativeFromRGBA(rgba interface{}) *gobject.GValue {
	if r, ok := rgba.(RGBA); ok {
		gv := gobject.CreateCGValue(GdkType.RGBA, r.ToNative())
		return gv
	}
	return nil
}

// To be BoxedLike
func (self RGBA) GetBoxType() gobject.GType {
	return GdkType.RGBA
}

func (self RGBA) ToNative() unsafe.Pointer {
	crgba := C._new_rgba()

	crgba.red = C.gdouble(self.Red)
	crgba.green = C.gdouble(self.Green)
	crgba.blue = C.gdouble(self.Blue)
	crgba.alpha = C.gdouble(self.Alpha)

	return unsafe.Pointer(crgba)
}

// End RGBA type }}}

// GdkColor {{{

// GdkColor type
type Color struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
}

//Conversion func
func newColorFromNative(obj unsafe.Pointer) interface{} {
	cColor := *((*C.GdkColor)(obj))

	color := Color{}
	color.Pixel = uint32(cColor.pixel)
	color.Red = uint16(cColor.red)
	color.Green = uint16(cColor.green)
	color.Blue = uint16(cColor.blue)

	return color
}

func nativeFromColor(c interface{}) *gobject.GValue {
	if col, ok := c.(Color); ok {
		gv := gobject.CreateCGValue(GdkType.COLOR, col.ToNative())
		return gv
	}
	return nil
}

// To be BoxedLike
func (self Color) GetBoxType() gobject.GType {
	return GdkType.COLOR
}

func (self Color) ToNative() unsafe.Pointer {
	var cColor C.GdkColor

	cColor.pixel = C.guint32(self.Pixel)
	cColor.red = C.guint16(self.Red)
	cColor.green = C.guint16(self.Green)
	cColor.blue = C.guint16(self.Blue)

	return unsafe.Pointer(&cColor)
}

// End Color type }}}

// GdkAtom {{{

type Atom struct {
	object *C.GdkAtom
}

func (self Atom) ToNative() unsafe.Pointer {
	if self.object == nil {
		return nil
	}
	return unsafe.Pointer(self.object)
}

func (self Atom) Name() string {
	if self.object == nil {
		return ""
	}

	s := C.gdk_atom_name(*self.object)
	return gobject.GoString(unsafe.Pointer(s))
}

func NewAtomFromNative(obj unsafe.Pointer) Atom {
	a := Atom{}
	a.object = new(C.GdkAtom)
	*a.object = *((*C.GdkAtom)(obj))

	return a
}

func AtomIntern(atomName string, onlyIfExists bool) Atom {
	s := gobject.GString(atomName)
	defer s.Free()
	b := gobject.GBool(onlyIfExists)
	defer b.Free()

	o := C.gdk_atom_intern((*C.gchar)(s.GetPtr()), *((*C.gboolean)(b.GetPtr())))

	a := Atom{}
	a.object = new(C.GdkAtom)
	*a.object = o

	return a
}

// End GdkAtom }}}

// GdkModifier {{{

var GdkModifier gdkModifier

type gdkModifier struct {
	SHIFT_MASK    int
	LOCK_MASK     int
	CONTROL_MASK  int
	MOD1_MASK     int
	MOD2_MASK     int
	MOD3_MASK     int
	MOD4_MASK     int
	MOD5_MASK     int
	BUTTON1_MASK  int
	BUTTON2_MASK  int
	BUTTON3_MASK  int
	BUTTON4_MASK  int
	BUTTON5_MASK  int
	SUPER_MASK    int
	HYPER_MASK    int
	META_MASK     int
	RELEASE_MASK  int
	MODIFIER_MASK int
}

// Conversion func
func gdkModifierFromNative(m unsafe.Pointer) interface{} {
	return int(*((*C.gint)(m)))
}

// End GdkModifier }}}

// End basic type }}}

// Events implementation {{{

type EventLike interface {
	EventType() int
	GdkWindow() *Window
	SentExplicitly() bool
}

// GdkEventAny
type EventAny struct {
	eventType int
	window    *Window
	sendEvent bool
}

// To-be event like
func (self EventAny) EventType() int {
	return self.eventType
}

func (self EventAny) Window() *Window {
	return self.window
}

func (self EventAny) SentExplicitly() bool {
	return self.sendEvent
}

// GdkEventKey
type EventKey struct {
	EventAny
	Time            uint32
	State           uint
	KeyVal          uint
	HardwareKeycode uint16
	Group           byte
	IsModifier      uint
}

func createEventKey(obj unsafe.Pointer) EventKey {
	ke := (*C.GdkEventKey)(obj)
	keyEvent := EventKey{}
	keyEvent.Time = uint32(ke.time)
	keyEvent.State = uint(ke.state)
	keyEvent.KeyVal = uint(ke.keyval)
	keyEvent.HardwareKeycode = uint16(ke.hardware_keycode)
	keyEvent.Group = byte(ke.group)
	keyEvent.IsModifier = uint(C.getModifiers(ke))
	return keyEvent
}

// GdkEventButton
type EventButton struct {
	EventAny
	Time   uint32
	X      float64
	Y      float64
	State  uint
	Button uint
	Device *Device
	X_Root float64
	Y_Root float64
}

func createEventButton(obj unsafe.Pointer) EventButton {
	eb := (*C.GdkEventButton)(obj)
	eventButton := EventButton{}
	eventButton.Time = uint32(eb.time)
	eventButton.X = float64(eb.x)
	eventButton.Y = float64(eb.y)
	eventButton.State = uint(eb.state)
	eventButton.Button = uint(eb.button)
	eventButton.Device = nil
	if gd, err := gobject.ConvertToGo(unsafe.Pointer(eb.device), GdkType.DEVICE); err == nil {
		eventButton.Device = gd.(*Device)
	}
	eventButton.X_Root = float64(eb.x_root)
	eventButton.Y_Root = float64(eb.y_root)
	return eventButton
}

// Conversion function
func newEventFromNative(obj unsafe.Pointer) interface{} {
	// Cast to EventAny to extract type and send var
	var window *Window = nil
	ceventAny := (*C.GdkEventAny)(obj)
	t := int(ceventAny._type)
	if win, err := gobject.ConvertToGo(unsafe.Pointer(ceventAny.window)); err == nil {
		window = win.(*Window)
	}
	se := ceventAny.send_event
	send := gobject.GoBool(unsafe.Pointer(&se))

	// Create new EventAny Struct
	eventAny := EventAny{t, window, send}

	// test specific events
	switch t {
	case GdkEventType.KEY_PRESS, GdkEventType.KEY_RELEASE:
		ek := createEventKey(obj)
		ek.EventAny = eventAny
		return ek
	case GdkEventType.BUTTON_PRESS, GdkEventType.BUTTON_PRESS_2, GdkEventType.BUTTON_PRESS_3,
		GdkEventType.BUTTON_RELEASE:
		eb := createEventButton(obj)
		eb.EventAny = eventAny
		return eb
	}
	return eventAny
}

var GdkEventMask gdkEventMask

type gdkEventMask struct {
	EXPOSURE_MASK            int
	POINTER_MOTION_MASK      int
	POINTER_MOTION_HINT_MASK int
	BUTTON_MOTION_MASK       int
	BUTTON1_MOTION_MASK      int
	BUTTON2_MOTION_MASK      int
	BUTTON3_MOTION_MASK      int
	BUTTON_PRESS_MASK        int
	BUTTON_RELEASE_MASK      int
	KEY_PRESS_MASK           int
	KEY_RELEASE_MASK         int
	ENTER_NOTIFY_MASK        int
	LEAVE_NOTIFY_MASK        int
	FOCUS_CHANGE_MASK        int
	STRUCTURE_MASK           int
	PROPERTY_CHANGE_MASK     int
	VISIBILITY_CHANGE_MASK   int
	PROXIMITY_IN_MASK        int
	PROXIMITY_OUT_MASK       int
	SUBSTRUCTURE_MASK        int
	SCROLL_MASK              int
	ALL_EVENTS_MASK          int
}

func init_mask() {
	GdkEventMask.EXPOSURE_MASK = 1 << 1
	GdkEventMask.POINTER_MOTION_MASK = 1 << 2
	GdkEventMask.POINTER_MOTION_HINT_MASK = 1 << 3
	GdkEventMask.BUTTON_MOTION_MASK = 1 << 4
	GdkEventMask.BUTTON1_MOTION_MASK = 1 << 5
	GdkEventMask.BUTTON2_MOTION_MASK = 1 << 6
	GdkEventMask.BUTTON3_MOTION_MASK = 1 << 7
	GdkEventMask.BUTTON_PRESS_MASK = 1 << 8
	GdkEventMask.BUTTON_RELEASE_MASK = 1 << 9
	GdkEventMask.KEY_PRESS_MASK = 1 << 10
	GdkEventMask.KEY_RELEASE_MASK = 1 << 11
	GdkEventMask.ENTER_NOTIFY_MASK = 1 << 12
	GdkEventMask.LEAVE_NOTIFY_MASK = 1 << 13
	GdkEventMask.FOCUS_CHANGE_MASK = 1 << 14
	GdkEventMask.STRUCTURE_MASK = 1 << 15
	GdkEventMask.PROPERTY_CHANGE_MASK = 1 << 16
	GdkEventMask.VISIBILITY_CHANGE_MASK = 1 << 17
	GdkEventMask.PROXIMITY_IN_MASK = 1 << 18
	GdkEventMask.PROXIMITY_OUT_MASK = 1 << 19
	GdkEventMask.SUBSTRUCTURE_MASK = 1 << 20
	GdkEventMask.SCROLL_MASK = 1 << 21
	GdkEventMask.ALL_EVENTS_MASK = 0x3FFFFE
}

var GdkEventType gdkEventType

type gdkEventType struct {
	NOTHING           int
	DELETE            int
	DESTROY           int
	EXPOSE            int
	MOTION_NOTIFY     int
	BUTTON_PRESS      int
	BUTTON_PRESS_2    int
	BUTTON_PRESS_3    int
	BUTTON_RELEASE    int
	KEY_PRESS         int
	KEY_RELEASE       int
	ENTER_NOTIFY      int
	LEAVE_NOTIFY      int
	FOCUS_CHANGE      int
	CONFIGURE         int
	MAP               int
	UNMAP             int
	PROPERTY_NOTIFY   int
	SELECTION_CLEAR   int
	SELECTION_REQUEST int
	SELECTION_NOTIFY  int
	PROXIMITY_IN      int
	PROXIMITY_OUT     int
	DRAG_ENTER        int
	DRAG_LEAVE        int
	DRAG_MOTION       int
	DRAG_STATUS       int
	DRAG_START        int
	DRAG_FINISHED     int
	CLIENT_EVENT      int
	VISIBILITY_NOTIFY int
	SCROLL            int
	WINDOW_STATE      int
	SETTING           int
	OWNER_CHANGE      int
	GRAB_BROKEN       int
	DAMAGE            int
}

func init_events() {
	GdkEventType.NOTHING = -1
	GdkEventType.DELETE = 0
	GdkEventType.DESTROY = 1
	GdkEventType.EXPOSE = 2
	GdkEventType.MOTION_NOTIFY = 3
	GdkEventType.BUTTON_PRESS = 4
	GdkEventType.BUTTON_PRESS_2 = 5
	GdkEventType.BUTTON_PRESS_3 = 6
	GdkEventType.BUTTON_RELEASE = 7
	GdkEventType.KEY_PRESS = 8
	GdkEventType.KEY_RELEASE = 9
	GdkEventType.ENTER_NOTIFY = 10
	GdkEventType.LEAVE_NOTIFY = 11
	GdkEventType.FOCUS_CHANGE = 12
	GdkEventType.CONFIGURE = 13
	GdkEventType.MAP = 14
	GdkEventType.UNMAP = 15
	GdkEventType.PROPERTY_NOTIFY = 16
	GdkEventType.SELECTION_CLEAR = 17
	GdkEventType.SELECTION_REQUEST = 18
	GdkEventType.SELECTION_NOTIFY = 19
	GdkEventType.PROXIMITY_IN = 20
	GdkEventType.PROXIMITY_OUT = 21
	GdkEventType.DRAG_ENTER = 22
	GdkEventType.DRAG_LEAVE = 23
	GdkEventType.DRAG_MOTION = 24
	GdkEventType.DRAG_STATUS = 25
	GdkEventType.DRAG_START = 26
	GdkEventType.DRAG_FINISHED = 27
	GdkEventType.CLIENT_EVENT = 28
	GdkEventType.VISIBILITY_NOTIFY = 29
	GdkEventType.SCROLL = 31
	GdkEventType.WINDOW_STATE = 32
	GdkEventType.SETTING = 33
	GdkEventType.OWNER_CHANGE = 34
	GdkEventType.GRAB_BROKEN = 35
	GdkEventType.DAMAGE = 36
}

// End events implementation }}}

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
	C.gdk_init(nil, nil)
	GdkType.DEVICE = gobject.GType(C.gdk_device_get_type())
	GdkType.RECTANGLE = gobject.GType(C.gdk_rectangle_get_type())
	GdkType.RGBA = gobject.GType(C.gdk_rgba_get_type())
	GdkType.COLOR = gobject.GType(C.gdk_color_get_type())
	GdkType.EVENT = gobject.GType(C.gdk_event_get_type())
	GdkType.DISPLAY = gobject.GType(C.gdk_display_get_type())
	GdkType.SCREEN = gobject.GType(C.gdk_screen_get_type())
	GdkType.WINDOW = gobject.GType(C.gdk_window_get_type())
	GdkType.MODIFIER_TYPE = gobject.GType(C.gdk_modifier_type_get_type())

	// Register Rectangle
	gobject.RegisterCType(GdkType.RECTANGLE, newRectangleFromNative)
	gobject.RegisterGoType(GdkType.RECTANGLE, nativeFromRectangle)

	// Register RGBA type
	gobject.RegisterCType(GdkType.RGBA, newRGBAFromNative)
	gobject.RegisterGoType(GdkType.RGBA, nativeFromRGBA)

	// Register GdkModifier
	gobject.RegisterCType(GdkType.MODIFIER_TYPE, gdkModifierFromNative)

	// Initialize GdkModifier
	GdkModifier.SHIFT_MASK = 1 << 0
	GdkModifier.LOCK_MASK = 1 << 1
	GdkModifier.CONTROL_MASK = 1 << 2
	GdkModifier.MOD1_MASK = 1 << 3
	GdkModifier.MOD2_MASK = 1 << 4
	GdkModifier.MOD3_MASK = 1 << 5
	GdkModifier.MOD4_MASK = 1 << 6
	GdkModifier.MOD5_MASK = 1 << 7
	GdkModifier.BUTTON1_MASK = 1 << 8
	GdkModifier.BUTTON2_MASK = 1 << 9
	GdkModifier.BUTTON3_MASK = 1 << 10
	GdkModifier.BUTTON4_MASK = 1 << 11
	GdkModifier.BUTTON5_MASK = 1 << 12
	GdkModifier.SUPER_MASK = 1 << 26
	GdkModifier.HYPER_MASK = 1 << 27
	GdkModifier.META_MASK = 1 << 28
	GdkModifier.RELEASE_MASK = 1 << 30
	GdkModifier.MODIFIER_MASK = 0x5c001fff

	//Register GdkEvent to gobject (In Go)
	gobject.RegisterCType(GdkType.EVENT, newEventFromNative)
	init_events()
	init_mask()


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
