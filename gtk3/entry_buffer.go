package gtk3

/*
#include <gtk/gtk.h>

static inline GtkEntryBuffer* to_GtkEntryBuffer(void* obj) { return GTK_ENTRY_BUFFER(obj); }

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"


type EntryBuffer struct {
	object *C.GtkEntryBuffer
}

func NewEntryBuffer(initial_chars string) *EntryBuffer {
	e := EntryBuffer{}

	s := gobject.GString(initial_chars)
	defer s.Free()

	o := C.gtk_entry_buffer_new((*C.gchar)(s.GetPtr()), C.gint(len(initial_chars)))

	e.object = C.to_GtkEntryBuffer(unsafe.Pointer(o))

	return &e
}

// Conversion function for gobject registration map
func newEntryBufferFromNative(obj unsafe.Pointer) interface{} {
	var e EntryBuffer
	e.object = C.to_GtkEntryBuffer(obj)

	return &e
}

func nativeFromEntryBuffer(eb interface{}) *gobject.GValue {
	e, ok := eb.(*EntryBuffer)
	if ok {
		gv := gobject.CreateCGValue(GtkType.ENTRY_BUFFER, e.ToNative())
		return gv
	}

	return nil
}


func init() {
	// Register GtkEntryBuffer to gobject
	gobject.RegisterCType(GtkType.ENTRY_BUFFER, newEntryBufferFromNative)
	gobject.RegisterGoType(GtkType.ENTRY_BUFFER, nativeFromEntryBuffer)
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

func (self *EntryBuffer) GetText() (string) {
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
