package assistant

import "github.com/norisatir/go-gtk3/gtk3"
import "github.com/norisatir/go-gtk3/glib"
import "fmt"

var assistant *gtk3.Assistant = nil

func DoAssistant(w gtk3.WidgetLike) gtk3.WidgetLike {
	if assistant == nil {
		assistant = gtk3.NewAssistant()
		assistant.SetDefaultSize(-1, 300)

		createPage1(assistant)
		createPage2(assistant)
		createPage3(assistant)
		progress := createPage4(assistant)
		progress.SetShowText(true)
		progress.SetText(nil)

		assistant.Connect("destroy", func() { assistant = nil })
		assistant.Connect("cancel", func() { assistant.Destroy() })
		assistant.Connect("close", func() { assistant.Destroy() })
		assistant.Connect("apply", onAssistantApply, progress, assistant)
		assistant.Connect("prepare", onAssistantPrepare)
	}

	if !assistant.GetVisible() {
		assistant.Show()
	} else {
		assistant.Destroy()
		assistant = nil
		return nil
	}
	return assistant
}

func onAssistantApply(progressBar *gtk3.ProgressBar, assistant *gtk3.Assistant) {
	glib.GTimeoutAddFull(glib.GPriority.DEFAULT, 100, applyChangesGradually, progressBar, assistant)
}
func applyChangesGradually(progressBar *gtk3.ProgressBar, assistant *gtk3.Assistant) bool {
	fraction := progressBar.GetFraction()
	fraction = fraction + 0.05
	if fraction < 1.0 {
		progressBar.SetFraction(fraction)
		return true
	}

	assistant.Destroy()
	return false
}

func onAssistantPrepare(data ...interface{}) {
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

func onEntryChanged(assistant *gtk3.Assistant, data ...interface{}) {
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
func createPage1(assistant *gtk3.Assistant) {
	box := gtk3.NewHBox(12)
	box.SetBorderWidth(12)

	label := gtk3.NewLabel("You must fill out this entry to continue")
	box.PackStart(label, false, false, 0)

	entry := gtk3.NewEntry()
	entry.SetActivatesDefault(true)
	box.PackStart(entry, true, true, 0)
	entry.Connect("changed", onEntryChanged, assistant)

	box.ShowAll()

	assistant.AppendPage(box)
	assistant.SetPageTitle(box, "Page 1")
	assistant.SetPageType(box, gtk3.GtkAssistantPage.INTRO)
}

func createPage2(assistant *gtk3.Assistant) {
	box := gtk3.NewVBox(12)
	box.SetBorderWidth(12)

	checkButton := gtk3.NewCheckButtonWithLabel("This is optional data, you may continue even if you do not check this")
	box.PackStart(checkButton, false, false, 0)

	box.ShowAll()
	assistant.AppendPage(box)
	assistant.SetPageComplete(box, true)
	assistant.SetPageTitle(box, "Page 2")
}

func createPage3(assistant *gtk3.Assistant) {
	label := gtk3.NewLabel("This is a confirmation page, press 'Apply' to apply changes")

	label.Show()
	assistant.AppendPage(label)
	assistant.SetPageType(label, gtk3.GtkAssistantPage.CONFIRM)
	assistant.SetPageComplete(label, true)
	assistant.SetPageTitle(label, "Confirmation")
}

func createPage4(assistant *gtk3.Assistant) *gtk3.ProgressBar {
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
