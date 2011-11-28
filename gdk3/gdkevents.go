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
	EventType() GdkEventType
	//  GdkWindow() *GdkWindow
	SentExplicitly() bool
}

// GdkEventAny
type EventAny struct {
	Type GdkEventType
	// Window *GdkWindow
	SendEvent bool
}

// To-be event like
func (self EventAny) EventType() GdkEventType {
	return self.Type
}

func (self EventAny) SentExplicitly() bool {
	return self.SendEvent
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
	ceventAny := (*C.GdkEventAny)(obj)
	t := GdkEventType(ceventAny._type)
	se := ceventAny.send_event
	send := gobject.GoBool(unsafe.Pointer(&se))

	// Create new EventAny Struct
	eventAny := EventAny{t, send}

	// test specific events
	switch t {
	case EventType.KEY_PRESS, EventType.KEY_RELEASE:
		ek := createEventKey(obj)
		ek.EventAny = eventAny
		return ek
	case EventType.BUTTON_PRESS, EventType.BUTTON_PRESS_2, EventType.BUTTON_PRESS_3,
		EventType.BUTTON_RELEASE:
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

type GdkEventMask int

type eventMask struct {
	EXPOSURE_MASK            GdkEventMask
	POINTER_MOTION_MASK      GdkEventMask
	POINTER_MOTION_HINT_MASK GdkEventMask
	BUTTON_MOTION_MASK       GdkEventMask
	BUTTON1_MOTION_MASK      GdkEventMask
	BUTTON2_MOTION_MASK      GdkEventMask
	BUTTON3_MOTION_MASK      GdkEventMask
	BUTTON_PRESS_MASK        GdkEventMask
	BUTTON_RELEASE_MASK      GdkEventMask
	KEY_PRESS_MASK           GdkEventMask
	KEY_RELEASE_MASK         GdkEventMask
	ENTER_NOTIFY_MASK        GdkEventMask
	LEAVE_NOTIFY_MASK        GdkEventMask
	FOCUS_CHANGE_MASK        GdkEventMask
	STRUCTURE_MASK           GdkEventMask
	PROPERTY_CHANGE_MASK     GdkEventMask
	VISIBILITY_CHANGE_MASK   GdkEventMask
	PROXIMITY_IN_MASK        GdkEventMask
	PROXIMITY_OUT_MASK       GdkEventMask
	SUBSTRUCTURE_MASK        GdkEventMask
	SCROLL_MASK              GdkEventMask
	ALL_EVENTS_MASK          GdkEventMask
}

var EventMask eventMask

func init_mask() {
	EventMask = eventMask{}
	EventMask.EXPOSURE_MASK = GdkEventMask(1 << 1)
	EventMask.POINTER_MOTION_MASK = GdkEventMask(1 << 2)
	EventMask.POINTER_MOTION_HINT_MASK = GdkEventMask(1 << 3)
	EventMask.BUTTON_MOTION_MASK = GdkEventMask(1 << 4)
	EventMask.BUTTON1_MOTION_MASK = GdkEventMask(1 << 5)
	EventMask.BUTTON2_MOTION_MASK = GdkEventMask(1 << 6)
	EventMask.BUTTON3_MOTION_MASK = GdkEventMask(1 << 7)
	EventMask.BUTTON_PRESS_MASK = GdkEventMask(1 << 8)
	EventMask.BUTTON_RELEASE_MASK = GdkEventMask(1 << 9)
	EventMask.KEY_PRESS_MASK = GdkEventMask(1 << 10)
	EventMask.KEY_RELEASE_MASK = GdkEventMask(1 << 11)
	EventMask.ENTER_NOTIFY_MASK = GdkEventMask(1 << 12)
	EventMask.LEAVE_NOTIFY_MASK = GdkEventMask(1 << 13)
	EventMask.FOCUS_CHANGE_MASK = GdkEventMask(1 << 14)
	EventMask.STRUCTURE_MASK = GdkEventMask(1 << 15)
	EventMask.PROPERTY_CHANGE_MASK = GdkEventMask(1 << 16)
	EventMask.VISIBILITY_CHANGE_MASK = GdkEventMask(1 << 17)
	EventMask.PROXIMITY_IN_MASK = GdkEventMask(1 << 18)
	EventMask.PROXIMITY_OUT_MASK = GdkEventMask(1 << 19)
	EventMask.SUBSTRUCTURE_MASK = GdkEventMask(1 << 20)
	EventMask.SCROLL_MASK = GdkEventMask(1 << 21)
	EventMask.ALL_EVENTS_MASK = 0x3FFFFE
}

type GdkEventType int

type eventType struct {
	NOTHING           GdkEventType
	DELETE            GdkEventType
	DESTROY           GdkEventType
	EXPOSE            GdkEventType
	MOTION_NOTIFY     GdkEventType
	BUTTON_PRESS      GdkEventType
	BUTTON_PRESS_2    GdkEventType
	BUTTON_PRESS_3    GdkEventType
	BUTTON_RELEASE    GdkEventType
	KEY_PRESS         GdkEventType
	KEY_RELEASE       GdkEventType
	ENTER_NOTIFY      GdkEventType
	LEAVE_NOTIFY      GdkEventType
	FOCUS_CHANGE      GdkEventType
	CONFIGURE         GdkEventType
	MAP               GdkEventType
	UNMAP             GdkEventType
	PROPERTY_NOTIFY   GdkEventType
	SELECTION_CLEAR   GdkEventType
	SELECTION_REQUEST GdkEventType
	SELECTION_NOTIFY  GdkEventType
	PROXIMITY_IN      GdkEventType
	PROXIMITY_OUT     GdkEventType
	DRAG_ENTER        GdkEventType
	DRAG_LEAVE        GdkEventType
	DRAG_MOTION       GdkEventType
	DRAG_STATUS       GdkEventType
	DRAG_START        GdkEventType
	DRAG_FINISHED     GdkEventType
	CLIENT_EVENT      GdkEventType
	VISIBILITY_NOTIFY GdkEventType
	SCROLL            GdkEventType
	WINDOW_STATE      GdkEventType
	SETTING           GdkEventType
	OWNER_CHANGE      GdkEventType
	GRAB_BROKEN       GdkEventType
	DAMAGE            GdkEventType
}

// Global Event var, initialized in init()
var EventType eventType

func init_events() {
	EventType = eventType{}
	EventType.NOTHING = GdkEventType(-1)
	EventType.DELETE = GdkEventType(0)
	EventType.DESTROY = GdkEventType(1)
	EventType.EXPOSE = GdkEventType(2)
	EventType.MOTION_NOTIFY = GdkEventType(3)
	EventType.BUTTON_PRESS = GdkEventType(4)
	EventType.BUTTON_PRESS_2 = GdkEventType(5)
	EventType.BUTTON_PRESS_3 = GdkEventType(6)
	EventType.BUTTON_RELEASE = GdkEventType(7)
	EventType.KEY_PRESS = GdkEventType(8)
	EventType.KEY_RELEASE = GdkEventType(9)
	EventType.ENTER_NOTIFY = GdkEventType(10)
	EventType.LEAVE_NOTIFY = GdkEventType(11)
	EventType.FOCUS_CHANGE = GdkEventType(12)
	EventType.CONFIGURE = GdkEventType(13)
	EventType.MAP = GdkEventType(14)
	EventType.UNMAP = GdkEventType(15)
	EventType.PROPERTY_NOTIFY = GdkEventType(16)
	EventType.SELECTION_CLEAR = GdkEventType(17)
	EventType.SELECTION_REQUEST = GdkEventType(18)
	EventType.SELECTION_NOTIFY = GdkEventType(19)
	EventType.PROXIMITY_IN = GdkEventType(20)
	EventType.PROXIMITY_OUT = GdkEventType(21)
	EventType.DRAG_ENTER = GdkEventType(22)
	EventType.DRAG_LEAVE = GdkEventType(23)
	EventType.DRAG_MOTION = GdkEventType(24)
	EventType.DRAG_STATUS = GdkEventType(25)
	EventType.DRAG_START = GdkEventType(26)
	EventType.DRAG_FINISHED = GdkEventType(27)
	EventType.CLIENT_EVENT = GdkEventType(28)
	EventType.VISIBILITY_NOTIFY = GdkEventType(29)
	EventType.SCROLL = GdkEventType(31)
	EventType.WINDOW_STATE = GdkEventType(32)
	EventType.SETTING = GdkEventType(33)
	EventType.OWNER_CHANGE = GdkEventType(34)
	EventType.GRAB_BROKEN = GdkEventType(35)
	EventType.DAMAGE = GdkEventType(36)
}
