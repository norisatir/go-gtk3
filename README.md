go-gtk3 - GTK3 BINDINGS FOR GO
===================================

THIS IS:
--------

    Hopefully GTK3 Bindings using gobject binding for Go.

    Goal is to have fairly complete gtk3 widgets
    with pango, gdkpixbuf and cairo objects.

	
INSTALL:
--------

    git clone https://github.com/norisatir/go-gtk3
    cd go-gtk3
    gomake install
	
  or
  
    goinstall github.com/norisatir/go-gtk3/gtk3


  For demo:
  
    gomake demo
    cd demo_apps
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

#### BUTTONS AND TOGGLES:
- Button
- ToggleButton
- CheckButton

#### NUMERIC/TEXT DATA ENTRY:
- Entry
- EntryBuffer
- EntryCompletion

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

AUTHOR
------

- Matej Knific
