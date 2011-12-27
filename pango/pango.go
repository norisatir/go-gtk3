package pango

/*
#include <pango/pango.h>

static inline PangoContext* to_PangoContext(void* obj) { return PANGO_CONTEXT(obj); }
static inline PangoFont* to_PangoFont(void* obj) { return PANGO_FONT(obj); }
static inline PangoFontFamily* to_PangoFontFamily(void* obj) { return PANGO_FONT_FAMILY(obj); }
static inline PangoFontFace* to_PangoFontFace(void* obj) { return PANGO_FONT_FACE(obj); }
static inline PangoFontMap* to_PangoFontMap(void* obj) { return PANGO_FONT_MAP(obj); }
static inline PangoFontset* to_PangoFontset(void* obj) { return PANGO_FONTSET(obj); }
static inline PangoLayout* to_PangoLayout(void* obj) { return PANGO_LAYOUT(obj); }
*/
// #cgo pkg-config: pango
import "C"

import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type pangoTypes struct {
	CONTEXT gobject.GType
	FONT gobject.GType
	FONT_FAMILY gobject.GType
	FONT_FACE gobject.GType
	FONT_MAP gobject.GType
	FONTSET gobject.GType
	FONT_DESCRIPTION gobject.GType
	LAYOUT gobject.GType
}

var PangoTypes pangoTypes

// PangoWeight {{{

var PangoWeight pangoWeight

type pangoWeight struct {
	THIN       int
	ULTRALIGHT int
	LIGHT      int
	BOOK       int
	NORMAL     int
	MEDIUM     int
	SEMIBOLD   int
	BOLD       int
	ULTRABOLD  int
	HEAVY      int
	ULTRAHEAVY int
}

// PangoWeight }}}

// PangoStyle {{{

var PangoStyle pangoStyle

type pangoStyle struct {
	NORMAL  int
	OBLIQUE int
	ITALIC  int
}

// End PangoStyle }}}

// PangoScale {{{

var PangoScale pangoScale

type pangoScale struct {
	Scale    int
	XX_SMALL float64
	X_SMALL  float64
	SMALL    float64
	MEDIUM   float64
	LARGE    float64
	X_LARGE  float64
	XX_LARGE float64
}

// End PangoScale }}}

// PangoUnderline {{{

var PangoUnderline pangoUnderline

type pangoUnderline struct {
	NONE   int
	SINGLE int
	DOUBLE int
	LOW    int
	ERROR  int
}

// End PangoUnderline }}}

// PangoVariant {{{

var PangoVariant pangoVariant

type pangoVariant struct {
	NORMAL int
	SMALL_CAPS int
}
// End PangoVariant }}}

// PangoStretch {{{

var PangoStretch pangoStretch

type pangoStretch struct {
	ULTRA_CONDENSED int
	EXTRA_CONDENSED int
	CONDENSED int
	SEMI_CONDENSED int
	NORMAL int
	SEMI_EXPANDED int
	EXPANDED int
	EXTRA_EXPANDED int
	ULTRA_EXPANDED int
}
// End PangoStretch }}}

// PangoGravity {{{

var PangoGravity pangoGravity

type pangoGravity struct {
	SOUTH int
	EAST int
	NORTH int
	WEST int
	AUTO int
}
// End PangoGravity }}}

// PangoFontMask {{{

var PangoFontMask pangoFontMask

type pangoFontMask struct {
	FAMILY int
	STYLE int
	VARIANT int
	WEIGHT int
	STRETCH int
	SIZE int
	GRAVITY int
}
// End PangoFontMask }}}

// Rendering {{{

// PangoContext {{{
//////////////////////////////

// PangoContext Type
type Context struct {
	object *C.PangoContext
}

// PangoContext finalizer
func contextFinalizer(c *Context) {
	runtime.SetFinalizer(c, func(c *Context) { gobject.Unref(c) })
}

// Conversion funcs
func newContextFromNative(obj unsafe.Pointer) interface{} {
	c := &Context{}
	c.object = C.to_PangoContext(obj)

	if gobject.IsObjectFloating(c) {
		gobject.RefSink(c)
	} else {
		gobject.Ref(c)
	}
	contextFinalizer(c)

	return c
}

func nativeFromContext(pc interface{}) *gobject.GValue {
	if c, ok := pc.(*Context); ok {
		gv := gobject.CreateCGValue(PangoTypes.CONTEXT, c.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Context) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Context) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Context) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Context) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// PangoContext interface

//////////////////////////////
// End PangoContext
////////////////////////////// }}}

// End Rendering }}}

// Fonts {{{

// Font {{{
//////////////////////////////

// Font type
type Font struct {
	object *C.PangoFont
}

// Font finalizer
func fontFinalizer(f *Font) {
	runtime.SetFinalizer(f, func(f *Font) { gobject.Unref(f) })
}

// Conversion funcs
func newFontFromNative(obj unsafe.Pointer) interface{} {
	f := &Font{}
	f.object = C.to_PangoFont(obj)

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	} else {
		gobject.Ref(f)
	}
	fontFinalizer(f)

	return f
}

func nativeFromFont(font interface{}) *gobject.GValue {
	if f, ok := font.(*Font); ok {
		gv := gobject.CreateCGValue(PangoTypes.FONT, f.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Font) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Font) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Font) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Font) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Font interface

//////////////////////////////
// End Font
////////////////////////////// }}}

// FontFamily {{{
//////////////////////////////

// FontFamily type
type FontFamily struct {
	object *C.PangoFontFamily
}

// FontFamily finalizer
func fontFamilyFinalizer(f *FontFamily) {
	runtime.SetFinalizer(f, func(f *FontFamily) { gobject.Unref(f) })
}

// Conversion funcs
func newFontFamilyFromNative(obj unsafe.Pointer) interface{} {
	f := &FontFamily{}
	f.object = C.to_PangoFontFamily(obj)

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	} else {
		gobject.Ref(f)
	}
	fontFamilyFinalizer(f)

	return f
}

func nativeFromFontFamily(ff interface{}) *gobject.GValue {
	if f, ok := ff.(*FontFamily); ok {
		gv := gobject.CreateCGValue(PangoTypes.FONT_FAMILY, f.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self FontFamily) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self FontFamily) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self FontFamily) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self FontFamily) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// FontFamily interface

//////////////////////////////
// End FontFamily
////////////////////////////// }}}

// FontFace {{{
//////////////////////////////

// FontFace type
type FontFace struct {
	object *C.PangoFontFace
}

// FontFace finalizer
func fontFaceFinalizer(f *FontFace) {
	runtime.SetFinalizer(f, func(f *FontFace) { gobject.Unref(f) })
}

// Conversion funcs
func newFontFaceFromNative(obj unsafe.Pointer) interface{} {
	f := &FontFace{}
	f.object = C.to_PangoFontFace(obj)

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	} else {
		gobject.Ref(f)
	}
	fontFaceFinalizer(f)

	return f
}

func nativeFromFontFace(ff interface{}) *gobject.GValue {
	if f, ok := ff.(*FontFace); ok {
		gv := gobject.CreateCGValue(PangoTypes.FONT_FACE, f.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self FontFace) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self FontFace) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self FontFace) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self FontFace) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// FontFace interface

//////////////////////////////
// End FontFace
////////////////////////////// }}}

// FontMap {{{
//////////////////////////////

// FontMap type
type FontMap struct {
	object *C.PangoFontMap
}

// FontMap finalizer
func fontMapFinalizer(f *FontMap) {
	runtime.SetFinalizer(f, func(f *FontMap) { gobject.Unref(f) })
}

// Conversion funcs
func newFontMapFromNative(obj unsafe.Pointer) interface{} {
	f := &FontMap{}
	f.object = C.to_PangoFontMap(obj)

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	} else {
		gobject.Ref(f)
	}
	fontMapFinalizer(f)

	return f
}

func nativeFromFontMap(fm interface{}) *gobject.GValue {
	if f, ok := fm.(*FontMap); ok {
		gv := gobject.CreateCGValue(PangoTypes.FONT_MAP, f.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self FontMap) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self FontMap) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self FontMap) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self FontMap) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// FontMap interface

//////////////////////////////
// End FontMap
////////////////////////////// }}}

// Fontset {{{
//////////////////////////////

// Fontset type
type Fontset struct {
	object *C.PangoFontset
}

// Fontset finalizer
func fontsetFinalizer(f *Fontset) {
	runtime.SetFinalizer(f, func(f *Fontset) { gobject.Unref(f) })
}

// Conversion funcs
func newFontsetFromNative(obj unsafe.Pointer) interface{} {
	f := &Fontset{}
	f.object = C.to_PangoFontset(obj)

	if gobject.IsObjectFloating(f) {
		gobject.RefSink(f)
	} else {
		gobject.Ref(f)
	}
	fontsetFinalizer(f)

	return f
}

func nativeFromFontset(fs interface{}) *gobject.GValue {
	if f, ok := fs.(*Fontset); ok {
		gv := gobject.CreateCGValue(PangoTypes.FONTSET, f.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Fontset) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Fontset) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Fontset) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Fontset) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Fontset interface

//////////////////////////////
// End Fontset
////////////////////////////// }}}

// FontDescription {{{
//////////////////////////////

// FontDescription type
type FontDescription struct {
	object *C.PangoFontDescription
}

func NewFontDescription() *FontDescription {
	f := &FontDescription{}
	f.object = C.pango_font_description_new()
	FontDescriptionFinalizer(f)

	return f
}

func NewFontDescriptionFromString(str string) *FontDescription {
	s := gobject.GString(str)
	defer s.Free()

	f := &FontDescription{}
	f.object = C.pango_font_description_from_string((*C.char)(s.GetPtr()))
	FontDescriptionFinalizer(f)

	return f
}

func FontDescriptionFinalizer(fd *FontDescription) {
	runtime.SetFinalizer(fd, func(fd *FontDescription) { fd.Free() })
}

// Conversion funcs
func newFontDescriptionFromNative(obj unsafe.Pointer) interface{} {
	f := &FontDescription{}
	f.object = (*C.PangoFontDescription)(obj)

	return f
}

func nativeFromFontDescription(fd interface{}) *gobject.GValue {
	if f, ok := fd.(*FontDescription); ok {
		gv := gobject.CreateCGValue(PangoTypes.FONT_DESCRIPTION, f.ToNative())
		return gv
	}                                           
	return nil
}

// To be boxed-like
func (self FontDescription) GetBoxedType() gobject.GType {
	return PangoTypes.FONT_DESCRIPTION
}

// FontDescription interface

func (self *FontDescription) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self *FontDescription) Free() {
	C.pango_font_description_free(self.object)
}

func (self *FontDescription) SetFamily(family string) {
	s := gobject.GString(family)
	defer s.Free()
	C.pango_font_description_set_family(self.object, (*C.char)(s.GetPtr()))
}

func (self *FontDescription) GetFamily() string {
	s := C.pango_font_description_get_family(self.object)
	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

func (self *FontDescription) SetStyle(pango_Style int) {
	C.pango_font_description_set_style(self.object, C.PangoStyle(pango_Style))
}

func (self *FontDescription) GetStyle() int {
	return int(C.pango_font_description_get_style(self.object))
}

func (self *FontDescription) SetVariant(pango_Variant int) {
	C.pango_font_description_set_variant(self.object, C.PangoVariant(pango_Variant))
}

func (self *FontDescription) GetVariant() int {
	return int(C.pango_font_description_get_variant(self.object))
}

func (self *FontDescription) SetWeight(pango_Weight int) {
	C.pango_font_description_set_weight(self.object, C.PangoWeight(pango_Weight))
}

func (self *FontDescription) GetWeight() int {
	return int(C.pango_font_description_get_weight(self.object))
}

func (self *FontDescription) SetStretch(pango_Stretch int) {
	C.pango_font_description_set_stretch(self.object, C.PangoStretch(pango_Stretch))
}

func (self *FontDescription) GetStretch() int {
	return int(C.pango_font_description_get_stretch(self.object))
}

func (self *FontDescription) SetSize(size int) {
	C.pango_font_description_set_size(self.object, C.gint(size))
}

func (self *FontDescription) GetSize() int {
	return int(C.pango_font_description_get_size(self.object))
}

func (self *FontDescription) SetAbsoluteSize(size float64) {
	C.pango_font_description_set_absolute_size(self.object, C.double(size))
}

func (self *FontDescription) GetSizeIsAbsolute() bool {
	b := C.pango_font_description_get_size_is_absolute(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *FontDescription) SetGravity(pango_Gravity int) {
	C.pango_font_description_set_gravity(self.object, C.PangoGravity(pango_Gravity))
}

func (self *FontDescription) GetGravity() int {
	return int(C.pango_font_description_get_gravity(self.object))
}

func (self *FontDescription) GetSetFields() int {
	return int(C.pango_font_description_get_set_fields(self.object))
}

func (self *FontDescription) UnsetFields(pango_FontMask int) {
	C.pango_font_description_unset_fields(self.object, C.PangoFontMask(pango_FontMask))
}

func (self *FontDescription) Merget(descToMerge *FontDescription, replaceExisting bool) {
	b := gobject.GBool(replaceExisting)
	defer b.Free()

	if descToMerge == nil {
		C.pango_font_description_merge(self.object, nil, *((*C.gboolean)(b.GetPtr())))
		return
	}
	C.pango_font_description_merge(self.object, descToMerge.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *FontDescription) BetterMatch(oldMatch, newMatch *FontDescription) bool {
	var fDesc *C.PangoFontDescription = nil

	if oldMatch != nil {
		fDesc = oldMatch.object
	}

	b := C.pango_font_description_better_match(self.object, fDesc, newMatch.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *FontDescription) ToString() string {
	s := C.pango_font_description_to_string(self.object)
	
	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}

func (self *FontDescription) ToFilename() string {
	s := C.pango_font_description_to_filename(self.object)
	
	if s != nil {
		return gobject.GoString(unsafe.Pointer(s))
	}
	return ""
}
//////////////////////////////
// End FontDescription
////////////////////////////// }}}

// End Fonts  }}}

// Text Attributes {{{

// End Text Attributes }}}

// Layout Objects {{{

// Layout {{{
//////////////////////////////

// Layout type
type Layout struct {
	object *C.PangoLayout
}

// Layout finalizer
func layoutFinalizer(l *Layout) {
	runtime.SetFinalizer(l, func(l *Layout) { gobject.Unref(l) })
}

// Conversion funcs
func newLayoutFromNative(obj unsafe.Pointer) interface{} {
	l := &Layout{}
	l.object = C.to_PangoLayout(obj)

	if gobject.IsObjectFloating(l) {
		gobject.RefSink(l)
	} else {
		gobject.Ref(l)
	}
	layoutFinalizer(l)

	return l
}

func nativeFromLayout(layout interface{}) *gobject.GValue {
	if l, ok := layout.(*Layout); ok {
		gv := gobject.CreateCGValue(PangoTypes.LAYOUT, l.ToNative())
		return gv
	}
	return nil
}

// To be object-like
func (self Layout) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Layout) Connect(name string, f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Layout) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Layout) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}

// Layout interface

//////////////////////////////
// End Layout
////////////////////////////// }}}

// End Layout Objects }}}

// PANGO init funcs {{{
func init() {

	// Initialize PangoWeight
	PangoWeight.THIN = 100
	PangoWeight.ULTRALIGHT = 200
	PangoWeight.LIGHT = 300
	PangoWeight.BOOK = 380
	PangoWeight.NORMAL = 400
	PangoWeight.MEDIUM = 500
	PangoWeight.SEMIBOLD = 600
	PangoWeight.BOLD = 700
	PangoWeight.ULTRABOLD = 800
	PangoWeight.HEAVY = 900
	PangoWeight.ULTRAHEAVY = 1000

	// Initialize PangoStyle
	PangoStyle.NORMAL = 0
	PangoStyle.OBLIQUE = 1
	PangoStyle.ITALIC = 2

	// Initialize PangoScale
	PangoScale.Scale = 1024
	PangoScale.XX_SMALL = float64(0.5787037037037)
	PangoScale.X_SMALL = float64(0.6444444444444)
	PangoScale.SMALL = float64(0.8333333333333)
	PangoScale.MEDIUM = float64(1.0)
	PangoScale.LARGE = float64(1.2)
	PangoScale.X_LARGE = float64(1.4399999999999)
	PangoScale.XX_LARGE = float64(1.728)

	// Initialize PangoUnderline
	PangoUnderline.NONE = 0
	PangoUnderline.SINGLE = 1
	PangoUnderline.DOUBLE = 2
	PangoUnderline.LOW = 3
	PangoUnderline.ERROR = 4

	// Initialize PangoVariant
	PangoVariant.NORMAL = 0
	PangoVariant.SMALL_CAPS = 1

	// Initialize PangoStretch
	PangoStretch.ULTRA_CONDENSED = 0
	PangoStretch.EXTRA_CONDENSED = 1
	PangoStretch.CONDENSED = 2
	PangoStretch.SEMI_CONDENSED = 3
	PangoStretch.NORMAL = 4
	PangoStretch.SEMI_EXPANDED = 5
	PangoStretch.EXPANDED = 6
	PangoStretch.EXTRA_EXPANDED = 7
	PangoStretch.ULTRA_EXPANDED = 8

	// Initialize PangoGravity
	PangoGravity.SOUTH = 0
	PangoGravity.EAST = 1
	PangoGravity.NORTH = 2
	PangoGravity.WEST = 3
	PangoGravity.AUTO = 4

	// Initialize PangoFontMask
	PangoFontMask.FAMILY = 1 << 0
	PangoFontMask.STYLE = 1 << 1
	PangoFontMask.VARIANT = 1 << 2
	PangoFontMask.WEIGHT = 1 << 3
	PangoFontMask.STRETCH = 1 << 4
	PangoFontMask.SIZE = 1 << 5
	PangoFontMask.GRAVITY = 1 << 6

	// Init PangoTypes
	PangoTypes.CONTEXT = gobject.GType(C.pango_context_get_type())
	PangoTypes.FONT = gobject.GType(C.pango_font_get_type())
	PangoTypes.FONT_FAMILY = gobject.GType(C.pango_font_family_get_type())
	PangoTypes.FONT_FACE = gobject.GType(C.pango_font_face_get_type())
	PangoTypes.FONT_MAP = gobject.GType(C.pango_font_map_get_type())
	PangoTypes.FONTSET = gobject.GType(C.pango_fontset_get_type())
	PangoTypes.FONT_DESCRIPTION = gobject.GType(C.pango_font_description_get_type())
	PangoTypes.LAYOUT = gobject.GType(C.pango_layout_get_type())

	// Register PangoContext type
	gobject.RegisterCType(PangoTypes.CONTEXT, newContextFromNative)
	gobject.RegisterGoType(PangoTypes.CONTEXT, nativeFromContext)

	// Register Font type
	gobject.RegisterCType(PangoTypes.FONT, newFontFromNative)
	gobject.RegisterGoType(PangoTypes.FONT, nativeFromFont)

	// Register FontFamily type
	gobject.RegisterCType(PangoTypes.FONT_FAMILY, newFontFamilyFromNative)
	gobject.RegisterGoType(PangoTypes.FONT_FAMILY, nativeFromFontFamily)

	// Register FontFace type
	gobject.RegisterCType(PangoTypes.FONT_FACE, newFontFaceFromNative)
	gobject.RegisterGoType(PangoTypes.FONT_FACE, nativeFromFontFace)

	// Register FontMap type
	gobject.RegisterCType(PangoTypes.FONT_MAP, newFontMapFromNative)
	gobject.RegisterGoType(PangoTypes.FONT_MAP, nativeFromFontMap)

	// Register Fontset type
	gobject.RegisterCType(PangoTypes.FONTSET, newFontsetFromNative)
	gobject.RegisterGoType(PangoTypes.FONTSET, nativeFromFontset)

	// Register FontDescription type
	gobject.RegisterCType(PangoTypes.FONT_DESCRIPTION, newFontDescriptionFromNative)
	gobject.RegisterGoType(PangoTypes.FONT_DESCRIPTION, nativeFromFontDescription)

	// Register Layout type
	gobject.RegisterCType(PangoTypes.LAYOUT, newLayoutFromNative)
	gobject.RegisterGoType(PangoTypes.LAYOUT, nativeFromLayout)

} // }}}
