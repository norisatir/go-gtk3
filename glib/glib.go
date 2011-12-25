package glib

/*
#include <glib-object.h>
#include <stdlib.h>

extern gboolean _g_source_func(gpointer user_data);
extern void _g_destroy_notify(gpointer user_data);

static guint _g_timeout_add_full(gint priority, guint interval, gint64 id) {
	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
	return g_timeout_add_full(priority, interval, _g_source_func, (gpointer)uid, _g_destroy_notify);
}

static guint _g_timeout_add_seconds_full(gint priority, guint interval, gint64 id) {
	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
	return g_timeout_add_seconds_full(priority, interval, _g_source_func, (gpointer)uid, _g_destroy_notify);
}

static guint _g_idle_add_full(gint priority, gint64 id) {
	gint64* uid = (gint64*)malloc(sizeof(gint64));
	*uid = id;
	return g_idle_add_full(priority, _g_source_func, (gpointer)uid, _g_destroy_notify);
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
	gl.object = nil
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

func (self *GSList) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
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
	self.Free()
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

// GList {{{
//////////////////////////////

// GList type
// GList type which holds C GList.
// If GC_Free is true, then GC will call g_list_free on object
// IF GC_FreeFull it true then GC will call g_list_free_full
// even if GC_Free is true/false (it is ignored).
type GList struct {
	object         *C.GList
	GC_Free        bool
	GC_FreeFull    bool
	ConversionFunc ConverterFunc
	DestroyFunc    ListClosure
}

// NewGList returns new GList with first element initialized.
// GC_Free and GC_FreeFull will hold default bool values.
func NewGList() *GList {
	gl := &GList{}
	gl.object = nil
	gl.ConversionFunc = nil
	gl.DestroyFunc = nil

	return gl
}

func NewGListFromNative(list unsafe.Pointer) *GList {
	gl := &GList{}
	gl.object = (*C.GList)(list)
	gl.ConversionFunc = nil
	gl.DestroyFunc = nil

	return gl
}

// GList finalizer
func GListFinalizer(gl *GList) {
	runtime.SetFinalizer(gl, func(gl *GSList) {
		if gl.GC_FreeFull {
			gl.FreeFull()
			return
		}

		if gl.GC_Free {
			gl.Free()
		}
	})
}

func (self *GList) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self *GList) Free() {
	C.g_list_free(self.object)
}

func (self *GList) FreeFull() {
	if self.DestroyFunc == nil {
		return
	}
	var numElements int = int(self.Length())

	for i := 0; i < numElements; i++ {
		el := C.g_list_nth_data(self.object, C.guint(i))
		self.DestroyFunc(unsafe.Pointer(el))
	}
	self.Free()
}

func (self *GList) NthData(n uint) interface{} {
	data := C.g_list_nth_data(self.object, C.guint(n))

	if data == nil {
		return nil
	}

	if self.ConversionFunc != nil {
		return self.ConversionFunc(unsafe.Pointer(data))
	}
	return data
}

func (self *GList) Length() uint {
	return uint(C.g_list_length(self.object))
}

func (self *GList) Foreach(f interface{}, data ...interface{}) {
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
// End GList
////////////////////////////// }}}

func GTimeoutAddFull(priority int, interval uint, callback interface{}, data ...interface{}) uint {
	cl, id := gobject.CreateCustomClosure(callback, data...)
	_closures[id] = cl
	return uint(C._g_timeout_add_full(C.gint(priority), C.guint(interval), C.gint64(id)))
}

func GTimeoutAddSecondsFull(priority int, interval uint, callback interface{}, data ...interface{}) uint {
	cl, id := gobject.CreateCustomClosure(callback, data...)
	_closures[id] = cl
	return uint(C._g_timeout_add_seconds_full(C.gint(priority), C.guint(interval), C.gint64(id)))
}

func GIdleAddFull(priority int, callback interface{}, data ...interface{}) uint {
	cl, id := gobject.CreateCustomClosure(callback, data...)
	_closures[id] = cl
	return uint(C._g_idle_add_full(C.gint(priority), C.gint64(id)))
}

func GSourceRemove(tag uint) bool {
	b := C.g_source_remove(C.guint(tag))
	return gobject.GoBool(unsafe.Pointer(&b))
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
