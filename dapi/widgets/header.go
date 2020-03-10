package widgets

import "github.com/rivo/tview"

// Header is a widget displayed at the top of the app
type Header struct {
	Title *tview.TextView
}

func newHeader() *Header {
	header := Header{}
	header.Title = tview.NewTextView()
	header.Title.SetBackgroundColor(BackgroundColor)
	header.Title.SetDynamicColors(true)
	return &header
}

// Root returns the root element
func (h *Header) Root() tview.Primitive {
	return h.Title
}
