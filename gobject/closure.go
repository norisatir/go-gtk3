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

static inline gulong connect_to_signal(void* obj, gchar* name, gint64 id) {
	gint64 *pgint64 = (gint64*)malloc(sizeof(gint64));
	*pgint64 = id;
	GClosure* c = g_cclosure_new_swap(G_CALLBACK(func_handler),
									(gpointer)pgint64,
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

// This gets returned by RegisterHandler
type ClosureElement struct {
	id         int64
	handler_id uint32
}

var _closures map[int64]ClosureFunc

// RegisterHandler registers closure and returns ClosureElement with info for removal
func RegisterHandler(obj ObjectLike, name string, id int64, f ClosureFunc) *ClosureElement {
	// Register handler to our gobject system
	_closures[id] = f

	// Register handler in gobject system
	s := GString(name)
	defer s.Free()
	h_id := C.connect_to_signal(obj.ToNative(),
		(*C.gchar)(s.GetPtr()), C.gint64(id))

	return &ClosureElement{id, uint32(h_id)}
}

// RemoveHandler removes closure from our map and from gobject signal system
func RemoveHandler(obj ObjectLike, cel *ClosureElement) {
	if _, ok := _closures[cel.id]; ok {
		_closures[cel.id] = nil, false
		C.g_signal_handler_disconnect(C.gpointer(obj.ToNative()), C.gulong(cel.handler_id))
	}

}

func init() {
	_closures = make(map[int64]ClosureFunc)
}

// SimpleGoMarshal
// Our true callback which gets called
// for all connected signals.
// It call appropriate closure which it finds in closures map
//export simple_go_marshal
func simple_go_marshal(closure unsafe.Pointer,
	returnValue unsafe.Pointer,
	n_param_values C.guint,
	paramValues unsafe.Pointer,
	invocationHint unsafe.Pointer,
	marshalData unsafe.Pointer) {

	c := (*C.GClosure)(closure)
	id := int64(*((*C.gint64)(c.data)))

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
			argslice[i] = nil
		}
	}

	if h, ok := _closures[id]; ok {
		h(argslice)
	}

	C.free_array(array)
}
