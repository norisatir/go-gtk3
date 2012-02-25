#ifndef _GO_GOBJECT_H_
#define _GO_GOBJECT_H

#include <stdlib.h>
//#include <stdint.h>
#include <glib.h>
#include <glib-object.h>

typedef const GValue* constgvalue;

GObject* to_GObject(void* o);
void* new_GObject(GType typ);

const gchar* getTypeName(void * o);

GObjectClass* get_object_class(void* o);

GType _get_type_from_instance(void* o);

void destroy_data(gpointer data);

void _g_clear_object(void* object);

GBinding* to_GBinding(void* obj);

GParamSpec* to_GParamSpec(void* obj);
GType get_type(GParamSpec *spec);

GArray* g_array_from_GValues(constgvalue val, guint num_elements);

GValue get_index(GArray* ar, guint i);

void free_array(GArray* ar);

const gchar* get_type_name(GValue v);

GType getTypeId(void * o);

GType getGValueType(GValue* v);

const gchar* getGValueTypeName(GValue* v);

gboolean is_type_object(GType t);

gboolean _is_type_enum(GType t);

gboolean _is_type_flags(GType t);

gboolean is_type_boxed(GType t);

gboolean boxed_in_gvalue(GValue* v);

gboolean pointer_in_gvalue(GValue* v);

gboolean object_in_gvalue(GValue* v);

gboolean _enum_in_gvalue(GValue* v);

gboolean _flags_in_gvalue(GValue* v);

gboolean _pointer_in_gvalue(GValue* v);

extern void simple_go_marshal(GClosure *closure,
							  GValue* returnValue,
							  guint n_param_values,
							  constgvalue paramValues,
							  gpointer invocationHint,
							  gpointer marshalData);

extern void weak_ref_callback(gpointer data, GObject* object);

void _g_object_weak_ref(GObject* object, gint64 data);

gulong connect_to_signal(void* obj, gchar* name, gint64 id);

#endif
