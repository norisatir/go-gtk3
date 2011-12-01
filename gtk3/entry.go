package gtk3

/*
#include <gtk/gtk.h>

static inline GtkEntry* to_GtkEntry(void* obj) { return GTK_ENTRY(obj); }

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

*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"
import "github.com/norisatir/go-gtk3/gdk3"

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

func init() {
	// Register GtkEntry to gobject
	gobject.RegisterCType(GtkType.ENTRY, newEntryFromNative)
	gobject.RegisterGoType(GtkType.ENTRY, nativeFromEntry)
}

// To be Object-like
func (self Entry) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Entry) Connect(name string ,f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
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
