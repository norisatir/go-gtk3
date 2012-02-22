go-gtk3 - GTK3 BINDINGS FOR GO
===================================

THIS IS:
--------

    Hopefully GTK3 Bindings using gobject binding for Go.

    Goal is to have fairly complete gtk3 widgets
    with pango, gdkpixbuf and cairo objects.

	
INSTALL:
--------

    go get github.com/norisatir/go-gtk3/gtk3


  For demo:
  
    cd $GOPATH/src/github.com/norisatir/go-gtk3/demo
    go build
    ./demo
  
    (Demo apps are taken from gtk3 demo apps, translated to go)
LICENSE:
--------

	Under same terms and conditions as the Go language,
	BSD style license


IMPLEMENTATION LIST:
--------------------
##### Note:
	"Gtk" has been stripped from the names.
	
#### GTK3 CORE: 
- AccelGroups
- Clipboard

#### BASE OBJECT:
- Widget
- Container
- Bin
- Range

#### WINDOWS: 
- Window
- Dialog
- MessageDialog
- Invisible
- Assistant

#### DISPLAY WIDGETS:
- Label
- ProgressBar
- Image
- Statusbar
- Spinner
- InfoBar

#### BUTTONS AND TOGGLES:
- Button
- ToggleButton
- CheckButton
- RadioButton
- LinkButton
- Switch

#### NUMERIC/TEXT DATA ENTRY:
- Entry
- EntryBuffer
- EntryCompletion
- Scale
- SpinButton

#### MULTILINE TEXT EDITOR:
- TextTag
- TextTagTable
- TextIter
- TextBuffer
- TextChildAnchor
- TextView

#### TREE, LIST AND ICON GRID WIDGETS:
- TreePath
- TreeRowReference
- TreeIter
- TreeModel
- ListStore
- TreeStore
- CellArea
- CellAreaContext
- CellLayout
- CellView
- CellRenderer (and brothers)
- TreeViewColumn
- TreeSelection
- TreeView

#### MENUS, COMBOBOX, TOOLBAR:
- ComboBox
- ComboBoxText
- MenuShell
- Menu
- MenuBar
- MenuItem (and brothers)

#### SELECTORS (COLOR,FILE,FONT):
- ColorButton

#### ACTION-BASED MENUS AND TOOLBARS
- Action

#### LAYOUT CONTAINERS:
- Grid
- Box
- ButtonBox
- Paned
- Notebook

#### ORNAMENTS:
- Frame
- Separator

#### SCROLLING:
- ScrollBar
- ScrolledWindow

#### APP:
- Application
</p>

AUTHOR
-------
- Matej Knific
