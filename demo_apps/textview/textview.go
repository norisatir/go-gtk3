package main

import "github.com/norisatir/go-gtk3/gtk3"
import "github.com/norisatir/go-gtk3/gdkpixbuf"
import "github.com/norisatir/go-gtk3/pango"
import "path/filepath"
import "os"

func findFile(filename string) (string, error) {
	if filepath.IsAbs(filename) {
		return filename, nil
	}

	dir, _ := filepath.Split(filepath.Clean(os.Args[0]))

	fn := filepath.Join(dir, filename)

	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}

	f.Close()
	return fn, nil
}

func CreateTags(buffer *gtk3.TextBuffer) {
	buffer.CreateTag("heading", gtk3.P{"weight": pango.PangoWeight.BOLD,
		"size": 15 * pango.PangoScale.Scale})

	buffer.CreateTag("italic", gtk3.P{"style": pango.PangoStyle.ITALIC})
	buffer.CreateTag("bold", gtk3.P{"weight": pango.PangoWeight.BOLD})
	buffer.CreateTag("big", gtk3.P{"size": 20 * pango.PangoScale.Scale})
	buffer.CreateTag("xx-small", gtk3.P{"scale": pango.PangoScale.XX_SMALL})
	buffer.CreateTag("x-large", gtk3.P{"scale": pango.PangoScale.X_LARGE})
	buffer.CreateTag("monospace", gtk3.P{"family": "monospace"})
	buffer.CreateTag("blue_foreground", gtk3.P{"foreground": "blue"})
	buffer.CreateTag("red_background", gtk3.P{"background": "red"})
	buffer.CreateTag("big_gap_before_line", gtk3.P{"pixels_above_lines": 30})
	buffer.CreateTag("big_gap_after_line", gtk3.P{"pixels_below_lines": 30})
	buffer.CreateTag("double_spaced_line", gtk3.P{"pixels_inside_wrap": 10})
	buffer.CreateTag("not_editable", gtk3.P{"editable": false})
	buffer.CreateTag("word_wrap", gtk3.P{"wrap_mode": gtk3.GtkWrapMode.WORD})
	buffer.CreateTag("char_wrap", gtk3.P{"wrap_mode": gtk3.GtkWrapMode.CHAR})
	buffer.CreateTag("no_wrap", gtk3.P{"wrap_mode": gtk3.GtkWrapMode.NONE})
	buffer.CreateTag("center", gtk3.P{"justification": gtk3.GtkJustification.CENTER})
	buffer.CreateTag("right_justify", gtk3.P{"justification": gtk3.GtkJustification.RIGHT})
	buffer.CreateTag("wide_margins", gtk3.P{"left_margin": 50, "right_margin": 50})
	buffer.CreateTag("strikethrough", gtk3.P{"strikethrough": true})
	buffer.CreateTag("underline", gtk3.P{"underline": pango.PangoUnderline.SINGLE})
	buffer.CreateTag("double_underline", gtk3.P{"underline": pango.PangoUnderline.DOUBLE})
	buffer.CreateTag("superscript", gtk3.P{"rise": 10 * pango.PangoScale.Scale, "size": 8 * pango.PangoScale.Scale})
	buffer.CreateTag("subscript", gtk3.P{"rise": -10 * pango.PangoScale.Scale, "size": 8 * pango.PangoScale.Scale})
	buffer.CreateTag("rtl_quote", gtk3.P{"wrap_mode": gtk3.GtkWrapMode.WORD, "direction": gtk3.GtkTextDirection.RTL,
		"indent": 30, "left_margin": 20, "right_margin": 20})
}

func InsertText(buffer *gtk3.TextBuffer) {
	var iter, start, end gtk3.TextIter

	filename, err := findFile("gtk-logo-rgb.gif")
	if err != nil {
		panic(err)
	}
	pixbuf, err := gdkpixbuf.NewFromFile(filename)

	if err != nil {
		panic(err)
	}

	pixbuf = gdkpixbuf.ScaleSimple(pixbuf, 32, 32, gdkpixbuf.GdkInterp.BILINEAR)

	// Get the start of buffer; each insertion will revalidate the
	// iterator to point to just after the inserted text.
	buffer.GetIterAtOffset(&iter, 0)

	buffer.Insert(&iter, "The text widget can display text with all kinds of nifty attributes. "+
		"It also supports multiple views of the same buffer; this demo is "+
		"showing the same buffer in two places.\n\n")

	buffer.InsertWithTagsByName(&iter, "Font styles. ", "heading")

	buffer.Insert(&iter, "For example, you can have ")
	buffer.InsertWithTagsByName(&iter, "italic", "italic")

	buffer.Insert(&iter, ", ")
	buffer.InsertWithTagsByName(&iter, "bold", "bold")

	buffer.Insert(&iter, ", or ")
	buffer.InsertWithTagsByName(&iter, "monospace (typewriter)", "monospace")

	buffer.Insert(&iter, ", or ")
	buffer.InsertWithTagsByName(&iter, "big", "big")

	buffer.Insert(&iter, " text. ")
	buffer.Insert(&iter, "It's best not to hardcode specific text sizes; you can use relative "+
		"sizes as with CSS, such as ")
	buffer.InsertWithTagsByName(&iter, "xx-small", "xx-small")
	buffer.Insert(&iter, " or ")
	buffer.InsertWithTagsByName(&iter, "x-large", "x-large")

	buffer.Insert(&iter, " to ensure that your program properly adapts if the user changes the "+
		"default font size.\n\n")

	buffer.InsertWithTagsByName(&iter, "Colors. ", "heading")

	buffer.Insert(&iter, "Colors such as ")
	buffer.InsertWithTagsByName(&iter, "a blue foreground", "blue_foreground")
	buffer.Insert(&iter, " or ")
	buffer.InsertWithTagsByName(&iter, "a red background", "red_background")
	buffer.Insert(&iter, " or even ")
	buffer.InsertWithTagsByName(&iter, "a blue foreground on red background",
		"blue_foreground", "red_background")
	buffer.Insert(&iter, " (select that to read it) can be used.\n\n")

	buffer.InsertWithTagsByName(&iter, "Underline, strikethrough, and rise. ", "heading")
	buffer.InsertWithTagsByName(&iter, "Strikethrough", "strikethrough")
	buffer.Insert(&iter, ", ")
	buffer.InsertWithTagsByName(&iter, "underline", "underline")
	buffer.Insert(&iter, ", ")
	buffer.InsertWithTagsByName(&iter, "double underline", "double_underline")
	buffer.Insert(&iter, ", ")
	buffer.InsertWithTagsByName(&iter, "superscript", "superscript")
	buffer.Insert(&iter, ", ")
	buffer.InsertWithTagsByName(&iter, "subscript", "subscript")
	buffer.Insert(&iter, " are all supported.\n\n")

	buffer.InsertWithTagsByName(&iter, "Images. ", "heading")
	buffer.Insert(&iter, "The buffer can have images in it: ")
	buffer.InsertPixbuf(&iter, pixbuf)
	buffer.InsertPixbuf(&iter, pixbuf)
	buffer.InsertPixbuf(&iter, pixbuf)
	buffer.Insert(&iter, " for example.\n\n")

	buffer.InsertWithTagsByName(&iter, "Spacing. ", "heading")
	buffer.Insert(&iter, "You can adjust the amount of space before each line.\n")
	buffer.InsertWithTagsByName(&iter,
		"This line has a whole lot of space before it.\n", "big_gap_before_line", "wide_margins")
	buffer.InsertWithTagsByName(&iter,
		"You can also adjust the amount of space after each line; "+
			"this line has a whole lot of space after it.\n", "big_gap_after_line", "wide_margins")

	buffer.InsertWithTagsByName(&iter,
		"You can also adjust the amount of space between wrapped lines; "+
			"this line has extra space between each wrapped line in the same "+
			"paragraph. To show off wrapping, some filler text: the quick "+
			"brown fox jumped over the lazy dog. Blah blah blah blah blah "+
			"blah blah blah blah.\n", "double_spaced_line", "wide_margins")
	buffer.Insert(&iter, "Also note that those lines have extra-wide margins.\n\n")

	buffer.InsertWithTagsByName(&iter, "Editability. ", "heading")
	buffer.InsertWithTagsByName(&iter,
		"This line is 'locked down' and can't be edited by the user - just "+
			"try it! You can't delete this line.\n\n", "not_editable")

	buffer.InsertWithTagsByName(&iter, "Wraping ", "heading")
	buffer.Insert(&iter,
		"This line (and most of the others in this buffer) is word-wrapped, "+
			"using the proper Unicode algorithm. Word wrap should work in all "+
			"scripts and languages that GTK+ supports. Let's make this a long "+
			"paragraph to demonstrate: blah blah blah blah blah blah blah blah "+
			"blah blah blah blah blah blah blah blah blah blah blah\n\n")

	buffer.InsertWithTagsByName(&iter,
		"This line has character-based wrapping, and can wrap between any two "+
			"character glyphs. Let's make this a long paragraph to demonstrate: "+
			"blah blah blah blah blah blah blah blah blah blah blah blah blah blah "+
			"blah blah blah blah blah\n\n", "char_wrap")

	buffer.InsertWithTagsByName(&iter,
		"This line has all wrapping turned off, so it makes the horizontal "+
			"scrollbar appear.\n\n\n", "no_wrap")

	buffer.InsertWithTagsByName(&iter, "Justification. ", "heading")
	buffer.InsertWithTagsByName(&iter, "\nThis line has center justification.\n", "center")
	buffer.InsertWithTagsByName(&iter, "This line has right justification.\n", "right_justify")
	buffer.InsertWithTagsByName(&iter,
		"\nThis line has big wide margins. Text text text text text text text "+
			"text text text text text text text text text text text text text text "+
			"text text text text text text text text text text text text text text "+
			"text.\n", "wide_margins")

	buffer.InsertWithTagsByName(&iter, "Internationalization. ", "heading")
	buffer.Insert(&iter,
		"You can put all sorts of Unicode text in the buffer.\n\nGerman "+
			"(Deutsch S\303\274d) Gr\303\274\303\237 Gott\nGreek "+
			"(\316\225\316\273\316\273\316\267\316\275\316\271\316\272\316\254) "+
			"\316\223\316\265\316\271\316\254 \317\203\316\261\317\202\nHebrew      "+
			"\327\251\327\234\327\225\327\235\nJapanese "+
			"(\346\227\245\346\234\254\350\252\236)\n\nThe widget properly handles "+
			"bidirectional text, word wrapping, DOS/UNIX/Unicode paragraph separators, "+
			"grapheme boundaries, and so on using the Pango internationalization "+
			"framework.\n")
	buffer.Insert(&iter, "Here's a word-wrapped quaote in a right-to-left language:\n\n")
	buffer.InsertWithTagsByName(&iter,
		"\331\210\331\202\330\257 \330\250\330\257\330\243 "+
			"\330\253\331\204\330\247\330\253 \331\205\331\206 "+
			"\330\243\331\203\330\253\330\261 \330\247\331\204\331\205\330\244\330\263\330\263\330\247\330\252 "+
			"\330\252\331\202\330\257\331\205\330\247 \331\201\331\212 "+
			"\330\264\330\250\331\203\330\251 \330\247\331\203\330\263\331\212\331\210\331\206 "+
			"\330\250\330\261\330\247\331\205\330\254\331\207\330\247 "+
			"\331\203\331\205\331\206\330\270\331\205\330\247\330\252 "+
			"\331\204\330\247 \330\252\330\263\330\271\331\211 \331\204\331\204\330\261\330\250\330\255\330\214 "+
			"\330\253\331\205 \330\252\330\255\331\210\331\204\330\252 "+
			"\331\201\331\212 \330\247\331\204\330\263\331\206\331\210\330\247\330\252 "+
			"\330\247\331\204\330\256\331\205\330\263 \330\247\331\204\331\205\330\247\330\266\331\212\330\251 "+
			"\330\245\331\204\331\211 \331\205\330\244\330\263\330\263\330\247\330\252 "+
			"\331\205\330\247\331\204\331\212\330\251 \331\205\331\206\330\270\331\205\330\251\330\214 "+
			"\331\210\330\250\330\247\330\252\330\252 \330\254\330\262\330\241\330\247 "+
			"\331\205\331\206 \330\247\331\204\331\206\330\270\330\247\331\205 "+
			"\330\247\331\204\331\205\330\247\331\204\331\212 \331\201\331\212 "+
			"\330\250\331\204\330\257\330\247\331\206\331\207\330\247\330\214 "+
			"\331\210\331\204\331\203\331\206\331\207\330\247 \330\252\330\252\330\256\330\265\330\265 "+
			"\331\201\331\212 \330\256\330\257\331\205\330\251 \331\202\330\267\330\247\330\271 "+
			"\330\247\331\204\331\205\330\264\330\261\331\210\330\271\330\247\330\252 "+
			"\330\247\331\204\330\265\330\272\331\212\330\261\330\251. \331\210\330\243\330\255\330\257 "+
			"\330\243\331\203\330\253\330\261 \331\207\330\260\331\207 "+
			"\330\247\331\204\331\205\330\244\330\263\330\263\330\247\330\252 "+
			"\331\206\330\254\330\247\330\255\330\247 \331\207\331\210 "+
			"\302\273\330\250\330\247\331\206\331\203\331\210\330\263\331\210\331\204\302\253 "+
			"\331\201\331\212 \330\250\331\210\331\204\331\212\331\201\331\212\330\247.\n\n", "rtl_quote")

	buffer.Insert(&iter, "You can put widgets int the buffer: Here's a button: ")
	buffer.CreateChildAnchor(&iter)
	buffer.Insert(&iter, " and a menu: ")
	buffer.CreateChildAnchor(&iter)
	buffer.Insert(&iter, " and an animation: ")
	buffer.CreateChildAnchor(&iter)
	buffer.Insert(&iter, " finally a text entry: ")
	buffer.CreateChildAnchor(&iter)
	buffer.Insert(&iter, ".\n")

	buffer.Insert(&iter,
		"\n\nThis demo doesn't demonstrate all the GtkTextBuffer features; "+
			"it leaves out, for example: invisible/hidden text, tab stops, "+
			"application-drawn areas on the sides of the widget for displaying "+
			"breakpoints and such...")

	// Apply word_wrap tag to whole buffer
	buffer.GetBounds(&start, &end)
	buffer.ApplyTagByName("word_wrap", &start, &end)
}

func findAnchor(iter *gtk3.TextIter) bool {
	for iter.ForwardChar() {
		if iter.GetChildAnchor() != nil {
			return true
		}
	}
	return false
}

func attachWidgets(textView *gtk3.TextView) {
	var iter gtk3.TextIter

	buffer := textView.GetBuffer()
	buffer.GetStartIter(&iter)

	i := 0
	var widget gtk3.WidgetLike
	for findAnchor(&iter) {
		anchor := iter.GetChildAnchor()

		switch i {
		case 0:
			widget = gtk3.NewButtonWithLabel("Click Me")
			//widget.Connect("clicked",
		case 1:
			combo := gtk3.NewComboBoxText()
			combo.AppendText("Option 1")
			combo.AppendText("Option 2")
			combo.AppendText("Option 3")
			widget = combo
		case 2:
			filename, _ := findFile("floppybuddy.gif")
			widget = gtk3.NewImageFromFile(filename)
		case 3:
			widget = gtk3.NewEntry()
		default:
			break
		}
		textView.AddChildAtAnchor(widget, anchor)
		widget.W().ShowAll()
		i++
	}
}

func main() {
	gtk3.Init()

	window := gtk3.NewWindow(gtk3.GtkWindowType.TOPLEVEL)
	window.SetDefaultSize(450, 450)
	window.Connect("destroy", gtk3.MainQuit)
	window.SetTitle("TextView")
	window.SetBorderWidth(0)

	vpaned := gtk3.NewVPaned()
	vpaned.SetBorderWidth(5)
	window.Add(vpaned)

	view1 := gtk3.NewTextView()
	buffer := view1.GetBuffer()
	view2 := gtk3.NewTextViewWithBuffer(buffer)

	sw := gtk3.NewScrolledWindow(nil, nil)
	sw.SetPolicy(gtk3.GtkPolicy.AUTOMATIC, gtk3.GtkPolicy.AUTOMATIC)

	vpaned.Add1(sw)
	sw.Add(view1)

	sw = gtk3.NewScrolledWindow(nil, nil)
	sw.SetPolicy(gtk3.GtkPolicy.AUTOMATIC, gtk3.GtkPolicy.AUTOMATIC)

	vpaned.Add2(sw)
	sw.Add(view2)

	CreateTags(buffer)
	InsertText(buffer)

	attachWidgets(view1)
	attachWidgets(view2)

	window.ShowAll()

	gtk3.Main()
}
