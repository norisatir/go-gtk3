package gobject

/*
#include <glib-object.h>

static inline GType get_type(GParamSpec *spec) {
	return G_PARAM_SPEC_VALUE_TYPE(spec);
}

*/
import "C"
import "unsafe"

// Definition of GParamSpec Type
type GParamSpec struct {
	object *C.GParamSpec
}

func (self GParamSpec) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self GParamSpec) ValueSetDefaults(val *GValue) {
	C.g_param_value_set_default((*C.GParamSpec)(self.ToNative()),
		(*C.GValue)(val.ToCGValue()))
}

func (self GParamSpec) ValueDefaults(val *GValue) bool {
	v := C.g_param_value_defaults((*C.GParamSpec)(self.ToNative()),
		(*C.GValue)(val.ToCGValue()))
	return GoBool(unsafe.Pointer(&v))
}

func (self GParamSpec) ValueValidate(val *GValue) bool {
	v := C.g_param_value_validate((*C.GParamSpec)(self.ToNative()),
		(*C.GValue)(val.ToCGValue()))
	return GoBool(unsafe.Pointer(&v))
}

func (self GParamSpec) ValueConvert(src, dest *GValue, strict_validation bool) bool {
	sv := GBool(strict_validation)
	res := C.g_param_value_convert((*C.GParamSpec)(self.ToNative()),
		(*C.GValue)(src.ToCGValue()),
		(*C.GValue)(dest.ToCGValue()),
		*((*C.gboolean)(sv.GetPtr())))
	return GoBool(unsafe.Pointer(&res))
}

func (self GParamSpec) GetType() GType {
	return GType(C.get_type(self.object))
}
