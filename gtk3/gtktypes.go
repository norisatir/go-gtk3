// Types of all gtk3 objects
package gtk3

/*
#include <gtk/gtk.h>
#include <gtk/gtkx.h>

*/
import "C"

import g "github.com/norisatir/go-gtk3/gobject"

var GtkType gtkTypes

type gtkTypes struct {
	ABOUT_DIALOG              g.GType
	ACCEL_GROUP               g.GType
	ACCEL_LABEL               g.GType
	ACCEL_MAP                 g.GType
	ACCESSIBLE                g.GType
	ACTION_GROUP              g.GType
	ACTION                    g.GType
	ACTIVATABLE               g.GType
	ADJUSTMENT                g.GType
	ALIGNMENT                 g.GType
	APP_CHOOSER_BUTTON        g.GType
	APP_CHOOSER_DIALOG        g.GType
	APP_CHOOSER               g.GType
	APP_CHOOSER_WIDGET        g.GType
	APPLICATION               g.GType
	ARROW                     g.GType
	ASPECT_FRAME              g.GType
	ASSISTANT                 g.GType
	BUTTON_BOX                g.GType
	BIN                       g.GType
	BORDER                    g.GType
	BOX                       g.GType
	BUILDABLE                 g.GType
	BUILDER                   g.GType
	BUTTON                    g.GType
	CALENDAR                  g.GType
	CELL_AREA_BOX             g.GType
	CELL_AREA_CONTEXT         g.GType
	CELL_AREA                 g.GType
	CELL_EDITABLE             g.GType
	CELL_LAYOUT               g.GType
	CELL_RENDERER_ACCEL       g.GType
	CELL_RENDERER_COMBO       g.GType
	CELL_RENDERER             g.GType
	CELL_RENDERER_PIXBUF      g.GType
	CELL_RENDERER_PROGRESS    g.GType
	CELL_RENDERER_SPIN        g.GType
	CELL_RENDERER_SPINNER     g.GType
	CELL_RENDERER_TEXT        g.GType
	CELL_RENDERER_TOGGLE      g.GType
	CELL_VIEW                 g.GType
	CHECK_BUTTON              g.GType
	CHECK_MENU_ITEM           g.GType
	CLIPBOARD                 g.GType
	COLOR_BUTTON              g.GType
	COLOR_SELECTION_DIALOG    g.GType
	COLOR_SELECTION           g.GType
	COMBO_BOX                 g.GType
	COMBO_BOX_TEXT            g.GType
	CONTAINER                 g.GType
	CSS_PROVIDER              g.GType
	DIALOG                    g.GType
	DRAWING_AREA              g.GType
	EDITABLE                  g.GType
	ENTRY_BUFFER              g.GType
	ENTRY_COMPLETION          g.GType
	ENTRY                     g.GType
	EVENT_BOX                 g.GType
	EXPANDER                  g.GType
	FILE_CHOOSER_BUTTON       g.GType
	FILE_CHOOSER_DIALOG       g.GType
	FILE_CHOOSER              g.GType
	FILE_CHOOSER_WIDGET       g.GType
	FILE_FILTER               g.GType
	FIXED                     g.GType
	FONT_BUTTON               g.GType
	FONT_SELECTION            g.GType
	FONT_SELECTION_DIALOG     g.GType
	FRAME                     g.GType
	GRADIENT                  g.GType
	GRID                      g.GType
	HANDLE_BOX                g.GType
	HBUTTON_BOX               g.GType
	HBOX                      g.GType
	HPANED                    g.GType
	HSCALE                    g.GType
	HSCROLLBAR                g.GType
	HSEPARATOR                g.GType
	HSV                       g.GType
	ICON_FACTORY              g.GType
	ICON_SET                  g.GType
	ICON_SOURCE               g.GType
	ICON_INFO                 g.GType
	ICON_THEME                g.GType
	ICON_VIEW                 g.GType
	IMAGE                     g.GType
	IMAGE_MENU_ITEM           g.GType
	IM_CONTEXT                g.GType
	IM_CONTEXT_SIMPLE         g.GType
	IM_MULTICONTEXT           g.GType
	INFO_BAR                  g.GType
	INVISIBLE                 g.GType
	LABEL                     g.GType
	LAYOUT                    g.GType
	LINK_BUTTON               g.GType
	LIST_STORE                g.GType
	MENU_BAR                  g.GType
	MENU                      g.GType
	MENU_ITEM                 g.GType
	MENU_SHELL                g.GType
	MENU_TOOL_BUTTON          g.GType
	MESSAGE_DIALOG            g.GType
	MISC                      g.GType
	MOUNT_OPERATION           g.GType
	NOTEBOOK                  g.GType
	NUMERABLE_ICON            g.GType
	OFFSCREEN_WINDOW          g.GType
	ORIENTABLE                g.GType
	PAGE_SETUP                g.GType
	PANED                     g.GType
	PAPER_SIZE                g.GType
	PLUG                      g.GType
	PRINT_CONTEXT             g.GType
	PRINT_OPERATION           g.GType
	PRINT_OPERATION_PREVIEW   g.GType
	PRINT_SETTINGS            g.GType
	PROGRESS_BAR              g.GType
	RADIO_ACTION              g.GType
	RADIO_BUTTON              g.GType
	RADIO_MENU_ITEM           g.GType
	RADIO_TOOL_BUTTON         g.GType
	RANGE                     g.GType
	RC_STYLE                  g.GType
	RECENT_ACTION             g.GType
	RECENT_CHOOSER_DIALOG     g.GType
	RECENT_CHOOSER            g.GType
	RECENT_CHOOSER_MENU       g.GType
	RECENT_CHOOSER_WIDGET     g.GType
	RECENT_FILTER             g.GType
	RECENT_INFO               g.GType
	RECENT_MANAGER            g.GType
	SCALE_BUTTON              g.GType
	SCALE                     g.GType
	SCROLLABLE                g.GType
	SCROLLBAR                 g.GType
	SCROLLED_WINDOW           g.GType
	SELECTION_DATA            g.GType
	TARGET_LIST               g.GType
	SEPARATOR                 g.GType
	SEPARATOR_MENU_ITEM       g.GType
	SEPARATOR_TOOL_ITEM       g.GType
	SETTINGS                  g.GType
	SIZE_GROUP                g.GType
	SOCKET                    g.GType
	SPIN_BUTTON               g.GType
	SPINNER                   g.GType
	STATUSBAR                 g.GType
	STATUS_ICON               g.GType
	STYLE_CONTEXT             g.GType
	STYLE                     g.GType
	STYLE_PROPERTIES          g.GType
	STYLE_PROVIDER            g.GType
	SWITCH                    g.GType
	SYMBOLIC_COLOR            g.GType
	TABLE                     g.GType
	TEAROFF_MENU_ITEM         g.GType
	TEXT_ATTRIBUTES           g.GType
	TEXT_BUFFER               g.GType
	TEXT_CHILD_ANCHOR         g.GType
	TEXT_ITER                 g.GType
	TEXT_MARK                 g.GType
	TEXT_TAG                  g.GType
	TEXT_TAG_TABLE            g.GType
	TEXT_VIEW                 g.GType
	THEMING_ENGINE            g.GType
	TOGGLE_ACTION             g.GType
	TOGGLE_BUTTON             g.GType
	TOGGLE_TOOL_BUTTON        g.GType
	TOOLBAR                   g.GType
	TOOL_BUTTON               g.GType
	TOOL_ITEM_GROUP           g.GType
	TOOL_ITEM                 g.GType
	TOOL_PALETTE              g.GType
	TOOL_SHELL                g.GType
	TOOLTIP                   g.GType
	TREE_DRAG_SOURCE          g.GType
	TREE_DRAG_DEST            g.GType
	TREE_MODEL_FILTER         g.GType
	TREE_MODEL                g.GType
	TREE_ITER                 g.GType
	TREE_PATH                 g.GType
	TREE_ROW_REFERENCE        g.GType
	TREE_MODEL_SORT           g.GType
	TREE_SELECTION            g.GType
	TREE_SORTABLE             g.GType
	TREE_STORE                g.GType
	TREE_VIEW_COLUMN          g.GType
	TREE_VIEW                 g.GType
	UI_MANAGER                g.GType
	VBUTTON_BOX               g.GType
	VBOX                      g.GType
	VIEWPORT                  g.GType
	VOLUME_BUTTON             g.GType
	VPANED                    g.GType
	VSCALE                    g.GType
	VSCROLLBAR                g.GType
	VSEPARATOR                g.GType
	WIDGET                    g.GType
	REQUISITION               g.GType
	WIDGET_PATH               g.GType
	WINDOW                    g.GType
	WINDOW_GROUP              g.GType
	LICENSE                   g.GType
	ACCEL_FLAGS               g.GType
	ASSISTANT_PAGE_TYPE       g.GType
	BUILDER_ERROR             g.GType
	CALENDAR_DISPLAY_OPTIONS  g.GType
	CELL_RENDERER_STATE       g.GType
	CELL_RENDERER_MODE        g.GType
	CELL_RENDERER_ACCEL_MODE  g.GType
	CSS_PROVIDER_ERROR        g.GType
	DEBUG_FLAG                g.GType
	DIALOG_FLAGS              g.GType
	RESPONSE_TYPE             g.GType
	DEST_DEFAULTS             g.GType
	TARGET_FLAGS              g.GType
	ENTRY_ICON_POSITION       g.GType
	ALIGN                     g.GType
	ARROW_PLACEMENT           g.GType
	ARROW_TYPE                g.GType
	ATTACH_OPTIONS            g.GType
	BUTTON_BOX_STYLE          g.GType
	DELETE_TYPE               g.GType
	DIRECTION_TYPE            g.GType
	EXPANDER_STYLE            g.GType
	ICON_SIZE                 g.GType
	SENSITIVITY_TYPE          g.GType
	TEXT_DIRECTION            g.GType
	JUSTIFICATION             g.GType
	MENU_DIRECTION_TYPE       g.GType
	MESSAGE_TYPE              g.GType
	MOVEMENT_STEP             g.GType
	SCROLL_STEP               g.GType
	ORIENTATION               g.GType
	CORNER_TYPE               g.GType
	PACK_TYPE                 g.GType
	PATH_PRIORITY_TYPE        g.GType
	PATH_TYPE                 g.GType
	POLICY_TYPE               g.GType
	POSITION_TYPE             g.GType
	RELIEF_STYLE              g.GType
	RESIZE_MODE               g.GType
	SCROLL_TYPE               g.GType
	SELECTION_MODE            g.GType
	SHADOW_TYPE               g.GType
	STATE_TYPE                g.GType
	TOOLBAR_STYLE             g.GType
	WINDOW_POSITION           g.GType
	WINDOW_TYPE               g.GType
	WRAP_MODE                 g.GType
	SORT_TYPE                 g.GType
	IM_PREEDIT_STYLE          g.GType
	IM_STATUS_STYLE           g.GType
	PACK_DIRECTION            g.GType
	PRINT_PAGES               g.GType
	PAGE_SET                  g.GType
	NUMBER_UP_LAYOUT          g.GType
	PAGE_ORIENTATION          g.GType
	PRINT_QUALITY             g.GType
	PRINT_DUPLEX              g.GType
	UNIT                      g.GType
	TREE_VIEW_GRID_LINES      g.GType
	DRAG_RESULT               g.GType
	SIZE_REQUEST_MODE         g.GType
	SCROLLABLE_POLICY         g.GType
	STATE_FLAGS               g.GType
	REGION_FLAGS              g.GType
	JUNCTION_SIDES            g.GType
	BORDER_STYLE              g.GType
	FILE_CHOOSER_ACTION       g.GType
	FILE_CHOOSER_CONFIRMATION g.GType
	FILE_CHOOSER_ERROR        g.GType
	FILE_FILTER_FLAGS         g.GType
	ICON_LOOKUP_FLAGS         g.GType
	ICON_THEME_ERROR          g.GType
	ICON_VIEW_DROP_POSITION   g.GType
	IMAGE_TYPE                g.GType
	BUTTONS_TYPE              g.GType
	NOTEBOOK_TAB              g.GType
	PRINT_STATUS              g.GType
	PRINT_OPERATION_RESULT    g.GType
	PRINT_OPERATION_ACTION    g.GType
	PRINT_ERROR               g.GType
	RC_FLAGS                  g.GType
	RC_TOKEN_TYPE             g.GType
	RECENT_SORT_TYPE          g.GType
	RECENT_CHOOSER_ERROR      g.GType
	RECENT_FILTER_FLAGS       g.GType
	RECENT_MANAGER_ERROR      g.GType
	SIZE_GROUP_MODE           g.GType
	SPIN_BUTTON_UPDATE_POLICY g.GType
	SPIN_TYPE                 g.GType
	TEXT_BUFFER_TARGET_INFO   g.GType
	TEXT_SEARCH_FLAGS         g.GType
	TEXT_WINDOW_TYPE          g.GType
	TOOLBAR_SPACE_STYLE       g.GType
	TOOL_PALETTE_DRAG_TARGETS g.GType
	TREE_MODEL_FLAGS          g.GType
	TREE_VIEW_DROP_POSITION   g.GType
	TREE_VIEW_COLUMN_SIZING   g.GType
	UI_MANAGER_ITEM_TYPE      g.GType
	WIDGET_HELP_TYPE          g.GType
}

// GtkShadowType
var GtkShadow gtkShadow

type gtkShadow struct {
	NONE       int
	IN         int
	OUT        int
	ETCHED_IN  int
	ETCHED_OUT int
}
// End GtkShadow

// string
var GtkStock gtkStock

type gtkStock struct {
	ABOUT                         string
	ADD                           string
	APPLY                         string
	BOLD                          string
	CANCEL                        string
	CAPS_LOCK_WARNING             string
	CDROM                         string
	CLEAR                         string
	CLOSE                         string
	COLOR_PICKER                  string
	CONNECT                       string
	CONVERT                       string
	COPY                          string
	CUT                           string
	DELETE                        string
	DIALOG_AUTHENTICATION         string
	DIALOG_INFO                   string
	DIALOG_WARNING                string
	DIALOG_ERROR                  string
	DIALOG_QUESTION               string
	DIRECTORY                     string
	DISCARD                       string
	DISCONNECT                    string
	DND                           string
	DND_MULTIPLE                  string
	EDIT                          string
	EXECUTE                       string
	FILE                          string
	FIND                          string
	FIND_AND_REPLACE              string
	FLOPPY                        string
	FULLSCREEN                    string
	GOTO_BOTTOM                   string
	GOTO_FIRST                    string
	GOTO_LAST                     string
	GOTO_TOP                      string
	GO_BACK                       string
	GO_DOWN                       string
	GO_FORWARD                    string
	GO_UP                         string
	HARDDISK                      string
	HELP                          string
	HOME                          string
	INDEX                         string
	INDENT                        string
	INFO                          string
	ITALIC                        string
	JUMP_TO                       string
	JUSTIFY_CENTER                string
	JUSTIFY_FILL                  string
	JUSTIFY_LEFT                  string
	JUSTIFY_RIGHT                 string
	LEAVE_FULLSCREEN              string
	MISSING_IMAGE                 string
	MEDIA_FORWARD                 string
	MEDIA_NEXT                    string
	MEDIA_PAUSE                   string
	MEDIA_PLAY                    string
	MEDIA_PREVIOUS                string
	MEDIA_RECORD                  string
	MEDIA_REWIND                  string
	MEDIA_STOP                    string
	NETWORK                       string
	NEW                           string
	NO                            string
	OK                            string
	OPEN                          string
	ORIENTATION_PORTRAIT          string
	ORIENTATION_LANDSCAPE         string
	ORIENTATION_REVERSE_LANDSCAPE string
	ORIENTATION_REVERSE_PORTRAIT  string
	PAGE_SETUP                    string
	PASTE                         string
	PREFERENCES                   string
	PRINT                         string
	PRINT_ERROR                   string
	PRINT_PAUSED                  string
	PRINT_PREVIEW                 string
	PRINT_REPORT                  string
	PRINT_WARNING                 string
	PROPERTIES                    string
	QUIT                          string
	REDO                          string
	REFRESH                       string
	REMOVE                        string
	REVERT_TO_SAVED               string
	SAVE                          string
	SAVE_AS                       string
	SELECT_ALL                    string
	SELECT_COLOR                  string
	SELECT_FONT                   string
	SORT_ASCENDING                string
	SORT_DESCENDING               string
	SPELL_CHECK                   string
	STOP                          string
	STRIKETHROUGH                 string
	UNDELETE                      string
	UNDERLINE                     string
	UNDO                          string
	UNINDENT                      string
	YES                           string
	ZOOM_100                      string
	ZOOM_FIT                      string
	ZOOM_IN                       string
	ZOOM_OUT                      string
}
// End GtkStock

// GtkPositionType
var GtkPosition gtkPosition

type gtkPosition struct {
	LEFT   int
	RIGHT  int
	TOP    int
	BOTTOM int
}
// End GtkPositionType

// GtkDialogFlags
var GtkDialogFlags gtkDialogFlags

type gtkDialogFlags struct {
	MODAL               int
	DESTROY_WITH_PARENT int
}
// End GtkDialogFlags

// GtkResponse
var GtkResponse gtkResponse

type gtkResponse struct {
	NONE         int
	REJECT       int
	ACCEPT       int
	DELETE_EVENT int
	OK           int
	CANCEL       int
	CLOSE        int
	YES          int
	NO           int
	APPLY        int
	HELP         int
}
// End GtkResponse

// GtkMessageType
var GtkMessage gtkMessage

type gtkMessage struct {
	INFO     int
	WARNING  int
	QUESTION int
	ERROR    int
	OTHER    int
}
// End GtkMessageType

// GtkButtons
var GtkButtons gtkButtons

type gtkButtons struct {
	NONE      int
	OK        int
	CLOSE     int
	CANCEL    int
	YES_NO    int
	OK_CANCEL int
}
// End GtkButtons

// GtkOrientation
var GtkOrientation gtkOrientation

type gtkOrientation struct {
	HORIZONTAL int
	VERTICAL   int
}
// End GtkOrientation

// GtkPackType
var GtkPack gtkPack

type gtkPack struct {
	PACK_START int
	PACK_END   int
}
// End GtkPackType

// GtkButtonBoxStyle
var GtkButtonBoxStyle gtkButtonBoxStyle

type gtkButtonBoxStyle struct {
	SPREAD int
	EDGE   int
	START  int
	END    int
	CENTER int
}
// End GtkButtonBoxStyle

// GtkIconSize
var GtkIconSize gtkIconSize

type gtkIconSize struct {
	INVALID       int
	MENU          int
	SMALL_TOOLBAR int
	LARGE_TOOLBAR int
	BUTTON        int
	DND           int
	DIALOG        int
}
// End GtkIconSize

// GApplicatonFlags
var GApplicationFlags gApplicationFlags

type gApplicationFlags struct {
	FLAGS_NONE           int
	IS_SERVICE           int
	IS_LAUNCHER          int
	HANDLES_OPEN         int
	HANDLES_COMMAND_LINE int
	SEND_ENVIRONMENT     int
	NON_UNIQUE           int
}
// End GApplicationFlags

// GtkJustification
var GtkJustification gtkJustification

type gtkJustification struct {
	LEFT   int
	RIGHT  int
	CENTER int
	FILL   int
}
// End GtkJustification

// GtkWindowType
var GtkWindowType gtkWindowType

type gtkWindowType struct {
	TOPLEVEL int
	POPUP    int
}
// End GtkWindowType

// GtkSensitivityType
var GtkSensitivity gtkSensitivity

type gtkSensitivity struct {
	AUTO int
	ON   int
	OFF  int
}
// End GtkSensitivity

// GtkPolicyType
var GtkPolicy gtkPolicy

type gtkPolicy struct {
	ALWAYS    int
	AUTOMATIC int
	NEVER     int
}
// End GtkPolicyType

// GtkCornerType
var GtkCorner gtkCorner

type gtkCorner struct {
	TOP_LEFT     int
	BOTTOM_LEFT  int
	TOP_RIGHT    int
	BOTTOM_RIGHT int
}

func init() {
	GtkType.ABOUT_DIALOG = g.GType(C.gtk_about_dialog_get_type())
	GtkType.ACCEL_GROUP = g.GType(C.gtk_accel_group_get_type())
	GtkType.ACCEL_LABEL = g.GType(C.gtk_accel_label_get_type())
	GtkType.ACCEL_MAP = g.GType(C.gtk_accel_map_get_type())
	GtkType.ACCESSIBLE = g.GType(C.gtk_accessible_get_type())
	GtkType.ACTION_GROUP = g.GType(C.gtk_action_group_get_type())
	GtkType.ACTION = g.GType(C.gtk_action_get_type())
	GtkType.ACTIVATABLE = g.GType(C.gtk_activatable_get_type())
	GtkType.ADJUSTMENT = g.GType(C.gtk_adjustment_get_type())
	GtkType.ALIGNMENT = g.GType(C.gtk_alignment_get_type())
	GtkType.APP_CHOOSER_BUTTON = g.GType(C.gtk_app_chooser_button_get_type())
	GtkType.APP_CHOOSER_DIALOG = g.GType(C.gtk_app_chooser_dialog_get_type())
	GtkType.APP_CHOOSER = g.GType(C.gtk_app_chooser_get_type())
	GtkType.APP_CHOOSER_WIDGET = g.GType(C.gtk_app_chooser_widget_get_type())
	GtkType.APPLICATION = g.GType(C.gtk_application_get_type())
	GtkType.ARROW = g.GType(C.gtk_arrow_get_type())
	GtkType.ASPECT_FRAME = g.GType(C.gtk_aspect_frame_get_type())
	GtkType.ASSISTANT = g.GType(C.gtk_assistant_get_type())
	GtkType.BUTTON_BOX = g.GType(C.gtk_button_box_get_type())
	GtkType.BIN = g.GType(C.gtk_bin_get_type())
	GtkType.BORDER = g.GType(C.gtk_border_get_type())
	GtkType.BOX = g.GType(C.gtk_box_get_type())
	GtkType.BUILDABLE = g.GType(C.gtk_buildable_get_type())
	GtkType.BUILDER = g.GType(C.gtk_builder_get_type())
	GtkType.BUTTON = g.GType(C.gtk_button_get_type())
	GtkType.CALENDAR = g.GType(C.gtk_calendar_get_type())
	GtkType.CELL_AREA_BOX = g.GType(C.gtk_cell_area_box_get_type())
	GtkType.CELL_AREA_CONTEXT = g.GType(C.gtk_cell_area_context_get_type())
	GtkType.CELL_AREA = g.GType(C.gtk_cell_area_get_type())
	GtkType.CELL_EDITABLE = g.GType(C.gtk_cell_editable_get_type())
	GtkType.CELL_LAYOUT = g.GType(C.gtk_cell_layout_get_type())
	GtkType.CELL_RENDERER_ACCEL = g.GType(C.gtk_cell_renderer_accel_get_type())
	GtkType.CELL_RENDERER_COMBO = g.GType(C.gtk_cell_renderer_combo_get_type())
	GtkType.CELL_RENDERER = g.GType(C.gtk_cell_renderer_get_type())
	GtkType.CELL_RENDERER_PIXBUF = g.GType(C.gtk_cell_renderer_pixbuf_get_type())
	GtkType.CELL_RENDERER_PROGRESS = g.GType(C.gtk_cell_renderer_progress_get_type())
	GtkType.CELL_RENDERER_SPIN = g.GType(C.gtk_cell_renderer_spin_get_type())
	GtkType.CELL_RENDERER_SPINNER = g.GType(C.gtk_cell_renderer_spinner_get_type())
	GtkType.CELL_RENDERER_TEXT = g.GType(C.gtk_cell_renderer_text_get_type())
	GtkType.CELL_RENDERER_TOGGLE = g.GType(C.gtk_cell_renderer_toggle_get_type())
	GtkType.CELL_VIEW = g.GType(C.gtk_cell_view_get_type())
	GtkType.CHECK_BUTTON = g.GType(C.gtk_check_button_get_type())
	GtkType.CHECK_MENU_ITEM = g.GType(C.gtk_check_menu_item_get_type())
	GtkType.CLIPBOARD = g.GType(C.gtk_clipboard_get_type())
	GtkType.COLOR_BUTTON = g.GType(C.gtk_color_button_get_type())
	GtkType.COLOR_SELECTION_DIALOG = g.GType(C.gtk_color_selection_dialog_get_type())
	GtkType.COLOR_SELECTION = g.GType(C.gtk_color_selection_get_type())
	GtkType.COMBO_BOX = g.GType(C.gtk_combo_box_get_type())
	GtkType.COMBO_BOX_TEXT = g.GType(C.gtk_combo_box_text_get_type())
	GtkType.CONTAINER = g.GType(C.gtk_container_get_type())
	GtkType.CSS_PROVIDER = g.GType(C.gtk_css_provider_get_type())
	GtkType.DIALOG = g.GType(C.gtk_dialog_get_type())
	GtkType.DRAWING_AREA = g.GType(C.gtk_drawing_area_get_type())
	GtkType.EDITABLE = g.GType(C.gtk_editable_get_type())
	GtkType.ENTRY_BUFFER = g.GType(C.gtk_entry_buffer_get_type())
	GtkType.ENTRY_COMPLETION = g.GType(C.gtk_entry_completion_get_type())
	GtkType.ENTRY = g.GType(C.gtk_entry_get_type())
	GtkType.EVENT_BOX = g.GType(C.gtk_event_box_get_type())
	GtkType.EXPANDER = g.GType(C.gtk_expander_get_type())
	GtkType.FILE_CHOOSER_BUTTON = g.GType(C.gtk_file_chooser_button_get_type())
	GtkType.FILE_CHOOSER_DIALOG = g.GType(C.gtk_file_chooser_dialog_get_type())
	GtkType.FILE_CHOOSER = g.GType(C.gtk_file_chooser_get_type())
	GtkType.FILE_CHOOSER_WIDGET = g.GType(C.gtk_file_chooser_widget_get_type())
	GtkType.FILE_FILTER = g.GType(C.gtk_file_filter_get_type())
	GtkType.FIXED = g.GType(C.gtk_fixed_get_type())
	GtkType.FONT_BUTTON = g.GType(C.gtk_font_button_get_type())
	GtkType.FONT_SELECTION = g.GType(C.gtk_font_selection_get_type())
	GtkType.FONT_SELECTION_DIALOG = g.GType(C.gtk_font_selection_dialog_get_type())
	GtkType.FRAME = g.GType(C.gtk_frame_get_type())
	GtkType.GRADIENT = g.GType(C.gtk_gradient_get_type())
	GtkType.GRID = g.GType(C.gtk_grid_get_type())
	GtkType.HANDLE_BOX = g.GType(C.gtk_handle_box_get_type())
	GtkType.HBUTTON_BOX = g.GType(C.gtk_hbutton_box_get_type())
	GtkType.HBOX = g.GType(C.gtk_hbox_get_type())
	GtkType.HPANED = g.GType(C.gtk_hpaned_get_type())
	GtkType.HSCALE = g.GType(C.gtk_hscale_get_type())
	GtkType.HSCROLLBAR = g.GType(C.gtk_hscrollbar_get_type())
	GtkType.HSEPARATOR = g.GType(C.gtk_hseparator_get_type())
	GtkType.HSV = g.GType(C.gtk_hsv_get_type())
	GtkType.ICON_FACTORY = g.GType(C.gtk_icon_factory_get_type())
	GtkType.ICON_SET = g.GType(C.gtk_icon_set_get_type())
	GtkType.ICON_SOURCE = g.GType(C.gtk_icon_source_get_type())
	GtkType.ICON_INFO = g.GType(C.gtk_icon_info_get_type())
	GtkType.ICON_THEME = g.GType(C.gtk_icon_theme_get_type())
	GtkType.ICON_VIEW = g.GType(C.gtk_icon_view_get_type())
	GtkType.IMAGE = g.GType(C.gtk_image_get_type())
	GtkType.IMAGE_MENU_ITEM = g.GType(C.gtk_image_menu_item_get_type())
	GtkType.IM_CONTEXT = g.GType(C.gtk_im_context_get_type())
	GtkType.IM_CONTEXT_SIMPLE = g.GType(C.gtk_im_context_simple_get_type())
	GtkType.IM_MULTICONTEXT = g.GType(C.gtk_im_multicontext_get_type())
	GtkType.INFO_BAR = g.GType(C.gtk_info_bar_get_type())
	GtkType.INVISIBLE = g.GType(C.gtk_invisible_get_type())
	GtkType.LABEL = g.GType(C.gtk_label_get_type())
	GtkType.LAYOUT = g.GType(C.gtk_layout_get_type())
	GtkType.LINK_BUTTON = g.GType(C.gtk_link_button_get_type())
	GtkType.LIST_STORE = g.GType(C.gtk_list_store_get_type())
	GtkType.MENU_BAR = g.GType(C.gtk_menu_bar_get_type())
	GtkType.MENU = g.GType(C.gtk_menu_get_type())
	GtkType.MENU_ITEM = g.GType(C.gtk_menu_item_get_type())
	GtkType.MENU_SHELL = g.GType(C.gtk_menu_shell_get_type())
	GtkType.MENU_TOOL_BUTTON = g.GType(C.gtk_menu_tool_button_get_type())
	GtkType.MESSAGE_DIALOG = g.GType(C.gtk_message_dialog_get_type())
	GtkType.MISC = g.GType(C.gtk_misc_get_type())
	GtkType.MOUNT_OPERATION = g.GType(C.gtk_mount_operation_get_type())
	GtkType.NOTEBOOK = g.GType(C.gtk_notebook_get_type())
	GtkType.NUMERABLE_ICON = g.GType(C.gtk_numerable_icon_get_type())
	GtkType.OFFSCREEN_WINDOW = g.GType(C.gtk_offscreen_window_get_type())
	GtkType.ORIENTABLE = g.GType(C.gtk_orientable_get_type())
	GtkType.PAGE_SETUP = g.GType(C.gtk_page_setup_get_type())
	GtkType.PANED = g.GType(C.gtk_paned_get_type())
	GtkType.PAPER_SIZE = g.GType(C.gtk_paper_size_get_type())
	GtkType.PLUG = g.GType(C.gtk_plug_get_type())
	GtkType.PRINT_CONTEXT = g.GType(C.gtk_print_context_get_type())
	GtkType.PRINT_OPERATION = g.GType(C.gtk_print_operation_get_type())
	GtkType.PRINT_OPERATION_PREVIEW = g.GType(C.gtk_print_operation_preview_get_type())
	GtkType.PRINT_SETTINGS = g.GType(C.gtk_print_settings_get_type())
	GtkType.PROGRESS_BAR = g.GType(C.gtk_progress_bar_get_type())
	GtkType.RADIO_ACTION = g.GType(C.gtk_radio_action_get_type())
	GtkType.RADIO_BUTTON = g.GType(C.gtk_radio_button_get_type())
	GtkType.RADIO_MENU_ITEM = g.GType(C.gtk_radio_menu_item_get_type())
	GtkType.RADIO_TOOL_BUTTON = g.GType(C.gtk_radio_tool_button_get_type())
	GtkType.RANGE = g.GType(C.gtk_range_get_type())
	GtkType.RC_STYLE = g.GType(C.gtk_rc_style_get_type())
	GtkType.RECENT_ACTION = g.GType(C.gtk_recent_action_get_type())
	GtkType.RECENT_CHOOSER_DIALOG = g.GType(C.gtk_recent_chooser_dialog_get_type())
	GtkType.RECENT_CHOOSER = g.GType(C.gtk_recent_chooser_get_type())
	GtkType.RECENT_CHOOSER_MENU = g.GType(C.gtk_recent_chooser_menu_get_type())
	GtkType.RECENT_CHOOSER_WIDGET = g.GType(C.gtk_recent_chooser_widget_get_type())
	GtkType.RECENT_FILTER = g.GType(C.gtk_recent_filter_get_type())
	GtkType.RECENT_INFO = g.GType(C.gtk_recent_info_get_type())
	GtkType.RECENT_MANAGER = g.GType(C.gtk_recent_manager_get_type())
	GtkType.SCALE_BUTTON = g.GType(C.gtk_scale_button_get_type())
	GtkType.SCALE = g.GType(C.gtk_scale_get_type())
	GtkType.SCROLLABLE = g.GType(C.gtk_scrollable_get_type())
	GtkType.SCROLLBAR = g.GType(C.gtk_scrollbar_get_type())
	GtkType.SCROLLED_WINDOW = g.GType(C.gtk_scrolled_window_get_type())
	GtkType.SELECTION_DATA = g.GType(C.gtk_selection_data_get_type())
	GtkType.TARGET_LIST = g.GType(C.gtk_target_list_get_type())
	GtkType.SEPARATOR = g.GType(C.gtk_separator_get_type())
	GtkType.SEPARATOR_MENU_ITEM = g.GType(C.gtk_separator_menu_item_get_type())
	GtkType.SEPARATOR_TOOL_ITEM = g.GType(C.gtk_separator_tool_item_get_type())
	GtkType.SETTINGS = g.GType(C.gtk_settings_get_type())
	GtkType.SIZE_GROUP = g.GType(C.gtk_size_group_get_type())
	GtkType.SOCKET = g.GType(C.gtk_socket_get_type())
	GtkType.SPIN_BUTTON = g.GType(C.gtk_spin_button_get_type())
	GtkType.SPINNER = g.GType(C.gtk_spinner_get_type())
	GtkType.STATUSBAR = g.GType(C.gtk_statusbar_get_type())
	GtkType.STATUS_ICON = g.GType(C.gtk_status_icon_get_type())
	GtkType.STYLE_CONTEXT = g.GType(C.gtk_style_context_get_type())
	GtkType.STYLE = g.GType(C.gtk_style_get_type())
	GtkType.STYLE_PROPERTIES = g.GType(C.gtk_style_properties_get_type())
	GtkType.STYLE_PROVIDER = g.GType(C.gtk_style_provider_get_type())
	GtkType.SWITCH = g.GType(C.gtk_switch_get_type())
	GtkType.SYMBOLIC_COLOR = g.GType(C.gtk_symbolic_color_get_type())
	GtkType.TABLE = g.GType(C.gtk_table_get_type())
	GtkType.TEAROFF_MENU_ITEM = g.GType(C.gtk_tearoff_menu_item_get_type())
	GtkType.TEXT_ATTRIBUTES = g.GType(C.gtk_text_attributes_get_type())
	GtkType.TEXT_BUFFER = g.GType(C.gtk_text_buffer_get_type())
	GtkType.TEXT_CHILD_ANCHOR = g.GType(C.gtk_text_child_anchor_get_type())
	GtkType.TEXT_ITER = g.GType(C.gtk_text_iter_get_type())
	GtkType.TEXT_MARK = g.GType(C.gtk_text_mark_get_type())
	GtkType.TEXT_TAG = g.GType(C.gtk_text_tag_get_type())
	GtkType.TEXT_TAG_TABLE = g.GType(C.gtk_text_tag_table_get_type())
	GtkType.TEXT_VIEW = g.GType(C.gtk_text_view_get_type())
	GtkType.THEMING_ENGINE = g.GType(C.gtk_theming_engine_get_type())
	GtkType.TOGGLE_ACTION = g.GType(C.gtk_toggle_action_get_type())
	GtkType.TOGGLE_BUTTON = g.GType(C.gtk_toggle_button_get_type())
	GtkType.TOGGLE_TOOL_BUTTON = g.GType(C.gtk_toggle_tool_button_get_type())
	GtkType.TOOLBAR = g.GType(C.gtk_toolbar_get_type())
	GtkType.TOOL_BUTTON = g.GType(C.gtk_tool_button_get_type())
	GtkType.TOOL_ITEM_GROUP = g.GType(C.gtk_tool_item_group_get_type())
	GtkType.TOOL_ITEM = g.GType(C.gtk_tool_item_get_type())
	GtkType.TOOL_PALETTE = g.GType(C.gtk_tool_palette_get_type())
	GtkType.TOOL_SHELL = g.GType(C.gtk_tool_shell_get_type())
	GtkType.TOOLTIP = g.GType(C.gtk_tooltip_get_type())
	GtkType.TREE_DRAG_SOURCE = g.GType(C.gtk_tree_drag_source_get_type())
	GtkType.TREE_DRAG_DEST = g.GType(C.gtk_tree_drag_dest_get_type())
	GtkType.TREE_MODEL_FILTER = g.GType(C.gtk_tree_model_filter_get_type())
	GtkType.TREE_MODEL = g.GType(C.gtk_tree_model_get_type())
	GtkType.TREE_ITER = g.GType(C.gtk_tree_iter_get_type())
	GtkType.TREE_PATH = g.GType(C.gtk_tree_path_get_type())
	GtkType.TREE_ROW_REFERENCE = g.GType(C.gtk_tree_row_reference_get_type())
	GtkType.TREE_MODEL_SORT = g.GType(C.gtk_tree_model_sort_get_type())
	GtkType.TREE_SELECTION = g.GType(C.gtk_tree_selection_get_type())
	GtkType.TREE_SORTABLE = g.GType(C.gtk_tree_sortable_get_type())
	GtkType.TREE_STORE = g.GType(C.gtk_tree_store_get_type())
	GtkType.TREE_VIEW_COLUMN = g.GType(C.gtk_tree_view_column_get_type())
	GtkType.TREE_VIEW = g.GType(C.gtk_tree_view_get_type())
	GtkType.UI_MANAGER = g.GType(C.gtk_ui_manager_get_type())
	GtkType.VBUTTON_BOX = g.GType(C.gtk_vbutton_box_get_type())
	GtkType.VBOX = g.GType(C.gtk_vbox_get_type())
	GtkType.VIEWPORT = g.GType(C.gtk_viewport_get_type())
	GtkType.VOLUME_BUTTON = g.GType(C.gtk_volume_button_get_type())
	GtkType.VPANED = g.GType(C.gtk_vpaned_get_type())
	GtkType.VSCALE = g.GType(C.gtk_vscale_get_type())
	GtkType.VSCROLLBAR = g.GType(C.gtk_vscrollbar_get_type())
	GtkType.VSEPARATOR = g.GType(C.gtk_vseparator_get_type())
	GtkType.WIDGET = g.GType(C.gtk_widget_get_type())
	GtkType.REQUISITION = g.GType(C.gtk_requisition_get_type())
	GtkType.WIDGET_PATH = g.GType(C.gtk_widget_path_get_type())
	GtkType.WINDOW = g.GType(C.gtk_window_get_type())
	GtkType.WINDOW_GROUP = g.GType(C.gtk_window_group_get_type())
	GtkType.LICENSE = g.GType(C.gtk_license_get_type())
	GtkType.ACCEL_FLAGS = g.GType(C.gtk_accel_flags_get_type())
	GtkType.ASSISTANT_PAGE_TYPE = g.GType(C.gtk_assistant_page_type_get_type())
	GtkType.BUILDER_ERROR = g.GType(C.gtk_builder_error_get_type())
	GtkType.CALENDAR_DISPLAY_OPTIONS = g.GType(C.gtk_calendar_display_options_get_type())
	GtkType.CELL_RENDERER_STATE = g.GType(C.gtk_cell_renderer_state_get_type())
	GtkType.CELL_RENDERER_MODE = g.GType(C.gtk_cell_renderer_mode_get_type())
	GtkType.CELL_RENDERER_ACCEL_MODE = g.GType(C.gtk_cell_renderer_accel_mode_get_type())
	GtkType.CSS_PROVIDER_ERROR = g.GType(C.gtk_css_provider_error_get_type())
	GtkType.DEBUG_FLAG = g.GType(C.gtk_debug_flag_get_type())
	GtkType.DIALOG_FLAGS = g.GType(C.gtk_dialog_flags_get_type())
	GtkType.RESPONSE_TYPE = g.GType(C.gtk_response_type_get_type())
	GtkType.DEST_DEFAULTS = g.GType(C.gtk_dest_defaults_get_type())
	GtkType.TARGET_FLAGS = g.GType(C.gtk_target_flags_get_type())
	GtkType.ENTRY_ICON_POSITION = g.GType(C.gtk_entry_icon_position_get_type())
	GtkType.ALIGN = g.GType(C.gtk_align_get_type())
	GtkType.ARROW_PLACEMENT = g.GType(C.gtk_arrow_placement_get_type())
	GtkType.ARROW_TYPE = g.GType(C.gtk_arrow_type_get_type())
	GtkType.ATTACH_OPTIONS = g.GType(C.gtk_attach_options_get_type())
	GtkType.BUTTON_BOX_STYLE = g.GType(C.gtk_button_box_style_get_type())
	GtkType.DELETE_TYPE = g.GType(C.gtk_delete_type_get_type())
	GtkType.DIRECTION_TYPE = g.GType(C.gtk_direction_type_get_type())
	GtkType.EXPANDER_STYLE = g.GType(C.gtk_expander_style_get_type())
	GtkType.ICON_SIZE = g.GType(C.gtk_icon_size_get_type())
	GtkType.SENSITIVITY_TYPE = g.GType(C.gtk_sensitivity_type_get_type())
	GtkType.TEXT_DIRECTION = g.GType(C.gtk_text_direction_get_type())
	GtkType.JUSTIFICATION = g.GType(C.gtk_justification_get_type())
	GtkType.MENU_DIRECTION_TYPE = g.GType(C.gtk_menu_direction_type_get_type())
	GtkType.MESSAGE_TYPE = g.GType(C.gtk_message_type_get_type())
	GtkType.MOVEMENT_STEP = g.GType(C.gtk_movement_step_get_type())
	GtkType.SCROLL_STEP = g.GType(C.gtk_scroll_step_get_type())
	GtkType.ORIENTATION = g.GType(C.gtk_orientation_get_type())
	GtkType.CORNER_TYPE = g.GType(C.gtk_corner_type_get_type())
	GtkType.PACK_TYPE = g.GType(C.gtk_pack_type_get_type())
	GtkType.PATH_PRIORITY_TYPE = g.GType(C.gtk_path_priority_type_get_type())
	GtkType.PATH_TYPE = g.GType(C.gtk_path_type_get_type())
	GtkType.POLICY_TYPE = g.GType(C.gtk_policy_type_get_type())
	GtkType.POSITION_TYPE = g.GType(C.gtk_position_type_get_type())
	GtkType.RELIEF_STYLE = g.GType(C.gtk_relief_style_get_type())
	GtkType.RESIZE_MODE = g.GType(C.gtk_resize_mode_get_type())
	GtkType.SCROLL_TYPE = g.GType(C.gtk_scroll_type_get_type())
	GtkType.SELECTION_MODE = g.GType(C.gtk_selection_mode_get_type())
	GtkType.SHADOW_TYPE = g.GType(C.gtk_shadow_type_get_type())
	GtkType.STATE_TYPE = g.GType(C.gtk_state_type_get_type())
	GtkType.TOOLBAR_STYLE = g.GType(C.gtk_toolbar_style_get_type())
	GtkType.WINDOW_POSITION = g.GType(C.gtk_window_position_get_type())
	GtkType.WINDOW_TYPE = g.GType(C.gtk_window_type_get_type())
	GtkType.WRAP_MODE = g.GType(C.gtk_wrap_mode_get_type())
	GtkType.SORT_TYPE = g.GType(C.gtk_sort_type_get_type())
	GtkType.IM_PREEDIT_STYLE = g.GType(C.gtk_im_preedit_style_get_type())
	GtkType.IM_STATUS_STYLE = g.GType(C.gtk_im_status_style_get_type())
	GtkType.PACK_DIRECTION = g.GType(C.gtk_pack_direction_get_type())
	GtkType.PRINT_PAGES = g.GType(C.gtk_print_pages_get_type())
	GtkType.PAGE_SET = g.GType(C.gtk_page_set_get_type())
	GtkType.NUMBER_UP_LAYOUT = g.GType(C.gtk_number_up_layout_get_type())
	GtkType.PAGE_ORIENTATION = g.GType(C.gtk_page_orientation_get_type())
	GtkType.PRINT_QUALITY = g.GType(C.gtk_print_quality_get_type())
	GtkType.PRINT_DUPLEX = g.GType(C.gtk_print_duplex_get_type())
	GtkType.UNIT = g.GType(C.gtk_unit_get_type())
	GtkType.TREE_VIEW_GRID_LINES = g.GType(C.gtk_tree_view_grid_lines_get_type())
	GtkType.DRAG_RESULT = g.GType(C.gtk_drag_result_get_type())
	GtkType.SIZE_REQUEST_MODE = g.GType(C.gtk_size_request_mode_get_type())
	GtkType.SCROLLABLE_POLICY = g.GType(C.gtk_scrollable_policy_get_type())
	GtkType.STATE_FLAGS = g.GType(C.gtk_state_flags_get_type())
	GtkType.REGION_FLAGS = g.GType(C.gtk_region_flags_get_type())
	GtkType.JUNCTION_SIDES = g.GType(C.gtk_junction_sides_get_type())
	GtkType.BORDER_STYLE = g.GType(C.gtk_border_style_get_type())
	GtkType.FILE_CHOOSER_ACTION = g.GType(C.gtk_file_chooser_action_get_type())
	GtkType.FILE_CHOOSER_CONFIRMATION = g.GType(C.gtk_file_chooser_confirmation_get_type())
	GtkType.FILE_CHOOSER_ERROR = g.GType(C.gtk_file_chooser_error_get_type())
	GtkType.FILE_FILTER_FLAGS = g.GType(C.gtk_file_filter_flags_get_type())
	GtkType.ICON_LOOKUP_FLAGS = g.GType(C.gtk_icon_lookup_flags_get_type())
	GtkType.ICON_THEME_ERROR = g.GType(C.gtk_icon_theme_error_get_type())
	GtkType.ICON_VIEW_DROP_POSITION = g.GType(C.gtk_icon_view_drop_position_get_type())
	GtkType.IMAGE_TYPE = g.GType(C.gtk_image_type_get_type())
	GtkType.BUTTONS_TYPE = g.GType(C.gtk_buttons_type_get_type())
	GtkType.NOTEBOOK_TAB = g.GType(C.gtk_notebook_tab_get_type())
	GtkType.PRINT_STATUS = g.GType(C.gtk_print_status_get_type())
	GtkType.PRINT_OPERATION_RESULT = g.GType(C.gtk_print_operation_result_get_type())
	GtkType.PRINT_OPERATION_ACTION = g.GType(C.gtk_print_operation_action_get_type())
	GtkType.PRINT_ERROR = g.GType(C.gtk_print_error_get_type())
	GtkType.RC_FLAGS = g.GType(C.gtk_rc_flags_get_type())
	GtkType.RC_TOKEN_TYPE = g.GType(C.gtk_rc_token_type_get_type())
	GtkType.RECENT_SORT_TYPE = g.GType(C.gtk_recent_sort_type_get_type())
	GtkType.RECENT_CHOOSER_ERROR = g.GType(C.gtk_recent_chooser_error_get_type())
	GtkType.RECENT_FILTER_FLAGS = g.GType(C.gtk_recent_filter_flags_get_type())
	GtkType.RECENT_MANAGER_ERROR = g.GType(C.gtk_recent_manager_error_get_type())
	GtkType.SIZE_GROUP_MODE = g.GType(C.gtk_size_group_mode_get_type())
	GtkType.SPIN_BUTTON_UPDATE_POLICY = g.GType(C.gtk_spin_button_update_policy_get_type())
	GtkType.SPIN_TYPE = g.GType(C.gtk_spin_type_get_type())
	GtkType.TEXT_BUFFER_TARGET_INFO = g.GType(C.gtk_text_buffer_target_info_get_type())
	GtkType.TEXT_SEARCH_FLAGS = g.GType(C.gtk_text_search_flags_get_type())
	GtkType.TEXT_WINDOW_TYPE = g.GType(C.gtk_text_window_type_get_type())
	GtkType.TOOLBAR_SPACE_STYLE = g.GType(C.gtk_toolbar_space_style_get_type())
	GtkType.TOOL_PALETTE_DRAG_TARGETS = g.GType(C.gtk_tool_palette_drag_targets_get_type())
	GtkType.TREE_MODEL_FLAGS = g.GType(C.gtk_tree_model_flags_get_type())
	GtkType.TREE_VIEW_DROP_POSITION = g.GType(C.gtk_tree_view_drop_position_get_type())
	GtkType.TREE_VIEW_COLUMN_SIZING = g.GType(C.gtk_tree_view_column_sizing_get_type())
	GtkType.UI_MANAGER_ITEM_TYPE = g.GType(C.gtk_ui_manager_item_type_get_type())
	GtkType.WIDGET_HELP_TYPE = g.GType(C.gtk_widget_help_type_get_type())

	// Initialize GtkShadowType
	GtkShadow.NONE = 0
	GtkShadow.IN = 1
	GtkShadow.OUT = 2
	GtkShadow.ETCHED_IN = 3
	GtkShadow.ETCHED_OUT = 4

	// Initialize GtkStock
	GtkStock.ABOUT = "gtk-about"
	GtkStock.ADD = "gtk-add"
	GtkStock.APPLY = "gtk-apply"
	GtkStock.BOLD = "gtk-bold"
	GtkStock.CANCEL = "gtk-cancel"
	GtkStock.CAPS_LOCK_WARNING = "gtk-caps-lock-warning"
	GtkStock.CDROM = "gtk-cdrom"
	GtkStock.CLEAR = "gtk-clear"
	GtkStock.CLOSE = "gtk-close"
	GtkStock.COLOR_PICKER = "gtk-color-picker"
	GtkStock.CONNECT = "gtk-connect"
	GtkStock.CONVERT = "gtk-convert"
	GtkStock.COPY = "gtk-copy"
	GtkStock.CUT = "gtk-cut"
	GtkStock.DELETE = "gtk-delete"
	GtkStock.DIALOG_AUTHENTICATION = "gtk-dialog-authentication"
	GtkStock.DIALOG_INFO = "gtk-dialog-info"
	GtkStock.DIALOG_WARNING = "gtk-dialog-warning"
	GtkStock.DIALOG_ERROR = "gtk-dialog-error"
	GtkStock.DIALOG_QUESTION = "gtk-dialog-question"
	GtkStock.DIRECTORY = "gtk-directory"
	GtkStock.DISCARD = "gtk-discard"
	GtkStock.DISCONNECT = "gtk-disconnect"
	GtkStock.DND = "gtk-dnd"
	GtkStock.DND_MULTIPLE = "gtk-dnd-multiple"
	GtkStock.EDIT = "gtk-edit"
	GtkStock.EXECUTE = "gtk-execute"
	GtkStock.FILE = "gtk-file"
	GtkStock.FIND = "gtk-find"
	GtkStock.FIND_AND_REPLACE = "gtk-find-and-replace"
	GtkStock.FLOPPY = "gtk-floppy"
	GtkStock.FULLSCREEN = "gtk-fullscreen"
	GtkStock.GOTO_BOTTOM = "gtk-goto-bottom"
	GtkStock.GOTO_FIRST = "gtk-goto-first"
	GtkStock.GOTO_LAST = "gtk-goto-last"
	GtkStock.GOTO_TOP = "gtk-goto-top"
	GtkStock.GO_BACK = "gtk-go-back"
	GtkStock.GO_DOWN = "gtk-go-down"
	GtkStock.GO_FORWARD = "gtk-go-forward"
	GtkStock.GO_UP = "gtk-go-up"
	GtkStock.HARDDISK = "gtk-harddisk"
	GtkStock.HELP = "gtk-help"
	GtkStock.HOME = "gtk-home"
	GtkStock.INDEX = "gtk-index"
	GtkStock.INDENT = "gtk-indent"
	GtkStock.INFO = "gtk-info"
	GtkStock.ITALIC = "gtk-italic"
	GtkStock.JUMP_TO = "gtk-jump-to"
	GtkStock.JUSTIFY_CENTER = "gtk-justify-center"
	GtkStock.JUSTIFY_FILL = "gtk-justify-fill"
	GtkStock.JUSTIFY_LEFT = "gtk-justify-left"
	GtkStock.JUSTIFY_RIGHT = "gtk-justify-right"
	GtkStock.LEAVE_FULLSCREEN = "gtk-leave-fullscreen"
	GtkStock.MISSING_IMAGE = "gtk-missing-image"
	GtkStock.MEDIA_FORWARD = "gtk-media-forward"
	GtkStock.MEDIA_NEXT = "gtk-media-next"
	GtkStock.MEDIA_PAUSE = "gtk-media-pause"
	GtkStock.MEDIA_PLAY = "gtk-media-play"
	GtkStock.MEDIA_PREVIOUS = "gtk-media-previous"
	GtkStock.MEDIA_RECORD = "gtk-media-record"
	GtkStock.MEDIA_REWIND = "gtk-media-rewind"
	GtkStock.MEDIA_STOP = "gtk-media-stop"
	GtkStock.NETWORK = "gtk-network"
	GtkStock.NEW = "gtk-new"
	GtkStock.NO = "gtk-no"
	GtkStock.OK = "gtk-ok"
	GtkStock.OPEN = "gtk-open"
	GtkStock.ORIENTATION_PORTRAIT = "gtk-orientation-portrait"
	GtkStock.ORIENTATION_LANDSCAPE = "gtk-orientation-landscape"
	GtkStock.ORIENTATION_REVERSE_LANDSCAPE = "gtk-orientation-reverse-landscape"
	GtkStock.ORIENTATION_REVERSE_PORTRAIT = "gtk-orientation-reverse-portrait"
	GtkStock.PAGE_SETUP = "gtk-page-setup"
	GtkStock.PASTE = "gtk-paste"
	GtkStock.PREFERENCES = "gtk-preferences"
	GtkStock.PRINT = "gtk-print"
	GtkStock.PRINT_ERROR = "gtk-print-error"
	GtkStock.PRINT_PAUSED = "gtk-print-paused"
	GtkStock.PRINT_PREVIEW = "gtk-print-preview"
	GtkStock.PRINT_REPORT = "gtk-print-report"
	GtkStock.PRINT_WARNING = "gtk-print-warning"
	GtkStock.PROPERTIES = "gtk-properties"
	GtkStock.QUIT = "gtk-quit"
	GtkStock.REDO = "gtk-redo"
	GtkStock.REFRESH = "gtk-refresh"
	GtkStock.REMOVE = "gtk-remove"
	GtkStock.REVERT_TO_SAVED = "gtk-revert-to-saved"
	GtkStock.SAVE = "gtk-save"
	GtkStock.SAVE_AS = "gtk-save-as"
	GtkStock.SELECT_ALL = "gtk-select-all"
	GtkStock.SELECT_COLOR = "gtk-select-color"
	GtkStock.SELECT_FONT = "gtk-select-font"
	GtkStock.SORT_ASCENDING = "gtk-sort-ascending"
	GtkStock.SORT_DESCENDING = "gtk-sort-descending"
	GtkStock.SPELL_CHECK = "gtk-spell-check"
	GtkStock.STOP = "gtk-stop"
	GtkStock.STRIKETHROUGH = "gtk-strikethrough"
	GtkStock.UNDELETE = "gtk-undelete"
	GtkStock.UNDERLINE = "gtk-underline"
	GtkStock.UNDO = "gtk-undo"
	GtkStock.UNINDENT = "gtk-unindent"
	GtkStock.YES = "gtk-yes"
	GtkStock.ZOOM_100 = "gtk-zoom-100"
	GtkStock.ZOOM_FIT = "gtk-zoom-fit"
	GtkStock.ZOOM_IN = "gtk-zoom-in"
	GtkStock.ZOOM_OUT = "gtk-zoom-out"

	// Initialize GtkPosition
	GtkPosition.LEFT = 0
	GtkPosition.RIGHT = 1
	GtkPosition.TOP = 2
	GtkPosition.BOTTOM = 3

	// Initialize GtkDialogFlags
	GtkDialogFlags.MODAL = 1 << 0
	GtkDialogFlags.DESTROY_WITH_PARENT = 1 << 1

	// Initialize GtkResponse
	GtkResponse.NONE = -1
	GtkResponse.REJECT = -2
	GtkResponse.ACCEPT = -3
	GtkResponse.DELETE_EVENT = -4
	GtkResponse.OK = -5
	GtkResponse.CANCEL = -6
	GtkResponse.CLOSE = -7
	GtkResponse.YES = -8
	GtkResponse.NO = -9
	GtkResponse.APPLY = -10
	GtkResponse.HELP = -11

	// Initialize GtkMessage
	GtkMessage.INFO = 0
	GtkMessage.WARNING = 1
	GtkMessage.QUESTION = 2
	GtkMessage.ERROR = 3
	GtkMessage.OTHER = 4

	// Initialize GtkButtons
	GtkButtons.NONE = 0
	GtkButtons.OK = 1
	GtkButtons.CLOSE = 2
	GtkButtons.CANCEL = 3
	GtkButtons.YES_NO = 4
	GtkButtons.OK_CANCEL = 5

	// Initialize GtkOrientation
	GtkOrientation.HORIZONTAL = 0
	GtkOrientation.VERTICAL = 1

	// Initialize GtkPackType
	GtkPack.PACK_START = 0
	GtkPack.PACK_END = 1

	// Initialize GtkButtonBoxStyle
	GtkButtonBoxStyle.SPREAD = 1
	GtkButtonBoxStyle.EDGE = 2
	GtkButtonBoxStyle.START = 3
	GtkButtonBoxStyle.END = 4
	GtkButtonBoxStyle.CENTER = 5

	// Initialize GtkIconSize
	GtkIconSize.INVALID = 0
	GtkIconSize.MENU = 1
	GtkIconSize.SMALL_TOOLBAR = 2
	GtkIconSize.LARGE_TOOLBAR = 3
	GtkIconSize.BUTTON = 4
	GtkIconSize.DND = 5
	GtkIconSize.DIALOG = 6

	// Initialize GApplicationFlags
	GApplicationFlags.FLAGS_NONE = 0
	GApplicationFlags.IS_SERVICE = 1 << 0
	GApplicationFlags.IS_LAUNCHER = 1 << 1
	GApplicationFlags.HANDLES_OPEN = 1 << 2
	GApplicationFlags.HANDLES_COMMAND_LINE = 1 << 3
	GApplicationFlags.SEND_ENVIRONMENT = 1 << 4
	GApplicationFlags.NON_UNIQUE = 1 << 5

	// Initialize GtkJustification
	GtkJustification.LEFT = 0
	GtkJustification.RIGHT = 1
	GtkJustification.CENTER = 2
	GtkJustification.FILL = 3

	// Initialize GtkWindowType
	GtkWindowType.TOPLEVEL = 0
	GtkWindowType.POPUP = 1

	// Initialize GtkSensitivity
	GtkSensitivity.AUTO = 0
	GtkSensitivity.ON = 1
	GtkSensitivity.OFF = 2

	// Initialize GtkPolicy
	GtkPolicy.ALWAYS = 0
	GtkPolicy.AUTOMATIC = 1
	GtkPolicy.NEVER = 2

	// Initialize GtkCorner
	GtkCorner.TOP_LEFT = 0
	GtkCorner.BOTTOM_LEFT = 1
	GtkCorner.TOP_RIGHT = 2
	GtkCorner.BOTTOM_RIGHT = 3
}
