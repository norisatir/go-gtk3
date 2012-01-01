include $(GOROOT)/src/Make.inc

DEMOS=\
	assistant.go \
	button_box.go \
	combobox.go \
	entry_buffer.go \
	entry_completion.go \
	search_entry.go \
	dialog.go \
	menu.go \
	spinner.go \
	textview.go \
	list_store.go \
	tree_store.go \


PREREQ=$(patsubst %.go,%.$O,$(DEMOS))
TARG=demo
GOFILES=\
	demo.go \


include $(GOROOT)/src/Make.cmd

%.$O: %.go
	$(GC) -o $*.$O $*.go
