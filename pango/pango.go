package pango

/*
#include <pango/pango.h>
*/
// #cgo pkg-config: pango
import "C"

/*import "unsafe"
import "runtime"
import "github.com/norisatir/go-gtk3/gobject"
*/

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
} // }}}
