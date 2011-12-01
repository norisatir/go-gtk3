package gtk3

/*
#include <gtk/gtk.h>

static inline GtkDialog* to_GtkDialog(void* obj) { return GTK_DIALOG(obj); }

static inline GtkDialog* _dialog_new_with_buttons(const gchar* title,
												  GtkWindow* parent,
												  GtkDialogFlags flags,
												  const gchar* firstbutt,
												  gint resid) {
	GtkWidget* w = gtk_dialog_new_with_buttons(title, parent, flags, firstbutt, resid, NULL);
	return to_GtkDialog(w);
}



*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type DialogLike interface {
	D() *Dialog
}

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

func init() {
	// Register GtkDialog to gobject
	gobject.RegisterCType(GtkType.DIALOG, newDialogFromNative)
	gobject.RegisterGoType(GtkType.DIALOG, nativeFromDialog)
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
