#ifndef _GO_GTK3_H
#define _GO_GTK3_H

#include <stdlib.h>
#include <gtk/gtkx.h>

// Typedefs
typedef const gchar* constgchar;
typedef const GValue* constgvalue;
typedef GtkTreeModel* goTreeModel;
typedef GtkTreePath* goTreePath;
typedef GtkCellLayout* goCellLayout;

extern void _gtk_callback(GtkWidget* widget, gpointer data);

extern gboolean _gtk_entry_completion_match_func(GtkEntryCompletion* completion,
											 constgchar key,
											 GtkTreeIter* iter,
											 gpointer user_data);

extern void _gtk_text_tag_table_foreach_func(GtkTextTag* tag, gpointer data);

extern gboolean _gtk_text_char_predicate(gunichar ch, gpointer data);

extern gboolean _gtk_cell_callback(GtkCellRenderer* renderer, gpointer data);

extern void _gtk_cell_layout_data_func(GtkCellLayout* cell_layout,
									GtkCellRenderer* cell,
									GtkTreeModel* tree_model,
									GtkTreeIter* iter,
									gpointer data);

extern void _g_gtk_destroy_notify(gpointer data);

extern gboolean _gtk_tree_view_row_separator_func(GtkTreeModel* model,
									GtkTreeIter* iter,
									gpointer data);

extern void _gtk_closure_destroy_id(gpointer data, GClosure* closure);

extern void _gtk_marshal(GClosure *closure,
							  GValue* returnValue,
							  guint n_param_values,
							  constgvalue paramValues,
							  gpointer invocationHint,
							  gpointer marshalData);

extern void _gtk_menu_position_func(GtkMenu* menu, gint* x, gint* y, gboolean* push_in, gpointer user_data);

extern void _gtk_menu_detach_func(GtkWidget* attach_widget, GtkMenu* menu);

extern gboolean _gtk_tree_selection_func(GtkTreeSelection*, goTreeModel, goTreePath,
	gboolean, gpointer);

// Funcs defined in go-gtk3.c
void _gtk_entry_completion_set_match_func(GtkEntryCompletion* completion, gint64 func_data);

#endif
