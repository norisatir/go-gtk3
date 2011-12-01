package gtk3

/*
#include <gtk/gtk.h>

static inline GtkGrid* to_GtkGrid(void* obj) { return GTK_GRID(obj); }

static GtkWidget* _gtk_grid_get_child_at(GtkGrid* grid, gint left, gint top) {
#if GTK_CHECK_VERSION(3,2,0)
	return gtk_grid_get_child_at(grid, left, top);
#else
	return NULL;
#endif
}

static void _gtk_grid_insert_row(GtkGrid* grid, gint position) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_row(grid, position);
#endif
}

static void _gtk_grid_insert_column(GtkGrid* grid, gint position) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_column(grid, position);
#endif
}

static void _gtk_grid_insert_next_to(GtkGrid* grid, GtkWidget* sibling, GtkPositionType side) {
#if GTK_CHECK_VERSION(3,2,0)
	gtk_grid_insert_next_to(grid, sibling, side);
#endif
}

*/
import "C"
import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"

type Grid struct {
	object *C.GtkGrid
	*Container
}

func NewGrid() *Grid {
	grid := &Grid{}
	o := C.gtk_grid_new()
	grid.object = C.to_GtkGrid(unsafe.Pointer(o))
	grid.Container = NewContainer(unsafe.Pointer(o))

	if gobject.IsObjectFloating(grid) {
		gobject.RefSink(grid)
	}
	gridFinalizer(grid)

	return grid
}

// Clear Grid struct when it goes out of reach
func gridFinalizer(g *Grid) {
	runtime.SetFinalizer(g, func(g *Grid) { gobject.Unref(g) })
}

// Conversion function for gobject registration map
func newGridFromNative(obj unsafe.Pointer) interface{} {
	grid := &Grid{}
	grid.object = C.to_GtkGrid(obj)
	grid.Container = NewContainer(obj)

	if gobject.IsObjectFloating(grid) {
		gobject.RefSink(grid)
	} else {
		gobject.Ref(grid)
	}
	gridFinalizer(grid)

	return grid
}

func nativeFromGrid(g interface{}) *gobject.GValue {
	grid, ok := g.(*Grid)
	if ok {
		gv := gobject.CreateCGValue(GtkType.GRID, grid.ToNative())
		return gv
	}

	return nil
}

func init() {
	// Register GtkGrid to gobject
	gobject.RegisterCType(GtkType.GRID, newGridFromNative)
	gobject.RegisterGoType(GtkType.GRID, nativeFromGrid)
}

// To be Object-like
func (self Grid) ToNative() unsafe.Pointer {
	return unsafe.Pointer(self.object)
}

func (self Grid) Connect(name string ,f interface{}, data ...interface{}) (*gobject.ClosureElement, *gobject.SignalError) {
	return gobject.Connect(self, name, f, data...)
}

func (self Grid) Set(properties map[string]interface{}) {
	gobject.Set(self, properties)
}

func (self Grid) Get(properties []string) map[string]interface{} {
	return gobject.Get(self, properties)
}
// To be container-like
func (self Grid) C() *Container {
	return self.Container
}

// Grid interface
func (self *Grid) Attach(w WidgetLike, left, top, width, height int) {
	C.gtk_grid_attach(self.object, w.W().object, C.gint(left), C.gint(top),
		C.gint(width), C.gint(height))
}

func (self *Grid) AttachNextTo(child WidgetLike, sibling WidgetLike, side, width, height int) {
	C.gtk_grid_attach_next_to(self.object, child.W().object, sibling.W().object, C.GtkPositionType(side),
		C.gint(width), C.gint(height))
}

func (self *Grid) GetChildAt(left, top int) WidgetLike {
	c := C._gtk_grid_get_child_at(self.object, C.gint(left), C.gint(top))
	if c == nil {
		return nil
	}

	child, err := gobject.ConvertToGo(unsafe.Pointer(c))
	if err == nil {
		return child.(WidgetLike)
	}

	return nil
}

func (self *Grid) InsertRow(position int) {
	C._gtk_grid_insert_row(self.object, C.gint(position))
}

func (self *Grid) InsertColumn(position int) {
	C._gtk_grid_insert_column(self.object, C.gint(position))
}

func (self *Grid) InsertNextTo(sibling WidgetLike, side int) {
	C._gtk_grid_insert_next_to(self.object, sibling.W().object, C.GtkPositionType(side))
}

func (self *Grid) SetRowHomogeneous(homogeneous bool) {
	b := gobject.GBool(homogeneous)
	defer b.Free()
	C.gtk_grid_set_row_homogeneous(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Grid) GetRowHomogeneous() bool {
	b := C.gtk_grid_get_row_homogeneous(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Grid) SetRowSpacing(spacing uint) {
	C.gtk_grid_set_row_spacing(self.object, C.guint(spacing))
}

func (self *Grid) GetRowSpacing() uint {
	return uint(C.gtk_grid_get_row_spacing(self.object))
}

func (self *Grid) SetColumnHomogeneous(homogeneous bool) {
	b := gobject.GBool(homogeneous)
	defer b.Free()
	C.gtk_grid_set_column_homogeneous(self.object, *((*C.gboolean)(b.GetPtr())))
}

func (self *Grid) GetColumnHomogeneous() bool {
	b := C.gtk_grid_get_column_homogeneous(self.object)
	return gobject.GoBool(unsafe.Pointer(&b))
}

func (self *Grid) SetColumnSpacing(spacing uint) {
	C.gtk_grid_set_column_spacing(self.object, C.guint(spacing))
}

func (self *Grid) GetColumnSpacing() uint {
	return uint(C.gtk_grid_get_column_spacing(self.object))
}
