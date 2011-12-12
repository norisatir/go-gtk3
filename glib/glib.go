package glib

/*
#include <glib-object.h>

extern void _g_func(gpointer data, gpointer user_data);

static void _g_list_foreach(GSList* list, gint64 user_data) {
	gint64* data = &user_data;
	g_slist_foreach(list, _g_func, (gpointer)(data));
}


*/
// #cgo pkg-config: gobject-2.0
import "C"
import "unsafe"
import "runtime"
import "reflect"
import "github.com/norisatir/go-gtk3/gobject"

type ListClosure func(unsafe.Pointer)
type ConverterFunc func(unsafe.Pointer)interface{}

var _closures map[int64]ListClosure


func createListClosure(f interface{}, converterFunc ConverterFunc, data ...interface{}) ListClosure {
	cfunc := reflect.ValueOf(f)

	// cfunc is not function, then there is nothing to do
	if cfunc.Kind() != reflect.Func {
		return nil
	}

	// Do we have variadic function for additional arguments
	var argslice bool = false
	if cfunc.Type().IsVariadic() {
		argslice = true
	}

	return func(aditionalData unsafe.Pointer) {
		// Create slice of reflect.Value data
		var args = make([]reflect.Value, 0)
		if len(data) > 0 {
			for _, d := range data {
				args = append(args, reflect.ValueOf(d))
			}
		}
		if argslice == true {
			if converterFunc != nil {
				val := converterFunc(aditionalData)
				args = append(args, reflect.ValueOf(val))
			} else {
				args = append(args, reflect.ValueOf(aditionalData))
			}
		}
		cfunc.Call(args)
	}
}

type ListLike interface {
	NativeList() unsafe.Pointer
}

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

// To be list-like
func (self *GSList) NativeList() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self *GSList) Free() {
	C.g_slist_free(self.object)
}

func (self *GSList) FreeFull() {
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
	id := gobject.GetUniqueID()
	cl := createListClosure(f, self.ConversionFunc, data...)

	// Add Closure to closure map
	_closures[id] = cl

	// Call C function
	C._g_list_foreach(self.object, C.gint64(id))

	// Delete closure in closures
	delete(_closures, id)
}

//////////////////////////////
// End GSList
////////////////////////////// }}}

// Exported functions
//export _g_func
func _g_func(data, user_data unsafe.Pointer) {
	id := *((*C.gint64)(user_data))

	if f, ok := _closures[int64(id)]; ok {
		f(data)
	}

}

func init() {
	_closures = make(map[int64]ListClosure)
}
