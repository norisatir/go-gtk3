#include "go-gobject.h"

GObject* to_GObject(void* o) { return G_OBJECT(o); }

void* new_GObject(GType typ) { return g_object_new(typ, NULL); }

const gchar* getTypeName(void * o) { return G_OBJECT_TYPE_NAME(o); }

GObjectClass* get_object_class(void* o) {
	return G_OBJECT_GET_CLASS(o);
}

GType _get_type_from_instance(void* o) {
    return G_TYPE_FROM_INSTANCE(o);
}

void destroy_data(gpointer data) {
	free(data);
}

void _g_clear_object(void* object) {
	GObject* o = G_OBJECT(object);
	g_clear_object(&o);
}

GBinding* to_GBinding(void* obj) { return G_BINDING(obj); }

GParamSpec* to_GParamSpec(void* obj) { return G_PARAM_SPEC(obj); }
GType get_type(GParamSpec *spec) {
	return G_PARAM_SPEC_VALUE_TYPE(spec);
}

GArray* g_array_from_GValues(constgvalue val, guint num_elements) {
	GArray* na = g_array_new(TRUE, TRUE, sizeof(GValue));
	guint i;
	for(i = 0; i < num_elements; i++) {
		g_array_append_val(na, *(val + i));
	}
	return na;
}

GValue get_index(GArray* ar, guint i) {
	return g_array_index(ar, GValue, i);
}

void free_array(GArray* ar) {
	g_array_free(ar, TRUE);
}

const gchar* get_type_name(GValue v) { return G_VALUE_TYPE_NAME(&v); }

GType getTypeId(void * o) { return G_OBJECT_TYPE(o); }

GType getGValueType(GValue* v) {
	return G_VALUE_TYPE(v);
}

const gchar* getGValueTypeName(GValue* v) {
	return G_VALUE_TYPE_NAME(v);
}

gboolean is_type_object(GType t) {
	return G_TYPE_IS_OBJECT(t);
}

gboolean _is_type_enum(GType t) {
	return G_TYPE_IS_ENUM(t);
}

gboolean _is_type_flags(GType t) {
	return G_TYPE_IS_FLAGS(t);
}

gboolean is_type_boxed(GType t) {
	return G_TYPE_IS_BOXED(t);
}

gboolean boxed_in_gvalue(GValue* v) {
	return G_VALUE_HOLDS_BOXED(v);
}

gboolean pointer_in_gvalue(GValue* v) {
	return G_VALUE_HOLDS_POINTER(v);
}

gboolean object_in_gvalue(GValue* v) {
	return G_VALUE_HOLDS_OBJECT(v);
}

gboolean _enum_in_gvalue(GValue* v) {
	return G_VALUE_HOLDS_ENUM(v);
}

gboolean _flags_in_gvalue(GValue* v) {
	return G_VALUE_HOLDS_FLAGS(v);
}

gboolean _pointer_in_gvalue(GValue* v) {
	return G_VALUE_HOLDS_POINTER(v);
}

void _g_object_weak_ref(GObject* object, gint64 data) {
	gint64* intdata = (gint64*)malloc(sizeof(gint64));
	*intdata = data;
	g_object_weak_ref(object, weak_ref_callback, (gpointer)intdata);
}

void destroy_id(gpointer data, GClosure* closure) {
	free(data);
}

// Dummy callback function
// needed for simple_go_marshal function
// which is our true callback
gboolean func_handler(void* id, ...) {
	return TRUE;
}

gulong connect_to_signal(void* obj, gchar* name, gint64 id) {
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
