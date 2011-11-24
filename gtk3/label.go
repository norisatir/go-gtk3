package gtk3

/*
#include <gtk/gtk.h>

static inline GtkLabel* to_GtkLabel(void* obj) { return GTK_LABEL(obj); }

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"

type GtkJustification int

const (
	JUSTIFY_LEFT   GtkJustification = 0
	JUSTIFY_RIGHT  GtkJustification = 1
	JUSTIFY_CENTER GtkJustification = 2
	JUSTIFY_FILL   GtkJustification = 3
)

type Label struct {
	object *C.GtkLabel
	*Widget
}

func NewLabel(str string) *Label {
	l := &Label{}

	s := gobject.GString(str)
	defer s.Free()

	o := C.gtk_label_new((*C.gchar)(s.GetPtr()))

	l.Widget = NewWidget(unsafe.Pointer(o))
	l.object = C.to_GtkLabel(unsafe.Pointer(o))

	return l
}

func NewLabelWithMnemonic(str string) *Label {
	l := &Label{}

	s := gobject.GString(str)
	defer s.Free()

	o := C.gtk_label_new_with_mnemonic((*C.gchar)(s.GetPtr()))

	l.Widget = NewWidget(unsafe.Pointer(o))
	l.object = C.to_GtkLabel(unsafe.Pointer(o))

	return l
}
// Conversion function for gobject registration map
func newLabelFromNative(obj unsafe.Pointer) interface{} {
	var l Label
	l.object = C.to_GtkLabel(obj)
	l.Widget = NewWidget(unsafe.Pointer(l.object))
	return &l
}

func nativeFromLabel(label interface{}) *gobject.GValue {
	l, ok := label.(Label)
	if ok {
		gv := gobject.CreateCGValue(GtkType.LABEL, l.ToNative())
		return gv
	}

	return nil
}

func init() {
	// Register GtkLabel to gobject
	gobject.RegisterCType(GtkType.LABEL, newLabelFromNative)
	gobject.RegisterGoType(GtkType.LABEL, nativeFromLabel)
}

// To be object like
func (self Label) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Label) Connect(name string, f interface{}, data ...interface{}) {
	gobject.Connect(self, name, f, data...)
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

func (self *Label) SetJustify(jtype GtkJustification) {
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

func (self *Label) GetJustify() GtkJustification {
	j := C.gtk_label_get_justify(self.object)
	return GtkJustification(j)
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
