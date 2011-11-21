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

type ClosureFunc func(args []interface{}) bool

var _closures map[uint64][]ClosureFunc

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
	if handlers, ok := _closures[sig.id]; ok {
		for _, h := range handlers {
			if ok := h(sig.args); !ok {
				break
			}
		}
	}
	done <- true
}

func RegisterHandler(obj ObjectLike, name string, id uint64, f ClosureFunc) {
	// if id exists, then add closure to the end
	if _, ok := _closures[id]; ok {
		_closures[id] = append(_closures[id], f)
		return
	}

	// Not found, so create new slice of closures
	_closures[id] = make([]ClosureFunc, 1)
	_closures[id][0] = f

	// Register handler in gobject system
	s := GString(name)
	defer s.Free()
	C.connect_to_signal(obj.ToNative(),
		(*C.gchar)(s.GetPtr()), C.guint64(id))
}
func init() {
	_closures = make(map[uint64][]ClosureFunc)
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
