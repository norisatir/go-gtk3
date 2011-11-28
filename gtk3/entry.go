package gtk3

/*
#include <gtk/gtk.h>

static inline GtkEntry* to_GtkEntry(void* obj) { return GTK_ENTRY(obj); }

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"
import "github.com/norisatir/go-gtk3/gdk3"


type Entry struct {
	object *C.GtkEntry
	*Widget
}

func NewEntry() *Entry {
	e := &Entry{}

	o := C.gtk_entry_new()

	e.Widget = NewWidget(unsafe.Pointer(o))
	e.object = C.to_GtkEntry(unsafe.Pointer(o))

	return e
}

func NewEntryWithBuffer(eb *EntryBuffer) *Entry {
	e := &Entry{}

	o := C.gtk_entry_new_with_buffer(eb.object)

	e.Widget = NewWidget(unsafe.Pointer(o))
	e.object = C.to_GtkEntry(unsafe.Pointer(o))

	return e
}

// Conversion function for gobject registration map
func newEntryFromNative(obj unsafe.Pointer) interface{} {
	e := &Entry{}
	e.object = C.to_GtkEntry(obj)
	e.Widget = NewWidget(obj)

	return &e
}

func nativeFromEntry(entry interface{}) *gobject.GValue {
	e, ok := entry.(*Entry)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ENTRY, e.ToNative())
		return gv
	}
	return nil
}

func init() {
	// Register GtkEntry to gobject
	gobject.RegisterCType(GtkType.ENTRY, newEntryFromNative)
	gobject.RegisterGoType(GtkType.ENTRY, nativeFromEntry)
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

func (self *Entry) GetText() (string) {
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
