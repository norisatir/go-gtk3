package glib

/*
#include <glib-object.h>
#include <stdlib.h>

extern gboolean _g_source_func(gpointer user_data);
extern void _g_destroy_notify(gpointer user_data);

static void _g_timeout_add_full(gint priority, guint interval, gint64 id) {
	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
	g_timeout_add_full(priority, interval, _g_source_func, (gpointer)uid, _g_destroy_notify);
}

*/
// #cgo pkg-config: gobject-2.0
import "C"
import "unsafe"
import "runtime"
//import "reflect"
import "github.com/norisatir/go-gtk3/gobject"

type ListClosure func(unsafe.Pointer)
type ConverterFunc func(unsafe.Pointer) interface{}

var _closures map[int64]gobject.ClosureFunc

// GSList {{{
//////////////////////////////

// GSList type
// GSList type which holds C GSList.
// If GC_Free is true, then GC will call g_slist_free on object
// IF GC_FreeFull it true then GC will call g_slist_free_full
// even if GC_Free is true/false (it is ignored).
type GSList struct {
	object         *C.GSList
	GC_Free        bool
	GC_FreeFull    bool
	ConversionFunc ConverterFunc
	DestroyFunc    ListClosure
}

// NewGSList returns new GSList with first element initialized.
// GC_Free and GC_FreeFull will hold default bool values.
func NewGSList() *GSList {
	gl := &GSList{}
	gl.object = C.g_slist_alloc()
	gl.ConversionFunc = nil
	gl.DestroyFunc = nil

	return gl
}

func NewGSListFromNative(list unsafe.Pointer) *GSList {
	gl := &GSList{}
	gl.object = (*C.GSList)(list)
	gl.ConversionFunc = nil
	gl.DestroyFunc = nil

	return gl
}

// GSList finalizer
func GSListFinalizer(gsl *GSList) {
	runtime.SetFinalizer(gsl, func(gsl *GSList) {
		if gsl.GC_FreeFull {
			gsl.FreeFull()
			return
		}

		if gsl.GC_Free {
			gsl.Free()
		}
	})
}

func (self *GSList) Free() {
	C.g_slist_free(self.object)
}

func (self *GSList) FreeFull() {
	if self.DestroyFunc == nil {
		return
	}
	var numElements int = int(self.Length())

	for i := 0; i < numElements; i++ {
		el := C.g_slist_nth_data(self.object, C.guint(i))
		self.DestroyFunc(unsafe.Pointer(el))
	}
}

func (self *GSList) NthData(n uint) interface{} {
	data := C.g_slist_nth_data(self.object, C.guint(n))

	if data == nil {
		return nil
	}

	if self.ConversionFunc != nil {
		return self.ConversionFunc(unsafe.Pointer(data))
	}
	return data
}

func (self *GSList) Length() uint {
	return uint(C.g_slist_length(self.object))
}

func (self *GSList) Foreach(f interface{}, data ...interface{}) {
	// Create id and closure
	cl, _ := gobject.CreateCustomClosure(f, data...)
	listLength := int(self.Length())

	for i := 0; i < listLength; i++ {
		data := self.NthData(uint(i))
		if data != nil {
			cl([]interface{}{data})
		}
	}
}

//////////////////////////////
// End GSList
////////////////////////////// }}}

func GTimeoutAddFull(priority int, interval uint, callback interface{}, data ...interface{}) {
	cl, id := gobject.CreateCustomClosure(callback, data...)
	_closures[id] = cl
	C._g_timeout_add_full(C.gint(priority), C.guint(interval), C.gint64(id))
}

// Exported functions
//export _g_source_func
func _g_source_func(user_data unsafe.Pointer) C.gboolean {
	id := *((*C.gint64)(user_data))
	var res bool

	if f, ok := _closures[int64(id)]; ok {
		res = f([]interface{}{})
	}
	b := gobject.GBool(res)
	defer b.Free()
	return *((*C.gboolean)(b.GetPtr()))
}

//export _g_destroy_notify
func _g_destroy_notify(user_data unsafe.Pointer) {
	id := *((*C.gint64)(user_data))
	if _, ok := _closures[int64(id)]; ok {
		delete(_closures, int64(id))
	}
	C.free(user_data)
}

var GPriority gPriority

type gPriority struct {
	HIGH         int
	DEFAULT      int
	HIGH_IDLE    int
	DEFAULT_IDLE int
	LOW          int
}

func init() {
	_closures = make(map[int64]gobject.ClosureFunc)

	// Initialize GPriority struct
	GPriority.HIGH = -100
	GPriority.DEFAULT = 0
	GPriority.HIGH_IDLE = 100
	GPriority.DEFAULT_IDLE = 200
	GPriority.LOW = 300
}
