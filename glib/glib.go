package glib

/*
#include <glib-object.h>

*/
// #cgo pkg-config: gobject-2.0
import "C"
import "unsafe"
import "runtime"
//import "reflect"
import "github.com/norisatir/go-gtk3/gobject"

type ListClosure func(unsafe.Pointer)
type ConverterFunc func(unsafe.Pointer)interface{}

var _closures map[int64]gobject.ClosureFunc

// GSList {{{
//////////////////////////////

// GSList type
// GSList type which holds C GSList.
// If GC_Free is true, then GC will call g_slist_free on object
// IF GC_FreeFull it true then GC will call g_slist_free_full
// even if GC_Free is true/false (it is ignored).
type GSList struct {
	object *C.GSList
	GC_Free bool
	GC_FreeFull bool
	ConversionFunc ConverterFunc
	DestroyFunc ListClosure
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
	cl,_ := gobject.CreateCustomClosure(f, data...)
	listLength := int(self.Length())

	for i := 0; i < listLength; i++ {
		data := self.NthData(uint(i))
		if data != nil {
			cl([]interface{}{self.NthData(uint(i))})
		}
	}
}

//////////////////////////////
// End GSList
////////////////////////////// }}}

func GTimeoutAddFull(priority int, interval uint, callback interface{}) {

}
// Exported functions

func init() {
	_closures = make(map[int64]gobject.ClosureFunc)
}
