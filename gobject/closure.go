package gobject

/*
#include <glib-object.h>
#include <stdlib.h>

extern void simple_go_marshal(GClosure *closure,
							  GValue* returnValue,
							  guint n_param_values,
							  const GValue *paramValues,
							  gpointer invocationHint,
							  gpointer marshalData);

// Dummy callback function
// needed for simple_go_marshal function
// which is our true callback
static gboolean func_handler(void* id, ...) {
	return TRUE;
}

static inline void destroy_id(gpointer data, GClosure* closure) {
	free(data);
}

static inline gulong connect_to_signal(void* obj, gchar* name, guint64 id) {
	guint64 *pguint64 = (guint64*)malloc(sizeof(guint64));
	*pguint64 = id;
	GClosure* c = g_cclosure_new_swap(G_CALLBACK(func_handler),
									(gpointer)pguint64,
									destroy_id);

	g_closure_set_marshal(c, simple_go_marshal);

	gulong handler_id = g_signal_connect_closure((gpointer)obj,
										   name,
										   c,
										   TRUE);
	return handler_id;
}

static inline GArray* g_array_from_GValues(void* val, guint num_elements) {
	GValue* values = (GValue*)val;
	GArray* na = g_array_new(TRUE, TRUE, sizeof(GValue));
	guint i;
	for(i = 0; i < num_elements; i++) {
		g_array_append_val(na, *(values + i));
	}
	return na;
}

static inline GValue get_index(GArray* ar, guint i) {
	return g_array_index(ar, GValue, i);
}

static inline void free_array(GArray* ar) {
	g_array_free(ar, TRUE);
}



*/
import "C"
import "unsafe"
import "container/list"

type ClosureFunc func(args []interface{}) bool

// This gets returned by RegisterHandler
type ClosureElement struct {
	id      uint64
	closure *list.Element
}

//var _closures map[uint64][]ClosureFunc
var _closures map[uint64]*list.List

// Signal struct
type Signal struct {
	id   uint64
	args []interface{}
}

var emitedSignals chan Signal
var doneMarshal chan bool

func queueManager() {
	var done = make(chan bool)
	for s := range emitedSignals {
		go run(done, s)
		<-done
		doneMarshal <- true
	}
}
func run(done chan bool, sig Signal) {
	if l, ok := _closures[sig.id]; ok {
		for h := l.Front(); h != nil; h = h.Next() {
			handler := h.Value.(ClosureFunc)
			if ok := handler(sig.args); !ok {
				break
			}
		}
	}
	done <- true
}

// Return list element which holds our closure (and with which we can disconnect it)
func RegisterHandler(obj ObjectLike, name string, id uint64, f ClosureFunc) *ClosureElement {
	// if id exists, then add closure to the end
	if _, ok := _closures[id]; ok {
		el := _closures[id].PushBack(f)
		return &ClosureElement{id, el}
	}

	// Not found, so create new list of closures
	_closures[id] = list.New()
	el := _closures[id].PushBack(f)

	// Register handler in gobject system
	s := GString(name)
	defer s.Free()
	C.connect_to_signal(obj.ToNative(),
		(*C.gchar)(s.GetPtr()), C.guint64(id))

	return &ClosureElement{id, el}
}

// Removes handler from list (and entire map element,
// if list is empty - no closures anymore)
func RemoveHandler(cel *ClosureElement) {
	if l, ok := _closures[cel.id]; ok {
		l.Remove(cel.closure)

		// if list now empty, remove map element
		if l.Len() == 0 {
			_closures[cel.id] = nil, false
		}
	}
}

func init() {
	_closures = make(map[uint64]*list.List)
	emitedSignals = make(chan Signal, 100)
	doneMarshal = make(chan bool)
	go queueManager()
}

// SimpleGoMarshal
// Our true callback which gets called
// for all connected signals.
// It emits unique id of closure to queue manager.
//export simple_go_marshal
func simple_go_marshal(closure unsafe.Pointer,
	returnValue unsafe.Pointer,
	n_param_values C.guint,
	paramValues unsafe.Pointer,
	invocationHint unsafe.Pointer,
	marshalData unsafe.Pointer) {

	c := (*C.GClosure)(closure)
	id := uint64(*((*C.guint64)(c.data)))

	argslice := make([]interface{}, int(n_param_values))
	array := C.g_array_from_GValues(paramValues, n_param_values)
	for i := 0; i < int(n_param_values); i++ {
		v := C.get_index(array, C.guint(i))
		gv := GValue{0, &v}
		gv.ReInitializeType()
		t, e := ConvertToGo(gv.GetPtr(), gv.GetTypeID())
		if e == nil {
			argslice[i] = t
		} else {
			argslice = nil
		}
	}

	emitedSignals <- Signal{id, argslice}
	<-doneMarshal

	C.free_array(array)
}
