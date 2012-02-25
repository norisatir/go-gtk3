#include "go-gtk3.h"

void _gtk_entry_completion_set_match_func(GtkEntryCompletion* completion, gint64 func_data) {
	gint64* data = (gint64*)malloc(sizeof(gint64));
	*data = func_data;

	gtk_entry_completion_set_match_func(completion, _gtk_entry_completion_match_func, (gpointer)data, _g_gtk_destroy_notify);
}
