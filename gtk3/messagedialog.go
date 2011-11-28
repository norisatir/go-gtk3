package gtk3

/*
#include <gtk/gtk.h>

static inline GtkMessageDialog* to_GtkMessageDialog(void* obj) { return GTK_MESSAGE_DIALOG(obj); }

static inline GtkMessageDialog* _new_message_dialog(GtkWindow* parent,
													GtkDialogFlags flags,
													GtkMessageType typ,
													GtkButtonsType buttons,
													const gchar* message) {
	GtkWidget* w = gtk_message_dialog_new(parent, flags, typ, buttons, message, NULL);
	return to_GtkMessageDialog((void*)w);
}

static inline void _gtk_message_dialog_format_secondary_text(GtkMessageDialog* md, const gchar* message) {
	gtk_message_dialog_format_secondary_text(md, message, NULL);
}
static inline void _gtk_message_dialog_format_secondary_markup(GtkMessageDialog* md, const gchar* message) {
	gtk_message_dialog_format_secondary_markup(md, message, NULL);
}

*/
import "C"
import "unsafe"
import "github.com/norisatir/go-gtk3/gobject"
import "fmt"

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

	mdiag.object = C._new_message_dialog(pwin, C.GtkDialogFlags(flags), C.GtkMessageType(mtype), C.GtkButtonsType(buttons),
		(*C.gchar)(cmessage.GetPtr()))
	mdiag.Dialog = newDialogFromNative(unsafe.Pointer(mdiag.object)).(*Dialog)

	return mdiag
}

// Conversion function for gobject registration map
func newMessageDialogFromNative(obj unsafe.Pointer) interface{} {
	var dialog MessageDialog
	dialog.object = C.to_GtkMessageDialog(obj)
	dialog.Dialog = newDialogFromNative(obj).(*Dialog)
	return &dialog
}

func nativeFromMessageDialog(d interface{}) *gobject.GValue {
	dialog, ok := d.(MessageDialog)
	if ok {
		gv := gobject.CreateCGValue(GtkType.MESSAGE_DIALOG, dialog.ToNative())
		return gv
	}

	return nil
}

func init() {
	// Register GtkMessageDialog to gobject
	gobject.RegisterCType(GtkType.MESSAGE_DIALOG, newMessageDialogFromNative)
	gobject.RegisterGoType(GtkType.MESSAGE_DIALOG, nativeFromMessageDialog)
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
