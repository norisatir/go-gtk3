package main

import "github.com/norisatir/go-gtk3/gtk3"
import "github.com/norisatir/go-gtk3/glib"
import "fmt"

func main() {
	gtk3.Init()

	assistant := gtk3.NewAssistant()
	assistant.SetDefaultSize(-1, 300)

	CreatePage1(assistant)
	CreatePage2(assistant)
	CreatePage3(assistant)
	progress := CreatePage4(assistant)
	progress.SetShowText(true)
	progress.SetText(nil)

	assistant.Connect("destroy", gtk3.MainQuit)
	assistant.Connect("cancel", func() { assistant.Destroy()})
	assistant.Connect("close", func() { assistant.Destroy()})
	assistant.Connect("apply", OnAssistantApply, progress, assistant) 
	assistant.Connect("prepare", OnAssistantPrepare)

	assistant.Show()
	gtk3.Main()
}

func OnAssistantApply(progressBar *gtk3.ProgressBar, assistant *gtk3.Assistant) {
	glib.GTimeoutAddFull(glib.GPriority.DEFAULT, 100, ApplyChangesGradually, progressBar, assistant)
}
func ApplyChangesGradually(progressBar *gtk3.ProgressBar, assistant *gtk3.Assistant) bool {
	fraction := progressBar.GetFraction()
	fraction = fraction + 0.05
	if fraction < 1.0 {
		progressBar.SetFraction(fraction)
		return true
	}

	assistant.Destroy()
	return false
}

func OnAssistantPrepare(data ...interface{}) {
	assistant := data[0].(*gtk3.Assistant)
	currentPage := assistant.GetCurrentPage()
	nPages := assistant.GetNPages()
	assistant.SetTitle(fmt.Sprintf("Sample assistant (%d of %d)", currentPage+1, nPages))

  // The fourth page (counting from zero) is the progress page.  The
  // user clicked Apply to get here so we tell the assistant to commit,
  // which means the changes up to this point are permanent and cannot
  // be cancelled or revisited.
  if currentPage == 3 {
	  assistant.Commit()
  }
}

func OnEntryChanged(assistant *gtk3.Assistant, data ...interface{}) {
	pageNumber := assistant.GetCurrentPage()
	currentPage := assistant.GetNthPage(pageNumber)

	entry := data[0].(*gtk3.Entry)
	textLen := entry.GetTextLength()

	if textLen > 0 {
		assistant.SetPageComplete(currentPage, true)
	} else {
		assistant.SetPageComplete(currentPage, false)
	}
}
func CreatePage1(assistant *gtk3.Assistant) {
	box := gtk3.NewHBox(12)
	box.SetBorderWidth(12)

	label := gtk3.NewLabel("You must fill out this entry to continue")
	box.PackStart(label, false, false, 0)

	entry := gtk3.NewEntry()
	entry.SetActivatesDefault(true)
	box.PackStart(entry, true, true, 0)
	entry.Connect("changed", OnEntryChanged, assistant)

	box.ShowAll()

	assistant.AppendPage(box)
	assistant.SetPageTitle(box, "Page 1")
	assistant.SetPageType(box, gtk3.GtkAssistantPage.INTRO)
}

func CreatePage2(assistant *gtk3.Assistant) {
	box := gtk3.NewVBox(12)
	box.SetBorderWidth(12)

	checkButton := gtk3.NewCheckButtonWithLabel("This is optional data, you may continue even if you do not check this")
	box.PackStart(checkButton, false, false, 0)

	box.ShowAll()
	assistant.AppendPage(box)
	assistant.SetPageComplete(box, true)
	assistant.SetPageTitle(box, "Page 2")
}

func CreatePage3(assistant *gtk3.Assistant) {
	label := gtk3.NewLabel("This is a confirmation page, press 'Apply' to apply changes")

	label.Show()
	assistant.AppendPage(label)
	assistant.SetPageType(label, gtk3.GtkAssistantPage.CONFIRM)
	assistant.SetPageComplete(label, true)
	assistant.SetPageTitle(label, "Confirmation")
}

func CreatePage4(assistant *gtk3.Assistant) *gtk3.ProgressBar {
	progressBar := gtk3.NewProgressBar()
	progressBar.SetHalign(gtk3.GtkAlign.CENTER)
	progressBar.SetValign(gtk3.GtkAlign.CENTER)

	progressBar.Show()
	assistant.AppendPage(progressBar)
	assistant.SetPageType(progressBar, gtk3.GtkAssistantPage.PROGRESS)
	assistant.SetPageTitle(progressBar, "Applying changes")

	// this prevents assistant window from being closed while we're "busy" applying changes
	assistant.SetPageComplete(progressBar, false)
	return progressBar
}
