package gdk3

/*
#include <gdk/gdk.h>
static inline guint getModifiers(GdkEventKey* ke) {
	return ke->is_modifier;
}

*/
import "C"
import "github.com/norisatir/go-gtk3/gobject"
import "unsafe"

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
	gd, _ := gobject.ConvertToGo(unsafe.Pointer(eb.device), GdkType.DEVICE)
	eventButton.Device = gd.(*Device)
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

func init() {
	//Register GdkEvent to gobject (In Go)
	gobject.RegisterCType(GdkType.EVENT, newEventFromNative)
	init_events()
	init_mask()
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
